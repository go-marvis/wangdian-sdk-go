package virtual_warehouse

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

// Query 获取ERP虚拟仓、店铺、实体仓关联信息
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=setting.strategy.VirtualWarehouse.query
func (s *Service) Query(ctx context.Context, req *QueryReq, options ...core.ReqOptionFunc) (*QueryResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "setting.strategy.VirtualWarehouse.query"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &QueryResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// OrderSearch 获取ERP虚拟仓单据信息
//
// 时间跨度：start_time与end_time时间跨度不超过30天
//
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=setting.strategy.VirtualWarehouse.orderSearch
func (s *Service) OrderSearch(ctx context.Context, req *OrderSearchReq, options ...core.ReqOptionFunc) (*OrderSearchResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "setting.strategy.VirtualWarehouse.orderSearch"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &OrderSearchResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
