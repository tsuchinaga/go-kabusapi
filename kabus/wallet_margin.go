package kabus

import (
	"context"
	"fmt"
)

// WalletMarginRequest - 取引余力（信用）のリクエストパラメータ
type WalletMarginRequest struct{}

// WalletMarginSymbolRequest - 取引余力（信用）（銘柄指定）のリクエストパラメータ
type WalletMarginSymbolRequest struct {
	Symbol   string        // 銘柄コード
	Exchange StockExchange // 市場コード
}

// WalletMarginResponse - 取引余力（信用）のレスポンス
type WalletMarginResponse struct {
	MarginAccountWallet          float64 `json:"MarginAccountWallet"`          // 信用買付可能額
	DepositkeepRate              float64 `json:"DepositkeepRate"`              // 保証金維持率
	ConsignmentDepositRate       float64 `json:"ConsignmentDepositRate"`       // 委託保証金率
	CashOfConsignmentDepositRate float64 `json:"CashOfConsignmentDepositRate"` // 現金委託保証金率
}

// walletMarginRequester - 取引余力（信用）リクエスタの生成
func NewWalletMarginRequester(token string, isProd bool) WalletMarginRequester {
	return &walletMarginRequester{httpClient{token: token, url: createURL("/wallet/margin", isProd)}}
}

// WalletMarginRequester - 取引余力（信用）のリクエスタインターフェース
type WalletMarginRequester interface {
	Exec() (*WalletMarginResponse, error)
	ExecWithContext(ctx context.Context) (*WalletMarginResponse, error)
}

// walletMarginRequester - 取引余力（信用）のリクエスタ
type walletMarginRequester struct {
	httpClient
}

// Exec - 取引余力（信用）リクエストの実行
func (r *walletMarginRequester) Exec() (*WalletMarginResponse, error) {
	return r.ExecWithContext(context.Background())
}

// ExecWithContext - 取引余力（信用）リクエストの実行(contextあり)
func (r *walletMarginRequester) ExecWithContext(ctx context.Context) (*WalletMarginResponse, error) {
	code, b, err := r.httpClient.get(ctx, "", "")
	if err != nil {
		return nil, err
	}

	var res WalletMarginResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// NewWalletMarginSymbolRequester - 取引余力（信用）（銘柄指定）リクエスタの生成
func NewWalletMarginSymbolRequester(token string, isProd bool) WalletMarginSymbolRequester {
	return &walletMarginSymbolRequester{httpClient{token: token, url: createURL("/wallet/margin", isProd)}}
}

// WalletMarginSymbolRequester - 取引余力（信用）（銘柄指定）のリクエスタインターフェース
type WalletMarginSymbolRequester interface {
	Exec(request WalletMarginSymbolRequest) (*WalletMarginResponse, error)
	ExecWithContext(ctx context.Context, request WalletMarginSymbolRequest) (*WalletMarginResponse, error)
}

// walletMarginRequester - 取引余力（信用）（銘柄指定）のリクエスタ
type walletMarginSymbolRequester struct {
	httpClient
}

// Exec - 取引余力（信用）（銘柄指定）リクエストの実行
func (r *walletMarginSymbolRequester) Exec(request WalletMarginSymbolRequest) (*WalletMarginResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 取引余力（信用）（銘柄指定）リクエストの実行(contextあり)
func (r *walletMarginSymbolRequester) ExecWithContext(ctx context.Context, request WalletMarginSymbolRequest) (*WalletMarginResponse, error) {
	pathParam := fmt.Sprintf("%s@%d", request.Symbol, request.Exchange)
	code, b, err := r.httpClient.get(ctx, pathParam, "")
	if err != nil {
		return nil, err
	}

	var res WalletMarginResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
