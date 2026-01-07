package aftersales

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/aftersales/refund"
)

type Service struct {
	Refund *refund.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		Refund: refund.NewService(config),
	}
}
