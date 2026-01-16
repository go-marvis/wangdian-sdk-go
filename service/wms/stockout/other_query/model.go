package other_query

import (
	"fmt"

	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/model"
)

type QueryReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryReqBuilder() *QueryReqBuilder {
	return &QueryReqBuilder{core.NewApiReq()}
}

func (builder *QueryReqBuilder) PageSize(pageSize int) *QueryReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

func (builder *QueryReqBuilder) PageNo(pageNo int) *QueryReqBuilder {
	builder.apiReq.QueryParams.Set("page_no", fmt.Sprint(pageNo))
	return builder
}

func (builder *QueryReqBuilder) CalcTotal(calcTotal int) *QueryReqBuilder {
	builder.apiReq.QueryParams.Set("calc_total", fmt.Sprint(calcTotal))
	return builder
}

func (builder *QueryReqBuilder) Body(body *QueryBody) *QueryReqBuilder {
	builder.apiReq.Body = body
	return builder

}

func (builder *QueryReqBuilder) Build() *QueryReq {
	return &QueryReq{builder.apiReq}
}

type QueryReq struct {
	apiReq *core.ApiReq
}

type QueryBody struct {
	StartTime   string `json:"start_time"`             // 开始时间, 入库单修改时间, yyyy-MM-dd HH:mm:ss格式
	EndTime     string `json:"end_time"`               // 结束时间
	TimeType    int    `json:"time_type,omitempty"`    // 时间条件类型 0:修改时间; 1:入库时间 不传默认为0
	WarehouseNo string `json:"warehouse_no,omitempty"` // 仓库编号
	SrcOrderNo  string `json:"src_order_no,omitempty"` // 业务单号
	StockoutNo  string `json:"stockout_no,omitempty"`  // 入库单号
	Status      int    `json:"status,omitempty"`       // 业务单据状态
	ReasonName  string `json:"reason_name,omitempty"`  // 出库原因
}

type QueryResp struct {
	*core.ApiResp
	core.CodeError
	Data *QueryData `json:"data"`
}

type QueryData struct {
	Order      []*Order `json:"order"`
	TotalCount int64    `json:"total_count"`
}

type Order struct {
	StockoutId            int     `json:"stockout_id"`              // 出库单id
	OrderNo               string  `json:"order_no"`                 // 出库单编号
	SrcOrderNo            string  `json:"src_order_no"`             // 业务单号
	WarehouseNo           string  `json:"warehouse_no"`             // 出库的仓库编号
	ConsignTime           int64   `json:"consign_time"`             // 出库时间
	OrderType             int     `json:"order_type"`               // 源单据类别
	Status                int     `json:"status"`                   // 状态
	GoodsCount            float64 `json:"goods_count"`              // 出库数量
	PostFee               float64 `json:"post_fee"`                 // 邮费
	LogisticsNo           string  `json:"logistics_no"`             // 物流单号
	ReceiverName          string  `json:"receiver_name"`            // 收件人姓名
	ReceiverProvince      string  `json:"receiver_province"`        // 省
	ReceiverCity          string  `json:"receiver_city"`            // 城市
	ReceiverDistrict      string  `json:"receiver_district"`        // 地区
	ReceiverAddress       string  `json:"receiver_address"`         // 收件地址
	ReceiverMobile        string  `json:"receiver_mobile"`          // 收件人手机号
	Remark                string  `json:"remark"`                   // 出库单备注
	Weight                float64 `json:"weight"`                   // 实际称重重量(Kg)
	OperatorName          string  `json:"operator_name"`            // 制单人
	GoodsTotalCost        string  `json:"goods_total_cost"`         // 总成本
	GoodsTotalAmount      float64 `json:"goods_total_amount"`       // 总货款
	Modified              int64   `json:"modified"`                 // 最后修改时间
	Reason                string  `json:"reason"`                   // 出库原因
	CheckedGoodsTotalCost float64 `json:"checked_goods_total_cost"` // 瞬时成本总额
	LogisticsCompanyNo    string  `json:"logistics_company_no"`     // 物流公司编号
	WarehouseName         string  `json:"warehouse_name"`           // 仓库名称
	DetailList            []struct {
		RecId               int                    `json:"rec_id"`                 // 出库单详情id
		StockoutId          int                    `json:"stockout_id"`            // 出库单id
		GoodsCount          float64                `json:"goods_count"`            // 货品数量
		TotalAmount         float64                `json:"total_amount,omitempty"` // 总成本
		ExpireAate          string                 `json:"expire_date"`            // 有效期
		Remark              string                 `json:"remark,omitempty"`       // 出库单详情备注
		BrandNo             string                 `json:"brand_no"`               // 品牌编号
		BrandName           string                 `json:"brand_name"`             // 品牌名称
		GoodsName           string                 `json:"goods_name"`             // 货品名称
		GoodsNo             string                 `json:"goods_no"`               // 货品编码
		SpecNo              string                 `json:"spec_no,omitempty"`      // 商家编码
		SpecName            string                 `json:"spec_name"`              // 规格名称
		SpecCode            string                 `json:"spec_code,omitempty"`    // 规格码
		Defect              bool                   `json:"defect,omitempty"`       // 是否残次品
		CostPrice           float64                `json:"cost_price,omitempty"`   // 成本价
		Weight              float64                `json:"weight,omitempty"`       // 总重量
		GoodsType           int                    `json:"goods_type"`             // 货品类型
		UnitName            string                 `json:"unit_name,omitempty"`    // 单位
		BaseUnitId          int                    `json:"base_unit_id"`           // 单位id
		BatchNo             string                 `json:"batch_no"`               // 批次号
		PositionId          int                    `json:"position_id,omitempty"`  // 货位id
		PositionNo          string                 `json:"position_no,omitempty"`  // 货位编号
		PositionDetailsList []model.PositionDetail `json:"position_details_list"`  // 货位明细
	} `json:"detail_list"  gorm:"serializer:json"` // 出库单明细
}
