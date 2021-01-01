package kabus

import (
	"context"
	"fmt"
)

// WalletOptionRequest - 取引余力（オプション）のリクエストパラメータ
type WalletOptionRequest struct{}

// WalletOptionSymbolRequest - 取引余力（オプション）（銘柄指定）のリクエストパラメータ
type WalletOptionSymbolRequest struct {
	Symbol   string         // 銘柄コード
	Exchange FutureExchange // 市場コード
}

// WalletOptionResponse - 取引余力（オプション）のレスポンス
type WalletOptionResponse struct {
	OptionBuyTradeLimit  float64 `json:"OptionBuyTradeLimit"`  // 買新規建玉可能額
	OptionSellTradeLimit float64 `json:"OptionSellTradeLimit"` // 売新規建玉可能額
	MarginRequirement    float64 `json:"MarginRequirement"`    // 必要証拠金額
}

// WalletOption - 取引余力（オプション）リクエスト
func (c *restClient) WalletOption(token string) (*WalletOptionResponse, error) {
	return c.WalletOptionWithContext(context.Background(), token)
}

// WalletOptionWithContext - 取引余力（オプション）リクエスト(contextあり)
func (c *restClient) WalletOptionWithContext(ctx context.Context, token string) (*WalletOptionResponse, error) {
	code, b, err := c.get(ctx, token, "wallet/option", "")
	if err != nil {
		return nil, err
	}

	var res WalletOptionResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// WalletOptionSymbol - 取引余力（オプション）（銘柄指定）リクエスト
func (c *restClient) WalletOptionSymbol(token string, request WalletOptionSymbolRequest) (*WalletOptionResponse, error) {
	return c.WalletOptionSymbolWithContext(context.Background(), token, request)
}

// WalletOptionSymbolWithContext - 取引余力（オプション）（銘柄指定）リクエスト(contextあり)
func (c *restClient) WalletOptionSymbolWithContext(ctx context.Context, token string, request WalletOptionSymbolRequest) (*WalletOptionResponse, error) {
	path := fmt.Sprintf("wallet/option/%s@%d", request.Symbol, request.Exchange)
	code, b, err := c.get(ctx, token, path, "")
	if err != nil {
		return nil, err
	}

	var res WalletOptionResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
