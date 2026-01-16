package wangdian

import (
	"context"
	"os"
	"testing"

	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/aftersales/refund/refund"
	"github.com/go-marvis/wangdian-sdk-go/service/goods/goods"
	"github.com/go-marvis/wangdian-sdk-go/service/sales/trade_query"
	"github.com/go-marvis/wangdian-sdk-go/service/setting/shop"
	"github.com/go-marvis/wangdian-sdk-go/service/setting/strategy/virtual_warehouse"
	"github.com/go-marvis/wangdian-sdk-go/service/setting/warehouse"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/outer/outer_in"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/outer/outer_out"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockin/other"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockin/purchase"
	stockin_refund "github.com/go-marvis/wangdian-sdk-go/service/wms/stockin/refund"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockout/other_query"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockout/sales"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockpd"
	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

var (
	ctx = context.Background()
	api = NewClient(&core.Config{
		Sid:     os.Getenv("SID"),
		Key:     os.Getenv("KEY"),
		Secret:  os.Getenv("SECRET"),
		Salt:    os.Getenv("SALT"),
		BaseUrl: "http://wdt.wangdian.cn/openapi",
		Version: "1.0",
		Debug:   true,
		Retry:   5,
		Limiter: rate.NewLimiter(rate.Limit(1), 1),
	})
)

func Test_aftersales_refund_refund(t *testing.T) {
	resp, err := api.AfterSales.Refund.Refund.Search(ctx, refund.NewSearchReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&refund.SearchReqBody{
			SettleFrom: "2025-12-30 20:00:00",
			SettleTo:   "2025-12-31 23:59:59",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_goods_query(t *testing.T) {
	resp, err := api.Goods.Goods.QueryWithSpec(ctx, goods.NewQueryReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&goods.QueryReqBody{
			StartTime: "2025-12-30 12:00:00",
			EndTime:   "2025-12-30 13:00:00"}).
		Build())

	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, len(resp.Data.GoodsList) > 0)
}

func Test_sales_trade_query(t *testing.T) {
	resp, err := api.Sales.TradeQuery.QueryWithDetail(ctx, trade_query.NewQueryReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&trade_query.QueryBody{
			StartTime: "2025-12-30 12:00:00",
			EndTime:   "2025-12-30 13:00:00",
		}).
		Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_setting_shop(t *testing.T) {
	resp, err := api.Setting.Shop.QueryShop(ctx, shop.NewQueryShopReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&shop.QueryShopBody{}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_setting_stratety_virtual_warehouse(t *testing.T) {
	resp, err := api.Setting.Strategy.VirtualWarehouse.Query(ctx, virtual_warehouse.NewQueryReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&virtual_warehouse.QueryBody{}).
		Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_setting_stratety_virtual_warehouse_order(t *testing.T) {
	resp, err := api.Setting.Strategy.VirtualWarehouse.OrderSearch(ctx, virtual_warehouse.NewOrderSearchReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&virtual_warehouse.OrderSearchBody{
			StartTime: "2025-12-30",
			EndTime:   "2025-12-31",
		}).
		Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_setting_warehouse(t *testing.T) {
	resp, err := api.Setting.Warehouse.QueryWarehouse(ctx, warehouse.NewQueryWarehouseReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&warehouse.QueryWarehouseBody{}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_outer_in_order(t *testing.T) {
	resp, err := api.Wms.Outer.OuterIn.QueryWithDetail(ctx, outer_in.NewQueryReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&outer_in.QueryBody{
			StartTime: "2025-01-17",
			EndTime:   "2025-01-18",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_outer_out_order(t *testing.T) {
	resp, err := api.Wms.Outer.OuterOut.QueryWithDetail(ctx, outer_out.NewQueryReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&outer_out.QueryBody{
			StartTime: "2025-01-17",
			EndTime:   "2025-01-18",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockin_other_order(t *testing.T) {
	resp, err := api.Wms.StockIn.Other.QueryWithDetail(ctx, other.NewQueryReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&other.QueryBody{
			StartTime: "2025-12-30",
			EndTime:   "2025-12-31",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockin_purchase_order(t *testing.T) {
	resp, err := api.Wms.StockIn.Purchase.QueryWithDetail(ctx, purchase.NewQueryReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&purchase.QueryBody{
			StartTime: "2025-12-30",
			EndTime:   "2025-12-31",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockin_refund_order(t *testing.T) {
	resp, err := api.Wms.StockIn.Refund.QueryWithDetail(ctx, stockin_refund.NewQueryReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&stockin_refund.QueryBody{
			StartTime: "2025-12-28",
			EndTime:   "2025-12-31",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockout_other_order(t *testing.T) {
	resp, err := api.Wms.StockOut.Other.QueryWithDetail(ctx, other_query.NewQueryReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&other_query.QueryBody{
			StartTime: "2025-12-30",
			EndTime:   "2025-12-31",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockout_sales_order(t *testing.T) {
	resp, err := api.Wms.StockOut.Sales.QueryWithDetail(ctx, sales.NewQueryReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&sales.QueryBody{
			StartTime: "2025-12-31 11:00:00",
			EndTime:   "2025-12-31 12:00:00",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockpd_in_order(t *testing.T) {
	resp, err := api.Wms.StockPd.QueryStockPdInDetail(ctx, stockpd.NewQueryStockPdInDetailReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&stockpd.QueryStockPdInDetailBody{
			StartTime: "2025-02-05",
			EndTime:   "2025-02-12",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockpd_out_order(t *testing.T) {
	resp, err := api.Wms.StockPd.QueryStockPdOutDetail(ctx, stockpd.NewQueryStockPdOutDetailReqBuilder().
		PageNo(0).PageSize(1).CalcTotal(1).
		Body(&stockpd.QueryStockPdOutDetailBody{
			StartTime: "2025-03-26",
			EndTime:   "2025-04-02",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}
