package controller

import (
	"IrisManager/db"
	"log"

	"github.com/kataras/iris"
)

// DeptController operate dept
//type TbUserController struct {
//	DeptDao db.TbUser
//}
//var DeptDao db.TbUser

// L.JSON(c.Writer, http.StatusOK, depts)
func List(ctx iris.Context) {
	DeptDao := new(db.TbUser)

	id, _ := ctx.PostValueInt64("id")
	user, err := DeptDao.List(id)
	if err != nil {
		log.Fatal(err)
	}
	json(ctx, user)
}
