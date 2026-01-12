package stockin

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type refund struct {
	config *core.Config
}

// QueryWithDetail 退货入库单查询
//
// 时间跨度：start_time和end_time最大跨度为30天。
//
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.stockin.Refund.queryWithDetail
func (s *refund) QueryWithDetail(ctx context.Context, req *RefundQueryReq, options ...core.ReqOptionFunc) (*RefundQueryResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.stockin.Refund.queryWithDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &RefundQueryResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
