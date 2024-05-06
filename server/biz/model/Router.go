package model

type Router struct {
	BaseModel
	Name   string
	Path   string
	Title  string
	IsMenu bool
	Pid    uint
	Sort   int
}
