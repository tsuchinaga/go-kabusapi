package kabus

import (
	"context"
	"fmt"
)

// WalletFutureRequest - 取引余力（先物）のリクエストパラメータ
type WalletFutureRequest struct{}

// WalletFutureSymbolRequest - 取引余力（先物）（銘柄指定）のリクエストパラメータ
type WalletFutureSymbolRequest struct {
	Symbol   string         // 銘柄コード
	Exchange FutureExchange // 市場コード
}

// WalletFutureResponse - 取引余力（先物）のレスポンス
type WalletFutureResponse struct {
	FutureTradeLimit  float64 `json:"FutureTradeLimit"`  // 新規建玉可能額
	MarginRequirement float64 `json:"MarginRequirement"` // 必要証拠金額
}

// WalletFuture - 取引余力（先物）リクエスト
func (c *restClient) WalletFuture(token string) (*WalletFutureResponse, error) {
	return c.WalletFutureWithContext(context.Background(), token)
}

// WalletFutureWithContext - 取引余力（先物）リクエスト(contextあり)
func (c *restClient) WalletFutureWithContext(ctx context.Context, token string) (*WalletFutureResponse, error) {
	code, b, err := c.get(ctx, token, "wallet/future", "")
	if err != nil {
		return nil, err
	}

	var res WalletFutureResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// WalletFutureSymbol - 取引余力（先物）（銘柄指定）リクエスト
func (c *restClient) WalletFutureSymbol(token string, request WalletFutureSymbolRequest) (*WalletFutureResponse, error) {
	return c.WalletFutureSymbolWithContext(context.Background(), token, request)
}

// WalletFutureSymbolWithContext - 取引余力（先物）（銘柄指定）リクエスト(contextあり)
func (c *restClient) WalletFutureSymbolWithContext(ctx context.Context, token string, request WalletFutureSymbolRequest) (*WalletFutureResponse, error) {
	path := fmt.Sprintf("wallet/future/%s@%d", request.Symbol, request.Exchange)
	code, b, err := c.get(ctx, token, path, "")
	if err != nil {
		return nil, err
	}

	var res WalletFutureResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
