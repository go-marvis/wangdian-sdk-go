package stockin

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type other struct {
	config *core.Config
}

// QueryWithDetail 其他入库单查询
//
// 时间跨度：start_time和end_time最大跨度为30天
//
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.stockin.Other.queryWithDetail
func (s *other) QueryWithDetail(ctx context.Context, req *QueryOtherReq, options ...core.ReqOptionFunc) (*QueryOtherResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.stockin.Other.queryWithDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &QueryOtherResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
