package goods

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/goods/goods"
)

type Service struct {
	Goods *goods.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		Goods: goods.NewService(config),
	}
}
