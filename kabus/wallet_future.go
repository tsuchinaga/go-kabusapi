package kabus

import (
	"context"
	"fmt"
)

// WalletFutureRequest - 取引余力（先物）のリクエストパラメータ
type WalletFutureRequest struct{}

// WalletFutureSymbolRequest - 取引余力（先物）（銘柄指定）のリクエストパラメータ
type WalletFutureSymbolRequest struct {
	Symbol   string         // 銘柄コード
	Exchange FutureExchange // 市場コード
}

// WalletFutureResponse - 取引余力（先物）のレスポンス
type WalletFutureResponse struct {
	FutureTradeLimit  float64 `json:"FutureTradeLimit"`  // 新規建玉可能額
	MarginRequirement float64 `json:"MarginRequirement"` // 必要証拠金額
}

// walletFutureRequester - 取引余力（先物）リクエスタの生成
func NewWalletFutureRequester(token string, isProd bool) WalletFutureRequester {
	return &walletFutureRequester{httpClient{token: token, url: createURL("/wallet/future", isProd)}}
}

// WalletFutureRequester - 取引余力（先物）のリクエスタインターフェース
type WalletFutureRequester interface {
	Exec() (*WalletFutureResponse, error)
	ExecWithContext(ctx context.Context) (*WalletFutureResponse, error)
}

// walletFutureRequester - 取引余力（先物）のリクエスタ
type walletFutureRequester struct {
	httpClient
}

// Exec - 取引余力（先物）リクエストの実行
func (r *walletFutureRequester) Exec() (*WalletFutureResponse, error) {
	return r.ExecWithContext(context.Background())
}

// ExecWithContext - 取引余力（先物）リクエストの実行(contextあり)
func (r *walletFutureRequester) ExecWithContext(ctx context.Context) (*WalletFutureResponse, error) {
	code, b, err := r.httpClient.get(ctx, "", "")
	if err != nil {
		return nil, err
	}

	var res WalletFutureResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// NewWalletFutureSymbolRequester - 取引余力（先物）（銘柄指定）リクエスタの生成
func NewWalletFutureSymbolRequester(token string, isProd bool) WalletFutureSymbolRequester {
	return &walletFutureSymbolRequester{httpClient{token: token, url: createURL("/wallet/future", isProd)}}
}

// WalletFutureSymbolRequester - 取引余力（先物）（銘柄指定）のリクエスタインターフェース
type WalletFutureSymbolRequester interface {
	Exec(request WalletFutureSymbolRequest) (*WalletFutureResponse, error)
	ExecWithContext(ctx context.Context, request WalletFutureSymbolRequest) (*WalletFutureResponse, error)
}

// walletFutureRequester - 取引余力（先物）（銘柄指定）のリクエスタ
type walletFutureSymbolRequester struct {
	httpClient
}

// Exec - 取引余力（先物）（銘柄指定）リクエストの実行
func (r *walletFutureSymbolRequester) Exec(request WalletFutureSymbolRequest) (*WalletFutureResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 取引余力（先物）（銘柄指定）リクエストの実行(contextあり)
func (r *walletFutureSymbolRequester) ExecWithContext(ctx context.Context, request WalletFutureSymbolRequest) (*WalletFutureResponse, error) {
	pathParam := fmt.Sprintf("%s@%d", request.Symbol, request.Exchange)
	code, b, err := r.httpClient.get(ctx, pathParam, "")
	if err != nil {
		return nil, err
	}

	var res WalletFutureResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
