package model

type Page struct {
	BaseModel
	Title  string      `json:"title"`
	Schema *PageSchema `json:"schema"`
}
