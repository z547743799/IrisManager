package controller

import (
	"github.com/kataras/iris"
	"gitlab.com/z547743799/irismanager/service"
	"gitlab.com/z547743799/iriscommon/utils"
)

type TbContentCategoryController struct {
	Ctx     iris.Context
	Service service.TbContentCategoryService
}

func (c TbContentCategoryController) GetList() {
	//	parentId := c.Ctx.PostValueInt64Default("id", 0)
	id := c.Ctx.URLParamInt64Default("id", 0)

	//	parentId, _ := c.Ctx.Po

	Node := c.Service.GetList(id)

	utils.Json(c.Ctx, Node)

}
