package kabus

import (
	"context"
	"fmt"
	"strings"
)

// PositionsRequest - 残高照会のリクエストパラメータ
type PositionsRequest struct {
	Product Product         // 取得する商品
	Symbol  string          // 銘柄コード
	Side    Side            // 売買区分フィルタ
	AddInfo GetPositionInfo // 追加情報出力フラグ
}

func (r *PositionsRequest) toQuery() string {
	var params []string
	params = append(params, fmt.Sprintf("product=%d", r.Product))

	if r.Symbol != "" {
		params = append(params, fmt.Sprintf("symbol=%s", r.Symbol))
	}

	if r.Side != SideUnspecified {
		params = append(params, fmt.Sprintf("side=%s", r.Side))
	}

	if r.AddInfo != GetSymbolInfoUnspecified {
		params = append(params, fmt.Sprintf("addinfo=%s", r.AddInfo))
	}

	return strings.Join(params, "&")
}

// PositionsResponse - 残高照会のレスポンス
type PositionsResponse []Position

// Position - 残高照会で返されるポジションの情報
type Position struct {
	ExecutionID     string          `json:"ExecutionID"`     // 約定番号
	AccountType     AccountType     `json:"AccountType"`     // 口座種別
	Symbol          string          `json:"Symbol"`          // 銘柄コード
	SymbolName      string          `json:"SymbolName"`      // 銘柄名
	Exchange        Exchange        `json:"Exchange"`        // 市場コード
	ExchangeName    string          `json:"ExchangeName"`    // 市場名
	SecurityType    SecurityType    `json:"SecurityType"`    // 銘柄種別
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

// Positions - 残高照会リクエスト
func (c *restClient) Positions(token string, request PositionsRequest) (*PositionsResponse, error) {
	return c.PositionsWithContext(context.Background(), token, request)
}

// PositionsWithContext - 残高照会リクエスト(contextあり)
func (c *restClient) PositionsWithContext(ctx context.Context, token string, request PositionsRequest) (*PositionsResponse, error) {
	code, b, err := c.get(ctx, token, "positions", request.toQuery())
	if err != nil {
		return nil, err
	}

	var res PositionsResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
