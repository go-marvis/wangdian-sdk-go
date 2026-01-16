package shop

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type Service struct {
	config *core.Config
}

func NewService(config *core.Config) *Service {
	return &Service{config}
}

// QueryShop 获取ERP的店铺档案资料
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=setting.Shop.queryShop
func (s *Service) QueryShop(ctx context.Context, req *QueryShopReq, options ...core.ReqOptionFunc) (*QueryShopResp, error) {
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
