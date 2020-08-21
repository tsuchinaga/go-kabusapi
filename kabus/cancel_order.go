package kabus

import (
	"context"
	"encoding/json"
)

// CancelOrderRequest - 注文取消のリクエストパラメータ
type CancelOrderRequest struct {
	OrderID  string `json:"OrderId"`  // 注文番号
	Password string `json:"Password"` // 注文パスワード
}

// CancelOrderResponse - 注文取消のレスポンス
type CancelOrderResponse struct {
	Result  int    `json:"Result"`  // 結果コード 0が成功、それ以外はエラー
	OrderID string `json:"OrderId"` // 受付注文番号
}

// NewCancelOrderRequester - 注文取消リクエスタの生成
func NewCancelOrderRequester(token string, isProd bool) *cancelOrderRequester {
	return &cancelOrderRequester{httpClient{url: createURL("/cancelorder", isProd), token: token}}
}

// cancelOrderRequester - 注文取消のリクエスタ
type cancelOrderRequester struct {
	httpClient
}

// Exec - 注文取消リクエストの実行
func (r *cancelOrderRequester) Exec(request CancelOrderRequest) (*CancelOrderResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 注文取消リクエストの実行(contextあり)
func (r *cancelOrderRequester) ExecWithContext(ctx context.Context, request CancelOrderRequest) (*CancelOrderResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := r.httpClient.put(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	var res CancelOrderResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
