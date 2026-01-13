package wangdian

import (
	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/aftersales"
	"github.com/go-marvis/wangdian-sdk-go/service/sales"
	"github.com/go-marvis/wangdian-sdk-go/service/setting"
	"github.com/go-marvis/wangdian-sdk-go/service/wms"
)

type Client struct {
	AfterSales *aftersales.Service
	Sales      *sales.Service
	Setting    *setting.Service
	Wms        *wms.Service
}

func NewClient(config *core.Config) *Client {
	core.Init(config)

	return &Client{
		AfterSales: aftersales.NewService(config),
		Sales:      sales.NewService(config),
		Setting:    setting.NewService(config),
		Wms:        wms.NewService(config),
	}
}
