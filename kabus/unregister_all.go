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

// NewUnregisterAllRequester - 銘柄登録全解除リクエスタの生成
func NewUnregisterAllRequester(token string, isProd bool) UnregisterAllRequester {
	return &unregisterAllRequester{httpClient: httpClient{token: token, url: createURL("/unregister/all", isProd)}}
}

// UnregisterAllRequester - 銘柄登録全解除のリクエスタインターフェース
type UnregisterAllRequester interface {
	Exec() (*UnregisterAllResponse, error)
	ExecWithContext(ctx context.Context) (*UnregisterAllResponse, error)
}

// unregisterAllRequester - 銘柄登録全解除のリクエスタ
type unregisterAllRequester struct {
	httpClient
}

// Exec - 銘柄登録全解除のリクエスト実行
func (r *unregisterAllRequester) Exec() (*UnregisterAllResponse, error) {
	return r.ExecWithContext(context.Background())
}

// ExecWithContext - 銘柄登録全解除のリクエスト実行(contextあり)
func (r *unregisterAllRequester) ExecWithContext(ctx context.Context) (*UnregisterAllResponse, error) {
	code, b, err := r.httpClient.put(ctx, []byte(""))
	if err != nil {
		return nil, err
	}

	var res UnregisterAllResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
