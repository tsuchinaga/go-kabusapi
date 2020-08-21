package kabus

import (
	"context"
	"encoding/json"
	"net/http"
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

// NewTokenRequester - トークン発行リクエスタの生成
func NewTokenRequester() *tokenRequester {
	return &tokenRequester{client: client{url: "http://localhost:18080/kabusapi/token"}}
}

// tokenRequester - トークン発行のリクエスタ
type tokenRequester struct {
	client
}

// Exec - トークン発行リクエストの実行
func (r *tokenRequester) Exec(request TokenRequest) (*TokenResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - トークン発行リクエストの実行(contextあり)
func (r *tokenRequester) ExecWithContext(ctx context.Context, request TokenRequest) (*TokenResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := r.client.post(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	if code == http.StatusOK {
		var res TokenResponse
		if err := json.Unmarshal(b, &res); err != nil {
			return nil, err
		}
		return &res, nil
	} else {
		var errRes ErrorResponse
		if err := json.Unmarshal(b, &errRes); err != nil {
			return nil, err
		}
		errRes.StatusCode = code
		errRes.Body = string(b)
		return nil, errRes
	}
}
