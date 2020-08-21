package kabus

import (
	"context"
	"encoding/json"
)

// UnregisterRequest - 銘柄登録解除のリクエストパラメータ
type UnregisterRequest struct {
	Symbols []UnregistSymbol `json:"Symbols"` // 登録解除する銘柄のリスト
}

// UnregistSymbol - 銘柄登録解除で解除する銘柄
type UnregistSymbol struct {
	Symbol   string   `json:"Symbol"`   // 銘柄コード
	Exchange Exchange `json:"Exchange"` // 市場コード
}

// UnregisterResponse - 銘柄登録解除のレスポンス
type UnregisterResponse struct {
	RegistList []RegisteredSymbol `json:"RegistList"` // 現在登録されている銘柄のリスト
}

// NewUnregisterRequester - 銘柄登録解除リクエスタの生成
func NewUnregisterRequester(token string) *unregisterRequester {
	return &unregisterRequester{
		client{token: token, url: "http://localhost:18080/kabusapi/unregister"},
	}
}

// unregisterRequester - 銘柄登録解除のリクエスタ
type unregisterRequester struct {
	client
}

// Exec - 銘柄登録解除リクエストの実行
func (r *unregisterRequester) Exec(request UnregisterRequest) (*UnregisterResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 銘柄登録解除リクエストの実行(contextあり)
func (r *unregisterRequester) ExecWithContext(ctx context.Context, request UnregisterRequest) (*UnregisterResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := r.client.put(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	var res UnregisterResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
