package kabus

import (
	"context"
	"fmt"
)

// SymbolNameFutureRequest - 先物銘柄コード取得のリクエストパラメータ
type SymbolNameFutureRequest struct {
	FutureCode FutureCode // 先物コード
	DerivMonth YmNUM      // 限月
}

// SymbolNameFutureResponse - 先物銘柄コード取得のレスポンス
type SymbolNameFutureResponse struct {
	Symbol     string `json:"Symbol"`     // 銘柄コード
	SymbolName string `json:"SymbolName"` // 銘柄名称
}

// SymbolNameFuture - 先物銘柄コード取得リクエスト
func (c *restClient) SymbolNameFuture(token string, request SymbolNameFutureRequest) (*SymbolNameFutureResponse, error) {
	return c.SymbolNameFutureWithContext(context.Background(), token, request)
}

// SymbolNameFutureWithContext - 先物銘柄コード取得リクエスト(contextあり)
func (c *restClient) SymbolNameFutureWithContext(ctx context.Context, token string, request SymbolNameFutureRequest) (*SymbolNameFutureResponse, error) {
	query := fmt.Sprintf("FutureCode=%s&DerivMonth=%s", request.FutureCode, request.DerivMonth.String())
	code, b, err := c.get(ctx, token, "symbolname/future", query)
	if err != nil {
		return nil, err
	}

	var res SymbolNameFutureResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
