package model

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"sessionkey"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
