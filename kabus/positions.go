package kabus

import (
	"context"
	"fmt"
)

// PositionsRequest - 残高照会のリクエストパラメータ
type PositionsRequest struct {
	Product Product // 取得する商品
}

// PositionsResponse - 残高照会のレスポンス
type PositionsResponse []Position

// Position - 残高照会で返されるポジションの情報
type Position struct {
	ExecutionID     string          `json:"ExecutionID"`     // 約定番号
	AccountType     AccountType     `json:"AccountType"`     // 口座種別
	Symbol          string          `json:"Symbol"`          // 銘柄コード
	SymbolName      string          `json:"SymbolName"`      // 銘柄名
	Exchange        StockExchange   `json:"Exchange"`        // 市場コード
	ExchangeName    string          `json:"ExchangeName"`    // 市場名
	ExecutionDay    YmdNUM          `json:"ExecutionDay"`    // 約定日（建玉日）
	Price           float64         `json:"Price"`           // 値段
	LeavesQty       float64         `json:"LeavesQty"`       // 残数量
	HoldQty         float64         `json:"HoldQty"`         // 拘束数量（保有数量）
	Side            Side            `json:"Side"`            // 売買区分
	Expenses        float64         `json:"Expenses"`        // 諸経費
	Commission      float64         `json:"Commission"`      // 手数料
	CommissionTax   float64         `json:"CommissionTax"`   // 手数料消費税
	ExpireDay       YmdNUM          `json:"ExpireDay"`       // 返済期日
	MarginTradeType MarginTradeType `json:"MarginTradeType"` // 信用取引区分
	CurrentPrice    float64         `json:"CurrentPrice"`    // 現在値
	Valuation       float64         `json:"Valuation"`       // 評価金額
	ProfitLoss      float64         `json:"ProfitLoss"`      // 評価損益額
	ProfitLossRate  float64         `json:"ProfitLossRate"`  // 評価損益率
}

// NewPositionsRequester - 残高照会リクエスタの生成
func NewPositionsRequester(token string, isProd bool) *positionsRequester {
	return &positionsRequester{httpClient{token: token, url: createURL("/positions", isProd)}}
}

// positionsRequester - 残高照会のリクエスタ
type positionsRequester struct {
	httpClient
}

// Exec - 残高照会リクエストの実行
func (r *positionsRequester) Exec(request PositionsRequest) (*PositionsResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 残高照会リクエストの実行(contextあり)
func (r *positionsRequester) ExecWithContext(ctx context.Context, request PositionsRequest) (*PositionsResponse, error) {
	queryParam := fmt.Sprintf("product=%d", request.Product)

	code, b, err := r.httpClient.get(ctx, "", queryParam)
	if err != nil {
		return nil, err
	}

	var res PositionsResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
