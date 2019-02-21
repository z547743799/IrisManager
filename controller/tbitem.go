package controller

import (
	"fmt"

	"github.com/kataras/iris"
	"gitlab.com/z547743799/irismanager/models"
	"gitlab.com/z547743799/irismanager/service"
	"gitlab.com/z547743799/iriscommon/utils"
)

// DeptController operate dept
//type TbUserController struct {
//	DeptDao db.TbUser
//}
//var DeptDao db.TbUser

// L.JSON(c.Writer, http.StatusOK, depts)
type ItemController struct {
	Ctx     iris.Context
	Service service.TbItemService
}

func (c *ItemController) GetBy(id int64) {
	data := c.Service.Get(id)
	utils.Json(c.Ctx, data)
	ns := make(map[string]interface{})
	ns["ssd"] = 0
	ns["url"] = "---------"

	c.Ctx.JSON(ns)
}
func (c *ItemController) GetList() {
	pa := c.Ctx.Params().Get("page")
	fmt.Println(pa)

	page, _ := c.Ctx.URLParamInt("page")
	rows, _ := c.Ctx.URLParamInt("Rows")
	fmt.Println(rows + page)
	data := c.Service.GetList(page, rows)
	utils.Json(c.Ctx, data)

}

func (c ItemController) PostSave() {
	itme := c.Ctx.FormValues()
	ss := models.TbItem{}
	_ = c.Ctx.ReadForm(&ss)

	//	parentId, _ := c.Ctx.Po
	fmt.Println(itme)

}
