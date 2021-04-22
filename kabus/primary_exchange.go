package kabus

import (
	"context"
	"fmt"
)

// PrimaryExchangeRequest - 優先市場のリクエストパラメータ
type PrimaryExchangeRequest struct {
	Symbol string // 銘柄コード
}

// PrimaryExchangeResponse - 優先市場のレスポンス
type PrimaryExchangeResponse struct {
	Symbol          string        `json:"Symbol"`          // 銘柄
	PrimaryExchange StockExchange `json:"PrimaryExchange"` // 優先市場
}

// PrimaryExchange - 優先市場リクエスト
func (c *restClient) PrimaryExchange(token string, request PrimaryExchangeRequest) (*PrimaryExchangeResponse, error) {
	return c.PrimaryExchangeWithContext(context.Background(), token, request)
}

// PrimaryExchangeWithContext - 優先市場リクエスト(contextあり)
func (c *restClient) PrimaryExchangeWithContext(ctx context.Context, token string, request PrimaryExchangeRequest) (*PrimaryExchangeResponse, error) {
	path := fmt.Sprintf("primaryexchange/%s", request.Symbol)
	code, b, err := c.get(ctx, token, path, "")
	if err != nil {
		return nil, err
	}

	var res PrimaryExchangeResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
