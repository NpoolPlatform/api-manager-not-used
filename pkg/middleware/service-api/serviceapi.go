package serviceapi

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
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

func updater() { //nolint
	pageIndex := 1
	allRouters := []EntryPoint{}

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
			break
		}

		if len(routers) == 0 {
			break
		}

		allRouters = append(allRouters, routers...)
		pageIndex++
	}

	resp, err := crud.GetApis(context.Background(), &npool.GetApisRequest{})
	if err != nil {
		logger.Sugar().Errorf("fail get apis: %v", err)
		return
	}

	for _, router := range allRouters {
		domain := router.Domain()
		prefix, err := router.PathPrefix()
		if err != nil {
			logger.Sugar().Errorf("fail get prefix of %v: %v", router, err)
			continue
		}
		path, err := router.Path()
		if err != nil {
			logger.Sugar().Errorf("fail get path of %v: %v", router, err)
			continue
		}
		serviceKey, err := router.ServiceKey()
		if err != nil {
			logger.Sugar().Errorf("fail get service key of %v: %v", router, err)
			continue
		}

		for _, info := range resp.Infos {
			if !strings.Contains(serviceKey, info.ServiceName) || !strings.Contains(info.Path, path) {
				continue
			}

			info.PathPrefix = prefix
			info.Exported = true

			contains := false
			for _, dom := range info.Domains {
				if dom == domain {
					contains = true
					break
				}
			}

			if !contains {
				info.Domains = append(info.Domains, domain)
			}
		}
	}

	err = crud.Update(context.Background(), resp.Infos)
	if err != nil {
		logger.Sugar().Errorf("fail update apis: %v", err)
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
