package serviceapi

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/apimgr"
	"github.com/go-resty/resty/v2"

	crud "github.com/NpoolPlatform/api-manager/pkg/crud/service-api"

	"golang.org/x/xerrors"
)

type EntryPoint struct {
	Rule string `json:"rule"`
}

var notifier = make(chan struct{})

func updater() {
	pageIndex := 1
	for {
		url := fmt.Sprintf("http://traefik.kube-system.svc.cluster.local:38080/api/http/routers?page=%v", pageIndex)
		resp, err := resty.New().R().Get(url)
		if err != nil {
			logger.Sugar().Errorf("fail get routers: %v (%v)", err, resp.Body())
			return
		}

		routers := []EntryPoint{}
		err = json.Unmarshal(resp.Body(), &routers)
		if err != nil {
			logger.Sugar().Errorf("fail parse body: %v", err)
			return
		}

		for _, router := range routers {
			logger.Sugar().Info(router)
		}

		if len(routers) == 0 {
			break
		}

		pageIndex++
	}
}

func Watcher() {
	ticker := time.NewTicker(24 * time.Hour)
	for {
		select {
		case <-ticker.C:
			updater()
		case <-notifier:
			updater()
		}
	}
}

func Register(ctx context.Context, in *npool.RegisterRequest) (*npool.RegisterResponse, error) {
	resp, err := crud.Register(ctx, in)
	if err != nil {
		return nil, xerrors.Errorf("fail register apis: %v", err)
	}

	go func() { notifier <- struct{}{} }()

	return resp, nil
}

func ReliableRegister(apis *npool.ServiceApis) {
	for {
		_, err := Register(context.Background(), &npool.RegisterRequest{
			Info: apis,
		})
		if err != nil {
			logger.Sugar().Errorf("fail register apis: %v", err)
			time.Sleep(time.Minute)
			continue
		}

		break
	}
}
