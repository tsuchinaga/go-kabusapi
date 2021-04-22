package kabus

import (
	"context"
	"fmt"
)

// RegulationRequest - 規制情報のリクエストパラメータ
type RegulationRequest struct {
	Symbol   string        // 銘柄
	Exchange StockExchange // 市場
}

// RegulationResponse - 規制情報のレスポンス
type RegulationResponse struct {
	Symbol          string            `json:"Symbol"`          // 銘柄
	RegulationsInfo []RegulationsInfo `json:"RegulationsInfo"` // 規制情報
}

// RegulationsInfo - 規制情報
type RegulationsInfo struct {
	Exchange      RegulationExchange `json:"Exchange"`      // 規制市場
	Product       RegulationProduct  `json:"Product"`       // 規制取引区分
	Side          RegulationSide     `json:"Side"`          // 規制売買
	Reason        string             `json:"Reason"`        // 理由
	LimitStartDay YmdHmString        `json:"LimitStartDay"` // 制限開始日
	LimitEndDay   YmdHmString        `json:"LimitEndDay"`   // 制限終了日
	Level         RegulationLevel    `json:"Level"`         // コンプライアンスレベル
}

// Regulation - 規制情報リクエスト
func (c *restClient) Regulation(token string, request RegulationRequest) (*RegulationResponse, error) {
	return c.RegulationWithContext(context.Background(), token, request)
}

// RegulationWithContext - 規制情報リクエスト(contextあり)
func (c *restClient) RegulationWithContext(ctx context.Context, token string, request RegulationRequest) (*RegulationResponse, error) {
	path := fmt.Sprintf("regulations/%s@%d", request.Symbol, request.Exchange)
	code, b, err := c.get(ctx, token, path, "")
	if err != nil {
		return nil, err
	}

	var res RegulationResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
