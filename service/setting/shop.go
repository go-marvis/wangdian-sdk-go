package setting

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type shop struct {
	config *core.Config
}

func (s *shop) QueryShop(ctx context.Context, req *QueryShopReq, options ...core.ReqOptionFunc) (*QueryShopResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "setting.Shop.queryShop"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &QueryShopResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
