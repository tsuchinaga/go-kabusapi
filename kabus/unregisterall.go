package kabus

// UnregisterAllRequest - 銘柄登録全解除のリクエストパラメータ
type UnregisterAllRequest struct{}

// UnregisterAllResponse - 銘柄登録全解除のレスポンス
type UnregisterAllResponse struct {
	RegistList []RegisteredSymbol `json:"RegistList"` // 現在登録されている銘柄のリスト
}
