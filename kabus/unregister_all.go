package kabus

import (
	"context"
)

// UnregisterAllRequest - 銘柄登録全解除のリクエストパラメータ
type UnregisterAllRequest struct{}

// UnregisterAllResponse - 銘柄登録全解除のレスポンス
type UnregisterAllResponse struct {
	RegistList []RegisteredSymbol `json:"RegistList"` // 現在登録されている銘柄のリスト
}

// NewUnregisterAllRequester - 銘柄登録全解除リクエスタの生成
func NewUnregisterAllRequester(token string, isProd bool) *unregisterAllRequester {
	u := "http://localhost:18080/kabusapi/unregister/all"
	if !isProd {
		u = "http://localhost:18081/kabusapi/unregister/all"
	}
	return &unregisterAllRequester{client: client{token: token, url: u}}
}

// unregisterAllRequester - 銘柄登録全解除のリクエスタ
type unregisterAllRequester struct {
	client
}

// Exec - 銘柄登録全解除のリクエスト実行
func (r *unregisterAllRequester) Exec() (*UnregisterAllResponse, error) {
	return r.ExecWithContext(context.Background())
}

// ExecWithContext - 銘柄登録全解除のリクエスト実行(contextあり)
func (r *unregisterAllRequester) ExecWithContext(ctx context.Context) (*UnregisterAllResponse, error) {
	code, b, err := r.client.put(ctx, []byte(""))
	if err != nil {
		return nil, err
	}

	var res UnregisterAllResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
