package kabus

// TokenRequest - トークン発行のリクエストパラメータ
type TokenRequest struct {
	APIPassword string `json:"APIPassword"` // APIパスワード
}

// TokenResponse - トークン発行のレスポンス
type TokenResponse struct {
	Result int    `json:"Result"` // 結果コード
	Token  string `json:"Token"`  // APIトークン
}
