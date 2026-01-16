package stockout

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockout/other_query"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockout/sales"
)

type Service struct {
	Other *other_query.Service
	Sales *sales.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		Other: other_query.NewService(config),
		Sales: sales.NewService(config),
	}
}
