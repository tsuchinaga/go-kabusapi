package kabus

import (
	"context"
	"encoding/json"
)

// RegisterRequest - 銘柄登録のリクエストパラメータ
type RegisterRequest struct {
	Symbols []RegistSymbol `json:"Symbols"` // 登録する銘柄のリスト
}

// RegistSymbol - 銘柄登録で登録する銘柄
type RegistSymbol struct {
	Symbol   string   `json:"Symbol"`   // 銘柄コード
	Exchange Exchange `json:"Exchange"` // 市場コード
}

// RegisterResponse - 銘柄登録のレスポンス
type RegisterResponse struct {
	RegistList []RegisteredSymbol `json:"RegistList"` // 現在登録されている銘柄のリスト
}

// RegisteredSymbol - 銘柄登録によって登録された銘柄
type RegisteredSymbol struct {
	Symbol   string   `json:"Symbol"`   // 銘柄コード
	Exchange Exchange `json:"Exchange"` // 市場コード
}

// NewRegisterRequester - 銘柄登録のリクエスタの生成
func NewRegisterRequester(token string) *registerRequester {
	return &registerRequester{client: client{url: "http://localhost:18080/kabusapi/register", token: token}}
}

// registerRequester - 銘柄登録のリクエスタ
type registerRequester struct {
	client
}

// Exec - 銘柄登録リクエストの実行
func (r *registerRequester) Exec(request RegisterRequest) (*RegisterResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 銘柄登録リクエストの実行(contextあり)
func (r *registerRequester) ExecWithContext(ctx context.Context, request RegisterRequest) (*RegisterResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := r.client.put(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	var res RegisterResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
