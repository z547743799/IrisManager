package service

import (
	"github.com/go-xorm/xorm"
	"gitlab.com/z547743799/iriscommon/pojo"
	"gitlab.com/z547743799/irismanager/db"
	"gitlab.com/z547743799/irismanager/models"
)

type TbContentCategoryService interface {
	GetList(parentId int64) []*pojo.EasyUITreeNode
}

type tbContentCategoryService struct {
	engine *xorm.Engine
}

func NewTbContentCategoryService() TbContentCategoryService {
	return &tbContentCategoryService{
		engine: db.X,
	}
}

func (d *tbContentCategoryService) GetList(parentId int64) []*pojo.EasyUITreeNode {
	datalist := make([]models.TbContentCategory, 0)
	err := d.engine.Where("parent_id=?", parentId).Desc("id").Find(&datalist)

	if err != nil {
		return nil
	}

	nodelist := make([]*pojo.EasyUITreeNode, 0)

	for _, v := range datalist {
		node := new(pojo.EasyUITreeNode)
		node.Id = v.Id
		node.Text = v.Name
		if v.IsParent == 0 {
			node.State = "open"
		} else {
			node.State = "closed"
		}
		nodelist = append(nodelist, node)

	}

	return nodelist

}
