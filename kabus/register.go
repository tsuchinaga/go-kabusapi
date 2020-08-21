package kabus

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
func NewRegisterRequester() *registerRequester {
	return &registerRequester{
		url:    "http://localhost:18080/kabusapi/register",
		method: "PUT",
	}
}

// registerRequester - 銘柄登録のリクエスタ
type registerRequester struct {
	url    string
	method string
}

// Exec - 銘柄登録リクエストの実行
func (r *registerRequester) Exec(token string, request RegisterRequest) (*RegisterResponse, error) {
	return r.ExecWithContext(context.Background(), token, request)
}

// ExecWithContext - 銘柄登録リクエストの実行(contextあり)
func (r *registerRequester) ExecWithContext(ctx context.Context, token string, request RegisterRequest) (*RegisterResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, r.method, r.url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("X-API-KEY", token)

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
		var res RegisterResponse
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
