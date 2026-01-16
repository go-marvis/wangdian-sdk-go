package stockpd

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

// QueryStockPdInDetail 盘点入库单查询
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.StockPd.queryStockPdInDetail
func (s *Service) QueryStockPdInDetail(ctx context.Context, req *QueryStockPdInDetailReq, options ...core.ReqOptionFunc) (*QueryStockPdInDetailResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.StockPd.queryStockPdInDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &QueryStockPdInDetailResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}

// QueryStockPdOutDetail 盘点出库单查询
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.StockPd.queryStockPdOutDetail
func (s *Service) QueryStockPdOutDetail(ctx context.Context, req *QueryStockPdOutDetailReq, options ...core.ReqOptionFunc) (*QueryStockPdOutDetailResp, error) {
	apiReq := req.apiReq
	apiReq.HttpMethod = http.MethodPost
	apiReq.Method = "wms.StockPd.queryStockPdOutDetail"

	apiResp, err := core.Request(ctx, apiReq, s.config, options...)
	if err != nil {
		return nil, err
	}

	resp := &QueryStockPdOutDetailResp{ApiResp: apiResp}
	err = apiResp.UnmarshalBody(resp, s.config)
	return resp, err
}
