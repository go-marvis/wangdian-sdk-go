package trade_query

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
	StartTime          string `json:"start_time"`                      //开始时间
	EndTime            string `json:"end_time"`                        //结束时间
	WarehouseNo        string `json:"warehouse_no,omitempty"`          //仓库编号
	Status             string `json:"status,omitempty"`                //订单状态
	TradeNo            string `json:"trade_no,omitempty"`              //订单编号
	ShopNo             string `json:"shop_no,omitempty"`               //店铺编号
	LogisticsNo        string `json:"logistics_no,omitempty"`          //物流单号
	SrcTid             string `json:"src_tid,omitempty"`               //原始单号
	IsSlave            bool   `json:"is_slave,omitempty"`              //是否使用从库查询
	CalSharePostAmount bool   `json:"cal_share_post_amount,omitempty"` //是否计算分摊邮费
	TradeFrom          string `json:"trade_from,omitempty"`            //订单来源
	OrderType          int    `json:"order_type,omitempty"`            //排序类型
	TimeType           int    `json:"time_type,omitempty"`             //时间类型
	NeedGiftRelation   *bool  `json:"need_gift_relation,omitempty"`    //是否返回赠品关联关系
	AccurateQuery      bool   `json:"accurate_query,omitempty"`        //是否强制指定src_tid精准查询
	CutLogisticsNo     bool   `json:"cut_logistics_no,omitempty"`      //是否截取物流单号
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
	TradeId             int64   `json:"trade_id"`              // 订单唯一键
	TradeNo             string  `json:"trade_no"`              // 订单编号
	PlatformId          int     `json:"platform_id"`           // 平台ID
	WarehouseType       int     `json:"warehouse_type"`        // 仓库类型
	SrcTids             string  `json:"src_tids"`              // 原始单号
	PayAccount          string  `json:"pay_account"`           // 平台支付帐号
	TradeStatus         int     `json:"trade_status"`          // 订单状态
	TradeType           int     `json:"trade_type"`            // 订单类型
	DeliveryTerm        int     `json:"delivery_term"`         // 发货条件
	ReceiverRing        string  `json:"receiver_ring"`         // 京东几环
	FreezeReason        string  `json:"freeze_reason"`         // 冻结原因
	RefundStatus        int     `json:"refund_status"`         // 退款状态
	FenxiaoType         int     `json:"fenxiao_type"`          // 分销类别
	FenxiaoNick         string  `json:"fenxiao_nick"`          // 分销商昵称
	TradeTime           int64   `json:"trade_time"`            // 下单时间
	PayTime             string  `json:"pay_time"`              // 付款时间
	ConsignTime         int64   `json:"consign_time"`          // 发货时间
	BuyerNick           string  `json:"buyer_nick"`            // 客户网名
	ReceiverName        string  `json:"receiver_name"`         // 收货人/收件人
	ReceiverProvince    int     `json:"receiver_province"`     // 省份id
	ReceiverCity        int     `json:"receiver_city"`         // 城市id
	ReceiverDistrict    int     `json:"receiver_district"`     // 地区id
	ReceiverAddress     string  `json:"receiver_address"`      // 地址
	ReceiverMobile      string  `json:"receiver_mobile"`       // 手机
	ReceiverTelno       string  `json:"receiver_telno"`        // 固话
	ReceiverZip         string  `json:"receiver_zip"`          // 邮编
	ReceiverArea        string  `json:"receiver_area"`         // 地区
	ReceiverDtb         string  `json:"receiver_dtb"`          // 大头笔
	BadReason           int     `json:"bad_reason"`            // 异常订单原因
	LogisticsNo         string  `json:"logistics_no"`          // 物流单号
	BuyerMessage        string  `json:"buyer_message"`         // 买家留言
	CsRemark            string  `json:"cs_remark"`             // 客服备注
	RemarkFlag          int     `json:"remark_flag"`           // 标旗
	PrintRemark         string  `json:"print_remark"`          // 打印备注
	GoodsTypeCount      float64 `json:"goods_type_count"`      // 货品种类数
	GoodsCount          float64 `json:"goods_count"`           // 货品总量
	GoodsAmount         float64 `json:"goods_amount"`          // 总货款
	PostAmount          float64 `json:"post_amount"`           // 邮费
	OtherAmount         float64 `json:"other_amount"`          // 其他费用
	Discount            float64 `json:"discount"`              // 优惠
	Receivable          float64 `json:"receivable"`            // 应收
	CodAmount           float64 `json:"cod_amount"`            //COD金额
	ExtCodFee           float64 `json:"ext_cod_fee"`           // 买家COD费用
	GoodsCost           float64 `json:"goods_cost"`            // 预估货品成本
	PostCost            float64 `json:"post_cost"`             // 预估邮资成本
	Weight              float64 `json:"weight"`                // 预估重量
	Profit              float64 `json:"profit"`                // 预估毛利
	Tax                 float64 `json:"tax"`                   // 税额
	TaxRate             float64 `json:"tax_rate,omitempty"`    // 税率
	Commission          float64 `json:"commission"`            // 佣金
	InvoiceType         int     `json:"invoice_type"`          // 发票类型
	InvoiceTitle        string  `json:"invoice_title"`         // 发票抬头
	InvoiceContent      string  `json:"invoice_content"`       // 发票内容
	SalesmanName        string  `json:"salesman_name"`         // 业务员
	CheckerName         string  `json:"checker_name"`          // 审核人
	FcheckerName        string  `json:"fchecker_name"`         // 财审人
	CheckouterName      string  `json:"checkouter_name"`       // 签出人
	StockoutNo          string  `json:"stockout_no"`           // 出库单号
	FlagName            string  `json:"flag_name"`             // 标记名称
	TradeFrom           int     `json:"trade_from"`            // 订单来源
	SingleSpecNo        string  `json:"single_spec_no"`        // 货品商家编码
	RawGoodsCount       float64 `json:"raw_goods_count"`       // 原始货品数量
	RawGoodsTypeCount   int     `json:"raw_goods_type_count"`  // 原始货品种类数
	Currency            string  `json:"currency"`              // 币种
	InvoiceId           int     `json:"invoice_id"`            // 发票ID
	VersionId           int     `json:"version_id"`            // 版本号
	Modified            string  `json:"modified"`              // 修改时间
	Created             int64   `json:"created"`               // 递交时间
	CheckTime           string  `json:"check_time"`            // 审核时间
	IdCardType          int     `json:"id_card_type"`          // 证件类别
	ShopNo              string  `json:"shop_no"`               // 店铺编号
	ShopName            string  `json:"shop_name"`             // 店铺名称
	ShopRemark          string  `json:"shop_remark"`           // 店铺备注
	WarehouseNo         string  `json:"warehouse_no"`          // 仓库编号
	CustomerName        string  `json:"customer_name"`         // 客户姓名
	CustomerNo          string  `json:"customer_no"`           // 客户编码
	LogisticsName       string  `json:"logistics_name"`        // 物流公司名称
	LogisticsCode       string  `json:"logistics_code"`        // 物流公司编号
	LogisticsTypeName   string  `json:"logistics_type_name"`   // 物流类型
	ToDeliverTime       string  `json:"to_deliver_time"`       // 送货时间
	DelayToTime         string  `json:"delay_to_time"`         // 计划发货时间
	EstimateConsignTime string  `json:"estimate_consign_time"` // 最晚发货时间
	ShopId              int     `json:"shop_id"`               // 店铺id
	WarehouseId         int     `json:"warehouse_id"`          // 仓库id
	Volume              float64 `json:"volume"`                // 体积
	TradeLabel          string  `json:"trade_label"`           // 订单标签
	TradeMask           int     `json:"trade_mask"`            // 订单掩码
	ShopPlatformId      int     `json:"shop_platform_id"`      // 店铺平台id
	SubPlatformId       int     `json:"sub_platform_id"`       // 子平台id
	PackageName         string  `json:"package_name"`          // 包装
	PackageId           int     `json:"package_id"`            // 包装id
	PackageCost         string  `json:"package_cost"`          // 包装成本
	Paid                float64 `json:"paid"`                  // 已付
	LargeType           int     `json:"large_type"`            // 大件类型
	GiftMask            int     `json:"gift_mask"`             // 赠品标记
	CustomerId          int     `json:"customer_id"`           // 客户id
	OtherCost           float64 `json:"other_cost"`            // 其他成本
	IsSealed            bool    `json:"is_sealed"`             // 不可合并拆分
	CustomerType        int     `json:"customer_type"`         // 客户类型
	LogisticsId         int     `json:"logistics_id"`          // 物流公司id
	CancelReason        string  `json:"cancel_reason"`         // 取消原因
	RevertReason        string  `json:"revert_reason"`         // 驳回原因
	NewTradeLabel       string  `json:"new_trade_label"`       // 订单标签mask
	FenxiaoTid          string  `json:"fenxiao_tid"`           // 分销原始单号
	CustomerUniqueId    string  `json:"customer_unique_id"`    // 客户唯一编码
	PlatformLabel       string  `json:"platform_label"`        // 平台标签
	FenxiaoNickNo       string  `json:"fenxiao_nick_no"`       // 分销商编号
	PreSyncLogisticsNo  string  `json:"pre_sync_logistics_no"` // 预物流同步单号
	PreSyncTime         string  `json:"pre_sync_time"`         // 预物流同步时间
	EndTime             string  `json:"end_time"`              // 订单结束时间
	SettleTime          string  `json:"settle_time"`           // 订单结算时间
	Tasks               string  `json:"tasks"`                 // 订单便签
	DetailList          []struct {
		TradeId           int64   `json:"trade_id"`                      // 订单唯一键
		RecId             int64   `json:"rec_id"`                        // 订单明细唯一键
		PlatformId        int     `json:"platform_id"`                   // 平台ID
		SrcOid            string  `json:"src_oid"`                       // 原始子单号
		SrcTid            string  `json:"src_tid"`                       // 原始单号
		GiftType          int     `json:"gift_type,omitempty"`           // 赠品方式
		PayStatus         int     `json:"pay_status"`                    // 支付状态
		RefundStatus      int     `json:"refund_status,omitempty"`       // 退款状态
		GuaranteeMode     int     `json:"guarantee_mode"`                // 退款模式
		PlatformStatus    int     `json:"platform_status"`               // 子单平台状态
		DeliveryTerm      int     `json:"delivery_term"`                 // 发货条件
		Num               float64 `json:"num"`                           // 数量
		Price             float64 `json:"price,omitempty"`               // 标价
		RefundNum         float64 `json:"refund_num,omitempty"`          // 售后退款数量
		OrderPrice        float64 `json:"order_price,omitempty"`         // 成交价
		SharePrice        float64 `json:"share_price,omitempty"`         // 分摊后价格
		Adjust            float64 `json:"adjust,omitempty"`              // 手工调整价
		Discount          float64 `json:"discount,omitempty"`            // 优惠
		ShareAmount       float64 `json:"share_amount,omitempty"`        // 分摊后总价
		TaxRate           float64 `json:"tax_rate"`                      // 税率
		GoodsName         string  `json:"goods_name"`                    // 货品名称
		GoodsNo           string  `json:"goods_no"`                      // 货品编号
		SpecName          string  `json:"spec_name,omitempty"`           // 规格名称
		SpecNo            string  `json:"spec_no"`                       // 商家编码
		SpecCode          string  `json:"spec_code,omitempty"`           // 规格码
		SuiteNo           string  `json:"suite_no"`                      // 组合装编码
		SuiteName         string  `json:"suite_name"`                    // 组合装名称
		SuiteNum          float64 `json:"suite_num"`                     // 组合装数量
		SuiteAmount       float64 `json:"suite_amount"`                  // 组合装分摊后总价
		SuiteDiscount     float64 `json:"suite_discount,omitempty"`      // 组合装优惠
		ApiGoodsName      string  `json:"api_goods_name"`                // 平台货品名称
		ApiSpecName       string  `json:"api_spec_name"`                 // 平台规格名称
		ApiGoodsId        string  `json:"api_goods_id"`                  // 平台货品id
		ApiSpecId         string  `json:"api_spec_id"`                   // 平台规格id
		GoodsId           int     `json:"goods_id"`                      // 货品id
		SpecId            int     `json:"spec_id"`                       // 单品id
		Commission        float64 `json:"commission,omitempty"`          // 佣金
		GoodsType         int     `json:"goods_type"`                    // 货品类型
		FromMask          int     `json:"from_mask"`                     // 订单内部来源
		FromMaskExt       int     `json:"from_mask_ext,omitempty"`       // 订单内部来源扩展
		Remark            string  `json:"remark"`                        // 子单备注
		Modified          int64   `json:"modified"`                      // 修改时间
		Created           int64   `json:"created"`                       // 创建时间
		Prop1             string  `json:"prop1,omitempty"`               // 单品自定义属性1
		Prop2             string  `json:"prop2,omitempty"`               // 单品自定义属性2
		Prop3             string  `json:"prop3,omitempty"`               // 单品自定义属性3
		Prop4             string  `json:"prop4,omitempty"`               // 单品自定义属性4
		Prop5             string  `json:"prop5,omitempty"`               // 单品自定义属性5
		Prop6             string  `json:"prop6,omitempty"`               // 单品自定义属性6
		Weight            float64 `json:"weight,omitempty"`              // 货品重量
		ImgUrl            string  `json:"img_url,omitempty"`             // 图片路径
		ActualNum         float64 `json:"actual_num"`                    // 实发数量
		Barcode           string  `json:"barcode"`                       // 条码
		Paid              float64 `json:"paid,omitempty"`                // 已付
		SuiteId           int     `json:"suite_id"`                      // 组合装id
		BindOid           string  `json:"bind_oid,omitempty"`            //Bind_oid
		PrintSuiteMode    int     `json:"print_suite_mode,omitempty"`    // 打印组合装
		Flag              int     `json:"flag"`                          // 天猫物流升级信息
		StockState        int     `json:"stock_state,omitempty"`         // 库存保留情况
		IsConsigned       bool    `json:"is_consigned"`                  // 平台已发货
		IsReceived        int     `json:"is_received,omitempty"`         // 是否付款
		Cid               int     `json:"cid,omitempty"`                 // 平台类目主键
		ModifiedDate      string  `json:"modified_date"`                 // 最后更新时间
		CreatedDate       string  `json:"created_date"`                  // 创建时间
		SharePostPrice    float64 `json:"share_post_price"`              // 分摊邮费
		InvoiceContent    string  `json:"invoice_content"`               // 发票内容
		PayTime           string  `json:"pay_time"`                      // 支付时间
		ShortName         string  `json:"short_name"`                    // 货品简称
		GoodsGiftRelation string  `json:"goods_gift_relation,omitempty"` // 赠品关联关系
		AvgCost           float64 `json:"avg_cost,omitempty"`            // 成本价
		ReserveNum        float64 `json:"reserve_num,omitempty"`         // 已占用库存
	} `json:"detail_list" gorm:"serializer:json"` // 订单明细
}
