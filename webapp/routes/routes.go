package routes

import (
	"github.com/kataras/iris/mvc"
	"gitlab.com/z547743799/irismanager/bootstrap"
	"gitlab.com/z547743799/irismanager/controller"
	"gitlab.com/z547743799/irismanager/service"
	searchService "gitlab.com/z547743799/irissearch/service"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {

	index := mvc.New(b.Party("/"))
	index.Handle(new(controller.PageController))

	superstarService := service.NewTbItemService()
	item := mvc.New(b.Party("/item"))
	item.Register(superstarService)
	item.Handle(new(controller.ItemController))

	itemcatService := service.NewTbItemCatService()
	itemcat := mvc.New(b.Party("/item"))
	itemcat.Register(itemcatService)
	itemcat.Handle(new(controller.ItemCatController))

	pic := mvc.New(b.Party("/pic"))
	pic.Handle(new(controller.PictureController))

	contentcategoryService := service.NewTbContentCategoryService()
	contentcatergory := mvc.New(b.Party("/content/category"))
	contentcatergory.Register(contentcategoryService)
	contentcatergory.Handle(new(controller.TbContentCategoryController))

	SearchItemService := searchService.NewTbSearchItemService()
	searcitemcat := mvc.New(b.Party("/index/item"))
	searcitemcat.Register(SearchItemService)
	searcitemcat.Handle(new(controller.TbSearchItemController))

	//admin := mvc.New(b.Party("/admin"))
	//admin.Router.Use(middleware.BasicAuth)
	//admin.Register(superstarService)
	//admin.Handle(new(controllers.AdminController))

	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}
