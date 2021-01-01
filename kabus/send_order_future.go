package kabus

import (
	"context"
	"encoding/json"
)

// SendOrderFutureRequest - 注文発注(先物)のリクエストパラメータ
type SendOrderFutureRequest struct {
	Password           string               `json:"Password"`           // 注文パスワード
	Symbol             string               `json:"Symbol"`             // 銘柄コード
	Exchange           FutureExchange       `json:"Exchange"`           // 先物市場コード
	TradeType          TradeType            `json:"TradeType"`          // 取引区分
	TimeInForce        TimeInForce          `json:"TimeInForce"`        // 有効期間条件
	Side               Side                 `json:"Side"`               // 売買区分
	Qty                int                  `json:"Qty"`                // 注文数量
	ClosePositionOrder ClosePositionOrder   `json:"ClosePositionOrder"` // 決済順序
	ClosePositions     []ClosePosition      `json:"ClosePositions"`     // 返済建玉指定
	FrontOrderType     FutureFrontOrderType `json:"FrontOrderType"`     // 執行条件
	Price              float64              `json:"Price"`              // 注文価格
	ExpireDay          YmdNUM               `json:"ExpireDay"`          // 注文有効期限（年月日）
}

func (r *SendOrderFutureRequest) toJSON() ([]byte, error) {
	// エントリー
	if r.TradeType == TradeTypeEntry {
		return json.Marshal(sendOrderFutureEntryRequest{
			Password:       r.Password,
			Symbol:         r.Symbol,
			Exchange:       r.Exchange,
			TradeType:      r.TradeType,
			TimeInForce:    r.TimeInForce,
			Side:           r.Side,
			Qty:            r.Qty,
			FrontOrderType: r.FrontOrderType,
			Price:          r.Price,
			ExpireDay:      r.ExpireDay,
		})
	} else if r.TradeType == TradeTypeExit {
		// 返済建玉指定に指定があれば決済順序のリクエストにする
		if r.ClosePositions != nil && len(r.ClosePositions) > 0 {
			return json.Marshal(sendOrderFutureExitRequestWithClosePositions{
				Password:       r.Password,
				Symbol:         r.Symbol,
				Exchange:       r.Exchange,
				TradeType:      r.TradeType,
				TimeInForce:    r.TimeInForce,
				Side:           r.Side,
				Qty:            r.Qty,
				ClosePositions: r.ClosePositions,
				FrontOrderType: r.FrontOrderType,
				Price:          r.Price,
				ExpireDay:      r.ExpireDay,
			})
		} else {
			return json.Marshal(sendOrderFutureExitRequestWithClosePositionOrder{
				Password:           r.Password,
				Symbol:             r.Symbol,
				Exchange:           r.Exchange,
				TradeType:          r.TradeType,
				TimeInForce:        r.TimeInForce,
				Side:               r.Side,
				Qty:                r.Qty,
				ClosePositionOrder: r.ClosePositionOrder,
				FrontOrderType:     r.FrontOrderType,
				Price:              r.Price,
				ExpireDay:          r.ExpireDay,
			})
		}
	}

	// 上記以外はそのまま出す
	return json.Marshal(r)
}

// sendOrderFutureEntryRequest - 注文発注(先物)のエントリーリクエスト
type sendOrderFutureEntryRequest struct {
	Password       string               `json:"Password"`       // 注文パスワード
	Symbol         string               `json:"Symbol"`         // 銘柄コード
	Exchange       FutureExchange       `json:"Exchange"`       // 先物市場コード
	TradeType      TradeType            `json:"TradeType"`      // 取引区分
	TimeInForce    TimeInForce          `json:"TimeInForce"`    // 有効期間条件
	Side           Side                 `json:"Side"`           // 売買区分
	Qty            int                  `json:"Qty"`            // 注文数量
	FrontOrderType FutureFrontOrderType `json:"FrontOrderType"` // 執行条件
	Price          float64              `json:"Price"`          // 注文価格
	ExpireDay      YmdNUM               `json:"ExpireDay"`      // 注文有効期限（年月日）
}

// sendOrderFutureExitRequestWithClosePositionOrder - 注文発注(先物)のエグジットリクエスト(建玉指定)
type sendOrderFutureExitRequestWithClosePositionOrder struct {
	Password           string               `json:"Password"`           // 注文パスワード
	Symbol             string               `json:"Symbol"`             // 銘柄コード
	Exchange           FutureExchange       `json:"Exchange"`           // 先物市場コード
	TradeType          TradeType            `json:"TradeType"`          // 取引区分
	TimeInForce        TimeInForce          `json:"TimeInForce"`        // 有効期間条件
	Side               Side                 `json:"Side"`               // 売買区分
	Qty                int                  `json:"Qty"`                // 注文数量
	ClosePositionOrder ClosePositionOrder   `json:"ClosePositionOrder"` // 決済順序
	FrontOrderType     FutureFrontOrderType `json:"FrontOrderType"`     // 執行条件
	Price              float64              `json:"Price"`              // 注文価格
	ExpireDay          YmdNUM               `json:"ExpireDay"`          // 注文有効期限（年月日）
}

// sendOrderFutureExitRequestWithClosePositions - 注文発注(先物)のエグジットリクエスト(建玉順序指定)
type sendOrderFutureExitRequestWithClosePositions struct {
	Password       string               `json:"Password"`       // 注文パスワード
	Symbol         string               `json:"Symbol"`         // 銘柄コード
	Exchange       FutureExchange       `json:"Exchange"`       // 先物市場コード
	TradeType      TradeType            `json:"TradeType"`      // 取引区分
	TimeInForce    TimeInForce          `json:"TimeInForce"`    // 有効期間条件
	Side           Side                 `json:"Side"`           // 売買区分
	Qty            int                  `json:"Qty"`            // 注文数量
	ClosePositions []ClosePosition      `json:"ClosePositions"` // 返済建玉指定
	FrontOrderType FutureFrontOrderType `json:"FrontOrderType"` // 執行条件
	Price          float64              `json:"Price"`          // 注文価格
	ExpireDay      YmdNUM               `json:"ExpireDay"`      // 注文有効期限（年月日）
}

// SendOrderFutureResponse - 注文発注(先物)のレスポンス
type SendOrderFutureResponse struct {
	Result  int    `json:"Result"`  // 結果コード 0が成功、それ以外はエラー
	OrderID string `json:"OrderId"` // 受付注文番号
}

// SendOrderFuture - 注文発注(先物)リクエスト
func (c *restClient) SendOrderFuture(token string, request SendOrderFutureRequest) (*SendOrderFutureResponse, error) {
	return c.SendOrderFutureWithContext(context.Background(), token, request)
}

// SendOrderFutureWithContext - 注文発注(先物)リクエスト(contextあり)
func (c *restClient) SendOrderFutureWithContext(ctx context.Context, token string, request SendOrderFutureRequest) (*SendOrderFutureResponse, error) {
	reqBody, err := request.toJSON()
	if err != nil {
		return nil, err
	}

	code, b, err := c.post(ctx, token, "sendorder/future", reqBody)
	if err != nil {
		return nil, err
	}

	var res SendOrderFutureResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
