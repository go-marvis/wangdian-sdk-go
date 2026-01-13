package model

// GovSubsidyInfo 国补信息
type GovSubsidyInfo struct {
	Tid            string `json:"tid"`              // 原始单号
	Oid            string `json:"oid"`              // 原始子单号
	CorpEntityName string `json:"corp_entity_name"` // 公司主体名称
}
