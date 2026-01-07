package wms

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type stockpd struct {
	config *core.Config
}

type detailInfo struct {
	RecId      int     `json:"rec_id"`                // 明细id
	GoodsCount float64 `json:"goods_count"`           // 数量
	GoodsName  string  `json:"goods_name"`            // 货品名称
	GoodsNo    string  `json:"goods_no"`              // 货品编号
	Remark     string  `json:"remark,omitempty"`      // 备注
	BrandNo    string  `json:"brand_no"`              // 品牌编号
	BrandName  string  `json:"brand_name"`            // 品牌名称
	SpecNo     string  `json:"spec_no,omitempty"`     // 商家编码
	SpecName   string  `json:"spec_name,omitempty"`   // 规格名称
	SpecCode   string  `json:"spec_code,omitempty"`   // 规格码
	BatchNo    string  `json:"batch_no,omitempty"`    // 批次号
	PositionNo string  `json:"position_no,omitempty"` // 货位编号
	ExpireDate string  `json:"expire_date,omitempty"` // 有效期
	Defect     bool    `json:"defect,omitempty"`      // 是否残次品
}

type QueryStockPdInDetailReq struct {
	StartTime   string `json:"start_time"`             // 开始时间
	EndTime     string `json:"end_time"`               // 结束时间
	TimeType    int    `json:"time_type,omitempty"`    // 时间类型
	Status      string `json:"status,omitempty"`       // 入库单状态
	WarehouseNo string `json:"warehouse_no,omitempty"` // 仓库编号
	StockinNo   string `json:"stockin_no,omitempty"`   // 入库单号
	SrcOrderNo  string `json:"src_order_no,omitempty"` // 盘点单号
}

type QueryStockPdInDetailResp struct {
	Order      []StockPdInOrder `json:"order"`
	TotalCount int64            `json:"total_count"`
}

type StockPdInOrder struct {
	StockinId     int     `json:"stockin_id"`                // 入库单ID
	OrderNo       string  `json:"order_no"`                  // 入库单号
	WarehouseNo   string  `json:"warehouse_no"`              // 仓库编号
	WarehouseName string  `json:"warehouse_name"`            // 仓库名称
	StockinTime   int64   `json:"stockin_time"`              // 入库时间
	CreatedTime   int64   `json:"created_time"`              // 创建时间
	Remark        string  `json:"remark"`                    // 入库单备注
	Status        int     `json:"status"`                    // 入库单状态
	GoodsCount    float64 `json:"goods_count"`               // 货品总数
	CheckTime     int64   `json:"check_time"`                // 审核时间
	SrcOrderNo    string  `json:"src_order_no"`              // 源单号
	OperatorName  string  `json:"operator_name"`             // 操作人
	TotalPrice    float64 `json:"total_price"`               // 总成本价
	PdOrderRemark string  `json:"pd_order_remark,omitempty"` // 盘点单备注
	DetailList    []struct {
		detailInfo
		StockinId      int     `json:"stockin_id"`                // 入库单ID
		TotalCost      float64 `json:"total_cost,omitempty"`      // 成本价
		RightNum       float64 `json:"right_num"`                 // 调整后数量
		GoodsUnit      string  `json:"goods_unit"`                // 单位
		Prop2          string  `json:"prop2,omitempty"`           // 自定义属性2
		ProductionDate string  `json:"production_date,omitempty"` // 生产日期
	} `json:"detail_list" gorm:"serializer:json"` // 入库单明细
}

// QueryStockPdInDetail 盘点入库单查询
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.StockPd.queryStockPdInDetail
func (s *stockpd) QueryStockPdInDetail(ctx context.Context, req *QueryStockPdInDetailReq, options ...core.ReqOptionFunc) (*QueryStockPdInDetailResp, error) {
	apiReq := &core.ApiReq{
		HttpMethod: http.MethodPost,
		Method:     "wms.StockPd.queryStockPdInDetail",
		Body:       []any{req},
	}
	return core.Retry[QueryStockPdInDetailResp](ctx, s.config, func() (*core.ApiResp, error) {
		return core.Request(ctx, apiReq, s.config, options...)
	})
}

type QueryStockPdOutReq struct {
	StartTime   string `json:"start_time"`             // 开始时间
	EndTime     string `json:"end_time"`               // 结束时间
	TimeType    int    `json:"time_type,omitempty"`    // 时间类型
	Status      string `json:"status,omitempty"`       // 入库单状态
	WarehouseNo string `json:"warehouse_no,omitempty"` // 仓库编号
	StockoutNo  string `json:"stockout_no,omitempty"`  // 出库单号
	SrcOrderNo  string `json:"src_order_no,omitempty"` // 盘点单号
}

type QueryStockPdOutDetailResp struct {
	Order      []StockPdInOrder `json:"order"`
	TotalCount int64            `json:"total_count"`
}

type StockPdOutOrder struct {
	StockoutId       int     `json:"stockout_id"`       // 出库单ID
	OrderNo          string  `json:"order_no"`          // 出库单号
	SrcOrderId       int     `json:"src_order_id"`      // 源单号
	WarehouseNo      string  `json:"warehouse_no"`      // 仓库编号
	ConsignTime      int64   `json:"consign_time"`      // 出库时间
	OrderType        int     `json:"order_type"`        // 出库类型
	Status           int     `json:"status"`            // 出库单状态
	GoodsCount       float64 `json:"goods_count"`       // 出库数量
	LogisticsNo      string  `json:"logistics_no"`      // 物流单号
	ReceiverName     string  `json:"receiver_name"`     // 收件人姓名
	ReceiverCountry  int     `json:"receiver_country"`  // 国家ID
	ReceiverProvince int     `json:"receiver_province"` // 省ID
	ReceiverCity     int     `json:"receiver_city"`     // 城市ID
	ReceiverDistrict int     `json:"receiver_district"` // 地区ID
	ReceiverAddress  string  `json:"receiver_address"`  // 收件地址
	ReceiverMobile   string  `json:"receiver_mobile"`   // 收件人手机号
	ReceiverTelno    string  `json:"receiver_telno"`    // 收件固话号
	ReceiverZip      string  `json:"receiver_zip"`      // 邮编
	Remark           string  `json:"remark"`            // 出库单备注
	OperatorName     string  `json:"operator_name"`     // 制单人
	GoodsTotalCost   string  `json:"goods_total_cost"`  // 总成本
	PdOrderRemark    string  `json:"pd_order_remark"`   // 盘点单备注
	DetailList       []struct {
		detailInfo
		StockoutId int     `json:"stockout_id"`          // 出库单ID
		CostPrice  float64 `json:"cost_price,omitempty"` // 成本价
		Weight     float64 `json:"weight,omitempty"`     // 总重量
		GoodsType  int     `json:"goods_type,omitempty"` // 货品类型
		UnitName   string  `json:"unit_name"`            // 单位
	} `json:"detail_list" gorm:"serializer:json"` // 出库单明细
}

// QueryStockPdOutDetail 盘点出库单查询
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.StockPd.queryStockPdOutDetail
func (s *stockpd) QueryStockPdOutDetail(ctx context.Context, req *QueryStockPdOutReq, options ...core.ReqOptionFunc) (*QueryStockPdOutDetailResp, error) {
	apiReq := &core.ApiReq{
		HttpMethod: http.MethodPost,
		Method:     "wms.StockPd.queryStockPdOutDetail",
		Body:       []any{req},
	}
	return core.Retry[QueryStockPdOutDetailResp](ctx, s.config, func() (*core.ApiResp, error) {
		return core.Request(ctx, apiReq, s.config, options...)
	})
}
