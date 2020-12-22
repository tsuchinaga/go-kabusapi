package kabus

import (
	"context"
	"fmt"
)

// SymbolNameFutureRequest - 先物銘柄コード取得のリクエストパラメータ
type SymbolNameFutureRequest struct {
	FutureCode FutureCode // 先物コード
	DerivMonth YmNUM      // 限月
}

// SymbolNameFutureResponse - 先物銘柄コード取得のレスポンス
type SymbolNameFutureResponse struct {
	Symbol     string `json:"Symbol"`     // 銘柄コード
	SymbolName string `json:"SymbolName"` // 銘柄名称
}

// NewSymbolNameFutureRequester - 先物銘柄コード取得リクエスタの生成
func NewSymbolNameFutureRequester(token string, isProd bool) SymbolNameFutureRequester {
	return &symbolNameFutureRequester{httpClient{token: token, url: createURL("/symbolname/future", isProd)}}
}

// SymbolNameFutureRequester - 先物銘柄コード取得のリクエスタインターフェース
type SymbolNameFutureRequester interface {
	Exec(request SymbolNameFutureRequest) (*SymbolNameFutureResponse, error)
	ExecWithContext(ctx context.Context, request SymbolNameFutureRequest) (*SymbolNameFutureResponse, error)
}

// symbolNameFutureRequester - 先物銘柄コード取得のリクエスタ
type symbolNameFutureRequester struct {
	httpClient
}

// Exec - 先物銘柄コード取得リクエストの実行
func (r *symbolNameFutureRequester) Exec(request SymbolNameFutureRequest) (*SymbolNameFutureResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 先物銘柄コード取得リクエストの実行(contextあり)
func (r *symbolNameFutureRequester) ExecWithContext(ctx context.Context, request SymbolNameFutureRequest) (*SymbolNameFutureResponse, error) {
	queryParam := fmt.Sprintf("FutureCode=%s&DerivMonth=%s", request.FutureCode, request.DerivMonth.String())

	code, b, err := r.httpClient.get(ctx, "", queryParam)
	if err != nil {
		return nil, err
	}

	var res SymbolNameFutureResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
