package client

import (
	"context"
	"reflect"
	"time"
	"unsafe"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	config "github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	apimgr "github.com/NpoolPlatform/message/npool/apimgr"

	constant "github.com/NpoolPlatform/api-manager/pkg/message/const"
)

func reliableRegister(apis *apimgr.ServiceApis) {
	for {
		conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
		if err != nil {
			logger.Sugar().Errorf("fail get api manager connection: %v", err)
			time.Sleep(time.Minute)
			continue
		}

		cli := apimgr.NewApiManagerClient(conn)

		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)

		_, err = cli.Register(ctx, &apimgr.RegisterRequest{
			Info: apis,
		})
		if err == nil {
			cancel()
			conn.Close()
			return
		}

		logger.Sugar().Errorf("fail register apis: %v", err)
		time.Sleep(time.Minute)

		cancel()
		conn.Close()
	}
}

func MuxApis(mux *runtime.ServeMux) *apimgr.ServiceApis {
	apis := &apimgr.ServiceApis{
		ServiceName: config.GetStringValueWithNameSpace("", config.KeyHostname),
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
				Method: methIter.Key().String(),
				Path:   str,
			})
		}
	}

	return apis
}

func Register(mux *runtime.ServeMux) {
	apis := MuxApis(mux)
	go reliableRegister(apis)
}
