package controller

import (
	"IrisManager/service"
	"IrisManager/utils"

	"github.com/kataras/iris"
)

// DeptController operate dept
//type TbUserController struct {
//	DeptDao db.TbUser
//}
//var DeptDao db.TbUser

// L.JSON(c.Writer, http.StatusOK, depts)
type IndexController struct {
	Ctx     iris.Context
	Service service.TbItemService
}

func (c *IndexController) GetBy(id int64) {
	data := c.Service.Get(id)
	utils.Json(c.Ctx, data)

}
