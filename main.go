package main

import (
	R "IrisManager/controller"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	//var port string
	//flag.StringVar(&port, "port", "3001", "service listening at, default 3000")
	//flag.StringVar(&mode, "mode", "debug", "service running mode, default debug mode")

	//	flag.Parse()
	app := iris.New()
	app.Use(recover.New())

	app.Get("/:id", R.List)

	app.Run(iris.Addr(":3000"))
}
