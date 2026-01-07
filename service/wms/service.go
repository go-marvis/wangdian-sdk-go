package wms

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/outer"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockin"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockout"
)

type Service struct {
	Outer    *outer.Service
	StockIn  *stockin.Service
	StockOut *stockout.Service
	StockPd  *stockpd
}

func NewService(config *core.Config) *Service {
	return &Service{
		Outer:    outer.NewService(config),
		StockIn:  stockin.NewService(config),
		StockOut: stockout.NewService(config),
		StockPd:  &stockpd{config: config},
	}
}
