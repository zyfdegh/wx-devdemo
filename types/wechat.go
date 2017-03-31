package types

// WechatErrResp is the error response from wechat server
type WechatErrResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
