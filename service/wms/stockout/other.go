package stockout

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type other struct {
	config *core.Config
}

// QueryWithDetail 其他出库单查询
//
// 时间跨度：start_time和end_time最大跨度为30天
//
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.stockout.OtherQuery.queryWithDetail
func (s *other) QueryWithDetail(ctx context.Context, req *OtherQueryReq, options ...core.ReqOptionFunc) (*OtherQueryResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.stockout.OtherQuery.queryWithDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &OtherQueryResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
