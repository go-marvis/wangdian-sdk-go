package setting

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type warehouse struct {
	config *core.Config
}

type QueryWarehouseReq struct {
	WarehouseNo   string `json:"warehouse_no,omitempty"`   // 仓库编号
	WarehouseName string `json:"warehouse_name,omitempty"` // 仓库名称
	Type          int    `json:"type,omitempty"`           // 类型
	SubType       int    `json:"sub_type,omitempty"`       // 子类型
	StartTime     string `json:"start_time,omitempty"`     // 开始时间
	EndTime       string `json:"end_time,omitempty"`       // 结束时间
	HideDelete    int    `json:"hide_delete"`              // 是否隐藏已停用数据
}

type QueryWarehouseResp struct {
	Details    []WarehouseDetail `json:"details"`
	TotalCount int64             `json:"total_count"`
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

func (w warehouse) QueryWarehouse(ctx context.Context, req QueryWarehouseReq, options ...core.ReqOptionFunc) (*QueryWarehouseResp, error) {
	apiReq := &core.ApiReq{
		HttpMethod: http.MethodPost,
		Method:     "setting.Warehouse.queryWarehouse",
		Body:       []any{req},
	}
	return core.Retry[QueryWarehouseResp](ctx, w.config, func() (*core.ApiResp, error) {
		return core.Request(ctx, apiReq, w.config, options...)
	})
}
