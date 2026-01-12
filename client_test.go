package wangdian

import (
	"context"
	"testing"

	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/aftersales/refund"
	"github.com/go-marvis/wangdian-sdk-go/service/setting"
	"github.com/go-marvis/wangdian-sdk-go/service/setting/strategy"
	"github.com/go-marvis/wangdian-sdk-go/service/wms"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/outer"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockin"
	"github.com/go-marvis/wangdian-sdk-go/service/wms/stockout"
	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

var (
	ctx = context.Background()
	api = NewClient(&core.Config{
		Sid:     "", // 卖家账号，旺店通分配
		Key:     "", // 接口账号，旺店通分配
		Secret:  "",
		Salt:    "",
		BaseUrl: "http://wdt.wangdian.cn/openapi", // 服务地址
		Version: "1.0",
		Retry:   5,
		Limiter: rate.NewLimiter(rate.Limit(1), 1),
	})
)

func Test_aftersales_refund_refund(t *testing.T) {
	resp, err := api.AfterSales.Refund.Refund.Search(ctx, refund.NewRefundSearchReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&refund.RefundSearchReqBody{
			SettleFrom: "2025-12-30 20:00:00",
			SettleTo:   "2025-12-31 23:59:59",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_setting_shop(t *testing.T) {
	resp, err := api.Setting.Shop.QueryShop(ctx, setting.NewQueryShopReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&setting.QueryShopBody{}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_setting_stratety_virtual_warehouse(t *testing.T) {
	resp, err := api.Setting.Strategy.VirtualWarehouse.Query(ctx, strategy.NewVirtualWarehouseQueryReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&strategy.VirtualWarehouseQueryBody{}).
		Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_setting_stratety_virtual_warehouse_order(t *testing.T) {
	resp, err := api.Setting.Strategy.VirtualWarehouse.OrderSearch(ctx, strategy.NewVirtualWarehouseOrderSearchReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&strategy.VirtualWarehouseOrderSearchBody{
			StartTime: "2025-12-30",
			EndTime:   "2025-12-31",
		}).
		Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_setting_warehouse(t *testing.T) {
	resp, err := api.Setting.Warehouse.QueryWarehouse(ctx, setting.NewQueryWarehouseReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&setting.QueryWarehouseBody{}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_outer_in_order(t *testing.T) {
	resp, err := api.Wms.Outer.OuterIn.QueryWithDetail(ctx, outer.NewQueryOuterInReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&outer.QueryOuterInBody{
			StartTime: "2025-01-17",
			EndTime:   "2025-01-18",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_outer_out_order(t *testing.T) {
	resp, err := api.Wms.Outer.OuterOut.QueryWithDetail(ctx, outer.NewQueryOuterOutReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&outer.QueryOuterOutBody{
			StartTime: "2025-01-17",
			EndTime:   "2025-01-18",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockin_other_order(t *testing.T) {
	resp, err := api.Wms.StockIn.Other.QueryWithDetail(ctx, stockin.NewQueryOtherReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&stockin.QueryOtherBody{
			StartTime: "2025-12-30",
			EndTime:   "2025-12-31",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockin_purchase_order(t *testing.T) {
	resp, err := api.Wms.StockIn.Purchase.QueryWithDetail(ctx, stockin.NewPurchaseQueryReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&stockin.PurchaseQueryBody{
			StartTime: "2025-12-30",
			EndTime:   "2025-12-31",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockin_refund_order(t *testing.T) {
	resp, err := api.Wms.StockIn.Refund.QueryWithDetail(ctx, stockin.NewRefundQueryReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&stockin.RefundQueryBody{
			StartTime: "2025-12-28",
			EndTime:   "2025-12-31",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockout_other_order(t *testing.T) {
	resp, err := api.Wms.StockOut.Other.QueryWithDetail(ctx, stockout.NewOtherQueryReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&stockout.OtherQueryBody{
			StartTime: "2025-12-30",
			EndTime:   "2025-12-31",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockout_sales_order(t *testing.T) {
	resp, err := api.Wms.StockOut.Sales.QueryWithDetail(ctx, stockout.NewSalesQueryReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&stockout.SalesQueryBody{
			StartTime: "2025-12-31 11:00:00",
			EndTime:   "2025-12-31 12:00:00",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockpd_in_order(t *testing.T) {
	resp, err := api.Wms.StockPd.QueryStockPdInDetail(ctx, wms.NewQueryStockPdInDetailReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&wms.QueryStockPdInDetailBody{
			StartTime: "2025-02-05",
			EndTime:   "2025-02-12",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}

func Test_wms_stockpd_out_order(t *testing.T) {
	resp, err := api.Wms.StockPd.QueryStockPdOutDetail(ctx, wms.NewQueryStockPdOutDetailReqBuilder().
		PageNo(0).PageSize(200).CalcTotal(1).
		Body(&wms.QueryStockPdOutDetailBody{
			StartTime: "2025-03-26",
			EndTime:   "2025-04-02",
		}).Build())
	assert.Nil(t, err)
	t.Log(resp.Data.TotalCount)
	assert.True(t, resp.Data.TotalCount > 0)
}
