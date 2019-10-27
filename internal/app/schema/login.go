package schema

// LoginParam 登录参数
type LoginParam struct {
	NickName string `json:"nickname"`
	Password string `json:"password"`
}

// LoginTokenInfo 登录令牌信息
type LoginTokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpireAt    int64  `json:"expire_at"`
}
