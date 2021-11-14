package routes

import (
	"app/controllers"
	"github.com/kataras/iris/v12"
)

func RegisterRoute(app *iris.Application) error {
	//app.HandleDir("/", "views")
	//app.HandleDir("/static", "views/static")
	//app.RegisterView(iris.HTML("views", ".html"))

	index := app.Party("/")
	{
		_ctrl := controllers.Controllers{}
		//index.Get("/", _ctrl.Index)
		index.Get("/ai-marketplace/contract/{contractId:string}", _ctrl.Trade)
		index.Get("/ai-marketplace/model", _ctrl.GetModel)
		index.Get("/ai-marketplace/data", _ctrl.GetData)
	}

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.WriteString("404 NotFound")
	})

	return nil
}
