package stockout

// PositionDetail 货位明细
type PositionDetail struct {
	RecId              int     `json:"rec_id"`                // 货位明细id
	StockoutDetailId   int     `json:"stockout_detail_id"`    // 出库单明细id
	PositionId         int     `json:"position_id,omitempty"` // 货位id
	PositionNo         string  `json:"position_no,omitempty"` // 货位编号
	ExpireDate         string  `json:"expire_date"`           // 有效期
	BatchNo            string  `json:"batch_no"`              // 批次号
	PositionGoodsCount float64 `json:"position_goods_count"`  // 当前货位出库总货品数量
}

// LogisticsDetail 物流单
type LogisticsDetail struct {
	RecId         int     `json:"rec_id"`           // 物流单id
	StockoutId    int     `json:"stockout_id"`      // 出库单id
	LogisticsNo   string  `json:"logistics_no"`     // 物流单号
	CalcWeight    float64 `json:"calc_weight"`      // 估算重量
	Weight        float64 `json:"weight"`           // 称重重量
	PackageName   string  `json:"package_name"`     // 包装
	LogisticsName string  `json:"logistics_name"`   // 物流名称
	LogisticsId   int     `json:"logistics_id"`     // 物流ID
	Postage       string  `json:"postage"`          // 估算邮资
	Remark        string  `json:"remark,omitempty"` // 备注
	Length        float64 `json:"length,omitempty"` // 长
	Width         float64 `json:"width,omitempty"`  // 宽
	Height        float64 `json:"height,omitempty"` // 高
	Volume        string  `json:"volume,omitempty"` // 体积
}
