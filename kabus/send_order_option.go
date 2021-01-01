package kabus

import (
	"context"
	"encoding/json"
)

// SendOrderOptionRequest - 注文発注(オプション)のリクエストパラメータ
type SendOrderOptionRequest struct {
	Password           string               `json:"Password"`           // 注文パスワード
	Symbol             string               `json:"Symbol"`             // 銘柄コード
	Exchange           OptionExchange       `json:"Exchange"`           // オプション市場コード
	TradeType          TradeType            `json:"TradeType"`          // 取引区分
	TimeInForce        TimeInForce          `json:"TimeInForce"`        // 有効期間条件
	Side               Side                 `json:"Side"`               // 売買区分
	Qty                int                  `json:"Qty"`                // 注文数量
	ClosePositionOrder ClosePositionOrder   `json:"ClosePositionOrder"` // 決済順序
	ClosePositions     []ClosePosition      `json:"ClosePositions"`     // 返済建玉指定
	FrontOrderType     OptionFrontOrderType `json:"FrontOrderType"`     // 執行条件
	Price              float64              `json:"Price"`              // 注文価格
	ExpireDay          YmdNUM               `json:"ExpireDay"`          // 注文有効期限（年月日）
}

func (r *SendOrderOptionRequest) toJSON() ([]byte, error) {
	// エントリー
	if r.TradeType == TradeTypeEntry {
		return json.Marshal(sendOrderOptionEntryRequest{
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
			return json.Marshal(sendOrderOptionExitRequestWithClosePositions{
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
			return json.Marshal(sendOrderOptionExitRequestWithClosePositionOrder{
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

// sendOrderOptionEntryRequest - 注文発注(オプション)のエントリーリクエスト
type sendOrderOptionEntryRequest struct {
	Password       string               `json:"Password"`       // 注文パスワード
	Symbol         string               `json:"Symbol"`         // 銘柄コード
	Exchange       OptionExchange       `json:"Exchange"`       // オプション市場コード
	TradeType      TradeType            `json:"TradeType"`      // 取引区分
	TimeInForce    TimeInForce          `json:"TimeInForce"`    // 有効期間条件
	Side           Side                 `json:"Side"`           // 売買区分
	Qty            int                  `json:"Qty"`            // 注文数量
	FrontOrderType OptionFrontOrderType `json:"FrontOrderType"` // 執行条件
	Price          float64              `json:"Price"`          // 注文価格
	ExpireDay      YmdNUM               `json:"ExpireDay"`      // 注文有効期限（年月日）
}

// sendOrderOptionExitRequestWithClosePositions - 注文発注(オプション)のエグジットリクエスト(建玉指定)
type sendOrderOptionExitRequestWithClosePositions struct {
	Password       string               `json:"Password"`       // 注文パスワード
	Symbol         string               `json:"Symbol"`         // 銘柄コード
	Exchange       OptionExchange       `json:"Exchange"`       // オプション市場コード
	TradeType      TradeType            `json:"TradeType"`      // 取引区分
	TimeInForce    TimeInForce          `json:"TimeInForce"`    // 有効期間条件
	Side           Side                 `json:"Side"`           // 売買区分
	Qty            int                  `json:"Qty"`            // 注文数量
	ClosePositions []ClosePosition      `json:"ClosePositions"` // 返済建玉指定
	FrontOrderType OptionFrontOrderType `json:"FrontOrderType"` // 執行条件
	Price          float64              `json:"Price"`          // 注文価格
	ExpireDay      YmdNUM               `json:"ExpireDay"`      // 注文有効期限（年月日）
}

// sendOrderOptionExitRequestWithClosePositionOrder - 注文発注(オプション)のエグジットリクエスト(建玉順序指定)
type sendOrderOptionExitRequestWithClosePositionOrder struct {
	Password           string               `json:"Password"`           // 注文パスワード
	Symbol             string               `json:"Symbol"`             // 銘柄コード
	Exchange           OptionExchange       `json:"Exchange"`           // オプション市場コード
	TradeType          TradeType            `json:"TradeType"`          // 取引区分
	TimeInForce        TimeInForce          `json:"TimeInForce"`        // 有効期間条件
	Side               Side                 `json:"Side"`               // 売買区分
	Qty                int                  `json:"Qty"`                // 注文数量
	ClosePositionOrder ClosePositionOrder   `json:"ClosePositionOrder"` // 決済順序
	FrontOrderType     OptionFrontOrderType `json:"FrontOrderType"`     // 執行条件
	Price              float64              `json:"Price"`              // 注文価格
	ExpireDay          YmdNUM               `json:"ExpireDay"`          // 注文有効期限（年月日）
}

// SendOrderOptionResponse - 注文発注(オプション)のレスポンス
type SendOrderOptionResponse struct {
	Result  int    `json:"Result"`  // 結果コード 0が成功、それ以外はエラー
	OrderID string `json:"OrderId"` // 受付注文番号
}

// SendOrderOption - 注文発注(オプション)リクエスト
func (c *restClient) SendOrderOption(token string, request SendOrderOptionRequest) (*SendOrderOptionResponse, error) {
	return c.SendOrderOptionWithContext(context.Background(), token, request)
}

// SendOrderOptionWithContext - 注文発注(オプション)リクエスト(contextあり)
func (c *restClient) SendOrderOptionWithContext(ctx context.Context, token string, request SendOrderOptionRequest) (*SendOrderOptionResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	code, b, err := c.post(ctx, token, "sendorder/option", reqBody)
	if err != nil {
		return nil, err
	}

	var res SendOrderOptionResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
