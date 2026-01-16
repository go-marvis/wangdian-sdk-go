package sales

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/sales/trade_query"
)

type Service struct {
	TradeQuery *trade_query.Service
}

func NewService(config *core.Config) *Service {
	return &Service{
		TradeQuery: trade_query.NewService(config),
	}
}
