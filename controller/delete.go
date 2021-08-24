package controller

import (
	"iris-demo/repo"

	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)




func DeleteUser(ctx iris.Context){
	id, _ := ctx.Params().GetInt("id")
	userData, err := repo.QueryById(id)
	if err != nil {
		logger.Log(ctx, eris.Warning("User not found").NotFound())
	}
_, _ = ctx.Writef("Đã xóa user: %s", userData.User)
}