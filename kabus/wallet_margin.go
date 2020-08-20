package kabus

// WalletMarginRequest - 取引余力（信用）のリクエストパラメータ
type WalletMarginRequest struct{}

// WalletMarginSymbolRequest - 取引余力（信用）（銘柄指定）のリクエストパラメータ
type WalletMarginSymbolRequest struct {
	Symbol   string   // 銘柄コード
	Exchange Exchange // 市場コード
}

// WalletMarginResponse - 取引余力（信用）のレスポンス
type WalletMarginResponse struct {
	MarginAccountWallet          float64 `json:"MarginAccountWallet"`          // 信用買付可能額
	DepositkeepRate              float64 `json:"DepositkeepRate"`              // 保証金維持率
	ConsignmentDepositRate       float64 `json:"ConsignmentDepositRate"`       // 委託保証金率
	CashOfConsignmentDepositRate float64 `json:"CashOfConsignmentDepositRate"` // 現金委託保証金率
}
