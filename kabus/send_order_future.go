package kabus

import (
	"context"
	"encoding/json"
)

// SendOrderFutureRequest - 注文発注(先物)のリクエストパラメータ
type SendOrderFutureRequest struct {
	Password           string              `json:"Password"`           // 注文パスワード
	Symbol             string              `json:"Symbol"`             // 銘柄コード
	Exchange           FutureExchange      `json:"Exchange"`           // 先物市場コード
	TradeType          TradeType           `json:"TradeType"`          // 取引区分
	TimeInForce        TimeInForce         `json:"TimeInForce"`        // 有効期間条件
	Side               Side                `json:"Side"`               // 売買区分
	Qty                int                 `json:"Qty"`                // 注文数量
	ClosePositionOrder ClosePositionOrder  `json:"ClosePositionOrder"` // 決済順序
	ClosePositions     []ClosePosition     `json:"ClosePositions"`     // 返済建玉指定
	FrontOrderType     StockFrontOrderType `json:"FrontOrderType"`     // 執行条件
	Price              int                 `json:"Price"`              // 注文価格
	ExpireDay          YmdNUM              `json:"ExpireDay"`          // 注文有効期限（年月日）
}

// SendOrderFutureResponse - 注文発注(先物)のレスポンス
type SendOrderFutureResponse struct {
	Result  int    `json:"Result"`  // 結果コード 0が成功、それ以外はエラー
	OrderID string `json:"OrderId"` // 受付注文番号
}

// NewSendOrderFutureRequester - 注文発注(先物)リクエスタの生成
func NewSendOrderFutureRequester(token string, isProd bool) *sendOrderFutureRequester {
	return &sendOrderFutureRequester{httpClient{url: createURL("/sendorder/future", isProd), token: token}}
}

// sendOrderFutureRequester - 注文発注(先物)のリクエスタ
type sendOrderFutureRequester struct {
	httpClient
}

// Exec - 注文発注(先物)リクエストの実行
func (r *sendOrderFutureRequester) Exec(request SendOrderFutureRequest) (*SendOrderFutureResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 注文発注(先物)リクエストの実行(contextあり)
func (r *sendOrderFutureRequester) ExecWithContext(ctx context.Context, request SendOrderFutureRequest) (*SendOrderFutureResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := r.httpClient.post(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	var res SendOrderFutureResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
