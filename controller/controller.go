package controller

import (
	"fmt"
	"iris-demo/repo"

	"iris-demo/pmodel"

	"iris-demo/session"


	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type LoginRequest struct {
	Email string
	Pass  string
}

func ShowHomePage(ctx iris.Context) {
	ctx.ViewData("infolist", repo.GetUserInfoList())
	_ = ctx.View("index")
}

func ShowLoginForm(ctx iris.Context) {
	_ = ctx.View("form")
}

/*
Login thông qua form. Dành cho ứng dụng web server side renderings
*/
func Login(ctx iris.Context) {
	var loginReq LoginRequest

	if err := ctx.ReadForm(&loginReq); err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := repo.QueryByEmail(loginReq.Email)
	if err != nil { //Không tìm thấy user
		_, _ = ctx.WriteString("Login Failed")
		return
	}

	if user.Pass != loginReq.Pass {
		_, _ = ctx.WriteString("Wrong password")
		return
	}

	session.SetAuthenticated(ctx, pmodel.AuthenInfo{
		User:  user.User,
		Email: user.Email,
		Roles: user.Roles,
	})

	//Login thành công thì quay về trang chủ
	ctx.Redirect("/")
}

/*
Login thông qua axios.post dành cho ứng dụng Vue
Request.ContentType = 'application/json'
*/
func LoginJSON(ctx iris.Context) {
	var loginReq LoginRequest

	if err := ctx.ReadJSON(&loginReq); err != nil {
		logger.Log(ctx, eris.NewFrom(err).BadRequest())
		return
	}

	user, err := repo.QueryByEmail(loginReq.Email)
	if err != nil { //Không tìm thấy user
		logger.Log(ctx, eris.Warning("User not found").UnAuthorized())
		return
	}

	if user.Pass != loginReq.Pass {
		logger.Log(ctx, eris.Warning("Wrong password").UnAuthorized())
		return
	}

	session.SetAuthenticated(ctx, pmodel.AuthenInfo{
		User:  user.User,
		Email: user.Email,
		Roles: user.Roles,
	})

	sess := sessions.Get(ctx)
	sess.Set(session.SESS_AUTH, true)
	sess.Set(session.SESS_USER, pmodel.AuthenInfo{
		User:  user.User,
		Email: user.Email,
		Roles: user.Roles,
	})
	//Login thành công thì quay về trang chủ
	_, _ = ctx.JSON("Login successfully")
}

func LogoutFromWeb(ctx iris.Context) {
	logout(ctx)
	ctx.Redirect("/")
}

func LogoutFromREST(ctx iris.Context) {
	logout(ctx)
	_, _ = ctx.JSON("Logout success")
}

func logout(ctx iris.Context) {
	/*	if !session.IsLogin(ctx) {
		logger.Log(ctx, eris.Warning("Bạn chưa login").UnAuthorized())
	}*/
	//Xoá toàn bộ session và xoá luôn cả Cookie sessionid ở máy người dùng
	session.Sess.Destroy(ctx)
}

