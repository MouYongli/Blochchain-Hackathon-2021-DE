package main

import (
	"app/config"
	"app/routes"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"strconv"
)

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

func main() {
	app := iris.New()
	app.Use(Cors)
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