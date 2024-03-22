package model

import (
	"ginblog/utils/errormsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在

func CheckCategory(name string) int {
	var cate Category
	SqlDb.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errormsg.ERROR_CATENAME_USED
	}
	return errormsg.SUCCESS
}

//添加分类

func CreatCategory(data *Category) int {
	err := SqlDb.Create(&data).Error
	if err != nil {
		return errormsg.ERROR //500
	}
	return errormsg.SUCCESS
}

//查询分类列表

func GetCategory(pageSize int, pageNum int) ([]Category, int) {
	var cate []Category
	var total int
	//一般涉及到用户列表的都是要进行分页的
	//limit和offset传-1为不使用限制
	err := SqlDb.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	//我们要考虑到返回err的原因，查询错误，或者为空
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

//删除分类

func DeleteCategory(id int) int {
	var cate Category
	//执行的软删除
	err := SqlDb.Where("id=?", id).Delete(&cate).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

//查询分类下的所有文章

//编辑分类信息

func EditCategory(id int, data *Category) int {
	var cate Category
	//用map传参数，如果用struct为零的字段是不会进行更新的
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	SqlDb.Model(&cate).Where("id = ?", id).Update(maps)
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
