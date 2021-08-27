package controller

import (
	"iris-demo/repo"

	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)

func ShowAddUser(ctx iris.Context) {
	_ = ctx.View("addnew")
}

func AddUser(ctx iris.Context) {
	var newUser repo.MyUser
	if err := ctx.ReadForm(&newUser); err != nil && !iris.IsErrPath(err) {
		logger.Log(ctx, eris.NewFrom(err).BadRequest())
		return
	}
	repo.AddNewUser(newUser)
	id, e := repo.GetIdByEmail(newUser.Email)
	if e != nil {
		logger.Log(ctx, eris.Warning("Id not found").BadRequest())
	}
	ctx.Values().Set("id", id)
	ctx.Next()
	// ctx.Redirect("/")
}
