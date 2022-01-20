package serviceapi

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/api-manager/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/apimgr"

	"github.com/google/uuid"
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
	apis := npool.ServiceApis{
		ServiceName: fmt.Sprintf("test-app.npool.top-%v", uuid.New()),
		PathPrefix:  "/api/test-app",
		Paths: []*npool.Path{
			{Method: "GET", Path: "/v1/get/test", Exported: false},
			{Method: "POST", Path: "/v1/get/test1", Exported: false},
			{Method: "GET", Path: "/v1/get/test2", Exported: false},
			{Method: "GET", Path: "/v1/get/test3", Exported: true},
		},
	}

	resp, err := Register(context.Background(), &npool.RegisterRequest{
		Info: &apis,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp.Info, &apis)
	}
}
