package sales

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type tradeQuery struct {
	config *core.Config
}

// QueryWithDetail 获取ERP系统内订单信息
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=sales.TradeQuery.queryWithDetail
func (s *tradeQuery) QueryWithDetail(ctx context.Context, req *TradeQueryReq, options ...core.ReqOptionFunc) (*TradeQueryResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "sales.TradeQuery.queryWithDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &TradeQueryResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
