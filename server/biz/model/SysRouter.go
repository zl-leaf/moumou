package model

type SysRouter struct {
	BaseModel
	Name      string `json:"name"`
	Path      string `json:"path"`
	Title     string `json:"title"`
	IsMenu    bool   `json:"is_menu"`
	Pid       uint   `json:"pid"`
	Sort      int    `json:"sort"`
	Component string `json:"component"`
}
