package model

type SysRouter struct {
	BaseModel
	Name   string
	Path   string
	Title  string
	IsMenu bool
	Pid    uint
	Sort   int
}
