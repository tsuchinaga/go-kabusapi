package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_NewWalletOptionRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *walletOptionRequester
	}{
		{name: "本番用のURLが取れる",
			arg1: "token1", arg2: true,
			want: &walletOptionRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/wallet/option", token: "token1"}}},
		{name: "検証用のURLが取れる",
			arg1: "token2", arg2: false,
			want: &walletOptionRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/wallet/option", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := NewWalletOptionRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_walletOptionRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *WalletOptionResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"OptionBuyTradeLimit": 300000, "OptionSellTradeLimit": 300000, "MarginRequirement": null}`,
			want1:  &WalletOptionResponse{OptionBuyTradeLimit: 300000, OptionSellTradeLimit: 300000, MarginRequirement: 0},
			want2:  nil,
		},
		{name: "nullを含む正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"OptionBuyTradeLimit": null, "OptionSellTradeLimit": null, "MarginRequirement": null}`,
			want1:  &WalletOptionResponse{OptionBuyTradeLimit: 0, OptionSellTradeLimit: 0, MarginRequirement: 0},
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

			req := &walletOptionRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec()
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_NewWalletOptionSymbolRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *walletOptionSymbolRequester
	}{
		{name: "本番用のURLが取れる",
			arg1: "token1", arg2: true,
			want: &walletOptionSymbolRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/wallet/option", token: "token1"}}},
		{name: "検証用のURLが取れる",
			arg1: "token2", arg2: false,
			want: &walletOptionSymbolRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/wallet/option", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := NewWalletOptionSymbolRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_walletOptionSymbolRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *WalletOptionResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"OptionBuyTradeLimit": 900000, "OptionSellTradeLimit": 900000, "MarginRequirement": 900000}`,
			want1:  &WalletOptionResponse{OptionBuyTradeLimit: 900000, OptionSellTradeLimit: 900000, MarginRequirement: 900000},
			want2:  nil,
		},
		{name: "nullを含む正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"OptionBuyTradeLimit": null, "OptionSellTradeLimit": null, "MarginRequirement": null}`,
			want1:  &WalletOptionResponse{OptionBuyTradeLimit: 0, OptionSellTradeLimit: 0, MarginRequirement: 0},
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

			req := &walletOptionSymbolRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec(WalletOptionSymbolRequest{Symbol: "145124818", Exchange: FutureExchangeAll})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}
