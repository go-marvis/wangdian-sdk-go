package refund

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

// Search 获取ERP销售退货（换货）订单信息
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=aftersales.refund.Refund.search
func (s *Service) Search(ctx context.Context, req *SearchReq, options ...core.ReqOptionFunc) (*SearchResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "aftersales.refund.Refund.search"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &SearchResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
