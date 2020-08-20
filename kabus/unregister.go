package kabus

// UnregisterRequest - 銘柄登録解除のリクエストパラメータ
type UnregisterRequest struct {
	Symbols []UnregistSymbol `json:"Symbols "` // 登録解除する銘柄のリスト
}

// UnregistSymbol - 銘柄登録解除で解除する銘柄
type UnregistSymbol struct {
	Symbol   string   `json:"Symbol"`   // 銘柄コード
	Exchange Exchange `json:"Exchange"` // 市場コード
}

// UnregisterResponse - 銘柄登録解除のレスポンス
type UnregisterResponse struct {
	RegistList []RegisteredSymbol `json:"RegistList"` // 現在登録されている銘柄のリスト
}
