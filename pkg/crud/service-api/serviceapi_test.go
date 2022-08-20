package serviceapi

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/api-manager/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/apimgr"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	apis := npool.ServiceApis{
		ServiceName: "third-login-gateway.npool.top",
		Protocol:    "grpc",
		Paths: []*npool.Path{
			{Method: "NONE", Path: "third.logon.gateway.v1.ThirdLoginGateway/CreateAuths"},
			{Method: "NONE", Path: "third.logon.gateway.v1.ThirdLoginGateway/CreateAppAuth"},
			{Method: "NONE", Path: "third.logon.gateway.v1.ThirdLoginGateway/CreateAppAuths"},
			{Method: "NONE", Path: "third.logon.gateway.v1.ThirdLoginGateway/CreateThirdParty"},
			{Method: "NONE", Path: "third.logon.gateway.v1.ThirdLoginGateway/GetThirdPartyOnly"},
			{Method: "NONE", Path: "third.logon.gateway.v1.ThirdLoginGateway/Login"},
			{Method: "NONE", Path: "third.logon.gateway.v1.ThirdLoginGateway/Version"},
			{Method: "NONE", Path: "third.logon.gateway.v1.ThirdLoginGateway/GetAuths"},
			{Method: "NONE", Path: "third.logon.gateway.v1.ThirdLoginGateway/GetAppAuths"},
		},
	}

	resp, err := Register(context.Background(), &npool.RegisterRequest{
		Info: &apis,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info, &apis)
	}
}
