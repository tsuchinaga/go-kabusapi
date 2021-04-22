package kabus

import (
	"context"
	"fmt"
)

// ExchangeRequest - 為替情報のリクエストパラメータ
type ExchangeRequest struct {
	Symbol ExchangeSymbol // 通貨
}

// ExchangeResponse - 為替情報のレスポンス
type ExchangeResponse struct {
	Symbol   ExchangeSymbolDetail `json:"Symbol"`   // 通貨
	BidPrice float64              `json:"BidPrice"` // BID
	Spread   float64              `json:"Spread"`   // SP
	AskPrice float64              `json:"AskPrice"` // ASK
	Change   float64              `json:"Change"`   // 前日比
	Time     HmsString            `json:"Time"`     // 時刻 ※HH:mm:ss形式
}

// Exchange - 為替情報リクエスト
func (c *restClient) Exchange(token string, request ExchangeRequest) (*ExchangeResponse, error) {
	return c.ExchangeWithContext(context.Background(), token, request)
}

// ExchangeWithContext - 為替情報リクエスト(contextあり)
func (c *restClient) ExchangeWithContext(ctx context.Context, token string, request ExchangeRequest) (*ExchangeResponse, error) {
	path := fmt.Sprintf("exchange/%s", request.Symbol)
	code, b, err := c.get(ctx, token, path, "")
	if err != nil {
		return nil, err
	}

	var res ExchangeResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
