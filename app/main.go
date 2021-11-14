package main

import (
	"app/config"
	"app/routes"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"strconv"
)

func main() {
	app := iris.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Logger().SetLevel("debug")

	_conf := config.GetConfigInstance()
	_ = routes.RegisterRoute(app)
	app.Run(iris.Addr(_conf.Server.Host + ":" + strconv.Itoa(_conf.Server.Port), func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			println("Server Terminated...")
		})
	}))
}