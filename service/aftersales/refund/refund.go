package refund

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type refund struct {
	config *core.Config
}

type RefundSearchReq struct {
	ShopNos                   string `json:"shop_nos,omitempty"`                     // 店铺编号
	Tid                       string `json:"tid,omitempty"`                          // 原始单号
	BuyerNick                 string `json:"buyer_nick,omitempty"`                   // 客户网名
	TradeNo                   string `json:"trade_no,omitempty"`                     // 系统订单编号
	RefundNo                  string `json:"refund_no,omitempty"`                    // 退换单号
	ReturnLogisticsNo         string `json:"return_logistics_no,omitempty"`          // 物流单号
	ModifiedFrom              string `json:"modified_from,omitempty"`                // 修改起始时间
	ModifiedTo                string `json:"modified_to,omitempty"`                  // 修改结束时间
	SettleFrom                string `json:"settle_from,omitempty"`                  // 开始时间
	SettleTo                  string `json:"settle_to,omitempty"`                    // 结束时间
	AgreeFrom                 string `json:"agree_from,omitempty"`                   // 审核起始时间
	AgreeTo                   string `json:"agree_to,omitempty"`                     // 审核结束时间
	Status                    int    `json:"status,omitempty"`                       // 退换单状态
	StockinStatus             int    `json:"stockin_status,omitempty"`               // 入库状态
	ReturnWarehouseNos        string `json:"return_warehouse_nos,omitempty"`         // 退回仓库编号
	Type                      int    `json:"type,omitempty"`                         // 退换单类型
	FuzzyQuery                bool   `json:"fuzzy_query,omitempty"`                  // 模糊查询
	ReturnMaskManualintercept bool   `json:"return_mask_manual_intercept,omitempty"` // 是否为拦截件
	RefundMaskRejectReturn    bool   `json:"refund_mask_reject_return,omitempty"`    // 是否拒收
	OrderType                 int    `json:"order_type,omitempty"`                   // 排序类型
	RrStatus                  int    `json:"rr_status,omitempty"`                    // 平台状态
	CurrentPhaseTimeoutFrom   string `json:"current_phase_timeout_from,omitempty"`   // 退款成功时间起始
	CurrentPhaseTimeoutTo     string `json:"current_phase_timeout_to,omitempty"`     // 退款成功时间结束
	NeedGovSubsidyInfo        bool   `json:"need_gov_subsidy_info,omitempty"`        // 是否需要国补信息
	RawRefundNos              string `json:"raw_refund_nos,omitempty"`               // 原始退款单号
}

type RefundSearchResp struct {
	Order      []*RefundSearchOrder `json:"order"`
	TotalCount int64                `json:"total_count"`
}

type RefundSearchOrder struct {
	RefundId              int     `json:"refund_id"`               // 退换单id
	SrcTids               string  `json:"src_tids"`                // 原始单号
	RefundNo              string  `json:"refund_no"`               // 退换单号(唯一)
	Remark                string  `json:"remark"`                  // 备注
	Type                  int     `json:"type"`                    // 退换单类型
	StockinStatus         int     `json:"stockin_status"`          // 入库状态
	FlagName              string  `json:"flag_name"`               // 标记名称
	ReturnGoodsCount      float64 `json:"return_goods_count"`      // 退回货品数量
	ReceiverTelno         string  `json:"receiver_telno"`          // 退款订单中收件人电话
	ReceiverName          string  `json:"receiver_name"`           // 退款订单中收件人姓名
	Modified              int64   `json:"modified"`                // 修改时间
	NoteCount             int     `json:"note_count"`              // 便签数量
	ShopNo                string  `json:"shop_no"`                 // 店铺编号
	FromType              int     `json:"from_type"`               // 建单方式
	Created               int64   `json:"created"`                 // 建单时间
	SettleTime            string  `json:"settle_time"`             // 结算时间
	CheckTime             string  `json:"check_time"`              // 审核时间
	ReturnLogisticsNo     string  `json:"return_logistics_no"`     // 退货物流单号
	TradeNoList           string  `json:"trade_no_list"`           // 系统订单号列表
	GuaranteeRefundAmount float64 `json:"guarantee_refund_amount"` // 平台退款金额
	ReturnGoodsAmount     float64 `json:"return_goods_amount"`     // 退货金额
	ReturnLogisticsName   string  `json:"return_logistics_name"`   // 物流公司名称
	ReasonName            string  `json:"reason_name"`             // 退换说明
	RefundReason          string  `json:"refund_reason"`           // 退款原因
	BuyerNick             string  `json:"buyer_nick"`              // 客户网名
	OperatorName          string  `json:"operator_name"`           // 建单者
	ActualRefundAmount    float64 `json:"actual_refund_amount"`    // 实际退款金额
	RevertReasonName      string  `json:"revert_reason_name"`      // 驳回原因
	ReturnWarehouseNo     string  `json:"return_warehouse_no"`     // 退回仓库编号
	DirectRefundAmount    float64 `json:"direct_refund_amount"`    // 线下退款金额
	ReceiveAmount         float64 `json:"receive_amount"`          // 收款金额
	CustomerName          string  `json:"customer_name"`           // 客户姓名
	FenxiaoNickName       string  `json:"fenxiao_nick_name"`       // 分销商昵称
	Status                int     `json:"status"`                  // 退换单状态
	ShopId                int     `json:"shop_id"`                 // 店铺id
	TradeId               int     `json:"trade_id"`                // 订单id
	RawRefundNos          string  `json:"raw_refund_nos"`          // 原始退换单号
	PayId                 string  `json:"pay_id"`                  // 支付订单号
	ProviderRefundNo      string  `json:"provider_refund_no"`      // 分销退换单号
	ShopPlatformId        int     `json:"shop_platform_id"`        // 店铺平台id
	TidList               string  `json:"tid_list"`                // 原始单号
	SubPlatformId         int     `json:"sub_platform_id"`         // 子平台id
	ReturnWarehouseId     int     `json:"return_warehouse_id"`     // 退回仓库id
	PlatformId            string  `json:"platform_id"`             // 平台id
	WmsOwnerNo            string  `json:"wms_owner_no"`            // 奇门货主编号
	WarehouseType         int     `json:"warehouse_type"`          // 退回仓库类型
	BadReason             int     `json:"bad_reason"`              // 拦截原因
	ModifiedDate          string  `json:"modified_date"`           // 最后修改时间
	ReturnMaskInfo        string  `json:"return_mask_info"`        // 退换信息
	ProcessStatus         string  `json:"process_status"`          // 处理状态
	ReasonId              int     `json:"reason_id"`               // 退款原因id
	RevertReason          int     `json:"revert_reason"`           // 驳回原因id
	CustomerId            int     `json:"customer_id"`             // 客户id
	ConsignMode           int     `json:"consign_mode"`            // 发货方式
	RefundTime            string  `json:"refund_time"`             // 退款创建时间
	FenxiaoTid            string  `json:"fenxiao_tid"`             // 分销原始单号
	FenxiaoNickNo         string  `json:"fenxiao_nick_no"`         // 分销商编码
	WmsCode               string  `json:"wms_code"`                // wms单号
	RrStatus              int     `json:"rr_status"`               // 平台状态
	CurrentPhaseTimeout   string  `json:"current_phase_timeout"`   // 退款成功时间
	DetailList            []struct {
		RecId            int     `json:"rec_id"`                       // 退换单明细Id
		RefundId         int     `json:"refund_id"`                    // 退换单id
		Od               string  `json:"oid"`                          // 原始子单号
		TradeOrderId     int     `json:"trade_order_id"`               // 订单明细id
		PlatformId       int     `json:"platform_id"`                  // 平台id
		Tid              string  `json:"tid"`                          // 原始单号
		TradeNo          string  `json:"trade_no"`                     // 系统订单编号
		Num              float64 `json:"num"`                          // 数量
		Price            float64 `json:"price,omitempty"`              // 价格
		OriginalPrice    float64 `json:"original_price,omitempty"`     // 原价
		CheckedCostPrice float64 `json:"checked_cost_price,omitempty"` // 成本价
		RefundNum        float64 `json:"refund_num"`                   // 退款数量
		TotalAmount      float64 `json:"total_amount,omitempty"`       // 退款总额
		RefundAmount     float64 `json:"refund_amount,omitempty"`      // 退款金额
		IsGuarantee      bool    `json:"is_guarantee,omitempty"`       // 担保退款
		GoodsNo          string  `json:"goods_no"`                     // 货品编号
		GoodsName        string  `json:"goods_name"`                   // 货品名称
		SpecName         string  `json:"spec_name"`                    // 规格名
		SpecNo           string  `json:"spec_no"`                      // 商家编码
		GoodsId          string  `json:"goods_id,omitempty"`           // 平台货品id
		SpecId           string  `json:"spec_id,omitempty"`            // 平台规格id
		SysGoodsId       int     `json:"sys_goods_id"`                 // 系统货品id
		SysSpecId        int     `json:"sys_spec_id"`                  // 系统规格id
		SpecCode         string  `json:"spec_code,omitempty"`          // 规格码
		Barcode          string  `json:"barcode"`                      // 条码
		StockinNum       float64 `json:"stockin_num"`                  // 入库数量
		Remark           string  `json:"remark,omitempty"`             // 备注
		ApiSpecName      string  `json:"api_spec_name,omitempty"`      // 平台规格名称
		ApiGoodsName     string  `json:"api_goods_name,omitempty"`     // 平台货品名称
		Modified         int64   `json:"modified"`                     // 最后修改时间
		SuiteNo          string  `json:"suite_no,omitempty"`           // 组合装编号
		SuiteName        string  `json:"suite_name,omitempty"`         // 组合装名称
		RawRefundNos     string  `json:"raw_refund_nos,omitempty"`     // 原始退款单号
		RawRefundNo      string  `json:"raw_refund_no,omitempty"`      // 原始退款单号
		SalesTradeId     int     `json:"sales_trade_id,omitempty"`     // 订单id
		Discount         float64 `json:"discount,omitempty"`           // 总折扣金额
		Paid             float64 `json:"paid,omitempty"`               // 已支付金额
		SuiteId          int     `json:"suite_id,omitempty"`           // 组合装id
		SuiteNum         float64 `json:"suite_num,omitempty"`          // 组合装数量
		Created          string  `json:"created"`                      // 创建时间
		ModifiedDate     string  `json:"modified_date"`                // 最后修改时间
		GiftType         int     `json:"gift_type"`                    // 赠品类型
	} `json:"detail_list" gorm:"serializer:json"` // 退换单详情
	AmountDetailList []struct {
		RecId         int     `json:"rec_id"`                  // 金额明细记录id
		RefundId      int     `json:"refund_id"`               // 退换单id
		RefundType    int     `json:"refund_type"`             // 退换类型
		IsReturn      int     `json:"is_return"`               // 金额流向
		RefundAmount  float64 `json:"refund_amount,omitempty"` // 退款金额
		ReceiveAmount float64 `json:"receive_amount"`          // 收款金额
		IsGuarantee   bool    `json:"is_guarantee"`            // 是否担保支付
		AccountId     int     `json:"account_id,omitempty"`    // 支付账户
		PayAccount    string  `json:"pay_account,omitempty"`   // 买家账号
		AccountName   string  `json:"account_name,omitempty"`  // 买家开户人姓名
		AccountBank   string  `json:"account_bank,omitempty"`  // 开户银行
		IsAuto        bool    `json:"is_auto"`                 // 是否系统自动生成
		Remark        string  `json:"remark,omitempty"`        // 备注
	} `json:"amount_detail_list" gorm:"serializer:json"` // 金额明细
	SwapOrder struct {
		Tid                 string  `json:"tid,omitempty"`               // 换出订单原始单号
		ShopNo              string  `json:"shop_no,omitempty"`           // 店铺编号
		ShopName            string  `json:"shop_name,omitempty"`         // 店铺名称
		WarehouseNo         string  `json:"warehouse_no,omitempty"`      // 仓库名称
		ShopId              int     `json:"shop_id,omitempty"`           // 店铺id
		WarehouseId         int     `json:"warehouse_id,omitempty"`      // 仓库id
		SwapProvince        int     `json:"swap_province,omitempty"`     // 省份id
		SwapCity            int     `json:"swap_city,omitempty"`         // 城市id
		SwapArea            string  `json:"swap_area,omitempty"`         // 地区
		SwapDistrict        int     `json:"swap_district,omitempty"`     // 地区id
		SwapLogisticsId     int     `json:"swap_logistics_id,omitempty"` // 换货新订单物流公司id
		PostAmount          float64 `json:"post_amount,omitempty"`       // 换货邮费
		OtherAmount         float64 `json:"other_amount,omitempty"`      // 其他金额
		SwapOrderDetailList []struct {
			Oid         string  `json:"oid"`          // 原始子单号
			TargetType  int     `json:"target_type"`  // 货品类型
			TargetId    int     `json:"target_id"`    // 换出货品id
			Defect      bool    `json:"defect"`       // 是否残次品
			GoodsName   string  `json:"goods_name"`   // 货品名称
			GoodsNo     string  `json:"goods_no"`     // 货品编号
			SpecName    string  `json:"spec_name"`    // 规格名称
			SpecCode    string  `json:"spec_code"`    // 规格码
			MerchantNo  string  `json:"merchant_no"`  // 商家编码
			Price       float64 `json:"price"`        // 零售价
			TotalAmount float64 `json:"total_amount"` // 总价
			Num         float64 `json:"num"`          // 数量
			Remark      string  `json:"remark"`       // 备注
		} `json:"swap_order_detail_list,omitempty"` // 换出订单明细
	} `json:"swap_order" gorm:"serializer:json"` // 换出订单
	GovSubsidyInfo []struct {
		Tid            string `json:"tid"`              // 原始单号
		Oid            string `json:"oid"`              // 原始子单号
		CorpEntityName string `json:"corp_entity_name"` // 公司主体名称
	} `json:"gov_subsidy_info" gorm:"serializer:json"` // 国补信息
}

// Search 获取ERP销售退货（换货）订单信息
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=aftersales.refund.Refund.search
func (r *refund) Search(ctx context.Context, req *RefundSearchReq, options ...core.ReqOptionFunc) (*RefundSearchResp, error) {
	apiReq := &core.ApiReq{
		HttpMethod: http.MethodPost,
		Method:     "aftersales.refund.Refund.search",
		Body:       []any{req},
	}

	return core.Retry[RefundSearchResp](ctx, r.config, func() (*core.ApiResp, error) {
		return core.Request(ctx, apiReq, r.config, options...)
	})
}
