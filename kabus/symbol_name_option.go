package kabus

import (
	"context"
	"fmt"
)

// SymbolNameOptionRequest - オプション銘柄コード取得のリクエストパラメータ
type SymbolNameOptionRequest struct {
	DerivMonth  YmNUM     // 限月
	PutOrCall   PutOrCall // コール or プット
	StrikePrice int       // 権利行使価格
}

// SymbolNameOptionResponse - オプション銘柄コード取得のレスポンス
type SymbolNameOptionResponse struct {
	Symbol     string `json:"Symbol"`     // 銘柄コード
	SymbolName string `json:"SymbolName"` // 銘柄名称
}

// NewSymbolNameOptionRequester - オプション銘柄コード取得リクエスタの生成
func NewSymbolNameOptionRequester(token string, isProd bool) SymbolNameOptionRequester {
	return &symbolNameOptionRequester{httpClient{token: token, url: createURL("/symbolname/option", isProd)}}
}

// SymbolNameOptionRequester - オプション銘柄コード取得のリクエスタインターフェース
type SymbolNameOptionRequester interface {
	Exec(request SymbolNameOptionRequest) (*SymbolNameOptionResponse, error)
	ExecWithContext(ctx context.Context, request SymbolNameOptionRequest) (*SymbolNameOptionResponse, error)
}

// symbolNameOptionRequester - オプション銘柄コード取得のリクエスタ
type symbolNameOptionRequester struct {
	httpClient
}

// Exec - オプション銘柄コード取得リクエストの実行
func (r *symbolNameOptionRequester) Exec(request SymbolNameOptionRequest) (*SymbolNameOptionResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - オプション銘柄コード取得リクエストの実行(contextあり)
func (r *symbolNameOptionRequester) ExecWithContext(ctx context.Context, request SymbolNameOptionRequest) (*SymbolNameOptionResponse, error) {
	queryParam := fmt.Sprintf("DerivMonth=%s&PutOrCall=%s&StrikePrice=%d", request.DerivMonth.String(), request.PutOrCall, request.StrikePrice)

	code, b, err := r.httpClient.get(ctx, "", queryParam)
	if err != nil {
		return nil, err
	}

	var res SymbolNameOptionResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
