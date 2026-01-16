package sales

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

// QueryWithDetail 销售出库单查询
//
// 时间跨度：start_time和end_time最大跨度为60分钟。
//
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.stockout.Sales.queryWithDetail
func (s *Service) QueryWithDetail(ctx context.Context, req *QueryReq, options ...core.ReqOptionFunc) (*QueryResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.stockout.Sales.queryWithDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &QueryResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
