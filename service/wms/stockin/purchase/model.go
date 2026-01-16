package purchase

import (
	"fmt"

	"github.com/go-marvis/wangdian-sdk-go/core"
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

func (builder QueryReqBuilder) Build() *QueryReq {
	return &QueryReq{builder.apiReq}
}

type QueryReq struct {
	apiReq *core.ApiReq
}
type QueryBody struct {
	StartTime   string `json:"start_time"`             // 开始时间, 入库单修改时间, yyyy-MM-dd HH:mm:ss格式
	EndTime     string `json:"end_time"`               // 结束时间
	Status      string `json:"status,omitempty"`       // 10=已取消；20=编辑中；30=待审核；37=待质检；40=质检确认；80=已完成
	WarehouseNo string `json:"warehouse_no,omitempty"` // 仓库编号
	StockinNo   string `json:"stockin_no,omitempty"`   // 入库单号
	PurchaseNo  string `json:"purchase_no,omitempty"`  // 采购单号
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
	StockinId         int64   `json:"stockin_id"`          // 入库单唯一键
	OrderNo           string  `json:"order_no"`            // 入库单号
	WarehouseNo       string  `json:"warehouse_no"`        // 仓库编号
	Status            int     `json:"status"`              // 入库单状态
	Modified          int64   `json:"modified"`            // 最后修改时间
	CreatedTime       int64   `json:"created_time"`        // 制单时间
	Remark            string  `json:"remark"`              // 入库单备注
	LogisticsTypeName string  `json:"logistics_type_name"` // 入库单物流类型
	CheckTime         int64   `json:"check_time"`          // 入库单审核时间
	PurchaseId        int64   `json:"purchase_id"`         // 采购单唯一键
	PurchaseNo        string  `json:"purchase_no"`         // 采购单号
	GoodsCount        float64 `json:"goods_count"`         // 货品数量
	ProviderNo        string  `json:"provider_no"`         // 供应商编号
	ProviderName      string  `json:"provider_name"`       // 供应商名称
	LogisticsNo       string  `json:"logistics_no"`        // 物流单号
	LogisticsName     string  `json:"logistics_name"`      // 物流公司
	GoodsAmount       string  `json:"goods_amount"`        // 货品总价格，不包含优惠
	TotalPrice        float64 `json:"total_price"`         // 税前总货款（折后）
	TaxAmount         float64 `json:"tax_amount"`          // 税后总额（折后）
	TotalStockinPrice string  `json:"total_stockin_price"` // 明细入库价*数量之和
	FlagName          string  `json:"flag_name"`           // 标记名称
	OperatorName      string  `json:"operator_name"`       // 经办人
	DetailsList       []struct {
		RecId            int64   `json:"rec_id"`                    // 入库单明细id
		Num              float64 `json:"num"`                       // 数量
		Discount         float64 `json:"discount"`                  // 折扣
		CostPrice        float64 `json:"cost_price,omitempty"`      // 税前单价
		SrcPrice         float64 `json:"src_price,omitempty"`       // 税前折后单价
		TaxPrice         float64 `json:"tax_price,omitempty"`       // 税后单价（折后）
		TaxAmount        float64 `json:"tax_amount,omitempty"`      // 税后金额（税后总价）
		Tax              float64 `json:"tax,omitempty"`             // 税率
		TotalCost        float64 `json:"total_cost,omitempty"`      // 税前金额（税前总价）
		Remark           string  `json:"remark,omitempty"`          // 入库单明细备注
		GoodsName        string  `json:"goods_name"`                // 货品名称
		GoodsNo          string  `json:"goods_no"`                  // 货品编号
		SpecNo           string  `json:"spec_no,omitempty"`         // 商家编码
		SpecCode         string  `json:"spec_code,omitempty"`       // 规格码
		SpecName         string  `json:"spec_name,omitempty"`       // 规格名称
		Prop1            string  `json:"prop1,omitempty"`           // 采购明细自定义属性1
		Prop2            string  `json:"prop2,omitempty"`           // 采购明细自定义属性2
		Prop3            string  `json:"prop3,omitempty"`           // 采购明细自定义属性3
		Prop4            string  `json:"prop4,omitempty"`           // 采购明细自定义属性4
		BrandName        string  `json:"brand_name"`                // 品牌名称
		UnitName         string  `json:"unit_name"`                 // 基本单位
		BatchNo          string  `json:"batch_no,omitempty"`        // 批次，若无批次则返回空字符串
		ExpireDate       string  `json:"expire_data,omitempty"`     // 有效期，若无有效期则返回空字符串
		ProductionDate   string  `json:"production_date,omitempty"` // 生产日期，若无生产日期则返回空字符串
		Defect           bool    `json:"defect,omitempty"`          // 是否残品
		UnitRatio        float64 `json:"unit_ratio"`                // 换算系数
		PurchaseUnitName string  `json:"purchase_unit_name"`        // 辅助单位
		StockinPrice     float64 `json:"stockin_price,omitempty"`   // 入库价
	} `json:"details_list" gorm:"serializer:json"` //入库单明细
}
