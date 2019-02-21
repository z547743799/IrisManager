package controller

import (
	"github.com/kataras/iris"
	"gitlab.com/z547743799/irismanager/service"
	"gitlab.com/z547743799/iriscommon/utils"
)

type ItemCatController struct {
	Ctx     iris.Context
	Service service.TbItemCatService
}

func (c ItemCatController) PostCatList() {
	parentId := c.Ctx.PostValueInt64Default("id", 0)

	//	parentId, _ := c.Ctx.Po

	Node := c.Service.GetList(parentId)

	utils.Json(c.Ctx, Node)

}
