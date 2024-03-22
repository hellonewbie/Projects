package model

import (
	"ginblog/utils/errormsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 添加文章

func CreateArt(data *Article) int {
	err := SqlDb.Create(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

//todo 查询分类下的所有文章

func GetCateArt(cid int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArtList []Article
	var total int
	err := SqlDb.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", cid).Find(&cateArtList).Count(&total).Error
	if err != nil {
		return nil, errormsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errormsg.SUCCESS, total

}

//todo 查询单个文章

func GetArtInfo(id int) (Article, int) {
	var art Article
	err := SqlDb.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errormsg.ERROR_ART_EXIST
	}
	return art, errormsg.SUCCESS
}

//查询文章列表

func GetArt(pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var total int
	//limit和offset传-1为不使用限制
	err := SqlDb.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Count(&total).Error
	//我们要考虑到返回err的原因，查询错误，或者为空
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errormsg.ERROR, 0
	}
	return articleList, errormsg.SUCCESS, total
}

//编辑文章

func EditArt(id int, data *Article) int {
	var art Article
	//用map传参数，如果用struct为零的字段是不会进行更新的
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	SqlDb.Model(&art).Where("id = ?", id).Update(maps)
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
func DeleteArt(id int) int {
	var art Article
	err = SqlDb.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
