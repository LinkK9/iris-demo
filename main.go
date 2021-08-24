package main

import (
	"iris-demo/config"
	"iris-demo/rbac"
	"iris-demo/router"
	"iris-demo/session"
	"iris-demo/template"

	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	config.ReadConfig()

	logFile := logger.Init() //Cần phải có 2 file error.html và info.html ở /views
	if logFile != nil {
		defer logFile.Close()
	}

	app.Use(session.Sess.Handler())

	rbacConfig := rbac.NewConfig()
	rbac.Init(rbacConfig) //Khởi động với cấu hình mặc định

	//đặt hàm này trên các hàm đăng ký route - controller
	app.Use(rbac.CheckRoutePermission)

	app.HandleDir("/", iris.Dir("./static"))

	router.RegisterRoute(app)

	template.InitViewEngine(app)

	//Luôn để hàm này sau tất cả lệnh cấu hình đường dẫn với RBAC
	rbac.BuildPublicRoute(app)
	_ = app.Listen(config.Config.Port)
}
