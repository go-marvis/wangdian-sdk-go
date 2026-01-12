package strategy

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type virtualWarehouse struct {
	config *core.Config
}

// Query 获取ERP虚拟仓、店铺、实体仓关联信息
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=setting.strategy.VirtualWarehouse.query
func (s *virtualWarehouse) Query(ctx context.Context, req *VirtualWarehouseQueryReq, options ...core.ReqOptionFunc) (*VirtualWarehouseQueryResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "setting.strategy.VirtualWarehouse.query"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &VirtualWarehouseQueryResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// OrderSearch 获取ERP虚拟仓单据信息
//
// 时间跨度：start_time与end_time时间跨度不超过30天
//
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=setting.strategy.VirtualWarehouse.orderSearch
func (s *virtualWarehouse) OrderSearch(ctx context.Context, req *VirtualWarehouseOrderSearchReq, options ...core.ReqOptionFunc) (*VirtualWarehouseOrderSearchResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "setting.strategy.VirtualWarehouse.orderSearch"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &VirtualWarehouseOrderSearchResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
