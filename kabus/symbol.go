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
	DisplayName        string          `json:"DisplayName"`        // 銘柄略称
	Exchange           Exchange        `json:"Exchange"`           // 市場コード
	ExchangeName       string          `json:"ExchangeName"`       // 市場名称
	BisCategory        string          `json:"BisCategory"`        // 業種コード名 TODO 必要ならenumにする
	TotalMarketValue   float64         `json:"TotalMarketValue"`   // 時価総額
	TotalStocks        float64         `json:"TotalStocks"`        // 発行済み株式数（千株）
	TradingUnit        float64         `json:"TradingUnit"`        // 売買単位
	FiscalYearEndBasic float64         `json:"FiscalYearEndBasic"` // 決算期日
	PriceRangeGroup    PriceRangeGroup `json:"PriceRangeGroup"`    // 呼値グループ
	KCMarginBuy        bool            `json:"KCMarginBuy"`        // 一般信用買建フラグ
	KCMarginSell       bool            `json:"KCMarginSell"`       // 一般信用売建フラグ
	MarginBuy          bool            `json:"MarginBuy"`          // 制度信用買建フラグ
	MarginSell         bool            `json:"MarginSell"`         // 制度信用売建フラグ
	UpperLimit         float64         `json:"UpperLimit"`         // 値幅上限
	LowerLimit         float64         `json:"LowerLimit"`         // 値幅下限
}

// NewSymbolRequester - 銘柄情報リクエスタの生成
func NewSymbolRequester(token string, isProd bool) *symbolRequester {
	u := "http://localhost:18080/kabusapi/symbol"
	if !isProd {
		u = "http://localhost:18081/kabusapi/symbol"
	}
	return &symbolRequester{client{token: token, url: u}}
}

// symbolRequester - 銘柄情報リクエスタ
type symbolRequester struct {
	client
}

// Exec - 銘柄情報リクエストの実行
func (r *symbolRequester) Exec(request SymbolRequest) (*SymbolResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 銘柄情報リクエストの実行(contextあり)
func (r *symbolRequester) ExecWithContext(ctx context.Context, request SymbolRequest) (*SymbolResponse, error) {
	pathParam := fmt.Sprintf("%s@%d", request.Symbol, request.Exchange)
	code, b, err := r.client.get(ctx, pathParam, "")
	if err != nil {
		return nil, err
	}

	var res SymbolResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
