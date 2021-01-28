package kabus

import (
	"context"
	"fmt"
)

// RankingRequest - 詳細ランキングのリクエストパラメータ
type RankingRequest struct {
	Type             RankingType      // ランキング種別
	ExchangeDivision ExchangeDivision // 市場
}

// RankingResponse - 詳細ランキングのレスポンス
type RankingResponse struct {
	Type                 RankingType            `json:"Type"`             // ランキング種別
	ExchangeDivision     ExchangeDivision       `json:"ExchangeDivision"` // 市場
	PriceRanking         []PriceRanking         // 株価 ※Type=1~4
	TickRanking          []TickRanking          // TICK回数 ※Type=5
	VolumeRapidRanking   []VolumeRapidRanking   // 売買高急増 ※Type=6
	ValueRapidRanking    []ValueRapidRanking    // 売買代金急増 ※Type=7
	MarginRanking        []MarginRanking        // 信用残 ※Type=8~13
	CategoryPriceRanking []CategoryPriceRanking // 業種別株価 ※Type=14~15
}

// priceRankingResponse - 株価 ※Type=1~4
type priceRankingResponse struct {
	Type             RankingType      `json:"Type"`             // ランキング種別
	ExchangeDivision ExchangeDivision `json:"ExchangeDivision"` // 市場
	PriceRanking     []PriceRanking   `json:"Ranking"`          // 株価 ※Type=1~4
}

// tickRankingResponse - TICK回数 ※Type=5
type tickRankingResponse struct {
	Type             RankingType      `json:"Type"`             // ランキング種別
	ExchangeDivision ExchangeDivision `json:"ExchangeDivision"` // 市場
	TickRanking      []TickRanking    `json:"Ranking"`          // TICK回数 ※Type=5
}

// volumeRapidRankingResponse - 売買高急増 ※Type=6
type volumeRapidRankingResponse struct {
	Type               RankingType          `json:"Type"`             // ランキング種別
	ExchangeDivision   ExchangeDivision     `json:"ExchangeDivision"` // 市場
	VolumeRapidRanking []VolumeRapidRanking `json:"Ranking"`          // 売買高急増 ※Type=6
}

// valueRapidRankingResponse - 売買代金急増 ※Type=7
type valueRapidRankingResponse struct {
	Type              RankingType         `json:"Type"`             // ランキング種別
	ExchangeDivision  ExchangeDivision    `json:"ExchangeDivision"` // 市場
	ValueRapidRanking []ValueRapidRanking `json:"Ranking"`          // 売買代金急増 ※Type=7
}

// marginRankingResponse - 信用残 ※Type=8~13
type marginRankingResponse struct {
	Type             RankingType      `json:"Type"`             // ランキング種別
	ExchangeDivision ExchangeDivision `json:"ExchangeDivision"` // 市場
	MarginRanking    []MarginRanking  `json:"Ranking"`          // 信用残 ※Type=8~13
}

// industryPriceRankingResponse - 業種別株価 ※Type=14~15
type industryPriceRankingResponse struct {
	Type                 RankingType            `json:"Type"`             // ランキング種別
	ExchangeDivision     ExchangeDivision       `json:"ExchangeDivision"` // 市場
	IndustryPriceRanking []CategoryPriceRanking `json:"Ranking"`          // 業種別株価 ※Type=14~15
}

// PriceRanking - 株価 ※Type=1~4
type PriceRanking struct {
	No               int          `json:"No"`               // 順位
	Trend            RankingTrend `json:"Trend"`            // トレンド
	AverageRanking   float64      `json:"AverageRanking"`   // 平均順位 ※100位以下は「999」
	Symbol           string       `json:"Symbol"`           // 銘柄コード
	SymbolName       string       `json:"SymbolName"`       // 銘柄名称
	CurrentPrice     float64      `json:"CurrentPrice"`     // 現在値
	ChangeRatio      float64      `json:"ChangeRatio"`      // 前日比
	ChangePercentage float64      `json:"ChangePercentage"` // 騰落率(%)
	CurrentPriceTime HmString     `json:"CurrentPriceTime"` // 時刻(HH:mm)
	TradingVolume    float64      `json:"TradingVolume"`    // 売買高
	Turnover         float64      `json:"Turnover"`         // 売買代金
	ExchangeName     string       `json:"ExchangeName"`     // 市場名
	CategoryName     string       `json:"CategoryName"`     // 業種名
}

// TickRanking - 売買高急増 ※Type=6
type TickRanking struct {
	No               int          `json:"No"`               // 順位
	Trend            RankingTrend `json:"Trend"`            // トレンド
	AverageRanking   float64      `json:"AverageRanking"`   // 平均順位 ※100位以下は「999」
	Symbol           string       `json:"Symbol"`           // 銘柄コード
	SymbolName       string       `json:"SymbolName"`       // 銘柄名称
	CurrentPrice     float64      `json:"CurrentPrice"`     // 現在値
	ChangeRatio      float64      `json:"ChangeRatio"`      // 前日比
	TickCount        int          `json:"TickCount"`        // TICK回数
	UpCount          int          `json:"UpCount"`          // UP
	DownCount        int          `json:"DownCount"`        // Down
	ChangePercentage float64      `json:"ChangePercentage"` // 騰落率(%)
	TradingVolume    float64      `json:"TradingVolume"`    // 売買高
	Turnover         float64      `json:"Turnover"`         // 売買代金
	ExchangeName     string       `json:"ExchangeName"`     // 市場名
	CategoryName     string       `json:"CategoryName"`     // 業種名
}

// VolumeRapidRanking - 売買高急増 ※Type=6
type VolumeRapidRanking struct {
	No                   int          `json:"No"`                   // 順位
	Trend                RankingTrend `json:"Trend"`                // トレンド
	AverageRanking       float64      `json:"AverageRanking"`       // 平均順位 ※100位以下は「999」
	Symbol               string       `json:"Symbol"`               // 銘柄コード
	SymbolName           string       `json:"SymbolName"`           // 銘柄名称
	CurrentPrice         float64      `json:"CurrentPrice"`         // 現在値
	ChangeRatio          float64      `json:"ChangeRatio"`          // 前日比
	RapidTradePercentage float64      `json:"RapidTradePercentage"` // 売買高急増（％）
	TradingVolume        float64      `json:"TradingVolume"`        // 売買高
	CurrentPriceTime     HmString     `json:"CurrentPriceTime"`     // 時刻(HH:mm)
	ChangePercentage     float64      `json:"ChangePercentage"`     // 騰落率(%)
	ExchangeName         string       `json:"ExchangeName"`         // 市場名
	CategoryName         string       `json:"CategoryName"`         // 業種名
}

// ValueRapidRanking - 売買代金急増 ※Type=7
type ValueRapidRanking struct {
	No                     int          `json:"No"`                     // 順位
	Trend                  RankingTrend `json:"Trend"`                  // トレンド
	AverageRanking         float64      `json:"AverageRanking"`         // 平均順位 ※100位以下は「999」
	Symbol                 string       `json:"Symbol"`                 // 銘柄コード
	SymbolName             string       `json:"SymbolName"`             // 銘柄名称
	CurrentPrice           float64      `json:"CurrentPrice"`           // 現在値
	ChangeRatio            float64      `json:"ChangeRatio"`            // 前日比
	RapidPaymentPercentage float64      `json:"RapidPaymentPercentage"` // 代金急増（％）
	Turnover               float64      `json:"Turnover"`               // 売買代金
	CurrentPriceTime       HmString     `json:"CurrentPriceTime"`       // 時刻(HH:mm)
	ChangePercentage       float64      `json:"ChangePercentage"`       // 騰落率(%)
	ExchangeName           string       `json:"ExchangeName"`           // 市場名
	CategoryName           string       `json:"CategoryName"`           // 業種名
}

// MarginRanking - 信用残 ※Type=8~13
type MarginRanking struct {
	No                         int     `json:"No"`                         // 順位
	Symbol                     string  `json:"Symbol"`                     // 銘柄コード
	SymbolName                 string  `json:"SymbolName"`                 // 銘柄名称
	SellRapidPaymentPercentage float64 `json:"SellRapidPaymentPercentage"` // 売残（千株）
	SellLastWeekRatio          float64 `json:"SellLastWeekRatio"`          // 売残前週比
	BuyRapidPaymentPercentage  float64 `json:"BuyRapidPaymentPercentage"`  // 買残（千株）
	BuyLastWeekRatio           float64 `json:"BuyLastWeekRatio"`           // 買残前週比
	Ratio                      float64 `json:"Ratio"`                      // 倍率
	ExchangeName               string  `json:"ExchangeName"`               // 市場名
	CategoryName               string  `json:"CategoryName"`               // 業種名
}

// CategoryPriceRanking - 業種別株価 ※Type=14~15
type CategoryPriceRanking struct {
	No               int          `json:"No"`               // 順位
	Trend            RankingTrend `json:"Trend"`            // トレンド
	AverageRanking   float64      `json:"AverageRanking"`   // 平均順位 ※100位以下は「999」
	Category         string       `json:"Category"`         // 業種コード
	CategoryName     string       `json:"CategoryName"`     // 業種名
	CurrentPrice     float64      `json:"CurrentPrice"`     // 現在値
	ChangeRatio      float64      `json:"ChangeRatio"`      // 前日比
	CurrentPriceTime HmString     `json:"CurrentPriceTime"` // 時刻(HH:mm)
	ChangePercentage float64      `json:"ChangePercentage"` // 騰落率(%)
}

// Ranking - 時価情報・板情報リクエスト
func (c *restClient) Ranking(token string, request RankingRequest) (*RankingResponse, error) {
	return c.RankingWithContext(context.Background(), token, request)
}

// RankingWithContext - 時価情報・板情報リクエスト(contextあり)
func (c *restClient) RankingWithContext(ctx context.Context, token string, request RankingRequest) (*RankingResponse, error) {
	code, b, err := c.get(ctx, token, "ranking", fmt.Sprintf("Type=%s&ExchangeDivision=%s", request.Type, request.ExchangeDivision))
	if err != nil {
		return nil, err
	}

	var response RankingResponse
	switch request.Type {
	case "1", "2", "3", "4":
		var res priceRankingResponse
		if err := parseResponse(code, b, &res); err != nil {
			return nil, err
		}
		response = RankingResponse{
			Type:             res.Type,
			ExchangeDivision: res.ExchangeDivision,
			PriceRanking:     res.PriceRanking,
		}
	case "5":
		var res tickRankingResponse
		if err := parseResponse(code, b, &res); err != nil {
			return nil, err
		}
		response = RankingResponse{
			Type:             res.Type,
			ExchangeDivision: res.ExchangeDivision,
			TickRanking:      res.TickRanking,
		}
	case "6":
		var res volumeRapidRankingResponse
		if err := parseResponse(code, b, &res); err != nil {
			return nil, err
		}
		response = RankingResponse{
			Type:               res.Type,
			ExchangeDivision:   res.ExchangeDivision,
			VolumeRapidRanking: res.VolumeRapidRanking,
		}
	case "7":
		var res valueRapidRankingResponse
		if err := parseResponse(code, b, &res); err != nil {
			return nil, err
		}
		response = RankingResponse{
			Type:              res.Type,
			ExchangeDivision:  res.ExchangeDivision,
			ValueRapidRanking: res.ValueRapidRanking,
		}
	case "8", "9", "10", "11", "12", "13":
		var res marginRankingResponse
		if err := parseResponse(code, b, &res); err != nil {
			return nil, err
		}
		response = RankingResponse{
			Type:             res.Type,
			ExchangeDivision: res.ExchangeDivision,
			MarginRanking:    res.MarginRanking,
		}
	case "14", "15":
		var res industryPriceRankingResponse
		if err := parseResponse(code, b, &res); err != nil {
			return nil, err
		}
		response = RankingResponse{
			Type:                 res.Type,
			ExchangeDivision:     res.ExchangeDivision,
			CategoryPriceRanking: res.IndustryPriceRanking,
		}
	default:
		var res priceRankingResponse
		if err := parseResponse(code, b, &res); err != nil {
			return nil, err
		}
		response = RankingResponse{
			Type:             res.Type,
			ExchangeDivision: res.ExchangeDivision,
			PriceRanking:     res.PriceRanking,
		}
	}

	return &response, nil
}
