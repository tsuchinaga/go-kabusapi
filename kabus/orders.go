package kabus

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// OrdersRequest - 注文約定照会のリクエストパラメータ
type OrdersRequest struct {
	Product          Product          // 取得する商品
	ID               string           // 注文番号
	UpdateTime       time.Time        // 更新日時 ※指定された更新日時以降（指定日時含む）に更新された注文のみ返す
	IsGetOrderDetail IsGetOrderDetail // 注文詳細抑止
	Symbol           string           // 銘柄コード
	State            OrderState       // 注文状態
	Side             Side             // 売買区分
	CashMargin       CashMargin       // 取引区分 TODO 現物は指定できないようにする
}

func (r *OrdersRequest) toQuery() string {
	var params []string

	// 取得する商品
	params = append(params, fmt.Sprintf("product=%d", r.Product))

	// 注文番号
	if r.ID != "" {
		params = append(params, fmt.Sprintf("id=%s", r.ID))
	}

	// 更新日時
	if !r.UpdateTime.IsZero() {
		params = append(params, fmt.Sprintf("updtime=%s", r.UpdateTime.Format("20060102150405")))
	}

	// 注文詳細抑止
	if r.IsGetOrderDetail != IsGetOrderDetailUnspecified {
		params = append(params, fmt.Sprintf("details=%s", r.IsGetOrderDetail))
	}

	// 銘柄コード
	if r.Symbol != "" {
		params = append(params, fmt.Sprintf("symbol=%s", r.Symbol))
	}

	// 状態
	if r.State != OrderStateUnspecified {
		params = append(params, fmt.Sprintf("state=%d", r.State))
	}

	// 売買区分
	if r.Side != SideUnspecified {
		params = append(params, fmt.Sprintf("side=%s", r.Side))
	}

	// 取引区分
	if r.CashMargin != CashMarginUnspecified {
		params = append(params, fmt.Sprintf("cashmargin=%d", r.CashMargin))
	}

	return strings.Join(params, "&")
}

// OrdersResponse - 注文約定照会のレスポンス
type OrdersResponse []Order

// Order - 注文約定照会で返される注文の情報
type Order struct {
	ID              string          `json:"ID"`              // 注文番号
	State           State           `json:"State"`           // 状態
	OrderState      OrderState      `json:"OrderState"`      // 注文状態
	OrdType         OrdType         `json:"OrdType"`         // 執行条件
	RecvTime        time.Time       `json:"RecvTime"`        // 受注日時
	Symbol          string          `json:"Symbol"`          // 銘柄コード
	SymbolName      string          `json:"SymbolName"`      // 銘柄名
	Exchange        OrderExchange   `json:"Exchange"`        // 市場コード
	ExchangeName    string          `json:"ExchangeName"`    // 市場名
	Price           float64         `json:"Price"`           // 値段
	OrderQty        float64         `json:"OrderQty"`        // 発注数量
	CumQty          float64         `json:"CumQty"`          // 約定数量
	Side            Side            `json:"Side"`            // 売買区分
	CashMargin      CashMargin      `json:"CashMargin"`      // 現物信用区分
	AccountType     AccountType     `json:"AccountType"`     // 口座種別
	DelivType       DelivType       `json:"DelivType"`       // 受渡区分
	ExpireDay       YmdNUM          `json:"ExpireDay"`       // 注文有効期限
	MarginTradeType MarginTradeType `json:"MarginTradeType"` // 信用取引区分
	Details         []OrderDetail   `json:"Details"`         // 注文詳細
}

// OrderDetail - 注文詳細
type OrderDetail struct {
	SeqNum        int              `json:"SeqNum"`        // 連番
	ID            string           `json:"ID"`            // 注文詳細番号
	RecType       RecType          `json:"RecType"`       // 明細種別
	ExchangeID    string           `json:"ExchangeID"`    // 取引所番号
	State         OrderDetailState `json:"State"`         // 状態
	TransactTime  time.Time        `json:"TransactTime"`  // 処理時刻
	OrdType       OrdType          `json:"OrdType"`       // 執行条件
	Price         float64          `json:"Price"`         // 値段
	Qty           float64          `json:"Qty"`           // 数量
	ExecutionID   string           `json:"ExecutionID"`   // 約定番号
	ExecutionDay  time.Time        `json:"ExecutionDay"`  // 約定日時
	DelivDay      YmdNUM           `json:"DelivDay"`      // 受渡日
	Commission    float64          `json:"Commission"`    // 手数料
	CommissionTax float64          `json:"CommissionTax"` // 手数料消費税
}

// Orders - 注文約定照会リクエスト
func (c *restClient) Orders(token string, request OrdersRequest) (*OrdersResponse, error) {
	return c.OrdersWithContext(context.Background(), token, request)
}

// OrdersWithContext - 注文約定照会リクエスト(contextあり)
func (c *restClient) OrdersWithContext(ctx context.Context, token string, request OrdersRequest) (*OrdersResponse, error) {
	code, b, err := c.get(ctx, token, "orders", request.toQuery())
	if err != nil {
		return nil, err
	}

	var res OrdersResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
