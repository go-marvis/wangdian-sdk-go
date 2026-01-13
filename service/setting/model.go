package setting

import (
	"fmt"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type QueryShopReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryShopReqBuilder() *QueryShopReqBuilder {
	return &QueryShopReqBuilder{core.NewApiReq()}
}

func (builder *QueryShopReqBuilder) PageSize(pageSize int) *QueryShopReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

func (builder *QueryShopReqBuilder) PageNo(pageNo int) *QueryShopReqBuilder {
	builder.apiReq.QueryParams.Set("page_no", fmt.Sprint(pageNo))
	return builder
}

func (builder *QueryShopReqBuilder) CalcTotal(calcTotal int) *QueryShopReqBuilder {
	builder.apiReq.QueryParams.Set("calc_total", fmt.Sprint(calcTotal))
	return builder
}

func (builder *QueryShopReqBuilder) Body(body *QueryShopBody) *QueryShopReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder *QueryShopReqBuilder) Build() *QueryShopReq {
	return &QueryShopReq{builder.apiReq}
}

type QueryShopReq struct {
	apiReq *core.ApiReq
}

type QueryShopBody struct {
	ShopNo     string `json:"shop_no,omitempty"`     // 店铺编号
	PlatformId int    `json:"platform_id,omitempty"` // 平台id
}

type QueryShopResp struct {
	*core.ApiResp
	core.CodeError
	Data *QueryShopData `json:"data"`
}

type QueryShopData struct {
	Details    []*ShopDetail `json:"details"`
	TotalCount int64         `json:"total_count"`
}

type ShopDetail struct {
	ShopId        int    `json:"shop_id"`         // 店铺id
	ShopName      string `json:"shop_name"`       // 店铺名称
	ShopNo        string `json:"shop_no"`         // 店铺编号
	PlatformId    int    `json:"platform_id"`     // 平台id
	SubPlatformId int    `json:"sub_platform_id"` // 子平台id
	Contact       string `json:"contact"`         // 联系人
	Province      string `json:"province"`        // 省份
	City          string `json:"city"`            // 城市
	District      string `json:"district"`        // 区县
	Address       string `json:"address"`         // 地址
	Telno         string `json:"telno"`           // 电话
	Mobile        string `json:"mobile"`          // 手机
	Zip           string `json:"zip"`             // 邮编
	Email         string `json:"email"`           // Email
	Remark        string `json:"remark"`          // 备注
	Website       string `json:"website"`         // 网址
	GroupId       string `json:"group_id"`        // 分组
	AccountId     string `json:"account_id"`      // 平台的主键
	IsDisabled    bool   `json:"is_disabled"`     // 停用
	AuthState     int    `json:"auth_state"`      // 授权状态
	AuthTime      string `json:"auth_time"`       // 授权时间
	ReExpireTime  string `json:"re_expire_time"`  // refresh_token过期时间
	Modified      string `json:"modified"`        // 修改时间
	ExpireTime    string `json:"expire_time"`     // session过期时间
	Created       string `json:"created"`         // 创建时间
}

type QueryWarehouseReqBuilder struct {
	apiReq *core.ApiReq
}

func NewQueryWarehouseReqBuilder() *QueryWarehouseReqBuilder {
	return &QueryWarehouseReqBuilder{core.NewApiReq()}
}

func (builder *QueryWarehouseReqBuilder) PageSize(pageSize int) *QueryWarehouseReqBuilder {
	builder.apiReq.QueryParams.Set("page_size", fmt.Sprint(pageSize))
	return builder
}

func (builder *QueryWarehouseReqBuilder) PageNo(pageNo int) *QueryWarehouseReqBuilder {
	builder.apiReq.QueryParams.Set("page_no", fmt.Sprint(pageNo))
	return builder
}

func (builder *QueryWarehouseReqBuilder) CalcTotal(calcTotal int) *QueryWarehouseReqBuilder {
	builder.apiReq.QueryParams.Set("calc_total", fmt.Sprint(calcTotal))
	return builder
}

func (builder *QueryWarehouseReqBuilder) Body(body *QueryWarehouseBody) *QueryWarehouseReqBuilder {
	builder.apiReq.Body = body
	return builder
}

func (builder *QueryWarehouseReqBuilder) Build() *QueryWarehouseReq {
	return &QueryWarehouseReq{builder.apiReq}
}

type QueryWarehouseReq struct {
	apiReq *core.ApiReq
}

type QueryWarehouseBody struct {
	WarehouseNo   string `json:"warehouse_no,omitempty"`   // 仓库编号
	WarehouseName string `json:"warehouse_name,omitempty"` // 仓库名称
	Type          int    `json:"type,omitempty"`           // 类型
	SubType       int    `json:"sub_type,omitempty"`       // 子类型
	StartTime     string `json:"start_time,omitempty"`     // 开始时间
	EndTime       string `json:"end_time,omitempty"`       // 结束时间
	HideDelete    int    `json:"hide_delete"`              // 是否隐藏已停用数据
}

type QueryWarehouseResp struct {
	*core.ApiResp
	core.CodeError
	Data *QueryWarehouseData `json:"data"`
}

type QueryWarehouseData struct {
	Details    []*WarehouseDetail `json:"details"`
	TotalCount int64              `json:"total_count"`
}

type WarehouseDetail struct {
	WarehouseId     int    `json:"warehouse_id"`     //	仓库id
	WarehouseNo     string `json:"warehouse_no"`     //	仓库编号
	Name            string `json:"name"`             //	仓库名称
	Zip             string `json:"zip"`              //	邮编
	Address         string `json:"address"`          //	地址
	Province        string `json:"province"`         //	省份
	City            string `json:"city"`             //	城市
	District        string `json:"district"`         //	区县
	Mobile          string `json:"mobile"`           //	手机
	Remark          string `json:"remark"`           //	备注
	Type            int    `json:"type"`             //	类别
	Telno           string `json:"telno"`            //	固话
	SubType         int    `json:"sub_type"`         //	子类别
	Contact         string `json:"contact"`          //	联系人
	Modified        int64  `json:"modified"`         //	修改时间
	IsDisabled      bool   `json:"is_disabled"`      //	停用
	Created         string `json:"created"`          //	创建时间
	PurchaseContact string `json:"purchase_contact"` //	采购联系人
	PurchaseTelno   string `json:"purchase_telno"`   //	采购联系电话
	PurchaseAddress string `json:"purchase_address"` //	采购联系地址
}
