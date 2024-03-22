package CreateUpToken

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

// 自定义返回值结构体
type MyPutRet struct {
	Key    string `json:"key"`
	Hash   string `json:"hash"`
	Fsize  int    `json:"fsize"`
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

//简单的文件上传
func UpToQiNiu(file *multipart.FileHeader) (int, string) {
	var (
		accessKey = "fXBIBlHwHbcItG4B5XGTkeub4uIfUdxJeFEUl-6t"
		secretKey = "v0aXVRCNId1lBoCmkEfjgzd6qzBpdbLBztuVrL49"
		bucket    = "1project"
		ImgUrl    = "http://www.newbiewillfly.com/"
	)
	//到时候服务器直接直接那个地址加上读取的文件名称(路径：“./ 是upload的同级文件”)
	localFile := "./qiniu/picture/" + file.Filename
	//上传自定义的Key，可以指定上传目录及文件名和后缀
	key := "image/" + file.Filename
	fmt.Println(key)
	// 使用 returnBody 自定义回复格式
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	//获取上传凭证
	upToken := putPolicy.UploadToken(mac)
	//配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return 404, "PutFile failed"
	}
	fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
	url := ImgUrl + ret.Key
	return 200, url
}

//覆盖文件上传
func CoverPicture(file *multipart.FileHeader) (int, string) {
	var (
		accessKey = "fXBIBlHwHbcItG4B5XGTkeub4uIfUdxJeFEUl-6t"
		secretKey = "v0aXVRCNId1lBoCmkEfjgzd6qzBpdbLBztuVrL49"
		bucket    = "1project"
		ImgUrl    = "http://ri75r2lh4.hn-bkt.clouddn.com/"
	)
	//到时候服务器直接直接那个地址加上读取的文件名称(路径：“./ 是upload的同级文件”)
	localFile := "./qiniu/picture/" + file.Filename
	//上传自定义的Key，可以指定上传目录及文件名和后缀
	key := "image/" + file.Filename
	fmt.Println(key)
	// 使用 returnBody 自定义回复格式
	//覆盖上传
	keyToOverwrite := "image/" + file.Filename
	putPolicy := storage.PutPolicy{
		Scope:      fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	//获取上传凭证
	upToken := putPolicy.UploadToken(mac)
	//配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return 404, "PutFile failed"
	}
	fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
	url := ImgUrl + ret.Key
	return 200, url
}
