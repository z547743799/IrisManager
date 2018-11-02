package utils

import (
	"github.com/kataras/iris"
)

func RenderTemplate(ctx iris.Context, tmpl string, p interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.View(tmpl+".html", p)
}

func Json(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(data)
}

func Invalid(ctx iris.Context, what int) {
	if what == 0 {
		ctx.Redirect("/404")
	}
}
