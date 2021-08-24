package router

import (
	"iris-demo/controller"
	"iris-demo/rbac"

	"github.com/kataras/iris/v12"
)

func RegisterRoute(app *iris.Application) {
	app.Get("/", controller.ShowHomePage) //Không dùng rbac có nghĩa là public method
	app.Get("/logform", controller.ShowLoginForm) //Không dùng rbac có nghĩa là public method
	app.Post("/login", controller.Login)
	app.Post("/loginjson", controller.LoginJSON)

	rbac.Get(app, "/logout", rbac.AllowAll(), controller.LogoutFromWeb)
	rbac.Get(app, "/logoutjson", rbac.AllowAll(), controller.LogoutFromREST)

	modify := app.Party("/modify")
	{
		rbac.Get(modify, "/{id:int}", rbac.Allow(rbac.ROOT, rbac.EDITOR, rbac.STAFF) ,controller.ModifyUser)
	}

	delete := app.Party("/delete")
	{
		rbac.Get(delete, "/{id:int}", rbac.Allow(rbac.ROOT, rbac.EDITOR) ,controller.DeleteUser)
	}

	addnew := app.Party("/addnew")
	{
		rbac.Get(addnew, "", rbac.Allow(rbac.ROOT, rbac.EDITOR, rbac.STAFF) ,controller.ShowAddUser)
	}
	
}
