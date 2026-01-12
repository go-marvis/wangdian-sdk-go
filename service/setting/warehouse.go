package setting

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type warehouse struct {
	config *core.Config
}

func (w warehouse) QueryWarehouse(ctx context.Context, req *QueryWarehouseReq, options ...core.ReqOptionFunc) (*QueryWarehouseResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "setting.Warehouse.queryWarehouse"

	apiResp, err := core.Request(ctx, apiReq, w.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &QueryWarehouseResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, w.config)
	return resp, err
}
