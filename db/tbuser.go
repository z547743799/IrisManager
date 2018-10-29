package db

import (
	"IrisManager/models"
	"log"
)

// DeptDao operate dept db
type TbUser struct{}

// List query all dept
func (TbUser) List(id int64) (*models.TbUser, error) {

	tbuser := new(models.TbUser)
	_, err := x.Table("admin_panel").Id(id).Get(&tbuser)
	if err != nil {
		log.Fatal(err)
	}
	return tbuser, nil
}

//func (a *Admin_Panel) GetById(Id int64) *Admin_Panel {
//
//	admin_panel := new(Admin_Panel)
//	has, err := DB_Read.Table("admin_panel").Id(Id).Get(admin_panel)
//
//	if err != nil {
//		revel.WARN.Println(has)
//		revel.WARN.Printf("错误: %v", err)
//	}
//
//	return admin_panel
//}
