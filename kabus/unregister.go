package kabus

import (
	"context"
	"encoding/json"
)

// UnregisterRequest - 銘柄登録解除のリクエストパラメータ
type UnregisterRequest struct {
	Symbols []UnregisterSymbol `json:"Symbols"` // 登録解除する銘柄のリスト
}

// UnregisterSymbol - 銘柄登録解除で解除する銘柄
type UnregisterSymbol struct {
	Symbol   string   `json:"Symbol"`   // 銘柄コード
	Exchange Exchange `json:"Exchange"` // 市場コード
}

// UnregisterResponse - 銘柄登録解除のレスポンス
type UnregisterResponse struct {
	RegisterList []RegisteredSymbol `json:"RegistList"` // 現在登録されている銘柄のリスト
}

// Unregister - 銘柄登録解除リクエスト
func (c *restClient) Unregister(token string, request UnregisterRequest) (*UnregisterResponse, error) {
	return c.UnregisterWithContext(context.Background(), token, request)
}

// UnregisterWithContext - 銘柄登録解除リクエスト(contextあり)
func (c *restClient) UnregisterWithContext(ctx context.Context, token string, request UnregisterRequest) (*UnregisterResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := c.put(ctx, token, "unregister", reqBody)
	if err != nil {
		return nil, err
	}

	var res UnregisterResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
