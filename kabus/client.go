package kabus

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// NewRESTClient - 新しいREST APIのクライアントの生成
func NewRESTClient(isProd bool) RESTClient {
	return &restClient{isProd: isProd, url: baseURL(isProd)}
}

// RESTClient - REST APIのクライアントのインターフェース
type RESTClient interface {
	Token(request TokenRequest) (*TokenResponse, error)                                                                                // トークン発行
	TokenWithContext(ctx context.Context, request TokenRequest) (*TokenResponse, error)                                                // トークン発行(contextあり)
	Register(token string, request RegisterRequest) (*RegisterResponse, error)                                                         // 銘柄登録
	RegisterWithContext(ctx context.Context, token string, request RegisterRequest) (*RegisterResponse, error)                         // 銘柄登録(contextあり)
	Unregister(token string, request UnregisterRequest) (*UnregisterResponse, error)                                                   // 銘柄登録解除
	UnregisterWithContext(ctx context.Context, token string, request UnregisterRequest) (*UnregisterResponse, error)                   // 銘柄登録解除(contextあり)
	UnregisterAll(token string) (*UnregisterAllResponse, error)                                                                        // 銘柄登録全解除
	UnregisterAllWithContext(ctx context.Context, token string) (*UnregisterAllResponse, error)                                        // 銘柄登録全解除(contextあり)
	Board(token string, request BoardRequest) (*BoardResponse, error)                                                                  // 時価情報・板情報
	BoardWithContext(ctx context.Context, token string, request BoardRequest) (*BoardResponse, error)                                  // トークン発行(contextあり)
	Symbol(token string, request SymbolRequest) (*SymbolResponse, error)                                                               // 銘柄情報
	SymbolWithContext(ctx context.Context, token string, request SymbolRequest) (*SymbolResponse, error)                               // 銘柄情報(contextあり)
	Orders(token string, request OrdersRequest) (*OrdersResponse, error)                                                               // 注文約定照会
	OrdersWithContext(ctx context.Context, token string, request OrdersRequest) (*OrdersResponse, error)                               // 注文約定照会(contextあり)
	Positions(token string, request PositionsRequest) (*PositionsResponse, error)                                                      // 残高照会
	PositionsWithContext(ctx context.Context, token string, request PositionsRequest) (*PositionsResponse, error)                      // 残高照会(contextあり)
	SymbolNameFuture(token string, request SymbolNameFutureRequest) (*SymbolNameFutureResponse, error)                                 // 先物銘柄コード取得
	SymbolNameFutureWithContext(ctx context.Context, token string, request SymbolNameFutureRequest) (*SymbolNameFutureResponse, error) // 先物銘柄コード取得(contextあり)
	SymbolNameOption(token string, request SymbolNameOptionRequest) (*SymbolNameOptionResponse, error)                                 // オプション銘柄コード取得
	SymbolNameOptionWithContext(ctx context.Context, token string, request SymbolNameOptionRequest) (*SymbolNameOptionResponse, error) // オプション銘柄コード取得(contextあり)
	SendOrderStock(token string, request SendOrderStockRequest) (*SendOrderStockResponse, error)                                       // 注文発注(現物・信用)
	SendOrderStockWithContext(ctx context.Context, token string, request SendOrderStockRequest) (*SendOrderStockResponse, error)       // 注文発注(現物・信用)(contextあり)
	SendOrderFuture(token string, request SendOrderFutureRequest) (*SendOrderFutureResponse, error)                                    // 注文発注(先物)
	SendOrderFutureWithContext(ctx context.Context, token string, request SendOrderFutureRequest) (*SendOrderFutureResponse, error)    // 注文発注(先物)(contextあり)
	SendOrderOption(token string, request SendOrderOptionRequest) (*SendOrderOptionResponse, error)                                    // 注文発注(オプション)
	SendOrderOptionWithContext(ctx context.Context, token string, request SendOrderOptionRequest) (*SendOrderOptionResponse, error)    // 注文発注(オプション)(contextあり)
	CancelOrder(token string, request CancelOrderRequest) (*CancelOrderResponse, error)                                                // 注文取消
	CancelOrderWithContext(ctx context.Context, token string, request CancelOrderRequest) (*CancelOrderResponse, error)                // 注文取消(contextあり)
	WalletCash(token string) (*WalletCashResponse, error)                                                                              // 取引余力（現物）
	WalletCashWithContext(ctx context.Context, token string) (*WalletCashResponse, error)                                              // 取引余力（現物）(contextあり)
	WalletCashSymbol(token string, request WalletCashSymbolRequest) (*WalletCashResponse, error)                                       // 取引余力（現物）（銘柄指定）
	WalletCashSymbolWithContext(ctx context.Context, token string, request WalletCashSymbolRequest) (*WalletCashResponse, error)       // 取引余力（現物）（銘柄指定）(contextあり)
	WalletMargin(token string) (*WalletMarginResponse, error)                                                                          // 取引余力（信用）
	WalletMarginWithContext(ctx context.Context, token string) (*WalletMarginResponse, error)                                          // 取引余力（信用）(contextあり)
	WalletMarginSymbol(token string, request WalletMarginSymbolRequest) (*WalletMarginResponse, error)                                 // 取引余力（信用）（銘柄指定）
	WalletMarginSymbolWithContext(ctx context.Context, token string, request WalletMarginSymbolRequest) (*WalletMarginResponse, error) // 取引余力（信用）（銘柄指定）(contextあり)
	WalletFuture(token string) (*WalletFutureResponse, error)                                                                          // 取引余力（先物）
	WalletFutureWithContext(ctx context.Context, token string) (*WalletFutureResponse, error)                                          //  取引余力（先物）(contextあり)
	WalletFutureSymbol(token string, request WalletFutureSymbolRequest) (*WalletFutureResponse, error)                                 // 取引余力（先物）（銘柄指定）
	WalletFutureSymbolWithContext(ctx context.Context, token string, request WalletFutureSymbolRequest) (*WalletFutureResponse, error) // 取引余力（先物）（銘柄指定）(contextあり)
	WalletOption(token string) (*WalletOptionResponse, error)                                                                          // 取引余力（オプション）
	WalletOptionWithContext(ctx context.Context, token string) (*WalletOptionResponse, error)                                          // 取引余力（オプション）(contextあり)
	WalletOptionSymbol(token string, request WalletOptionSymbolRequest) (*WalletOptionResponse, error)                                 // 取引余力（オプション）（銘柄指定）
	WalletOptionSymbolWithContext(ctx context.Context, token string, request WalletOptionSymbolRequest) (*WalletOptionResponse, error) // 取引余力（オプション）（銘柄指定）(contextあり)
	Ranking(token string, request RankingRequest) (*RankingResponse, error)                                                            // 詳細ランキング
	RankingWithContext(ctx context.Context, token string, request RankingRequest) (*RankingResponse, error)                            // 詳細ランキング(contextあり)
	Exchange(token string, request ExchangeRequest) (*ExchangeResponse, error)                                                         // 為替情報
	ExchangeWithContext(ctx context.Context, token string, request ExchangeRequest) (*ExchangeResponse, error)                         // 為替情報(contextあり)
	Regulation(token string, request RegulationRequest) (*RegulationResponse, error)                                                   // 規制情報
	RegulationWithContext(ctx context.Context, token string, request RegulationRequest) (*RegulationResponse, error)                   // 規制情報(contextあり)
	PrimaryExchange(token string, request PrimaryExchangeRequest) (*PrimaryExchangeResponse, error)                                    // 優先市場
	PrimaryExchangeWithContext(ctx context.Context, token string, request PrimaryExchangeRequest) (*PrimaryExchangeResponse, error)    // 優先市場(contextあり)
	SoftLimit(token string, request SoftLimitRequest) (*SoftLimitResponse, error)                                                      // ソフトリミット
	SoftLimitWithContext(ctx context.Context, token string, request SoftLimitRequest) (*SoftLimitResponse, error)                      // ソフトリミット(contextあり)
}

// restClient - HTTPクライアント
type restClient struct {
	isProd bool
	url    string
}

// get - GETリクエスト
func (c *restClient) get(ctx context.Context, token string, path string, query string) (int, []byte, error) {
	u, _ := url.Parse(c.url)
	u.Path += "/" + path
	u.Path = strings.ReplaceAll(u.Path, "//", "/")
	u.RawQuery = query

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("X-API-KEY", token)

	// リクエスト送信
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, b, nil
}

// post - POSTリクエスト
func (c *restClient) post(ctx context.Context, token string, path string, request []byte) (int, []byte, error) {
	u, _ := url.Parse(c.url)
	u.Path += "/" + path
	u.Path = strings.ReplaceAll(u.Path, "//", "/")

	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bytes.NewReader(request))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("X-API-KEY", token)

	// リクエスト送信
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, b, nil
}

// put - PUTリクエスト
func (c *restClient) put(ctx context.Context, token string, path string, request []byte) (int, []byte, error) {
	u, _ := url.Parse(c.url)
	u.Path += "/" + path
	u.Path = strings.ReplaceAll(u.Path, "//", "/")

	req, err := http.NewRequestWithContext(ctx, "PUT", u.String(), bytes.NewReader(request))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("X-API-KEY", token)

	// リクエスト送信
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, b, nil
}

// parseResponse - レスポンスをパースする
func parseResponse(code int, body []byte, v interface{}) error {
	if code == http.StatusOK {
		if err := json.Unmarshal(body, v); err != nil {
			return err
		}
		return nil
	} else {
		var errRes ErrorResponse
		if err := json.Unmarshal(body, &errRes); err != nil {
			return err
		}
		errRes.StatusCode = code
		errRes.Body = string(body)
		return errRes
	}
}

// baseURL - リクエスト先の共通URLを生成する
func baseURL(isProd bool) string {
	return "http://" + host(isProd) + "/kabusapi/"
}

// wsClient - WSクライアント
type wsClient struct {
	url    string // URL
	isProd bool   // 本番用か
}

// createWS - リクエスト先のWS URLを生成する
func createWS(isProd bool) string {
	return "ws://" + host(isProd) + "/kabusapi/websocket"
}

// host - 本番か検証のホストを返す
func host(isProd bool) string {
	if isProd {
		return "localhost:18080"
	}
	return "localhost:18081"
}
