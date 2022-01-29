package kabus

import (
	"context"
	"fmt"
)

// MarginPremiumRequest - プレミアム料取得のリクエストパラメータ
type MarginPremiumRequest struct {
	Symbol string // 銘柄コード
}

// MarginPremiumResponse - プレミアム料取得のレスポンス
type MarginPremiumResponse struct {
	Symbol        string              `json:"Symbol"`        // 銘柄コード
	GeneralMargin MarginPremiumDetail `json:"GeneralMargin"` // 一般信用（長期）
	DayTrade      MarginPremiumDetail `json:"DayTrade"`      // 一般信用（デイトレ）
}

// MarginPremiumDetail - プレミアム料詳細
type MarginPremiumDetail struct {
	MarginPremiumType  MarginPremiumType `json:"MarginPremiumType"`  // プレミアム料入力区分
	MarginPremium      float64           `json:"MarginPremium"`      // 確定プレミアム料
	UpperMarginPremium float64           `json:"UpperMarginPremium"` // 上限プレミアム料
	LowerMarginPremium float64           `json:"LowerMarginPremium"` // 下限プレミアム料
	TickMarginPremium  float64           `json:"TickMarginPremium"`  // プレミアム料刻値
}

// MarginPremium - プレミアム料取得リクエスト
func (c *restClient) MarginPremium(token string, request MarginPremiumRequest) (*MarginPremiumResponse, error) {
	return c.MarginPremiumWithContext(context.Background(), token, request)
}

// MarginPremiumWithContext - プレミアム料取得リクエスト(contextあり)
func (c *restClient) MarginPremiumWithContext(ctx context.Context, token string, request MarginPremiumRequest) (*MarginPremiumResponse, error) {
	path := fmt.Sprintf("margin/marginpremium/%s", request.Symbol)
	code, b, err := c.get(ctx, token, path, "")
	if err != nil {
		return nil, err
	}

	res := MarginPremiumResponse{
		Symbol: "",
		GeneralMargin: MarginPremiumDetail{
			MarginPremiumType:  MarginPremiumTypeUnspecified,
			MarginPremium:      0,
			UpperMarginPremium: 0,
			LowerMarginPremium: 0,
			TickMarginPremium:  0,
		},
		DayTrade: MarginPremiumDetail{
			MarginPremiumType:  MarginPremiumTypeUnspecified,
			MarginPremium:      0,
			UpperMarginPremium: 0,
			LowerMarginPremium: 0,
			TickMarginPremium:  0,
		},
	}
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
