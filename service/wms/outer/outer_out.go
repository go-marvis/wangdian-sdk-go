package outer

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type outerOut struct {
	config *core.Config
}

// QueryWithDetail 外仓调整出库单查询
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.outer.OuterOut.queryWithDetail
func (s *outerOut) QueryWithDetail(ctx context.Context, req *QueryOuterOutReq, options ...core.ReqOptionFunc) (*QueryOuterOutResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.outer.OuterOut.queryWithDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &QueryOuterOutResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
