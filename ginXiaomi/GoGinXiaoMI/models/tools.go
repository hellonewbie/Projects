package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

//时间戳间戳转换成日期
func UnixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

//日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		beego.Info(err)
		return 0
	}
	return t.Unix()
}
func GetUnix() int64 {
	return time.Now().Unix()
}
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}
func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x\n", md5.Sum(data))
}
func Hello(in string) (out string) {

	out = in + "world"
	return
}
