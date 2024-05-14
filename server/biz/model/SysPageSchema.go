package model

type SysPageSchema struct {
	BaseModel
	PageID     uint                      `json:"page_id"`
	Attributes []*SysPageSchemaAttribute `json:"attributes"`
}

type SysPageSchemaAttribute struct {
	FieldAttribute *SysPageSchemaFieldAttribute `json:"field_attribute"`
}

type SysPageSchemaFieldAttribute struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	// TODO 例如字段类型
}
