package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/apimgr"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedApiManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterApiManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterApiManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
