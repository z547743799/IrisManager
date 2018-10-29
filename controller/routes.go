package controller

import (
	"github.com/kataras/iris"
)

func json(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(data)
}
