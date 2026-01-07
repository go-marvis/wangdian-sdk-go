package outer

import (
	"context"
	"net/http"

	"github.com/go-marvis/wangdian-sdk-go/core"
)

type outerIn struct {
	config *core.Config
}

type QueryOuterInReq struct {
	OuterInNo    string `json:"outer_in_no"`              // 入库单号
	WarehouseNo  string `json:"warehouse_no,omitempty"`   // 仓库编号
	StartTime    string `json:"start_time"`               // 开始时间
	EndTime      string `json:"end_time"`                 // 结束时间
	TimeType     int    `json:"time_type,omitempty"`      // 时间类型 1为建单时间，2为修改时间，默认为建单时间
	LogisticsNo  string `json:"logistics_no,omitempty"`   // 物流单号
	SrcOrderType string `json:"src_order_type,omitempty"` // 入库类型
	Status       string `json:"status,omitempty"`         // 状态
}

type QueryOuterInResp struct {
	Order      []*OuterInOrder `json:"order"`
	TotalCount int64           `json:"total_count"`
}

type OuterInOrder struct {
	RecId          int     `json:"rec_id"`           // 调整入库单id
	OuterInNo      string  `json:"outer_in_no"`      // 入库单号
	WarehouseName  string  `json:"warehouse_name"`   // 仓库名称
	WarehouseNo    string  `json:"warehouse_no"`     // 仓库编号
	LogisticsName  string  `json:"logistics_name"`   // 物流公司
	LogisticsNo    string  `json:"logistics_no"`     // 物流单号
	GoodsCount     float64 `json:"goods_count"`      // 货品数量
	GoodsTypeCount int     `json:"goods_type_count"` // 货品种类数
	Status         int     `json:"status"`           // 状态
	Reason         string  `json:"reason"`           // 入库原因
	CreatorName    string  `json:"creator_name"`     // 制单人
	SrcOrderType   int     `json:"src_order_type"`   // 入库类型
	SrcOrderNo     string  `json:"src_order_no"`     // 源单号
	Remark         string  `json:"remark"`           // 备注
	NoteCount      int     `json:"note_count"`       // 便签数量
	Created        string  `json:"created"`          // 创建时间
	Modified       string  `json:"modified"`         // 最后修改时间
	DetailList     []struct {
		RecId        int     `json:"rec_id"`              // 调整入库单明细id
		SpecNo       string  `json:"spec_no"`             // 商家编码
		GoodsNo      string  `json:"goods_no"`            // 货品编号
		GoodsName    string  `json:"goods_name"`          // 货品名称
		SpecCode     string  `json:"spec_code,omitempty"` // 规格码
		SpecName     string  `json:"spec_name"`           // 规格名称
		Barcode      string  `json:"barcode"`             // 条码
		Num          float64 `json:"num"`                 // 入库数量
		Num2         float64 `json:"num2"`                // 辅助数量
		Defect       bool    `json:"defect,omitempty"`    // 是否残次品
		Remark       string  `json:"remark,omitempty"`    // 备注
		UnitRatio    float64 `json:"unit_ratio"`          // 换算系数
		AuxUnitName  string  `json:"aux_unit_name"`       // 辅助单位
		BaseUnitName string  `json:"base_unit_name"`      // 单位
		BatchNo      string  `json:"batch_no"`            // 批次号
		ExpireDate   string  `json:"expire_date"`         // 有效期
		GoodsType    int     `json:"goods_type"`          // 货品类别
	} `json:"detail_list" gorm:"serializer:json"` // 入库单明细
}

// QueryWithDetail - 获取ERP外仓调整入库信息
// https://open.wangdian.cn/qjb/open/apidoc/doc?path=wms.outer.OuterIn.queryWithDetail
func (o *outerIn) QueryWithDetail(ctx context.Context, req *QueryOuterInReq, options ...core.ReqOptionFunc) (*QueryOuterInResp, error) {
	apiReq := &core.ApiReq{
		HttpMethod: http.MethodPost,
		Method:     "wms.outer.OuterIn.queryWithDetail",
		Body:       []any{req},
	}
	return core.Retry[QueryOuterInResp](ctx, o.config, func() (*core.ApiResp, error) {
		return core.Request(ctx, apiReq, o.config, options...)
	})
}
