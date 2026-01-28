package goods

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

func (builder *QueryReqBuilder) Body(body *QueryReqBody) *QueryReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder *QueryReqBuilder) Build() *QueryReq {
	return &QueryReq{builder.apiReq}
}

type QueryReq struct {
	apiReq *core.ApiReq
}

type QueryReqBody struct {
	SpecNo      string `json:"spec_no,omitempty"`      // 商家编码
	GoodsNo     string `json:"goods_no,omitempty"`     // 货品编号
	BrandName   string `json:"brand_name,omitempty"`   // 品牌名称
	ClassName   string `json:"class_name,omitempty"`   // 分类名称
	Barcode     string `json:"barcode,omitempty"`      // 条码
	HideDeleted *bool  `json:"hide_deleted,omitempty"` // 是否隐藏已删除的货品。
	StartTime   string `json:"start_time,omitempty"`   // 开始时间
	EndTime     string `json:"end_time,omitempty"`     // 结束时间
}

type QueryResp struct {
	*core.ApiResp
	core.CodeError
	Data *QueryData `json:"data"`
}
type QueryData struct {
	GoodsList  []*Goods `json:"goods_list"`
	TotalCount int64    `json:"total_count"`
}

type Goods struct {
	GoodsId       int64        `json:"goods_id"`                         // 货品id
	GoodsNo       string       `json:"goods_no"`                         // 货品编号
	GoodsName     string       `json:"goods_name"`                       // 货品名称
	ShortName     string       `json:"short_name"`                       // 货品简称
	Alias         string       `json:"alias"`                            // 货品别名
	GoodsType     int          `json:"goods_type"`                       // 货品类别
	SpecCount     int          `json:"spec_count"`                       // 规格数
	BrandName     string       `json:"brand_name"`                       // 品牌名称
	BrandId       int          `json:"brand_id"`                         // 品牌id
	Remark        string       `json:"remark"`                           // 备注
	Prop1         string       `json:"prop1"`                            // 自定义属性1
	Prop2         string       `json:"prop2"`                            // 自定义属性2
	Prop3         string       `json:"prop3"`                            // 自定义属性3
	Prop4         string       `json:"prop4"`                            // 自定义属性4
	Prop5         string       `json:"prop5"`                            // 自定义属性5
	Prop6         string       `json:"prop6"`                            // 自定义属性6
	Origin        string       `json:"origin"`                           // 产地
	ClassName     string       `json:"class_name"`                       // 分类名称
	ClassId       int          `json:"class_id"`                         // 分类ID
	UnitName      string       `json:"unit_name"`                        // 基本单位
	AuxUnitName   string       `json:"aux_unit_name"`                    // 辅助单位
	FlagName      string       `json:"flag_name"`                        // 标记名称
	Deleted       int          `json:"deleted"`                          // 货品已删除
	GoodsModified int64        `json:"goods_modified"`                   // 最后修改时间
	GoodsCreated  string       `json:"goods_created"`                    // 创建时间
	Modified      string       `json:"modified"`                         // 最后修改时间
	BrandNo       string       `json:"brand_no"`                         // 品牌编号
	WashingLabel  string       `json:"washing_label"`                    // 水洗标
	Unit          int          `json:"unit"`                             // 基本单位id
	AuxUnit       int          `json:"aux_unit"`                         // 辅助单位id
	FlagId        int          `json:"flag_id"`                          // 标记id
	SpecList      []*GoodsSpec `json:"spec_list" gorm:"serializer:json"` // 单品信息详情
}

type GoodsSpec struct {
	GoodsId          int64   `json:"goods_id"`                      // 货品id
	SpecId           int64   `json:"spec_id"`                       // 单品id
	SpecNo           string  `json:"spec_no"`                       // 商家编码
	SpecCode         string  `json:"spec_code,omitempty"`           // 规格码
	Barcode          string  `json:"barcode"`                       // 主条码
	SpecName         string  `json:"spec_name,omitempty"`           // 规格名称
	LowestPrice      float64 `json:"lowest_price,omitempty"`        // 最低价
	RetailPrice      float64 `json:"retail_price,omitempty"`        // 零售价
	WholesalePrice   float64 `json:"wholesale_price,omitempty"`     // 批发价
	MemberPrice      float64 `json:"member_price,omitempty"`        // 会员价
	MarketPrice      float64 `json:"market_price,omitempty"`        // 市场价
	ValidityDays     int     `json:"validity_days,omitempty"`       // 有效期天数
	SalesDays        int     `json:"sales_days,omitempty"`          // 最佳销售天数
	ReceiveDays      int     `json:"receive_days,omitempty"`        // 最佳收货天数
	Weight           float64 `json:"weight,omitempty"`              // 重量(kg)
	Length           float64 `json:"length,omitempty"`              // 长(cm)
	Width            float64 `json:"width,omitempty"`               // 宽(cm)
	Height           float64 `json:"height,omitempty"`              // 高(cm)
	SnType           int     `json:"sn_type,omitempty"`             // 启用序列号
	IsLowerCost      bool    `json:"is_lower_cost,omitempty"`       // 允许低于成本价
	IsNotUseAir      int     `json:"is_not_use_air,omitempty"`      // 航空禁运
	WmsProcessMask   int     `json:"wms_process_mask,omitempty"`    // 仓库流程
	TaxRate          float64 `json:"tax_rate,omitempty"`            // 税率
	LargeType        int     `json:"large_type,omitempty"`          // 拆分
	GoodsLabel       string  `json:"goods_label,omitempty"`         // 货品标签
	Deleted          int     `json:"deleted,omitempty"`             // 货品已删除
	Remark           string  `json:"remark,omitempty"`              // 备注
	SpecModified     int     `json:"spec_modified"`                 // 最后修改时间
	SpecCreated      string  `json:"spec_created"`                  // 创建时间
	Prop1            string  `json:"prop1,omitempty"`               // 自定义属性1
	Prop2            string  `json:"prop2,omitempty"`               // 自定义属性2
	Prop3            string  `json:"prop3,omitempty"`               // 自定义属性3
	Prop4            string  `json:"prop4,omitempty"`               // 自定义属性4
	Prop5            string  `json:"prop5,omitempty"`               // 自定义属性5
	Prop6            string  `json:"prop6,omitempty"`               // 自定义属性6
	CustomPrice1     float64 `json:"custom_price1,omitempty"`       // 自定义价格1
	CustomPrice2     float64 `json:"custom_price2,omitempty"`       // 自定义价格2
	ImgUrl           string  `json:"img_url,omitempty"`             // 图片URL
	SpecUnitName     string  `json:"spec_unit_name"`                // 基本单位
	SpecAuxUnitName  string  `json:"spec_aux_unit_name"`            // 辅助单位
	TaxCode          string  `json:"tax_code,omitempty"`            // 税务编码
	PackScore        float64 `json:"pack_score,omitempty"`          // 打包积分
	PickScore        float64 `json:"pick_score,omitempty"`          // 拣货积分
	ScanScore        float64 `json:"scan_score,omitempty"`          // 验货积分
	SpecFlagName     string  `json:"spec_flag_name"`                // 单品标记名称
	BarcodeCount     int     `json:"barcode_count"`                 // 条码个数
	SaleScore        float64 `json:"sale_score,omitempty"`          // 销售积分
	IsNotNeedExamine int     `json:"is_not_need_examine,omitempty"` // 是否出库不验货
	SpecUnit         int     `json:"spec_unit,omitempty"`           // 基本单位id
	SpecAuxUnit      int     `json:"spec_aux_unit,omitempty"`       // 辅助单位id
	ExchangeName     string  `json:"exchange_name,omitempty"`       // 备注换货货品的换货识别符
	MaxLimitNum      float64 `json:"max_limit_num,omitempty"`       // 备注换货货品的最大限制数量
	BarcodeList      []struct {
		Barcode  string `json:"barcode"`   // 条码
		Type     int    `json:"type"`      // 类型
		IsMaster int    `json:"is_master"` // 是否主条码
		Modified string `json:"modified"`  // 最后修改时间
	} `json:"barcode_list"` // 条码信息
}
