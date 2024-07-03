package model

type SysUser struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"-"`
}
