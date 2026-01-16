package setting

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/setting/shop"
	"github.com/go-marvis/wangdian-sdk-go/service/setting/strategy"
	"github.com/go-marvis/wangdian-sdk-go/service/setting/warehouse"
)

type Service struct {
	Strategy  *strategy.Service
	Shop      *shop.Service
	Warehouse *warehouse.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		Strategy:  strategy.NewService(config),
		Shop:      shop.NewService(config),
		Warehouse: warehouse.NewService(config),
	}
}
