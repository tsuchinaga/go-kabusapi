package kabus

// WalletCashRequest - 取引余力（現物）のリクエストパラメータ
type WalletCashRequest struct{}

// WalletCashSymbolRequest - 取引余力（現物）（銘柄指定）のリクエストパラメータ
type WalletCashSymbolRequest struct {
	Symbol   string   // 銘柄コード
	Exchange Exchange // 市場コード
}

// WalletCashResponse - 取引余力（現物）のレスポンス
type WalletCashResponse struct {
	StockAccountWallet float64 `json:"StockAccountWallet"` // 現物買付可能額
}
