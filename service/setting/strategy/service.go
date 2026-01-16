package strategy

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/setting/strategy/virtual_warehouse"
)

type Service struct {
	VirtualWarehouse *virtual_warehouse.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		VirtualWarehouse: virtual_warehouse.NewService(config),
	}
}
