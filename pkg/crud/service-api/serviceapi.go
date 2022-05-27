package serviceapi

import (
	"context"
	"time"

	"github.com/NpoolPlatform/api-manager/pkg/db"
	"github.com/NpoolPlatform/api-manager/pkg/db/ent"
	"github.com/NpoolPlatform/api-manager/pkg/db/ent/serviceapi"
	npool "github.com/NpoolPlatform/message/npool/apimgr"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	grpcTimeout = 5 * time.Second
)

func Register(ctx context.Context, in *npool.RegisterRequest) (*npool.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	tx, err := cli.Tx(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create transaction: %v", err)
	}

	for _, path := range in.GetInfo().GetPaths() {
		err := tx.
			ServiceAPI.
			Create().
			SetProtocol(in.GetInfo().GetProtocol()).
			SetServiceName(in.GetInfo().GetServiceName()).
			SetDomains([]string{}).
			SetMethod(path.GetMethod()).
			SetMethodName(path.GetMethodName()).
			SetPath(path.GetPath()).
			SetExported(path.GetExported()).
			SetPathPrefix(in.GetInfo().GetPathPrefix()).
			OnConflict().
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return nil, xerrors.Errorf("fail rollback create service api: %v (%v)", rerr, err)
			}
			return nil, xerrors.Errorf("fail create service api: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, xerrors.Errorf("fail commit service api create: %v", err)
	}

	return &npool.RegisterResponse{
		Info: in.GetInfo(),
	}, nil
}

func Update(ctx context.Context, apis []*npool.ServicePath) error {
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return xerrors.Errorf("fail get db client: %v", err)
	}

	tx, err := cli.Tx(ctx)
	if err != nil {
		return xerrors.Errorf("fail create transaction: %v", err)
	}

	for _, api := range apis {
		id, err := uuid.Parse(api.ID)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return xerrors.Errorf("fail rollback update service api: %v (%v)", rerr, err)
			}
			return xerrors.Errorf("fail parse service api id: %v", err)
		}

		_, err = tx.
			ServiceAPI.
			UpdateOneID(id).
			SetPathPrefix(api.PathPrefix).
			SetExported(api.Exported).
			SetDomains(api.Domains).
			Save(ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return xerrors.Errorf("fail rollback update service api: %v (%v)", rerr, err)
			}
			return xerrors.Errorf("fail update service api: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return xerrors.Errorf("fail commit service api update: %v", err)
	}

	return nil
}

func rowToObject(info *ent.ServiceAPI) *npool.ServicePath {
	return &npool.ServicePath{
		ID:          info.ID.String(),
		Domains:     info.Domains,
		ServiceName: info.ServiceName,
		Method:      info.Method,
		MethodName:  info.MethodName,
		Path:        info.Path,
		Exported:    info.Exported,
		PathPrefix:  info.PathPrefix,
		Protocol:    info.Protocol,
		CreateAt:    info.CreateAt,
		UpdateAt:    info.UpdateAt,
	}
}

func GetApis(ctx context.Context, in *npool.GetApisRequest) (*npool.GetApisResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		ServiceAPI.
		Query().
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query service api: %v", err)
	}

	apis := []*npool.ServicePath{}
	for _, info := range infos {
		apis = append(apis, rowToObject(info))
	}

	return &npool.GetApisResponse{
		Infos: apis,
	}, nil
}

func GetServiceMethodAPI(ctx context.Context, in *npool.GetServiceMethodAPIRequest) (*npool.GetServiceMethodAPIResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		ServiceAPI.
		Query().
		Where(
			serviceapi.And(
				serviceapi.ServiceName(in.GetServiceName()),
				serviceapi.MethodName(in.GetMethodName()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query service api: %v", err)
	}

	var api *npool.ServicePath
	for _, info := range infos {
		api = rowToObject(info)
		break
	}

	return &npool.GetServiceMethodAPIResponse{
		Info: api,
	}, nil
}
