package outer

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type outerIn struct {
	config *core.Config
}

// QueryWithDetail - 获取ERP外仓调整入库信息
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.outer.OuterIn.queryWithDetail
func (s *outerIn) QueryWithDetail(ctx context.Context, req *QueryOuterInReq, options ...core.ReqOptionFunc) (*QueryOuterInResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.outer.OuterIn.queryWithDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &QueryOuterInResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
