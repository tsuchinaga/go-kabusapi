package kabus

import (
	"context"
	"fmt"
)

// WalletOptionRequest - 取引余力（オプション）のリクエストパラメータ
type WalletOptionRequest struct{}

// WalletOptionSymbolRequest - 取引余力（オプション）（銘柄指定）のリクエストパラメータ
type WalletOptionSymbolRequest struct {
	Symbol   string         // 銘柄コード
	Exchange FutureExchange // 市場コード
}

// WalletOptionResponse - 取引余力（オプション）のレスポンス
type WalletOptionResponse struct {
	OptionBuyTradeLimit  float64 `json:"OptionBuyTradeLimit"`  // 買新規建玉可能額
	OptionSellTradeLimit float64 `json:"OptionSellTradeLimit"` // 売新規建玉可能額
	MarginRequirement    float64 `json:"MarginRequirement"`    // 必要証拠金額
}

// walletOptionRequester - 取引余力（オプション）リクエスタの生成
func NewWalletOptionRequester(token string, isProd bool) *walletOptionRequester {
	return &walletOptionRequester{httpClient{token: token, url: createURL("/wallet/option", isProd)}}
}

// walletOptionRequester - 取引余力（オプション）のリクエスタ
type walletOptionRequester struct {
	httpClient
}

// Exec - 取引余力（オプション）リクエストの実行
func (r *walletOptionRequester) Exec() (*WalletOptionResponse, error) {
	return r.ExecWithContext(context.Background())
}

// ExecWithContext - 取引余力（オプション）リクエストの実行(contextあり)
func (r *walletOptionRequester) ExecWithContext(ctx context.Context) (*WalletOptionResponse, error) {
	code, b, err := r.httpClient.get(ctx, "", "")
	if err != nil {
		return nil, err
	}

	var res WalletOptionResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// NewWalletOptionSymbolRequester - 取引余力（オプション）（銘柄指定）リクエスタの生成
func NewWalletOptionSymbolRequester(token string, isProd bool) *walletOptionSymbolRequester {
	return &walletOptionSymbolRequester{httpClient{token: token, url: createURL("/wallet/option", isProd)}}
}

// walletOptionRequester - 取引余力（オプション）（銘柄指定）のリクエスタ
type walletOptionSymbolRequester struct {
	httpClient
}

// Exec - 取引余力（オプション）（銘柄指定）リクエストの実行
func (r *walletOptionSymbolRequester) Exec(request WalletOptionSymbolRequest) (*WalletOptionResponse, error) {
	return r.ExecWithContext(context.Background(), request)
}

// ExecWithContext - 取引余力（オプション）（銘柄指定）リクエストの実行(contextあり)
func (r *walletOptionSymbolRequester) ExecWithContext(ctx context.Context, request WalletOptionSymbolRequest) (*WalletOptionResponse, error) {
	pathParam := fmt.Sprintf("%s@%d", request.Symbol, request.Exchange)
	code, b, err := r.httpClient.get(ctx, pathParam, "")
	if err != nil {
		return nil, err
	}

	var res WalletOptionResponse
	if err := parseResponse(code, b, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
