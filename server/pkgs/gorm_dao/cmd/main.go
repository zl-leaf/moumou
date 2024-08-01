package main

import (
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/pkgs/gorm_dao"
)

func main() {
	g := gorm_dao.NewGenerator("./biz/dao")
	g.Apply(model.User{}, model.Router{})
	err := g.Generate()
	if err != nil {
		panic(err)
	}
}
