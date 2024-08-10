package main

import (
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/pkgs/gorm_dao"
)

func main() {
	g := gorm_dao.NewGenerator("./gen/dao")
	g.Apply(model.User{}, model.Router{}, model.Role{}, model.Permission{})
	err := g.Generate()
	if err != nil {
		panic(err)
	}
}
