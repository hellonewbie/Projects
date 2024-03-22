package test

import (
	"bytes"
	"cloud-disk/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:ljy12345...@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Print("Mysql connect failed")
	}
	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		log.Print(err)
	}
	fmt.Print(data)
	//可导出的数据编码转成json格式字符串
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	dst := new(bytes.Buffer)
	//func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
	//目的就是使其更容易的嵌套到其它格式化的json数据中
	err = json.Indent(dst, b, "", "  ")
	if err != nil {
		log.Print(err)
	}
	fmt.Println(dst)
}
