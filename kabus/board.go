package kabus

import (
	"context"
	"fmt"
	"time"
)

// BoardRequest - 時価情報・板情報のリクエストパラメータ
type BoardRequest struct {
	Symbol   string   // 銘柄コード
	Exchange Exchange // 市場コード
}

// BoardResponse - 時価情報・板情報のレスポンス
type BoardResponse struct {
	Symbol                   string                   `json:"Symbol"`                   // 銘柄コード
	SymbolName               string                   `json:"SymbolName"`               // 銘柄名
	Exchange                 Exchange                 `json:"Exchange"`                 // 市場コード
	ExchangeName             string                   `json:"ExchangeName"`             // 市場名称
	CurrentPrice             float64                  `json:"CurrentPrice"`             // 現値
	CurrentPriceTime         time.Time                `json:"CurrentPriceTime"`         // 現値時刻
	CurrentPriceChangeStatus CurrentPriceChangeStatus `json:"CurrentPriceChangeStatus"` // 現値前値比較
	CurrentPriceStatus       CurrentPriceStatus       `json:"CurrentPriceStatus"`       // 現値ステータス
	CalcPrice                float64                  `json:"CalcPrice"`                // 計算用現値
	PreviousClose            float64                  `json:"PreviousClose"`            // 前日終値
	PreviousCloseTime        time.Time                `json:"PreviousCloseTime"`        // 前日終値日付
	ChangePreviousClose      float64                  `json:"ChangePreviousClose"`      // 前日比
	ChangePreviousClosePer   float64                  `json:"ChangePreviousClosePer"`   // 騰落率
	OpeningPrice             float64                  `json:"OpeningPrice"`             // 始値
	OpeningPriceTime         time.Time                `json:"OpeningPriceTime"`         // 始値時刻
	HighPrice                float64                  `json:"HighPrice"`                // 高値
	HighPriceTime            time.Time                `json:"HighPriceTime"`            // 高値時刻
	LowPrice                 float64                  `json:"LowPrice"`                 // 安値
	LowPriceTime             time.Time                `json:"LowPriceTime"`             // 安値時刻
	TradingVolume            float64                  `json:"TradingVolume"`            // 売買高
	TradingVolumeTime        time.Time                `json:"TradingVolumeTime"`        // 売買高時刻
	VWAP                     float64                  `json:"VWAP"`                     // 売買高加重平均価格（VWAP）
	TradingValue             float64                  `json:"TradingValue"`             // 売買代金
	BidQty                   float64                  `json:"BidQty"`                   // 最良売気配数量
	BidPrice                 float64                  `json:"BidPrice"`                 // 最良売気配値段
	BidTime                  time.Time                `json:"BidTime"`                  // 最良売気配時刻
	BidSign                  BidAskSign               `json:"BidSign"`                  // 最良売気配フラグ
	MarketOrderSellQty       float64                  `json:"MarketOrderSellQty"`       // 売成行数量
	Sell1                    FirstBoardSign           `json:"Sell1"`                    // 売気配数量1本目
	Sell2                    BoardSign                `json:"Sell2"`                    // 売気配数量2本目
	Sell3                    BoardSign                `json:"Sell3"`                    // 売気配数量3本目
	Sell4                    BoardSign                `json:"Sell4"`                    // 売気配数量4本目
	Sell5                    BoardSign                `json:"Sell5"`                    // 売気配数量5本目
	Sell6                    BoardSign                `json:"Sell6"`                    // 売気配数量6本目
	Sell7                    BoardSign                `json:"Sell7"`                    // 売気配数量7本目
	Sell8                    BoardSign                `json:"Sell8"`                    // 売気配数量8本目
	Sell9                    BoardSign                `json:"Sell9"`                    // 売気配数量9本目
	Sell10                   BoardSign                `json:"Sell10"`                   // 売気配数量10本目
	AskQty                   float64                  `json:"AskQty"`                   // 最良買気配数量
	AskPrice                 float64                  `json:"AskPrice"`                 // 最良買気配値段
	AskTime                  time.Time                `json:"AskTime"`                  // 最良買気配時刻
	AskSign                  BidAskSign               `json:"AskSign"`                  // 最良買気配フラグ
	MarketOrderBuyQty        float64                  `json:"MarketOrderBuyQty"`        // 買成行数量
	Buy1                     FirstBoardSign           `json:"Buy1"`                     // 買気配数量1本目
	Buy2                     BoardSign                `json:"Buy2"`                     // 買気配数量2本目
	Buy3                     BoardSign                `json:"Buy3"`                     // 買気配数量3本目
	Buy4                     BoardSign                `json:"Buy4"`                     // 買気配数量4本目
	Buy5                     BoardSign                `json:"Buy5"`                     // 買気配数量5本目
	Buy6                     BoardSign                `json:"Buy6"`                     // 買気配数量6本目
	Buy7                     BoardSign                `json:"Buy7"`                     // 買気配数量7本目
	Buy8                     BoardSign                `json:"Buy8"`                     // 買気配数量8本目
	Buy9                     BoardSign                `json:"Buy9"`                     // 買気配数量9本目
	Buy10                    BoardSign                `json:"Buy10"`                    // 買気配数量10本目
	OverSellQty              float64                  `json:"OverSellQty"`              // OVER気配数量
	UnderBuyQty              float64                  `json:"UnderBuyQty"`              // UNDER気配数量
	TotalMarketValue         float64                  `json:"TotalMarketValue"`         // 時価総額
	ClearingPrice            float64                  `json:"ClearingPrice"`            // 清算値
	IV                       float64                  `json:"IV"`                       // インプライド・ボラティリティ
	Gamma                    float64                  `json:"Gamma"`                    // ガンマ
	Theta                    float64                  `json:"Theta"`                    // セータ
	Vega                     float64                  `json:"Vega"`                     // ベガ
	Delta                    float64                  `json:"Delta"`                    // デルタ
	SecurityType             SecurityType             `json:"SecurityType"`             // 銘柄種別
}

// FirstBoardSign - 最良気配
type FirstBoardSign struct {
	Time  time.Time  `json:"Time"`  // 時刻
	Sign  BidAskSign `json:"Sign"`  // 気配フラグ
	Price float64    `json:"Price"` // 値段
	Qty   float64    `json:"Qty"`   // 数量
}

// BoardSign - 気配
type BoardSign struct {
	Price float64 `json:"Price"` // 値段
	Qty   float64 `json:"Qty"`   // 数量
}

// Board - 時価情報・板情報リクエスト
func (c *restClient) Board(token string, request BoardRequest) (*BoardResponse, error) {
	return c.BoardWithContext(context.Background(), token, request)
}

// BoardWithContext - 時価情報・板情報リクエスト(contextあり)
func (c *restClient) BoardWithContext(ctx context.Context, token string, request BoardRequest) (*BoardResponse, error) {
	path := fmt.Sprintf("board/%s@%d", request.Symbol, request.Exchange)
	code, b, err := c.get(ctx, token, path, "")
	if err != nil {
		return nil, err
	}

	var res BoardResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
