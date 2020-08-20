package kabus

// CancelOrderRequest - 注文取消のリクエストパラメータ
type CancelOrderRequest struct {
	OrderID  string `json:"OrderId"`  // 注文番号
	Password string `json:"Password"` // 注文パスワード
}

// CancelOrderResponse - 注文取消のレスポンス
type CancelOrderResponse struct {
	Result  int    `json:"Result"`  // 結果コード 0が成功、それ以外はエラー
	OrderID string `json:"OrderId"` // 受付注文番号
}
