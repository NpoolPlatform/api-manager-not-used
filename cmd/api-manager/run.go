package main

import (
	"github.com/NpoolPlatform/api-manager/api"
	db "github.com/NpoolPlatform/api-manager/pkg/db"

	servicecli "github.com/NpoolPlatform/api-manager/pkg/client"
	serviceapi "github.com/NpoolPlatform/api-manager/pkg/middleware/service-api"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	cli "github.com/urfave/cli/v2"

	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"r"},
	Usage:   "Run the daemon",
	After: func(ctx *cli.Context) error {
		if err := grpc2.HShutdown(); err != nil {
			logger.Sugar().Warnf("graceful shutdown http server error: %v", err)
		}
		grpc2.GShutdown()
		return nil
	},
	Action: func(c *cli.Context) error {
		if err := db.Init(); err != nil {
			return err
		}

		go func() {
			if err := grpc2.RunGRPC(rpcRegister); err != nil {
				logger.Sugar().Errorf("fail to run grpc server: %v", err)
			}
		}()

		return grpc2.RunGRPCGateWay(rpcGatewayRegister)
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)
	apis := servicecli.GrpcApis(server)
	serviceapi.ReliableRegister(apis)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return xerrors.Errorf("fail register gateway: %v", err)
	}
	apis := servicecli.MuxApis(mux)
	serviceapi.ReliableRegister(apis)

	return nil
}
