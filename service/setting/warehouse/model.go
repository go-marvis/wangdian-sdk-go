package warehouse

import (
	"fmt"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

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
