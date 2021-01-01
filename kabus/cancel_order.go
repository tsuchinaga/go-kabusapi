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

// CancelOrder - 注文取消リクエスト
func (c *restClient) CancelOrder(token string, request CancelOrderRequest) (*CancelOrderResponse, error) {
	return c.CancelOrderWithContext(context.Background(), token, request)
}

// CancelOrderWithContext - 注文取消リクエスト(contextあり)
func (c *restClient) CancelOrderWithContext(ctx context.Context, token string, request CancelOrderRequest) (*CancelOrderResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := c.put(ctx, token, "cancelorder", reqBody)
	if err != nil {
		return nil, err
	}

	var res CancelOrderResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
