package kabus

import (
	"context"
	"encoding/json"
)

// SendOrderRequest - 注文発注のリクエストパラメータ
type SendOrderRequest struct {
	Password           string             `json:"Password"`           // 注文パスワード
	Symbol             string             `json:"Symbol"`             // 銘柄コード
	Exchange           Exchange           `json:"Exchange"`           // 市場コード
	SecurityType       SecurityType       `json:"SecurityType"`       // 商品種別
	Side               Side               `json:"Side"`               // 売買区分
	CashMargin         CashMargin         `json:"CashMargin"`         // 現物信用区分
	MarginTradeType    MarginTradeType    `json:"MarginTradeType"`    // 信用取引区分 ※信用取引の場合必須
	DelivType          DelivType          `json:"DelivType"`          // 受渡区分 ※株式取引の場合必須で、現物売・信用返済では「指定なし」を指定する
	FundType           FundType           `json:"FundType"`           // 資産区分 ※株式取引の場合必須で、現物売・信用返済では「指定なし」を指定する
	AccountType        AccountType        `json:"AccountType"`        // 口座種別
	Qty                int                `json:"Qty"`                // 注文数量
	ClosePositionOrder ClosePositionOrder `json:"ClosePositionOrder"` // 決済順序 ※信用取引の場合必須
	ClosePositions     []ClosePosition    `json:"ClosePositions"`     // 返済建玉指定 ※信用取引の場合必須
	Price              int                `json:"Price"`              // 注文価格
	ExpireDay          YmdNUM             `json:"ExpireDay"`          // 注文有効期限（年月日）
	FrontOrderType     FrontOrderType     `json:"FrontOrderType"`     // 執行条件
}

func (r *SendOrderRequest) toJSON() ([]byte, error) {
	if r.CashMargin == CashMarginMarginEntry || r.CashMargin == CashMarginMarginExit {
		// 返済建玉指定に指定があれば決済順序なしのリクエストにする
		if r.ClosePositions != nil && len(r.ClosePositions) > 0 {
			return json.Marshal(sendOrderRequestWithoutClosePositionOrder{
				Password:        r.Password,
				Symbol:          r.Symbol,
				Exchange:        r.Exchange,
				SecurityType:    r.SecurityType,
				Side:            r.Side,
				CashMargin:      r.CashMargin,
				MarginTradeType: r.MarginTradeType,
				DelivType:       r.DelivType,
				FundType:        r.FundType,
				AccountType:     r.AccountType,
				Qty:             r.Qty,
				ClosePositions:  r.ClosePositions,
				Price:           r.Price,
				ExpireDay:       r.ExpireDay,
				FrontOrderType:  r.FrontOrderType,
			})
		} else {
			return json.Marshal(sendOrderRequestWithoutClosePositions{
				Password:           r.Password,
				Symbol:             r.Symbol,
				Exchange:           r.Exchange,
				SecurityType:       r.SecurityType,
				Side:               r.Side,
				CashMargin:         r.CashMargin,
				MarginTradeType:    r.MarginTradeType,
				DelivType:          r.DelivType,
				FundType:           r.FundType,
				AccountType:        r.AccountType,
				Qty:                r.Qty,
				ClosePositionOrder: r.ClosePositionOrder,
				Price:              r.Price,
				ExpireDay:          r.ExpireDay,
				FrontOrderType:     r.FrontOrderType,
			})
		}
	}

	// 上記以外はそのまま出す
	return json.Marshal(r)
}

// sendOrderRequestWithoutClosePositionOrder - 決済順序なしの注文発注のリクエストパラメータ
type sendOrderRequestWithoutClosePositionOrder struct {
	Password        string          `json:"Password"`        // 注文パスワード
	Symbol          string          `json:"Symbol"`          // 銘柄コード
	Exchange        Exchange        `json:"Exchange"`        // 市場コード
	SecurityType    SecurityType    `json:"SecurityType"`    // 商品種別
	Side            Side            `json:"Side"`            // 売買区分
	CashMargin      CashMargin      `json:"CashMargin"`      // 現物信用区分
	MarginTradeType MarginTradeType `json:"MarginTradeType"` // 信用取引区分 ※信用取引の場合必須
	DelivType       DelivType       `json:"DelivType"`       // 受渡区分 ※株式取引の場合必須で、現物売・信用返済では「指定なし」を指定する
	FundType        FundType        `json:"FundType"`        // 資産区分 ※株式取引の場合必須で、現物売・信用返済では「指定なし」を指定する
	AccountType     AccountType     `json:"AccountType"`     // 口座種別
	Qty             int             `json:"Qty"`             // 注文数量
	ClosePositions  []ClosePosition `json:"ClosePositions"`  // 返済建玉指定
	Price           int             `json:"Price"`           // 注文価格
	ExpireDay       YmdNUM          `json:"ExpireDay"`       // 注文有効期限（年月日）
	FrontOrderType  FrontOrderType  `json:"FrontOrderType"`  // 執行条件
}

// sendOrderRequestWithoutClosePositions - 返済建玉指定なしの注文発注のリクエストパラメータ
type sendOrderRequestWithoutClosePositions struct {
	Password           string             `json:"Password"`           // 注文パスワード
	Symbol             string             `json:"Symbol"`             // 銘柄コード
	Exchange           Exchange           `json:"Exchange"`           // 市場コード
	SecurityType       SecurityType       `json:"SecurityType"`       // 商品種別
	Side               Side               `json:"Side"`               // 売買区分
	CashMargin         CashMargin         `json:"CashMargin"`         // 現物信用区分
	MarginTradeType    MarginTradeType    `json:"MarginTradeType"`    // 信用取引区分 ※信用取引の場合必須
	DelivType          DelivType          `json:"DelivType"`          // 受渡区分 ※株式取引の場合必須で、現物売・信用返済では「指定なし」を指定する
	FundType           FundType           `json:"FundType"`           // 資産区分 ※株式取引の場合必須で、現物売・信用返済では「指定なし」を指定する
	AccountType        AccountType        `json:"AccountType"`        // 口座種別
	Qty                int                `json:"Qty"`                // 注文数量
	ClosePositionOrder ClosePositionOrder `json:"ClosePositionOrder"` // 決済順序 ※信用取引の場合必須
	Price              int                `json:"Price"`              // 注文価格
	ExpireDay          YmdNUM             `json:"ExpireDay"`          // 注文有効期限（年月日）
	FrontOrderType     FrontOrderType     `json:"FrontOrderType"`     // 執行条件
}

// ClosePosition - 返済建玉
type ClosePosition struct {
	HoldID string `json:"HoldID"` // 返済建玉ID
	Qty    int    `json:"Qty"`    // 返済建玉数量
}

// SendOrderResponse - 注文発注のレスポンス
type SendOrderResponse struct {
	Result  int    `json:"Result"`  // 結果コード 0が成功、それ以外はエラー
	OrderID string `json:"OrderId"` // 受付注文番号
}

// NewSendOrderRequester - 注文発注リクエスタの生成
func NewSendOrderRequester(token string, isProd bool) *sendOrderRequester {
	return &sendOrderRequester{httpClient{url: createURL("/sendorder", isProd), token: token}}
}

// sendOrderRequester - 注文発注のリクエスタ
type sendOrderRequester struct {
	httpClient
}

// Exec - 注文発注リクエストの実行
func (r *sendOrderRequester) Exec(request SendOrderRequest) (*SendOrderResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 注文発注リクエストの実行(contextあり)
func (r *sendOrderRequester) ExecWithContext(ctx context.Context, request SendOrderRequest) (*SendOrderResponse, error) {
	reqBody, err := request.toJSON()
	if err != nil {
		return nil, err
	}

	code, b, err := r.httpClient.post(ctx, reqBody)
	if err != nil {
		return nil, err
	}

	var res SendOrderResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
