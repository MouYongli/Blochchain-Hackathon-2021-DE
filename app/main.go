package main

import (
	"app/routes"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Logger().SetLevel("debug")

	_ = routes.RegisterRoute(app)
	app.Run(iris.Addr("localhost:8090", func(h *iris.Supervisor) {
		h.RegisterOnShutdown(func() {
			println("Server Terminated...")
		})
	}))
}