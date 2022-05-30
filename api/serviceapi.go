package api

import (
	"context"

	crud "github.com/NpoolPlatform/api-manager/pkg/crud/service-api"
	mw "github.com/NpoolPlatform/api-manager/pkg/middleware/service-api"
	npool "github.com/NpoolPlatform/message/npool/apimgr"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Register(ctx context.Context, in *npool.RegisterRequest) (*npool.RegisterResponse, error) {
	resp, err := mw.Register(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("api register error: %v", err)
		return &npool.RegisterResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetApis(ctx context.Context, in *npool.GetApisRequest) (*npool.GetApisResponse, error) {
	resp, err := crud.GetApis(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get apis error: %v", err)
		return &npool.GetApisResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetServiceMethodAPI(ctx context.Context, in *npool.GetServiceMethodAPIRequest) (*npool.GetServiceMethodAPIResponse, error) {
	resp, err := crud.GetServiceMethodAPI(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("get apis error: %v", err)
		return &npool.GetServiceMethodAPIResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
