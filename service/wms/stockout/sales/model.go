package sales

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
	StartTime          string `json:"start_time,omitempty"`            // 开始时间, 入库单修改时间, yyyy-MM-dd HH:mm:ss格式
	EndTime            string `json:"end_time,omitempty"`              // 结束时间
	StatusType         string `json:"status_type,omitempty"`           // 出库单状态
	Status             string `json:"status,omitempty"`                // 出库单状态详细
	WarehouseNo        string `json:"warehouse_no,omitempty"`          // 仓库编号
	StockoutNo         string `json:"stockout_no,omitempty"`           // 入库单号
	ShopNos            string `json:"shop_nos,omitempty"`              // 店铺编号
	SrcOrderNo         string `json:"src_order_no,omitempty"`          // 销售订单号
	LogisticsNo        string `json:"logistics_no,omitempty"`          // 物流单号
	NeedSn             *bool  `json:"need_sn,omitempty"`               // 是否返回sn信息
	Position           *int   `json:"position,omitempty"`              // 是否按照货位排序
	IsSlave            *bool  `json:"is_slave,omitempty"`              // 是否使用从库查询
	GetAnchor          *bool  `json:"get_anchor,omitempty"`            // 是否获取主播信息
	OrderType          *int   `json:"order_type,omitempty"`            // 排序类型
	NeedPickPosition   *int   `json:"need_pick_position,omitempty"`    // 是否需要拣货位明细
	NeedGovSubsidyInfo *bool  `json:"need_gov_subsidy_info,omitempty"` // 是否需要国补信息
	SrcTid             *bool  `json:"src_tid,omitempty"`               //  原始单号
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
	StockoutId           int                     `json:"stockout_id"`            // 出库单ID
	OrderNo              string                  `json:"order_no"`               // 出库单号
	VirtualWarehouseNo   string                  `json:"virtual_warehouse_no"`   // 虚拟仓编号
	VirtualWarehouseName string                  `json:"virtual_warehouse_name"` // 虚拟仓名称
	SrcOrderNo           string                  `json:"src_order_no"`           // 系统订单编号
	WarehouseNo          string                  `json:"warehouse_no"`           // 仓库编号
	WarehouseName        string                  `json:"warehouse_name"`         // 仓库名称
	ConsignTime          string                  `json:"consign_time"`           // 发货时间
	OrderType            int                     `json:"order_type"`             // 源单据类别
	GoodsCount           float64                 `json:"goods_count"`            // 货品数量
	LogisticsNo          string                  `json:"logistics_no"`           // 物流单号
	ReceiverName         string                  `json:"receiver_name"`          // 收件人姓名
	ReceiverCountry      int                     `json:"receiver_country"`       // 国家
	ReceiverProvince     int                     `json:"receiver_province"`      // 省份ID
	ReceiverCity         int                     `json:"receiver_city"`          // 城市ID
	ReceiverDistrict     int                     `json:"receiver_district"`      // 地区ID
	ReceiverAddress      string                  `json:"receiver_address"`       // 地址
	ReceiverMobile       string                  `json:"receiver_mobile"`        // 收件人手机
	ReceiverTelno        string                  `json:"receiver_telno"`         // 收件人固话
	ReceiverZip          string                  `json:"receiver_zip"`           // 收件人邮编
	ReceiverArea         string                  `json:"receiver_area"`          // 省市区
	Remark               string                  `json:"remark"`                 // 出库单备注
	Weight               float64                 `json:"weight"`                 // 重量
	BlockReason          int                     `json:"block_reason"`           // 截停原因
	LogisticsType        int                     `json:"logistics_type"`         // 物流方式
	LogisticsCode        string                  `json:"logistics_code"`         // 物流编号
	LogisticsName        string                  `json:"logistics_name"`         // 物流公司名称
	ShopId               int                     `json:"shop_id"`                // 店铺id
	WarehouseId          int                     `json:"warehouse_id"`           // 仓库id
	LogisticsId          int                     `json:"logistics_id"`           // 物流id
	BadReason            int                     `json:"bad_reason"`             // 异常原因
	ReceiverDtb          string                  `json:"receiver_dtb"`           // 大头笔
	RefundStatus         int                     `json:"refund_status"`          // 退款状态
	TradeType            int                     `json:"trade_type"`             // 销售类型
	SalesmanNo           string                  `json:"salesman_no"`            // 业务员编号
	Fullname             string                  `json:"fullname"`               // 业务员姓名
	PickerName           string                  `json:"picker_name"`            // 拣货员
	ExaminerName         string                  `json:"examiner_name"`          // 验货员
	ConsignerName        string                  `json:"consigner_name"`         // 发货员
	PrinterName          string                  `json:"printer_name"`           // 打单员
	PackagerName         string                  `json:"packager_name"`          // 打包员
	TradeStatus          int                     `json:"trade_status"`           // 订单状态
	TradeNo              string                  `json:"trade_no"`               // 订单编号
	SrcTradeNo           string                  `json:"src_trade_no"`           // 原始单号
	NickName             string                  `json:"nick_name"`              // 客户网名
	CustomerNo           string                  `json:"customer_no"`            // 客户编码
	CustomerName         string                  `json:"customer_name"`          // 客户姓名
	TradeTime            int64                   `json:"trade_time"`             // 下单时间
	PayTime              int64                   `json:"pay_time"`               // 支付时间
	FlagName             string                  `json:"flag_name"`              // 标记名称
	PostAmount           string                  `json:"post_amount"`            // 邮费
	IdCardType           int                     `json:"id_card_type"`           // 证件类别
	IdCard               string                  `json:"id_card"`                // 证件号码
	ShopName             string                  `json:"shop_name"`              // 店铺名称
	ShopNo               string                  `json:"shop_no"`                // 店铺编号
	ShopRemark           string                  `json:"shop_remark"`            // 店铺备注
	Status               int                     `json:"status"`                 // 出库单状态
	InvoiceType          int                     `json:"invoice_type"`           // 发票类型
	InvoiceId            int                     `json:"invoice_id"`             // 发票ID
	CodAmount            string                  `json:"cod_amount"`             // 货到付款金额
	DeliveryTerm         int                     `json:"delivery_term"`          // 发货条件
	PlatformId           int                     `json:"platform_id"`            // 平台ID
	TradeId              int                     `json:"trade_id"`               // 订单ID
	EmployeeNo           string                  `json:"employee_no"`            // 审核员编号
	Discount             string                  `json:"discount"`               // 优惠金额
	Tax                  string                  `json:"tax"`                    // 税额
	TaxRate              string                  `json:"tax_rate"`
	Currency             string                  `json:"currency"`                                // 币种
	Created              int64                   `json:"created"`                                 // 系统订单建单时间
	StockCheckTime       int64                   `json:"stock_check_time"`                        // 出库单建单时间
	PrintRemark          string                  `json:"print_remark"`                            // 打印备注
	BuyerMessage         string                  `json:"buyer_message"`                           // 买家留言
	CsRemark             string                  `json:"cs_remark"`                               // 客服备注
	InvoiceTitle         string                  `json:"invoice_title"`                           // 发票抬头
	InvoiceContent       string                  `json:"invoice_content"`                         // 发票内容
	PostFee              string                  `json:"post_fee"`                                // 称重预估邮资
	PackageFee           string                  `json:"package_fee"`                             // 包装成本
	Receivable           float64                 `json:"receivable"`                              // 已付金额
	GoodsTotalCost       string                  `json:"goods_total_cost"`                        // 总成本价
	GoodsTotalAmount     string                  `json:"goods_total_amount"`                      // 总货款
	Modified             string                  `json:"modified"`                                // 最后修改时间
	FenxiaoNick          string                  `json:"fenxiao_nick"`                            // 分销商昵称
	TradeLabel           string                  `json:"trade_label"`                             // 订单标签
	TradeFrom            int                     `json:"trade_from"`                              // 订单来源
	PicklistNo           string                  `json:"picklist_no"`                             // 分拣波次
	PicklistSeq          int                     `json:"picklist_seq"`                            // 分拣序号
	LogisticsPrintStatus int                     `json:"logistics_print_status"`                  // 物流单打印状态
	Paid                 string                  `json:"paid"`                                    // 已付
	ShopPlatformId       int                     `json:"shop_platform_id"`                        // 店铺平台id
	SubPlatformId        int                     `json:"sub_platform_id"`                         // 子平台id
	ErrorInfo            string                  `json:"error_info"`                              // 接口处理错误信息
	CustomType           int                     `json:"custom_type"`                             // 其他出库自定义子类别
	SendbillTemplateId   int                     `json:"sendbill_template_id"`                    // 发货单模板id
	CustomerId           int                     `json:"customer_id"`                             // 客户id
	WarehouseType        int                     `json:"warehouse_type"`                          // 仓库类别
	OperatorId           int                     `json:"operator_id"`                             // 制单人id（操作员）
	OuterNo              string                  `json:"outer_no"`                                // 外部单号
	ConsignStatus        int                     `json:"consign_status"`                          // 出库状态
	GoodsTypeCount       int                     `json:"goods_type_count"`                        // 货品种类
	CalcPostCost         string                  `json:"calc_post_cost"`                          // 预估邮资成本
	BatchNo              string                  `json:"batch_no"`                                // 打印批次
	CreatedDate          string                  `json:"created_date"`                            // 销售出库单创建时间
	FenxiaoTid           string                  `json:"fenxiao_tid"`                             // 分销原始单号
	FenxiaoNickNo        string                  `json:"fenxiao_nick_no"`                         // 分销商编号
	LogisticsList        []model.LogisticsDetail `json:"logistics_list" gorm:"serializer:json"`   // 物流单列表
	DetailsList          []OrderDetail           `json:"details_list" gorm:"serializer:json"`     // 销售出库单详情
	AnchorName           string                  `json:"anchor_name"`                             // 主播
	AssistAchorName      string                  `json:"assist_achor_name"`                       // 助播
	ControlAchorName     string                  `json:"control_achor_name"`                      // 场控
	OperationAnchorName  string                  `json:"operation_anchor_name"`                   // 运营
	PickGroupName        string                  `json:"pick_group_name"`                         // 拣货分组名称
	FenxiaoShopName      string                  `json:"fenxiao_shop_name"`                       // 分销店铺名称
	WmsCode              string                  `json:"wms_code"`                                // wms业务单号
	StockoutFlagName     string                  `json:"stockout_flag_name"`                      // 销售出库单标记名称
	GovSubsidyInfo       []*model.GovSubsidyInfo `json:"gov_subsidy_info" gorm:"serializer:json"` // 国补信息
}

type OrderDetail struct {
	RecId               int     `json:"rec_id"`                  // 销售出库单详情的id
	StockoutId          int     `json:"stockout_id"`             // 出库单id
	SrcOrderDetailId    int     `json:"src_order_detail_id"`     // 订单明细id
	SpecId              int     `json:"spec_id"`                 // 单品id
	SpecNo              string  `json:"spec_no"`                 // 商家编码
	GoodsCount          float64 `json:"goods_count"`             // 货品数量
	TotalAmount         float64 `json:"total_amount,omitempty"`  // 总成本
	SellPrice           float64 `json:"sell_price"`              // 成交价
	Remark              string  `json:"remark"`                  // 出库单明细备注
	GoodsName           string  `json:"goods_name"`              // 货品名称
	GoodsNo             string  `json:"goods_no"`                // 货品编号
	SpecName            string  `json:"spec_name"`               // 规格名称
	SpecCode            string  `json:"spec_code,omitempty"`     // 规格码
	CostPrice           float64 `json:"cost_price,omitempty"`    // 货品成本
	Weight              float64 `json:"weight,omitempty"`        // 总重量
	GoodsId             int     `json:"goods_id"`                // 货品id
	Prop1               string  `json:"prop1,omitempty"`         // 规格自定义属性1
	Prop2               string  `json:"prop2,omitempty"`         // 规格自定义属性2
	Prop3               string  `json:"prop3,omitempty"`         // 规格自定义属性3
	Prop4               string  `json:"prop4,omitempty"`         // 规格自定义属性4
	Prop5               string  `json:"prop5,omitempty"`         // 规格自定义属性5
	Prop6               string  `json:"prop6,omitempty"`         // 规格自定义属性6
	PlatformId          int     `json:"platform_id"`             // 平台id
	RefundStatus        int     `json:"refund_status,omitempty"` // 退款状态
	MarketPrice         float64 `json:"market_price"`            // 单价/货品原单价
	Discount            float64 `json:"discount"`                // 货品总优惠
	SharePrice          float64 `json:"share_price"`             // 货品成交价
	ShareAmount         float64 `json:"share_amount"`            // 总货款/货品成交总价
	TaxRate             float64 `json:"tax_rate,omitempty"`      // 税率
	Barcode             string  `json:"barcode"`                 // 主条码
	UnitName            string  `json:"unit_name,omitempty"`     // 基本单位名称
	SaleOrderId         int     `json:"sale_order_id"`           // 订单货品(子订单)id
	GiftType            int     `json:"gift_type"`               // 是否是赠品
	SrcOid              string  `json:"src_oid"`                 // 原始子订单号
	SrcTid              string  `json:"src_tid"`                 // 子单原始订单号
	FromMask            int     `json:"from_mask,omitempty"`     // 订单内部来源
	GoodsType           int     `json:"goods_type"`              // 货品类型
	GoodProp1           string  `json:"good_prop1,omitempty"`    // 货品自定义属性1
	GoodProp2           string  `json:"good_prop2,omitempty"`    // 货品自定义属性2
	GoodProp3           string  `json:"good_prop3,omitempty"`    // 货品自定义属性3
	GoodProp4           string  `json:"good_prop4,omitempty"`    // 货品自定义属性4
	GoodProp5           string  `json:"good_prop5,omitempty"`    // 货品自定义属性5
	GoodProp6           string  `json:"good_prop6,omitempty"`    // 货品自定义属性6
	SnList              string  `json:"sn_list,omitempty"`       // sn_list
	SuiteNo             string  `json:"suite_no,omitempty"`      // 组合装编码
	SuiteNum            float64 `json:"suite_num,omitempty"`     // 组合装数量
	SharePostAmount     float64 `json:"share_post_amount"`       // 分摊邮费
	Paid                float64 `json:"paid"`                    // 已付
	IsPackage           bool    `json:"is_package,omitempty"`    // 是否包装
	BrandNo             string  `json:"brand_no"`                // 品牌编号
	BrandName           string  `json:"brand_name"`              // 品牌名称
	SrcOrderType        int     `json:"src_order_type"`          // 源单据类别
	BaseUnitId          int     `json:"base_unit_id"`            // 基本单位id
	UnitId              int     `json:"unit_id,omitempty"`       // 辅助单位
	UnitRatio           float64 `json:"unit_ratio"`              // 单位换算
	Num2                float64 `json:"num2"`                    // 辅助数量
	Num                 float64 `json:"num"`                     // 货品数量
	PositionId          int     `json:"position_id,omitempty"`   // 出库货位id
	BatchId             int     `json:"batch_id,omitempty"`      // 指定出库批次
	IsExamined          int     `json:"is_examined,omitempty"`   // 是否验货
	ExpireDate          int64   `json:"expire_date,omitempty"`   // 有效期
	ScanType            int     `json:"scan_type,omitempty"`     // 扫描方式
	ModifiedDate        string  `json:"modified_date"`           // 最后修改时间
	CreatedDate         string  `json:"created_date"`            // 创建时间
	ClassName           string  `json:"class_name"`              // 分类
	ApiGoodsId          string  `json:"api_goods_id"`            // 平台货品id
	ApiSpecId           string  `json:"api_spec_id"`             // 平台规格id
	PackScore           float64 `json:"pack_score,omitempty"`    // 打包积分
	PickScore           float64 `json:"pick_score,omitempty"`    // 拣货积分
	ScanScore           float64 `json:"scan_score,omitempty"`    // 验货积分
	ApiBoodsName        string  `json:"api_goods_name"`          // 平台货品名称
	PositionDetailsList []struct {
		model.PositionDetail
		BatchRemark string `json:"batch_remark,omitempty"` // 批次备注
	} `json:"position_details_list,omitempty"` //出库货位明细
	PickPositionDetailsList []struct {
		model.PositionDetail
		SpecNo string `json:"spec_no,omitempty"` // 商家编码
	} `json:"pick_position_details_list,omitempty"` // 	拣货位明细
}
