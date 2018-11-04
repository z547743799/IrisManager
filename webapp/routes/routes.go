package routes

import (
	"IrisManager/bootstrap"
	"IrisManager/controller"
	"IrisManager/service"

	"github.com/kataras/iris/mvc"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	superstarService := service.NewTbItemService()

	item := mvc.New(b.Party("/item"))
	item.Register(superstarService)
	item.Handle(new(controller.ItemController))

	index := mvc.New(b.Party("/"))
	index.Handle(new(controller.PageController))

	//admin := mvc.New(b.Party("/admin"))
	//admin.Router.Use(middleware.BasicAuth)
	//admin.Register(superstarService)
	//admin.Handle(new(controllers.AdminController))

	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}
