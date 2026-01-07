package strategy

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type virtualWarehouse struct {
	config *core.Config
}

type VirtualWarehouseQueryReq struct {
	VirtualWarehouseNo string `json:"virtual_warehouse_no,omitempty"` // 虚拟仓编号
	SysWarehouseNo     string `json:"sys_warehouse_no,omitempty"`     // 实体仓编号
}

type VirtualWarehouseQueryResp struct {
	DetailList []*VirtualWarehouseQueryDetail `json:"detail_list"` // 数据
	TotalCount int64                          `json:"total_count"` // 总数
}

type VirtualWarehouseQueryDetail struct {
	VirtualWarehouseId int    `json:"virtual_warehouse_id"` // 虚拟仓id
	VirtualWarehouseNo string `json:"virtual_warehouse_no"` // 虚拟仓编号 (唯一)
	Name               string `json:"name"`                 // 虚拟仓名称
	SysWarehouseList   []struct {
		WarehouseId   int    `json:"warehouse_id"`   // 实体仓id
		WarehouseNo   string `json:"warehouse_no"`   // 实体仓编号
		WarehouseName string `json:"warehouse_name"` // 实体仓名称
	} `json:"sys_warehouse_list" gorm:"serializer:json"` // 实体仓列表
	ShopList []struct {
		ShopId   int    `json:"shop_id"`   // 店铺唯一键
		ShopName string `json:"shop_name"` // 店铺名称
		ShopNo   string `json:"shop_no"`   // 店铺编号
	} `json:"shop_list" gorm:"serializer:json"` // 店铺列表
}

// Query 获取ERP虚拟仓、店铺、实体仓关联信息
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=setting.strategy.VirtualWarehouse.query
func (v *virtualWarehouse) Query(ctx context.Context, req *VirtualWarehouseQueryReq, options ...core.ReqOptionFunc) (*VirtualWarehouseQueryResp, error) {
	apiReq := &core.ApiReq{
		HttpMethod: http.MethodPost,
		Method:     "setting.strategy.VirtualWarehouse.query",
		Body:       []any{req},
	}

	return core.Retry[VirtualWarehouseQueryResp](ctx, v.config, func() (*core.ApiResp, error) {
		return core.Request(ctx, apiReq, v.config, options...)
	})
}

type VirtualWarehouseOrderSearchReq struct {
	StartTime          string `json:"start_time"`                     // 开始时间
	EndTime            string `json:"end_time"`                       // 结束时间
	VirtualWarehouseNo string `json:"virtual_warehouse_no,omitempty"` // 虚拟仓编号
	WarehouseNo        string `json:"warehouse_no,omitempty"`         // 实体仓编号
	OrderType          int    `json:"order_type,omitempty"`           // 单据类型
}

type VirtualWarehouseOrderSearchResp struct {
	Order      []*VirtualWarehouseOrderSearchOrder `json:"order"`
	TotalCount int64                               `json:"total_count"`
}

type VirtualWarehouseOrderSearchOrder struct {
	OrderId              int     `json:"order_id"`                 // 单据id
	OrderNo              string  `json:"order_no"`                 // 单据编号 (唯一)
	OrderType            int     `json:"order_type"`               // 单据类型
	VirtualWarehouseNo   string  `json:"virtual_warehouse_no"`     // 虚拟仓仓库编号
	VirtualWarehouseId   int     `json:"virtual_warehouse_id"`     // 虚拟仓id
	ToVirtualWarehouseNo string  `json:"to_virtual_warehouse_no"`  // 目标虚拟仓编号
	ToVirtualId          int     `json:"to_virtual_id"`            // 目标虚拟仓id
	Num                  float64 `json:"num"`                      // 数量
	PreCheckTime         string  `json:"pre_check_time,omitempty"` // 预审核时间
	Status               int     `json:"status"`                   // 状态
	CheckTime            string  `json:"check_time"`               // 审核时间
	CheckedTime          int64   `json:"checked_time"`             // 审核时间
	OrderStatus          int     `json:"order_status"`             // 单据状态
	TypeNum              int     `json:"type_num"`                 // 货品种类
	TotalAmount          float64 `json:"total_amount"`             // 总金额
	Remark               string  `json:"remark"`                   // 备注
	WarehouseNames       string  `json:"warehouse_names"`          // 实体仓
	WarehouseIdStr       string  `json:"warehouse_id_str"`         // 实体仓id
	Modified             string  `json:"modified"`                 // 最后修改时间
	Created              string  `json:"created"`                  // 创建时间
	DetailList           []struct {
		RecId         int     `json:"rec_id"`                 // 明细唯一id
		OrderId       int     `json:"order_id"`               // 虚拟仓单据id
		SpecNo        string  `json:"spec_no"`                // 商家编码
		SpecCode      string  `json:"spec_code,omitempty"`    // 规格编码
		GoodsNo       string  `json:"goods_no"`               // 货品编号
		GoodsName     string  `json:"goods_name"`             // 货品名称
		WarehouseId   int     `json:"warehouse_id"`           // 实体仓id
		WarehouseNo   string  `json:"warehouse_no"`           // 实体仓编号
		WarehouseName string  `json:"warehouse_name"`         // 实体仓
		Num           float64 `json:"num,omitempty"`          // 入库数量
		ExpectedNum   float64 `json:"expected_num"`           // 预期数量
		FinalNum      float64 `json:"final_num,omitempty"`    // 实际数量
		Price         float64 `json:"price,omitempty"`        // 成本价
		PurchaseNum   float64 `json:"purchase_num,omitempty"` // 采购在途量
		FactoryNum    float64 `json:"factory_num,omitempty"`  // 工厂库存
		Created       string  `json:"created"`                // 创建时间
		Modified      string  `json:"modified"`               // 最后修改时间
	} `json:"detail_list"  gorm:"serializer:json"` // 明细信息
}

// OrderSearch 获取ERP虚拟仓单据信息
//
// 时间跨度：start_time与end_time时间跨度不超过30天
//
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=setting.strategy.VirtualWarehouse.orderSearch
func (v *virtualWarehouse) OrderSearch(ctx context.Context, req *VirtualWarehouseOrderSearchReq, options ...core.ReqOptionFunc) (*VirtualWarehouseOrderSearchResp, error) {
	apiReq := &core.ApiReq{
		HttpMethod: http.MethodPost,
		Method:     "setting.strategy.VirtualWarehouse.orderSearch",
		Body:       []any{req},
	}
	return core.Retry[VirtualWarehouseOrderSearchResp](ctx, v.config, func() (*core.ApiResp, error) {
		return core.Request(ctx, apiReq, v.config, options...)
	})
}
