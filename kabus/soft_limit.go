package kabus

import (
	"context"
)

// SoftLimitRequest - ソフトリミットのリクエストパラメータ
type SoftLimitRequest struct{}

// SoftLimitResponse - ソフトリミットのレスポンス
type SoftLimitResponse struct {
	Stock        float64 `json:"Stock"`        // 現物のワンショット上限
	Margin       float64 `json:"Margin"`       // 信用のワンショット上限
	Future       float64 `json:"Future"`       // 先物のワンショット上限
	FutureMini   float64 `json:"FutureMini"`   // 先物ミニのワンショット上限
	Option       float64 `json:"Option"`       // オプションのワンショット上限
	KabuSVersion string  `json:"KabuSVersion"` // kabuステーションのバージョン
}

// SoftLimit - ソフトリミットリクエスト
func (c *restClient) SoftLimit(token string, request SoftLimitRequest) (*SoftLimitResponse, error) {
	return c.SoftLimitWithContext(context.Background(), token, request)
}

// SoftLimitWithContext - ソフトリミットリクエスト(contextあり)
func (c *restClient) SoftLimitWithContext(ctx context.Context, token string, _ SoftLimitRequest) (*SoftLimitResponse, error) {
	code, b, err := c.get(ctx, token, "apisoftlimit", "")
	if err != nil {
		return nil, err
	}

	var res SoftLimitResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
