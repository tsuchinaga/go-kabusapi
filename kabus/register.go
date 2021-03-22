package kabus

import (
	"context"
	"encoding/json"
)

// RegisterRequest - 銘柄登録のリクエストパラメータ
type RegisterRequest struct {
	Symbols []RegisterSymbol `json:"Symbols"` // 登録する銘柄のリスト
}

// RegisterSymbol - 銘柄登録で登録する銘柄
type RegisterSymbol struct {
	Symbol   string   `json:"Symbol"`   // 銘柄コード
	Exchange Exchange `json:"Exchange"` // 市場コード
}

// RegisterResponse - 銘柄登録のレスポンス
type RegisterResponse struct {
	RegisterList []RegisteredSymbol `json:"RegistList"` // 現在登録されている銘柄のリスト
}

// RegisteredSymbol - 銘柄登録によって登録された銘柄
type RegisteredSymbol struct {
	Symbol   string   `json:"Symbol"`   // 銘柄コード
	Exchange Exchange `json:"Exchange"` // 市場コード
}

// Register - 銘柄登録リクエスト
func (c *restClient) Register(token string, request RegisterRequest) (*RegisterResponse, error) {
	return c.RegisterWithContext(context.Background(), token, request)
}

// RegisterWithContext - 銘柄登録リクエスト(contextあり)
func (c *restClient) RegisterWithContext(ctx context.Context, token string, request RegisterRequest) (*RegisterResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := c.put(ctx, token, "register", reqBody)
	if err != nil {
		return nil, err
	}

	var res RegisterResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
