package refund

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type refund struct {
	config *core.Config
}

// Search 获取ERP销售退货（换货）订单信息
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=aftersales.refund.Refund.search
func (r *refund) Search(ctx context.Context, req *RefundSearchReq, options ...core.ReqOptionFunc) (*RefundSearchResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "aftersales.refund.Refund.search"

	apiResp, err := core.Request(ctx, apiReq, r.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &RefundSearchResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, r.config)
	return resp, err
}
