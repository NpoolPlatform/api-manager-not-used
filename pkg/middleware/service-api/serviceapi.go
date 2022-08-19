package serviceapi

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/apimgr"

	crud "github.com/NpoolPlatform/api-manager/pkg/crud/service-api"
	"golang.org/x/xerrors"
)

type EntryPoint struct {
	Rule string `json:"rule"`
}

func (ep EntryPoint) Domain() string {
	// "Host(`api.xpool.top`) \u0026\u0026 PathPrefix(`/api/cloud-hashing-apis/version`)"
	regex := regexp.MustCompile(`(?:([a-z0-9-]+|\*)\.)?([a-z0-9-]{1,61})\.([a-z0-9]{2,7})`)
	domain := regex.Find([]byte(ep.Rule))
	return string(domain)
}

func (ep EntryPoint) ExportedPath() string {
	regex := regexp.MustCompile(`/[a-z0-9-/]+`)
	path := regex.Find([]byte(ep.Rule))
	return string(path)
}

func (ep EntryPoint) PathPrefix() (string, error) {
	path := ep.ExportedPath()
	paths := strings.Split(path, "/")
	if len(paths) < 3 {
		return "", xerrors.Errorf("invalid exported path")
	}

	return strings.Join(paths[0:3], "/"), nil
}

func (ep EntryPoint) ServiceKey() (string, error) {
	path := ep.ExportedPath()
	paths := strings.Split(path, "/")
	if len(paths) < 3 {
		return "", xerrors.Errorf("invalid exported path")
	}

	return paths[2], nil
}

func (ep EntryPoint) Path() (string, error) {
	path := ep.ExportedPath()
	paths := strings.Split(path, "/")
	if len(paths) < 3 {
		return "", xerrors.Errorf("invalid exported path")
	}

	return fmt.Sprintf("/%v", strings.Join(paths[3:], "/")), nil
}

func Register(ctx context.Context, in *npool.RegisterRequest) (*npool.RegisterResponse, error) {
	resp, err := crud.Register(ctx, in)
	if err != nil {
		return nil, xerrors.Errorf("fail register apis: %v", err)
	}

	return resp, nil
}

func ReliableRegister(apis *npool.ServiceApis) {
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()

	for range tick.C {
		_, err := Register(context.Background(), &npool.RegisterRequest{
			Info: apis,
		})
		if err != nil {
			logger.Sugar().Errorf("fail register apis: %v", err)
			continue
		}
		break
	}
}
