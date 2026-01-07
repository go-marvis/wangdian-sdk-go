package strategy

import "github.com/go-marvis/wangdian-sdk-go/core"

type Service struct {
	VirtualWarehouse *virtualWarehouse
}

func NewService(config *core.Config) *Service {
	return &Service{
		VirtualWarehouse: &virtualWarehouse{config},
	}
}
