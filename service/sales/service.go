package sales

import "github.com/go-marvis/wangdian-sdk-go/core"

type Service struct {
	TradeQuery *tradeQuery
}

func NewService(config *core.Config) *Service {
	return &Service{
		TradeQuery: &tradeQuery{config: config},
	}
}
