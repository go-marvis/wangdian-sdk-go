package setting

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/setting/strategy"
)

type Service struct {
	Strategy  *strategy.Service
	Shop      *shop
	Warehouse *warehouse
}

func NewService(config *core.Config) *Service {
	return &Service{
		Strategy:  strategy.NewService(config),
		Shop:      &shop{config},
		Warehouse: &warehouse{config},
	}
}
