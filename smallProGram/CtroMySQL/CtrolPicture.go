package CtroMySQL

import (
	"github.com/jinzhu/gorm"
	"log"
)

//把用户和图片绑定在一起
func BindIdAndPic(DB *gorm.DB, openid string, pictureName string, pictureUrl string) {
	sql := "insert into picture(openid,picturename,pictureurl) values (?,?,?)"
	DB.Debug().Exec(sql, openid, pictureName, pictureUrl)
}

type Picture struct {
	Picureurl string `form:"pictureurl"`
}

//获得指定用户的的照片
func GetIdPicture(DB *gorm.DB, openid string) []Picture {
	sql := "select pictureurl from picture where openid=?"
	rows, err := DB.Debug().Raw(sql, openid).Rows()
	if err != nil {
		log.Print("select pictureurl from picture failed")
	}
	AllPicture := make([]Picture, 0)
	for rows.Next() {
		var pitureurl Picture
		rows.Scan(&pitureurl.Picureurl)
		AllPicture = append(AllPicture, pitureurl)
	}
	return AllPicture
}

type CheckUrl struct {
	pictureurl string
}

//图片覆盖
func CoverPictureCheck(DB *gorm.DB, pictureurl string) string {
	var pictureUrl CheckUrl
	sql := "select pictureurl from picture where pictureurl=?"
	DB.Debug().Raw(sql, pictureurl).Scan(&pictureUrl)
	return pictureUrl.pictureurl
}

//图片删除
func DelPicture(DB *gorm.DB, openid string, pictureurl string) {
	sql := "delete from picture where pictureurl=? and openid=?"
	DB.Debug().Exec(sql, pictureurl, openid)
}
