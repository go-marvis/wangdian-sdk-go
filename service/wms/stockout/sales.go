package stockout

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type sales struct {
	config *core.Config
}

// QueryWithDetail 销售出库单查询
//
// 时间跨度：start_time和end_time最大跨度为60分钟。
//
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.stockout.Sales.queryWithDetail
func (s *sales) QueryWithDetail(ctx context.Context, req *SalesQueryReq, options ...core.ReqOptionFunc) (*SalesQueryResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.stockout.Sales.queryWithDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &SalesQueryResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
