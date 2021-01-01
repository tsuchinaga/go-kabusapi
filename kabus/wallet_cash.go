package kabus

import (
	"context"
	"fmt"
)

// WalletCashRequest - 取引余力（現物）のリクエストパラメータ
type WalletCashRequest struct{}

// WalletCashSymbolRequest - 取引余力（現物）（銘柄指定）のリクエストパラメータ
type WalletCashSymbolRequest struct {
	Symbol   string        // 銘柄コード
	Exchange StockExchange // 市場コード
}

// WalletCashResponse - 取引余力（現物）のレスポンス
type WalletCashResponse struct {
	StockAccountWallet float64 `json:"StockAccountWallet"` // 現物買付可能額
}

// WalletCash - 取引余力（現物）リクエスト
func (c *restClient) WalletCash(token string) (*WalletCashResponse, error) {
	return c.WalletCashWithContext(context.Background(), token)
}

// WalletCashWithContext - 取引余力（現物）リクエスト(contextあり)
func (c *restClient) WalletCashWithContext(ctx context.Context, token string) (*WalletCashResponse, error) {
	code, b, err := c.get(ctx, token, "wallet/cash", "")
	if err != nil {
		return nil, err
	}

	var res WalletCashResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// WalletCashSymbol - 取引余力（現物）（銘柄指定）リクエスト
func (c *restClient) WalletCashSymbol(token string, request WalletCashSymbolRequest) (*WalletCashResponse, error) {
	return c.WalletCashSymbolWithContext(context.Background(), token, request)
}

// WalletCashSymbolWithContext - 取引余力（現物）（銘柄指定）リクエスト(contextあり)
func (c *restClient) WalletCashSymbolWithContext(ctx context.Context, token string, request WalletCashSymbolRequest) (*WalletCashResponse, error) {
	path := fmt.Sprintf("wallet/cash/%s@%d", request.Symbol, request.Exchange)
	code, b, err := c.get(ctx, token, path, "")
	if err != nil {
		return nil, err
	}

	var res WalletCashResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
