package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_NewWalletFutureRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *walletFutureRequester
	}{
		{name: "本番用のURLが取れる",
			arg1: "token1", arg2: true,
			want: &walletFutureRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/wallet/future", token: "token1"}}},
		{name: "検証用のURLが取れる",
			arg1: "token2", arg2: false,
			want: &walletFutureRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/wallet/future", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := NewWalletFutureRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_walletFutureRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *WalletFutureResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"FutureTradeLimit": 300000, "MarginRequirement": null}`,
			want1:  &WalletFutureResponse{FutureTradeLimit: 300000, MarginRequirement: 0},
			want2:  nil,
		},
		{name: "nullを含む正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"FutureTradeLimit": null, "MarginRequirement": null}`,
			want1:  &WalletFutureResponse{FutureTradeLimit: 0, MarginRequirement: 0},
			want2:  nil,
		},
		{name: "異常レスポンスをパースして返せる",
			status: http.StatusBadRequest,
			body:   `{"Code": 4001001,"Message": "内部エラー"}`,
			want1:  nil,
			want2: ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `{"Code": 4001001,"Message": "内部エラー"}`,
				Code:       4001001,
				Message:    "内部エラー",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			}))
			defer ts.Close()

			req := &walletFutureRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec()
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_NewWalletFutureSymbolRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *walletFutureSymbolRequester
	}{
		{name: "本番用のURLが取れる",
			arg1: "token1", arg2: true,
			want: &walletFutureSymbolRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/wallet/future", token: "token1"}}},
		{name: "検証用のURLが取れる",
			arg1: "token2", arg2: false,
			want: &walletFutureSymbolRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/wallet/future", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := NewWalletFutureSymbolRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_walletFutureSymbolRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *WalletFutureResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"FutureTradeLimit": 900000, "MarginRequirement": 300000}`,
			want1:  &WalletFutureResponse{FutureTradeLimit: 900000, MarginRequirement: 300000},
			want2:  nil,
		},
		{name: "nullを含む正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"FutureTradeLimit": null, "MarginRequirement": null}`,
			want1:  &WalletFutureResponse{FutureTradeLimit: 0, MarginRequirement: 0},
			want2:  nil,
		},
		{name: "異常レスポンスをパースして返せる",
			status: http.StatusBadRequest,
			body:   `{"Code": 4001001,"Message": "内部エラー"}`,
			want1:  nil,
			want2: ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Body:       `{"Code": 4001001,"Message": "内部エラー"}`,
				Code:       4001001,
				Message:    "内部エラー",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			}))
			defer ts.Close()

			req := &walletFutureSymbolRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec(WalletFutureSymbolRequest{Symbol: "165120018", Exchange: FutureExchangeAll})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}
