package controller

import (
	"iris-demo/repo"
	"os"
	"path/filepath"

	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)

func GetUserInfo(ctx iris.Context) {
	id := ctx.Params().Get("id")
	userData, err := repo.QueryById(id)
	if err != nil {
		logger.Log(ctx, eris.Warning("User not found").NotFound())
	}
	ctx.ViewData("data", userData)
	_ = ctx.View("modify")
}

func ModifyUser(ctx iris.Context) {
	var user repo.MyUser
	if err := ctx.ReadForm(&user); err != nil && !iris.IsErrPath(err) {
		logger.Log(ctx, eris.NewFrom(err).BadRequest())
	}
	repo.ModifyById(user)
	ctx.Values().Set("id", user.ID)
	ctx.Next()
}

/*
POST /upload
Viết hàm upload ảnh vào đây
*/

func UploadPhoto(ctx iris.Context) {
	id := ctx.Values().Get("id").(string)
	uploadfiles, _, err := ctx.UploadFormFiles("./static")
	if err != nil {
		logger.Log(ctx, eris.NewFrom(err))
	}
	if len(uploadfiles) == 0 {
		ctx.Redirect("/")
		return
	}
	oldname := ""
	for _, upload := range uploadfiles {
		oldname = "static/" + upload.Filename
	}
	newname := "static/" + id + filepath.Ext(oldname)
	err = os.Rename(oldname, newname)
	if err != nil {
		logger.Log(ctx, eris.NewFrom(err))
	}
	err = repo.WriteAvatarFileName(id, id+filepath.Ext(oldname))
	if err != nil {
		logger.Log(ctx, eris.New("Id not found"))
	}
	ctx.Redirect("/")
}
