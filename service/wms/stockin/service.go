package stockin

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockin/other"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockin/purchase"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockin/refund"
)

type Service struct {
	Other    *other.Service
	Purchase *purchase.Service
	Refund   *refund.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		Other:    other.NewService(config),
		Purchase: purchase.NewService(config),
		Refund:   refund.NewService(config),
	}
}
