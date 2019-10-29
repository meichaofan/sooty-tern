package schema

// LoginParam 登录参数
type LoginParam struct {
	Code string `json:"code"`
}

// LoginTokenInfo
type LoginTokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpireAt    int64  `json:"expire_at"`
}

type LoginRes struct {
	LoginTokenInfo *LoginTokenInfo
	IsRegister     interface{} `json:"is_register"`
}
