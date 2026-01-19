package model

// PositionDetail 货位明细
type PositionDetail struct {
	RecId              int     `json:"rec_id"`                // 货位明细id
	StockoutDetailId   int     `json:"stockout_detail_id"`    // 出库单明细id
	PositionId         int     `json:"position_id,omitempty"` // 货位id
	PositionNo         string  `json:"position_no,omitempty"` // 货位编号
	ExpireDate         string  `json:"expire_date"`           // 有效期
	ProductionDate     string  `json:"production_date"`       // 生产日期
	BatchNo            string  `json:"batch_no"`              // 批次号
	PositionGoodsCount float64 `json:"position_goods_count"`  // 当前货位出库总货品数量
}
