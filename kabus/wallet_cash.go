package kabus

import (
	"context"
	"fmt"
)

// WalletCashRequest - 取引余力（現物）のリクエストパラメータ
type WalletCashRequest struct{}

// WalletCashSymbolRequest - 取引余力（現物）（銘柄指定）のリクエストパラメータ
type WalletCashSymbolRequest struct {
	Symbol   string        // 銘柄コード
	Exchange StockExchange // 市場コード
}

// WalletCashResponse - 取引余力（現物）のレスポンス
type WalletCashResponse struct {
	StockAccountWallet float64 `json:"StockAccountWallet"` // 現物買付可能額
}

// walletCashRequester - 取引余力（現物）リクエスタの生成
func NewWalletCashRequester(token string, isProd bool) WalletCashRequester {
	return &walletCashRequester{httpClient{token: token, url: createURL("/wallet/cash", isProd)}}
}

// WalletCashRequester - 取引余力（現物）のリクエスタインターフェース
type WalletCashRequester interface {
	Exec() (*WalletCashResponse, error)
	ExecWithContext(ctx context.Context) (*WalletCashResponse, error)
}

// walletCashRequester - 取引余力（現物）のリクエスタ
type walletCashRequester struct {
	httpClient
}

// Exec - 取引余力（現物）リクエストの実行
func (r *walletCashRequester) Exec() (*WalletCashResponse, error) {
	return r.ExecWithContext(context.Background())
}

// ExecWithContext - 取引余力（現物）リクエストの実行(contextあり)
func (r *walletCashRequester) ExecWithContext(ctx context.Context) (*WalletCashResponse, error) {
	code, b, err := r.httpClient.get(ctx, "", "")
	if err != nil {
		return nil, err
	}

	var res WalletCashResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// NewWalletCashSymbolRequester - 取引余力（現物）（銘柄指定）リクエスタの生成
func NewWalletCashSymbolRequester(token string, isProd bool) WalletCashSymbolRequester {
	return &walletCashSymbolRequester{httpClient{token: token, url: createURL("/wallet/cash", isProd)}}
}

// WalletCashSymbolRequester - 取引余力（現物）（銘柄指定）のリクエスタインターフェース
type WalletCashSymbolRequester interface {
	Exec(request WalletCashSymbolRequest) (*WalletCashResponse, error)
	ExecWithContext(ctx context.Context, request WalletCashSymbolRequest) (*WalletCashResponse, error)
}

// walletCashRequester - 取引余力（現物）（銘柄指定）のリクエスタ
type walletCashSymbolRequester struct {
	httpClient
}

// Exec - 取引余力（現物）（銘柄指定）リクエストの実行
func (r *walletCashSymbolRequester) Exec(request WalletCashSymbolRequest) (*WalletCashResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 取引余力（現物）（銘柄指定）リクエストの実行(contextあり)
func (r *walletCashSymbolRequester) ExecWithContext(ctx context.Context, request WalletCashSymbolRequest) (*WalletCashResponse, error) {
	pathParam := fmt.Sprintf("%s@%d", request.Symbol, request.Exchange)
	code, b, err := r.httpClient.get(ctx, pathParam, "")
	if err != nil {
		return nil, err
	}

	var res WalletCashResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
