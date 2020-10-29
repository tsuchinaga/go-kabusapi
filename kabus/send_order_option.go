package kabus

import (
	"context"
	"encoding/json"
)

// SendOrderOptionRequest - 注文発注(オプション)のリクエストパラメータ
type SendOrderOptionRequest struct {
	Password           string              `json:"Password"`           // 注文パスワード
	Symbol             string              `json:"Symbol"`             // 銘柄コード
	Exchange           FutureExchange      `json:"Exchange"`           // オプション市場コード
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

// SendOrderOptionResponse - 注文発注(オプション)のレスポンス
type SendOrderOptionResponse struct {
	Result  int    `json:"Result"`  // 結果コード 0が成功、それ以外はエラー
	OrderID string `json:"OrderId"` // 受付注文番号
}

// NewSendOrderOptionRequester - 注文発注(オプション)リクエスタの生成
func NewSendOrderOptionRequester(token string, isProd bool) *sendOrderOptionRequester {
	return &sendOrderOptionRequester{httpClient{url: createURL("/sendorder/option", isProd), token: token}}
}

// sendOrderOptionRequester - 注文発注(オプション)のリクエスタ
type sendOrderOptionRequester struct {
	httpClient
}

// Exec - 注文発注(オプション)リクエストの実行
func (r *sendOrderOptionRequester) Exec(request SendOrderOptionRequest) (*SendOrderOptionResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 注文発注(オプション)リクエストの実行(contextあり)
func (r *sendOrderOptionRequester) ExecWithContext(ctx context.Context, request SendOrderOptionRequest) (*SendOrderOptionResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := r.httpClient.post(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	var res SendOrderOptionResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
