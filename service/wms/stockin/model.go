package stockin

import (
	"fmt"

	"github.com/go-marvis/wangdian-sdk-go/core"
	"github.com/go-marvis/wangdian-sdk-go/service/model"
)

type QueryOtherReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryOtherReqBuilder() *QueryOtherReqBuilder {
	return &QueryOtherReqBuilder{core.NewApiReq()}
}

func (builder *QueryOtherReqBuilder) PageSize(pageSize int) *QueryOtherReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

func (builder *QueryOtherReqBuilder) PageNo(pageNo int) *QueryOtherReqBuilder {
	builder.apiReq.QueryParams.Set("page_no", fmt.Sprint(pageNo))
	return builder
}

func (builder *QueryOtherReqBuilder) CalcTotal(calcTotal int) *QueryOtherReqBuilder {
	builder.apiReq.QueryParams.Set("calc_total", fmt.Sprint(calcTotal))
	return builder
}

func (builder *QueryOtherReqBuilder) Body(body *QueryOtherBody) *QueryOtherReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder QueryOtherReqBuilder) Build() *QueryOtherReq {
	return &QueryOtherReq{builder.apiReq}
}

type QueryOtherReq struct {
	apiReq *core.ApiReq
}

type QueryOtherBody struct {
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
	*core.ApiResp
	core.CodeError
	Data QueryOtherData `json:"data"`
}

type QueryOtherData struct {
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

type PurchaseQueryReqBuilder struct {
	apiReq *core.ApiReq
}

func NewPurchaseQueryReqBuilder() *PurchaseQueryReqBuilder {
	return &PurchaseQueryReqBuilder{core.NewApiReq()}
}

func (builder *PurchaseQueryReqBuilder) PageSize(pageSize int) *PurchaseQueryReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

func (builder *PurchaseQueryReqBuilder) PageNo(pageNo int) *PurchaseQueryReqBuilder {
	builder.apiReq.QueryParams.Set("page_no", fmt.Sprint(pageNo))
	return builder
}

func (builder *PurchaseQueryReqBuilder) CalcTotal(calcTotal int) *PurchaseQueryReqBuilder {
	builder.apiReq.QueryParams.Set("calc_total", fmt.Sprint(calcTotal))
	return builder
}

func (builder *PurchaseQueryReqBuilder) Body(body *PurchaseQueryBody) *PurchaseQueryReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder PurchaseQueryReqBuilder) Build() *PurchaseQueryReq {
	return &PurchaseQueryReq{builder.apiReq}
}

type PurchaseQueryReq struct {
	apiReq *core.ApiReq
}
type PurchaseQueryBody struct {
	StartTime   string `json:"start_time"`             // 开始时间, 入库单修改时间, yyyy-MM-dd HH:mm:ss格式
	EndTime     string `json:"end_time"`               // 结束时间
	Status      string `json:"status,omitempty"`       // 10=已取消；20=编辑中；30=待审核；37=待质检；40=质检确认；80=已完成
	WarehouseNo string `json:"warehouse_no,omitempty"` // 仓库编号
	StockinNo   string `json:"stockin_no,omitempty"`   // 入库单号
	PurchaseNo  string `json:"purchase_no,omitempty"`  // 采购单号
}

type PurchaseQueryResp struct {
	*core.ApiResp
	core.CodeError
	Data *PurchaseQueryData `json:"data"`
}

type PurchaseQueryData struct {
	Order      []*PurchaseOrder `json:"order"`
	TotalCount int64            `json:"total_count"`
}

type PurchaseOrder struct {
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

type RefundQueryReqBuilder struct {
	apiReq *core.ApiReq
}

func NewRefundQueryReqBuilder() *RefundQueryReqBuilder {
	return &RefundQueryReqBuilder{core.NewApiReq()}
}

func (builder *RefundQueryReqBuilder) PageSize(pageSize int) *RefundQueryReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

func (builder *RefundQueryReqBuilder) PageNo(pageNo int) *RefundQueryReqBuilder {
	builder.apiReq.QueryParams.Set("page_no", fmt.Sprint(pageNo))
	return builder
}

func (builder *RefundQueryReqBuilder) CalcTotal(calcTotal int) *RefundQueryReqBuilder {
	builder.apiReq.QueryParams.Set("calc_total", fmt.Sprint(calcTotal))
	return builder
}

func (builder *RefundQueryReqBuilder) Body(body *RefundQueryBody) *RefundQueryReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder *RefundQueryReqBuilder) Build() *RefundQueryReq {
	return &RefundQueryReq{builder.apiReq}
}

type RefundQueryReq struct {
	apiReq *core.ApiReq
}

type RefundQueryBody struct {
	StartTime          string `json:"start_time"`                      // 开始时间, 入库单修改时间, yyyy-MM-dd HH:mm:ss格式
	EndTime            string `json:"end_time"`                        // 结束时间
	Status             string `json:"status,omitempty"`                // 10=已取消；20=编辑中；30=待审核；80=已完成
	WarehouseNo        string `json:"warehouse_no,omitempty"`          // 仓库编号
	StockinNo          string `json:"stockin_no,omitempty"`            // 入库单号
	RefundNo           string `json:"refund_no,omitempty"`             // 采购单号
	ShopNos            string `json:"shop_nos,omitempty"`              // 店铺编号
	TimeType           int    `json:"time_type,omitempty"`             // 时间条件类型 0:修改时间; 1:入库时间 不传默认为0
	IsSlave            bool   `json:"is_slave,omitempty"`              // 是否使用从库查询
	NeedSn             bool   `json:"need_sn,omitempty"`               // 是否返回sn信息
	NeedSummary        bool   `json:"need_summary,omitempty"`          // 是否需要汇总数据
	NeedGovSubsidyInfo int    `json:"need_gov_subsidy_info,omitempty"` // 是否需要国补信息
}

type RefundQueryResp struct {
	*core.ApiResp
	core.CodeError
	Data *RefundQueryData `json:"data"`
}

type RefundQueryData struct {
	Order      []*RefundOrder `json:"order"`
	TotalCount int64          `json:"total_count"`
}

type RefundSummary struct {
	WarehouseNo string `json:"warehouse_no"` // 仓库编号
	ShopData    []*struct {
		ShopNo   string `json:"shop_no"` // 店铺编号
		SpecData []*struct {
			SpecNo            string  `json:"spec_no"`             // 商家编码
			Num               float64 `json:"num"`                 // 数量
			TotalRefundAmount float64 `json:"total_refund_amount"` // 退款金额
		} `json:"spec_data"` // 单品数据汇总
	} `json:"shop_data"` // 店铺维度分组汇总
}

type RefundOrder struct {
	OrderNo              string                  `json:"order_no"`                                // 入库单号
	Status               int                     `json:"status"`                                  // 入库单状态
	AttachType           int                     `json:"attach_type"`                             // 关联类型
	WarehouseNo          string                  `json:"warehouse_no"`                            // 仓库编号
	WarehouseName        string                  `json:"warehouse_name"`                          // 仓库名称
	CreatedTime          int64                   `json:"created_time"`                            // 制单时间
	Remark               string                  `json:"remark"`                                  // 备注
	FenxiaoNick          string                  `json:"fenxiao_nick"`                            // 分销商昵称
	OperatorName         string                  `json:"operator_name"`                           // 入库人
	OperatorId           int                     `json:"operator_id"`                             // 入库人id
	LogisticsType        int                     `json:"logistics_type"`                          // 物流方式
	LogisticsName        string                  `json:"logistics_name"`                          // 物流公司
	LogisticsNo          string                  `json:"logistics_no"`                            // 物流单号
	LogisticsId          int                     `json:"logistics_id"`                            // 物流id
	CheckTime            int64                   `json:"check_time"`                              // 审核时间
	RefundNo             string                  `json:"refund_no"`                               // 退换单号
	GoodsCount           float64                 `json:"goods_count"`                             // 货品数量
	ActualRefundAmount   float64                 `json:"actual_refund_amount"`                    // 退换单实际退款金额
	CustomerNo           string                  `json:"customer_no"`                             // 客户编码
	CustomerName         string                  `json:"customer_name"`                           // 退货人姓名
	NickName             string                  `json:"nick_name"`                               // 客户网名
	ShopName             string                  `json:"shop_name"`                               // 店铺名称
	ShopNo               string                  `json:"shop_no"`                                 // 店铺编号
	ShopRemark           string                  `json:"shop_remark"`                             // 店铺备注
	FlagName             string                  `json:"flag_name"`                               // 标记名称
	TradeNoList          string                  `json:"trade_no_list"`                           // 系统订单编号
	TidList              string                  `json:"tid_list"`                                // 原始订单
	SrcOrderId           int                     `json:"src_order_id"`                            // 退换单id
	StockinId            int                     `json:"stockin_id"`                              // 入库单id
	ShopPlatformId       int                     `json:"shop_platform_id"`                        // 店铺平台id
	SubPlatformId        int                     `json:"sub_platform_id"`                         // 子平台id
	ShopId               string                  `json:"shop_id"`                                 // 店铺id
	WarehouseId          int                     `json:"warehouse_id"`                            // 仓库id
	TotalPrice           float64                 `json:"total_price"`                             // 入库单总成本
	TotalGoodsStockinNum string                  `json:"total_goods_stockin_num"`                 // 货品入库总数量
	ProcessStatus        int                     `json:"process_status"`                          // 退换单状态
	Modified             string                  `json:"modified"`                                // 修改时间
	CheckOperatorName    string                  `json:"check_operator_name"`                     // 审核人
	CheckOperatorId      int                     `json:"check_operator_id"`                       // 审核人id
	Reason               string                  `json:"reason"`                                  // 退换说明
	ReasonId             int                     `json:"reason_id"`                               // 退换说明id
	RefundAmount         float64                 `json:"refund_amount"`                           // 入库总金额
	AdjustNum            float64                 `json:"adjust_num"`                              // 调整数量
	WmsCode              string                  `json:"wms_code"`                                // wms业务单号
	Created              string                  `json:"created"`                                 // 创建时间
	FlagId               int                     `json:"flag_id"`                                 // 标记id
	GoodsTypeCount       int                     `json:"goods_type_count"`                        // 货品类型数量
	SrcOrderNo           string                  `json:"src_order_no"`                            // 退换单编号
	NoteCount            int                     `json:"note_count"`                              // 便签条数
	Prop3                string                  `json:"prop3"`                                   // 源预入库单号
	SrcOrderType         int                     `json:"src_order_type"`                          // 业务单类型
	Type                 int                     `json:"type"`                                    // 类型
	SumRefundAmount      float64                 `json:"sum_refund_amount"`                       // 货品退款金额
	ContactTime          string                  `json:"contact_time"`                            // 关联时间
	FenxiaoNickNo        string                  `json:"fenxiao_nick_no"`                         // 分销商编码
	TotalAmount          float64                 `json:"total_amount"`                            // 退货货品金额
	VirWarehouseNo       string                  `json:"vir_warehouse_no"`                        // 虚拟仓编码
	VirWarehouseName     string                  `json:"vir_warehouse_name"`                      // 虚拟仓名称
	DetailsList          []RefundOrderDetail     `json:"details_list" gorm:"serializer:json"`     // 入库单明细
	GovSubsidyInfo       []*model.GovSubsidyInfo `json:"gov_subsidy_info" gorm:"serializer:json"` // 国补信息
}

// RefundOrderDetail 入库单明细
type RefundOrderDetail struct {
	RecId                 int     `json:"rec_id"`                       // 入库单明细id(主键)
	StockinId             int     `json:"stockin_id"`                   // 入库单id
	RefundDetailId        string  `json:"refund_detail_id"`             // 退换单明细id
	Num                   float64 `json:"num"`                          // 数量
	TotalCost             float64 `json:"total_cost"`                   // 总成本
	Remark                string  `json:"remark"`                       // 备注
	RightNum              float64 `json:"right_num"`                    // 调整后数量
	GoodsName             string  `json:"goods_name"`                   // 货品名称
	GoodsNo               string  `json:"goods_no"`                     // 货品编码
	GoodsId               int     `json:"goods_id"`                     // 货品id
	SpecId                int     `json:"spec_id,omitempty"`            // 商品id
	SpecNo                string  `json:"spec_no"`                      // 商家编码
	Defect                bool    `json:"defect,omitempty"`             // 是否残次品
	Prop1                 string  `json:"prop1,omitempty"`              // 单品自定义属性1
	Prop2                 string  `json:"prop2,omitempty"`              // 单品自定义属性2
	Prop3                 string  `json:"prop3,omitempty"`              // 单品自定义属性3
	Prop4                 string  `json:"prop4,omitempty"`              // 单品自定义属性4
	Prop5                 string  `json:"prop5,omitempty"`              // 单品自定义属性5
	Prop6                 string  `json:"prop6,omitempty"`              // 单品自定义属性6
	SpecName              string  `json:"spec_name,omitempty"`          // 规格名称
	SpecCode              string  `json:"spec_code,omitempty"`          // 规格码
	BrandNo               string  `json:"brand_no"`                     // 品牌编号
	BrandName             string  `json:"brand_name"`                   // 品牌名称
	GoodsUnit             string  `json:"goods_unit"`                   // 辅助单位
	LogisticsName         string  `json:"logistics_name"`               // 物流名称
	LogisticsNo           string  `json:"logistics_no"`                 // 物流单号
	WarehouseId           int     `json:"warehouse_id"`                 // 仓库id
	SrcOrderId            int     `json:"src_order_id"`                 // 退换单id
	LogisticsId           int     `json:"logistics_id"`                 // 物流id
	BaseUnitName          string  `json:"base_unit_name"`               // 单位
	BatchNo               string  `json:"batch_no"`                     // 批次号
	ExpireDate            string  `json:"expire_date"`                  // 有效期
	ProductionDate        string  `json:"production_date"`              // 生产日期
	PositionNo            string  `json:"position_no,omitempty"`        // 货位
	ExpectNum             float64 `json:"expect_num,omitempty"`         // 预期数量
	StockinNum            string  `json:"stockin_num"`                  // 入库数量
	CheckedCostPrice      float64 `json:"checked_cost_price,omitempty"` // 入库单明细成本价
	RefundAmount          float64 `json:"refund_amount"`                // 退款总额
	SnList                string  `json:"sn_list,omitempty"`            // sn序列号信息
	UnitId                int     `json:"unit_id"`                      // 入库单位id
	BaseUnitId            int     `json:"base_unit_id"`                 // 基本单位id
	OrgStockinDetailId    int     `json:"org_stockin_detail_id"`        // 库存明细id
	BatchId               int     `json:"batch_id"`                     // 批次id
	PositionId            int     `json:"position_id"`                  // 货位id
	ValidityDays          int     `json:"validity_days"`                // 有效期天数
	Num2                  float64 `json:"num2"`                         // 辅助数量
	AdjustNum             float64 `json:"adjust_num"`                   // 调整数量
	UnitRatio             float64 `json:"unit_ratio"`                   // 单位换算关系
	Modified              string  `json:"modified"`                     // 最后修改时间
	Created               string  `json:"created"`                      // 创建时间
	SrcOrderType          int     `json:"src_order_type"`               // 业务单类型
	CustomerPrice1        float64 `json:"customer_price1,omitempty"`    // 自定义金额1
	CustomerPrice2        float64 `json:"customer_price2,omitempty"`    // 自定义金额2
	Tid                   string  `json:"tid,omitempty"`                // 原始单号
	ActualRefundAmount    string  `json:"actual_refund_amount"`         // 退款金额
	OriginalPrice         string  `json:"original_price"`               // 原价
	RefundOrderDetailList []struct {
		RefundOrderId         int     `json:"refund_order_id"`          // 退换单明细id
		StockinOrderDetailId  int     `json:"stockin_order_detail_id"`  // 入库单明细id
		Price                 float64 `json:"price"`                    // 价格
		SpecNo                string  `json:"spec_no"`                  // 商家编码
		StockinNum            float64 `json:"stockin_num"`              // 实际入库数量
		TradeType             int     `json:"trade_type"`               // 订单类型
		RefundOrderStockinNum float64 `json:"refund_order_stockin_num"` // 入库数量
		Oid                   string  `json:"oid"`                      // 原始子单号
		Tid                   string  `json:"tid"`                      // 原始单号
		ApiGoodsId            string  `json:"api_goods_id"`             // 平台货品id
		ApiSpecId             string  `json:"api_spec_id"`              // 平台规格id
		ApiGoodsName          string  `json:"api_goods_name"`           // 平台货品名称
		ApiSpecName           string  `json:"api_spec_name"`            // 平台规格名称
		TotalAmount           float64 `json:"total_amount"`             // 退款总额
		RefundAmount          float64 `json:"refund_amount"`            // 退款金额
	} `json:"refund_order_detail_list"` // 退换单明细
}
