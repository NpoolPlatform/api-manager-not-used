package serviceapi

import (
	"context"
	"time"

	"github.com/NpoolPlatform/api-manager/pkg/db"
	npool "github.com/NpoolPlatform/message/npool/apimgr"

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
		_, err := tx.
			ServiceAPI.
			Create().
			SetDomain(in.GetInfo().GetServiceName()).
			SetMethod(path.GetMethod()).
			SetPath(path.GetPath()).
			SetExported(path.GetExported()).
			SetPathPrefix(in.GetInfo().GetPathPrefix()).
			Save(ctx)
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

func GetApis(ctx context.Context, in *npool.GetApisRequest) (*npool.GetApisResponse, error) {
	return nil, nil
}
