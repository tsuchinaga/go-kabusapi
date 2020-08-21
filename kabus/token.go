package kabus

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
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
	return &tokenRequester{url: "http://localhost:18080/kabusapi/token", method: "POST"}
}

// tokenRequester - トークン発行のリクエスタ
type tokenRequester struct {
	url    string
	method string
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

	req, err := http.NewRequestWithContext(ctx, r.method, r.url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	// リクエスト送信
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
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
		errRes.StatusCode = res.StatusCode
		errRes.Body = string(b)
		return nil, errRes
	}
}
