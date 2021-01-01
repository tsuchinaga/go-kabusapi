package kabus

import (
	"context"
	"fmt"
)

// SymbolNameOptionRequest - オプション銘柄コード取得のリクエストパラメータ
type SymbolNameOptionRequest struct {
	DerivMonth  YmNUM     // 限月
	PutOrCall   PutOrCall // コール or プット
	StrikePrice int       // 権利行使価格
}

// SymbolNameOptionResponse - オプション銘柄コード取得のレスポンス
type SymbolNameOptionResponse struct {
	Symbol     string `json:"Symbol"`     // 銘柄コード
	SymbolName string `json:"SymbolName"` // 銘柄名称
}

// SymbolNameOption - オプション銘柄コード取得リクエスト
func (c *restClient) SymbolNameOption(token string, request SymbolNameOptionRequest) (*SymbolNameOptionResponse, error) {
	return c.SymbolNameOptionWithContext(context.Background(), token, request)
}

// SymbolNameOptionWithContext - オプション銘柄コード取得リクエスト(contextあり)
func (c *restClient) SymbolNameOptionWithContext(ctx context.Context, token string, request SymbolNameOptionRequest) (*SymbolNameOptionResponse, error) {
	query := fmt.Sprintf("DerivMonth=%s&PutOrCall=%s&StrikePrice=%d", request.DerivMonth.String(), request.PutOrCall, request.StrikePrice)
	code, b, err := c.get(ctx, token, "symbolname/option", query)
	if err != nil {
		return nil, err
	}

	var res SymbolNameOptionResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
