package stockin

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type purchase struct {
	config *core.Config
}

// QueryWithDetail 采购入库单查询 - 获取ERP内的采购入库单
//
// 时间跨度：start_time和end_time最大跨度为30天
//
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.stockin.Purchase.queryWithDetail
func (s *purchase) QueryWithDetail(ctx context.Context, req *PurchaseQueryReq, options ...core.ReqOptionFunc) (*PurchaseQueryResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.stockin.Purchase.queryWithDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &PurchaseQueryResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
