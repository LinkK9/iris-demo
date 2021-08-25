package controller

import (
	"iris-demo/pmodel"
	"iris-demo/repo"
	"os"
	"strconv"

	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)




func GetUserInfo(ctx iris.Context){
	id, _ := ctx.Params().GetInt("id")
	userData, err := repo.QueryById(id)
	if err != nil {
		logger.Log(ctx, eris.Warning("User not found").NotFound())
	}
	ctx.ViewData("data", userData)
	_ = ctx.View("modify")
}

func ModifyUser(ctx iris.Context){
	var user pmodel.User
if err := ctx.ReadJSON(&user); err != nil {
	logger.Log(ctx, eris.NewFrom(err).BadRequest())
}
	repo.ModifyById(user)
}

/*
POST /upload
Viết hàm upload ảnh vào đây
*/
// func UploadPhoto(ctx iris.Context) {
// 	// id, _ := ctx.Params().GetInt("id")
// 	uploadfiles, _, err := ctx.UploadFormFiles("./uploads")
// 	if err != nil {
// 		logger.Log(ctx, eris.NewFrom(err))
// 	}
// 	var filenames string
// 	for _, upload := range uploadfiles {
// 		filenames =  upload.Filename 
// 	}
// 	_, _ = ctx.WriteString(filenames)
// }

func UploadPhoto(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	// user, _ := repo.QueryById(id)
	uploadfiles, _, err := ctx.UploadFormFiles("./uploads")
	if err != nil {
		logger.Log(ctx, eris.NewFrom(err))
	}
	oldname := ""
	newname := "uploads/" + strconv.Itoa(id) + ".jpg"
	for _, upload := range uploadfiles {
		oldname = "uploads/" + upload.Filename
	}
	_, _ = ctx.WriteString(newname)
	e := os.Rename(oldname, newname)
	if e != nil {
		logger.Log(ctx, eris.NewFrom(err))
	}
}