package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/go-resty/resty/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	config "github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	apimgr "github.com/NpoolPlatform/message/npool/apimgr"

	constant "github.com/NpoolPlatform/api-manager/pkg/message/const"
	serviceapi "github.com/NpoolPlatform/api-manager/pkg/middleware/service-api"

	"google.golang.org/grpc"
)

const (
	protocolHTTP = "http"
	protocolGrpc = "grpc"
)

func reliableRegister(apis *apimgr.ServiceApis) {
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()

	for range tick.C {
		if err := func() error {
			conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
			if err != nil {
				return err
			}
			defer conn.Close()

			cli := apimgr.NewApiManagerClient(conn)

			ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cancel()

			_, err = cli.Register(ctx, &apimgr.RegisterRequest{
				Info: apis,
			})
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			logger.Sugar().Errorf("fail register apis: %v", err)
			continue
		}
		break
	}
}

func MuxApis(mux *runtime.ServeMux) *apimgr.ServiceApis {
	apis := &apimgr.ServiceApis{
		ServiceName: config.GetStringValueWithNameSpace("", config.KeyHostname),
		Protocol:    protocolHTTP,
	}

	valueOfMux := reflect.ValueOf(mux).Elem()
	handlers := valueOfMux.FieldByName("handlers")
	methIter := handlers.MapRange()
	for methIter.Next() {
		for i := 0; i < methIter.Value().Len(); i++ {
			pat := methIter.Value().Index(i).FieldByName("pat")
			tmp := reflect.NewAt(pat.Type(), unsafe.Pointer(pat.UnsafeAddr())).Elem()
			str := tmp.MethodByName("String").Call(nil)[0].String()
			apis.Paths = append(apis.Paths, &apimgr.Path{
				Method:     methIter.Key().String(),
				Path:       str,
				MethodName: "NONE",
			})
		}
	}

	return apis
}

func GrpcApis(server grpc.ServiceRegistrar) *apimgr.ServiceApis {
	srvInfo := server.(*grpc.Server).GetServiceInfo()
	apis := &apimgr.ServiceApis{
		ServiceName: config.GetStringValueWithNameSpace("", config.KeyHostname),
		Protocol:    protocolGrpc,
	}

	for key, info := range srvInfo {
		for _, method := range info.Methods {
			apis.Paths = append(apis.Paths, &apimgr.Path{
				Method:     "NONE",
				Path:       fmt.Sprintf("%v/%v", key, method.Name),
				MethodName: method.Name,
			})
		}
	}

	return apis
}

func getGatewayRouters(name string) ([]serviceapi.EntryPoint, error) {
	domain := strings.SplitN(name, ".", 2)
	if len(domain) < 2 {
		return nil, errors.New("service name must like example.npool.top")
	}

	// provider can use kubernetes or k8s
	url := fmt.Sprintf(
		"http://traefik.kube-system.svc.cluster.local:38080/api/http/routers?provider=kubernetes&search=%v",
		domain[0],
	)

	// internal already set timeout
	resp, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}

	routers := make([]serviceapi.EntryPoint, 0)
	err = json.Unmarshal(resp.Body(), &routers)
	if err != nil {
		return nil, err
	}

	return routers, nil
}

func Register(mux *runtime.ServeMux) error {
	apis := MuxApis(mux)
	pathMap := make(map[string]struct{})
	for _, api := range apis.Paths {
		// here for every host has one same path
		// host1 patha
		// host2 patha
		pathMap[api.Path] = struct{}{}
	}

	gatewayRouters, err := getGatewayRouters(apis.ServiceName)
	if err != nil {
		return err
	}

	for _, router := range gatewayRouters {
		prefix, err := router.PathPrefix()
		if err != nil {
			return err
		}

		// TODO router path check and in gateway already register
		// path, err := router.Path()
		// if err != nil {
		// 	return err
		// }
		// if _, ok := pathMap[path]; !ok {
		// 	return err
		// 	// logger.Sugar().Warn("some api not export")
		// }
		apis.PathPrefix = prefix
		apis.Exported = true
		apis.Domains = append(apis.Domains, router.Domain())
	}
	go reliableRegister(apis)
	return nil
}

func RegisterGRPC(server grpc.ServiceRegistrar) {
	apis := GrpcApis(server)
	go reliableRegister(apis)
}

func do(ctx context.Context, fn func(_ctx context.Context, cli apimgr.ApiManagerClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get stock connection: %v", err)
	}
	defer conn.Close()

	cli := apimgr.NewApiManagerClient(conn)

	return fn(_ctx, cli)
}

func GetServiceMethodAPI(ctx context.Context, serviceName, methodName string) (*apimgr.ServicePath, error) {
	info, err := do(ctx, func(_ctx context.Context, cli apimgr.ApiManagerClient) (cruder.Any, error) {
		resp, err := cli.GetServiceMethodAPI(ctx, &apimgr.GetServiceMethodAPIRequest{
			ServiceName: serviceName,
			MethodName:  methodName,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get api: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get api: %v", err)
	}
	return info.(*apimgr.ServicePath), nil
}
