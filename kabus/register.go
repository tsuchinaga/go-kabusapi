package kabus

// RegisterRequest - 銘柄登録のリクエストパラメータ
type RegisterRequest struct {
	Symbols []RegistSymbol `json:"Symbols"` // 登録する銘柄のリスト
}

// RegistSymbol - 銘柄登録で登録する銘柄
type RegistSymbol struct {
	Symbol   string   `json:"Symbol"`   // 銘柄コード
	Exchange Exchange `json:"Exchange"` // 市場コード
}

// RegisterResponse - 銘柄登録のレスポンス
type RegisterResponse struct {
	RegistList []RegisteredSymbol `json:"RegistList"` // 現在登録されている銘柄のリスト
}

// RegisteredSymbol - 銘柄登録によって登録された銘柄
type RegisteredSymbol struct {
	Symbol   string   `json:"Symbol"`   // 銘柄コード
	Exchange Exchange `json:"Exchange"` // 市場コード
}
