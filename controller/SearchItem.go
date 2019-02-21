package controller

import (
	"github.com/kataras/iris"
	"gitlab.com/z547743799/iriscommon/utils"
	"gitlab.com/z547743799/irissearch/service"
)

type TbSearchItemController struct {
	Ctx     iris.Context
	Service service.TbSearchItemService
}

func (c TbSearchItemController) PostImport() {
	E3Result := c.Service.ImportAllItems()

	utils.Json(c.Ctx, E3Result)
}
