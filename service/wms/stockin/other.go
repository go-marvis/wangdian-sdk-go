package stockin

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type other struct {
	config *core.Config
}

type QueryOtherReq struct {
	StartTime        string `json:"start_time"`                    // 开始时间, 入库单修改时间, yyyy-MM-dd HH:mm:ss格式
	EndTime          string `json:"end_time"`                      // 结束时间
	Status           string `json:"status,omitempty"`              // 入库单状态
	WarehouseNo      string `json:"warehouse_no,omitempty"`        // 仓库编号
	StockinNo        string `json:"stockin_no,omitempty"`          // 入库单号
	OtherStockinNo   string `json:"other_stockin_no,omitempty"`    // 其它入库业务单号
	ReasonName       string `json:"reason_name,omitempty"`         // 入库原因
	WmsStockChangeNo string `json:"wms_stock_change_no,omitempty"` // wms借调单号
}

type QueryOtherResp struct {
	Order      []*OtherOrder `json:"order"`
	TotalCount int64         `json:"total_count"`
}

type OtherOrder struct {
	StockinId          int     `json:"stockin_id"`                     // 入库单ID
	OrderNo            string  `json:"order_no"`                       // 入库单号
	Status             int     `json:"status"`                         // 状态
	WarehouseNo        string  `json:"warehouse_no"`                   // 仓库编号
	WarehouseName      string  `json:"warehouse_name"`                 // 仓库名称
	StockinTime        int64   `json:"stockin_time"`                   // 入库时间
	CreatedTime        int64   `json:"created_time"`                   // 建单时间
	Reason             string  `json:"reason"`                         // 其他入库原因
	Remark             string  `json:"remark"`                         // 备注
	GoodsCount         float64 `json:"goods_count"`                    // 货品总数
	LogisticsType      int     `json:"logistics_type"`                 // 物流类型
	CheckTime          int64   `json:"check_time"`                     // 审核时间
	SrcOrderNo         string  `json:"src_order_no"`                   // 业务单号
	OperatorName       string  `json:"operator_name"`                  // 操作员
	TotalPrice         float64 `json:"total_price"`                    // 总成本价
	TotalCost          float64 `json:"total_cost"`                     // 入库总金额
	LogisticsCompanyNo string  `json:"logistics_company_no,omitempty"` // 物流公司编号
	WmsStockChangeNo   string  `json:"wms_stock_change_no,omitempty"`  // wms借调单号
	DetailList         []struct {
		StockinId        int     `json:"stockin_id"`            // 入库单ID
		GoodsCount       float64 `json:"goods_count"`           // 数量
		TotalCost        float64 `json:"total_cost"`            // 总成本
		Remark           string  `json:"remark,omitempty"`      // 备注
		RightNum         float64 `json:"right_num"`             // 调整后数量
		GoodsUnit        string  `json:"goods_unit"`            // 单位
		BatchNo          string  `json:"batch_no"`              // 批次号
		RecId            int     `json:"rec_id"`                // 明细id
		ProductionDate   string  `json:"production_date"`       // 生产日期
		ExpireDate       string  `json:"expire_date"`           // 有效期
		GoodsName        string  `json:"goods_name"`            // 货品名称
		GoodsNo          string  `json:"goods_no"`              // 货品编号
		SpecNo           string  `json:"spec_no"`               // 商家编码
		Prop2            string  `json:"prop2,omitempty"`       // 自定义属性2
		SpecName         string  `json:"spec_name"`             // 规格名称
		SpecCode         string  `json:"spec_code,omitempty"`   // 规格码
		BrandNo          string  `json:"brand_no"`              // 品牌编号
		BrandName        string  `json:"brand_name"`            // 品牌名称
		Defect           bool    `json:"defect"`                // 是否残品
		CheckedCostPrice float64 `json:"checked_cost_price"`    // 入库价
		PositionNo       string  `json:"position_no,omitempty"` // 货位编号
	} `json:"detail_list" gorm:"serializer:json"` // 入库单明细
}

// QueryWithDetail 其他入库单查询
//
// 时间跨度：start_time和end_time最大跨度为30天
//
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.stockin.Other.queryWithDetail
func (o *other) QueryWithDetail(ctx context.Context, req *QueryOtherReq, options ...core.ReqOptionFunc) (*QueryOtherResp, error) {
	apiReq := &core.ApiReq{
		HttpMethod: http.MethodPost,
		Method:     "wms.stockin.Other.queryWithDetail",
		Body:       []any{req},
	}
	return core.Retry[QueryOtherResp](ctx, o.config, func() (*core.ApiResp, error) {
		return core.Request(ctx, apiReq, o.config, options...)
	})
}
