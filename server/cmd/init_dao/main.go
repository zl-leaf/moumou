package main

import (
	"github.com/moumou/server/biz/model"
	gormdao "github.com/zl-leaf/gorm-dao"
)

func main() {
	g := gormdao.NewGenerator("./gen/dao")
	g.Apply(model.User{}, model.Role{}, model.Permission{}, model.RolePermission{}, model.UserRelRole{})
	err := g.Generate()
	if err != nil {
		panic(err)
	}
}
