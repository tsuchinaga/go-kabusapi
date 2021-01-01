package kabus

import (
	"context"
)

// UnregisterAllRequest - 銘柄登録全解除のリクエストパラメータ
type UnregisterAllRequest struct{}

// UnregisterAllResponse - 銘柄登録全解除のレスポンス
type UnregisterAllResponse struct {
	RegistList []RegisteredSymbol `json:"RegisterList"` // 現在登録されている銘柄のリスト
}

// UnregisterAll - 銘柄登録全解除のリクエスト
func (c *restClient) UnregisterAll(token string) (*UnregisterAllResponse, error) {
	return c.UnregisterAllWithContext(context.Background(), token)
}

// UnregisterAllWithContext - 銘柄登録全解除のリクエスト(contextあり)
func (c *restClient) UnregisterAllWithContext(ctx context.Context, token string) (*UnregisterAllResponse, error) {
	code, b, err := c.put(ctx, token, "unregister/all", []byte(""))
	if err != nil {
		return nil, err
	}

	var res UnregisterAllResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
