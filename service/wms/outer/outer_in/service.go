package outer_in

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

// QueryWithDetail - 获取ERP外仓调整入库信息
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.outer.OuterIn.queryWithDetail
func (s *Service) QueryWithDetail(ctx context.Context, req *QueryReq, options ...core.ReqOptionFunc) (*QueryResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.outer.OuterIn.queryWithDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &QueryResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
