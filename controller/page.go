package controller

import (
	"github.com/kataras/iris"
)

type PageController struct {
	Ctx iris.Context
}

func (c *PageController) Get() {

	c.Ctx.View("index.html")
}

func (c *PageController) GetBy(page string) {

	c.Ctx.View(page + ".html")
}
