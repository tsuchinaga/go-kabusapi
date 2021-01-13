package kabus

import (
	"context"
	"fmt"
)

// SymbolRequest - 銘柄情報のリクエストパラメータ
type SymbolRequest struct {
	Symbol   string   // 銘柄コード
	Exchange Exchange // 市場コード
}

// SymbolResponse - 銘柄情報のレスポンス
type SymbolResponse struct {
	Symbol             string          `json:"Symbol"`             // 銘柄コード
	SymbolName         string          `json:"SymbolName"`         // 銘柄名
	DisplayName        string          `json:"DisplayName"`        // 銘柄略称 ※株式・先物・オプション
	Exchange           Exchange        `json:"Exchange"`           // 市場コード ※株式・先物・オプション
	ExchangeName       string          `json:"ExchangeName"`       // 市場名称 ※株式・先柄・オプション
	BisCategory        string          `json:"BisCategory"`        // 業種コード名 ※株式
	TotalMarketValue   float64         `json:"TotalMarketValue"`   // 時価総額 ※株式
	TotalStocks        float64         `json:"TotalStocks"`        // 発行済み株式数（千株） ※株式
	TradingUnit        float64         `json:"TradingUnit"`        // 売買単位 ※株式・先柄・オプション
	FiscalYearEndBasic YmdNUM          `json:"FiscalYearEndBasic"` // 決算期日 ※株式
	PriceRangeGroup    PriceRangeGroup `json:"PriceRangeGroup"`    // 呼値グループ ※株式・先柄・オプション
	KCMarginBuy        bool            `json:"KCMarginBuy"`        // 一般信用買建フラグ ※株式
	KCMarginSell       bool            `json:"KCMarginSell"`       // 一般信用売建フラグ ※株式
	MarginBuy          bool            `json:"MarginBuy"`          // 制度信用買建フラグ ※株式
	MarginSell         bool            `json:"MarginSell"`         // 制度信用売建フラグ ※株式
	UpperLimit         float64         `json:"UpperLimit"`         // 値幅上限 ※株式・先柄・オプション
	LowerLimit         float64         `json:"LowerLimit"`         // 値幅下限 ※株式・先柄・オプション
	Underlyer          Underlyer       `json:"Underlyer"`          // 原資産コード ※先柄・オプション
	DerivMonth         YmString        `json:"DerivMonth"`         // 限月-年月 ※先柄・オプション
	TradeStart         YmdNUM          `json:"TradeStart"`         // 取引開始 ※先柄・オプション
	TradeEnd           YmdNUM          `json:"TradeEnd"`           // 取引終了日 ※先柄・オプション
	StrikePrice        float64         `json:"StrikePrice"`        // 権利行使価格 ※オプション
	PutOrCall          PutOrCallNum    `json:"PutOrCall"`          // プット/コール区分 ※オプション
	ClearingPrice      float64         `json:"ClearingPrice"`      // 清算値 ※先物
}

// Symbol - 銘柄情報リクエスト
func (c *restClient) Symbol(token string, request SymbolRequest) (*SymbolResponse, error) {
	return c.SymbolWithContext(context.Background(), token, request)
}

// SymbolWithContext - 銘柄情報リクエスト(contextあり)
func (c *restClient) SymbolWithContext(ctx context.Context, token string, request SymbolRequest) (*SymbolResponse, error) {
	path := fmt.Sprintf("symbol/%s@%d", request.Symbol, request.Exchange)
	code, b, err := c.get(ctx, token, path, "")
	if err != nil {
		return nil, err
	}

	var res SymbolResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
