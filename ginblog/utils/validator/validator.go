package validator

import (
	"fmt"
	"ginblog/utils/errormsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	//转换为中文
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	//把标签映射出去
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		lable := field.Tag.Get("lable")
		return lable
	})

	//如果不知道数据类型是需要断言断出来的，但是我们知道我们接收的是一个结构体是所以直接下结构体
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errormsg.ERROR
		}
	}
	return "", errormsg.SUCCESS
}
