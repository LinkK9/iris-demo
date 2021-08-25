package controller

import (
	"iris-demo/pmodel"

	"iris-demo/repo"

	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)



func ShowAddUser(ctx iris.Context){
	_ = ctx.View("addnew")
}

func AddUser(ctx iris.Context) {
var newUser pmodel.User
if err := ctx.ReadJSON(&newUser); err != nil {
	logger.Log(ctx, eris.NewFrom(err).BadRequest())
	return
}
repo.AddNewUser(newUser)
}