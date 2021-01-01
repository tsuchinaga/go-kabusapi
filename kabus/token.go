package kabus

import (
	"context"
	"encoding/json"
)

// TokenRequest - トークン発行のリクエストパラメータ
type TokenRequest struct {
	APIPassword string `json:"APIPassword"` // APIパスワード
}

// TokenResponse - トークン発行のレスポンス
type TokenResponse struct {
	ResultCode int    `json:"ResultCode"` // 結果コード
	Token      string `json:"Token"`      // APIトークン
}

// Token - トークン発行リクエスト
func (c *restClient) Token(request TokenRequest) (*TokenResponse, error) {
	return c.TokenWithContext(context.Background(), request)
}

// TokenWithContext - トークン発行リクエスト(contextあり)
func (c *restClient) TokenWithContext(ctx context.Context, request TokenRequest) (*TokenResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := c.post(ctx, "", "token", reqBody)
	if err != nil {
		return nil, err
	}

	var res TokenResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
