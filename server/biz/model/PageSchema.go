package model

type PageSchema struct {
	BaseModel
	PageID     uint                   `json:"page_id"`
	Attributes []*PageSchemaAttribute `json:"attributes"`
}

type PageSchemaAttribute struct {
	FieldAttribute *PageSchemaFieldAttribute `json:"field_attribute"`
}

type PageSchemaFieldAttribute struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	// TODO 例如字段类型
}
