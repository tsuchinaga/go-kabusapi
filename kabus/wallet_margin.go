package kabus

import (
	"context"
	"fmt"
)

// WalletMarginRequest - 取引余力（信用）のリクエストパラメータ
type WalletMarginRequest struct{}

// WalletMarginSymbolRequest - 取引余力（信用）（銘柄指定）のリクエストパラメータ
type WalletMarginSymbolRequest struct {
	Symbol   string        // 銘柄コード
	Exchange StockExchange // 市場コード
}

// WalletMarginResponse - 取引余力（信用）のレスポンス
type WalletMarginResponse struct {
	MarginAccountWallet          float64 `json:"MarginAccountWallet"`          // 信用買付可能額
	DepositkeepRate              float64 `json:"DepositkeepRate"`              // 保証金維持率
	ConsignmentDepositRate       float64 `json:"ConsignmentDepositRate"`       // 委託保証金率
	CashOfConsignmentDepositRate float64 `json:"CashOfConsignmentDepositRate"` // 現金委託保証金率
}

// WalletMargin - 取引余力（信用）リクエスト
func (c *restClient) WalletMargin(token string) (*WalletMarginResponse, error) {
	return c.WalletMarginWithContext(context.Background(), token)
}

// WalletMarginWithContext - 取引余力（信用）リクエスト(contextあり)
func (c *restClient) WalletMarginWithContext(ctx context.Context, token string) (*WalletMarginResponse, error) {
	code, b, err := c.get(ctx, token, "wallet/margin", "")
	if err != nil {
		return nil, err
	}

	var res WalletMarginResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// WalletMarginSymbol - 取引余力（信用）（銘柄指定）リクエスト
func (c *restClient) WalletMarginSymbol(token string, request WalletMarginSymbolRequest) (*WalletMarginResponse, error) {
	return c.WalletMarginSymbolWithContext(context.Background(), token, request)
}

// WalletMarginSymbolWithContext - 取引余力（信用）（銘柄指定）リクエスト(contextあり)
func (c *restClient) WalletMarginSymbolWithContext(ctx context.Context, token string, request WalletMarginSymbolRequest) (*WalletMarginResponse, error) {
	path := fmt.Sprintf("wallet/margin/%s@%d", request.Symbol, request.Exchange)
	code, b, err := c.get(ctx, token, path, "")
	if err != nil {
		return nil, err
	}

	var res WalletMarginResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
