package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_NewWalletMarginRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *walletMarginRequester
	}{
		{name: "本番用のURLが取れる",
			arg1: "token1", arg2: true,
			want: &walletMarginRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/wallet/margin", token: "token1"}}},
		{name: "検証用のURLが取れる",
			arg1: "token2", arg2: false,
			want: &walletMarginRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/wallet/margin", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := NewWalletMarginRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_walletMarginRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *WalletMarginResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body: `{
  "MarginAccountWallet": 0,
  "DepositkeepRate": 0,
  "ConsignmentDepositRate": 0,
  "CashOfConsignmentDepositRate": 0
}`,
			want1: &WalletMarginResponse{MarginAccountWallet: 0, DepositkeepRate: 0, ConsignmentDepositRate: 0, CashOfConsignmentDepositRate: 0},
			want2: nil,
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

			req := &walletMarginRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec()
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_NewWalletMarginSymbolRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *walletMarginSymbolRequester
	}{
		{name: "本番用のURLが取れる",
			arg1: "token1", arg2: true,
			want: &walletMarginSymbolRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/wallet/margin", token: "token1"}}},
		{name: "検証用のURLが取れる",
			arg1: "token2", arg2: false,
			want: &walletMarginSymbolRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/wallet/margin", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := NewWalletMarginSymbolRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_walletMarginSymbolRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *WalletMarginResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body: `{
  "MarginAccountWallet": 0,
  "DepositkeepRate": 0,
  "ConsignmentDepositRate": 30,
  "CashOfConsignmentDepositRate": 0
}`,
			want1: &WalletMarginResponse{MarginAccountWallet: 0, DepositkeepRate: 0, ConsignmentDepositRate: 30, CashOfConsignmentDepositRate: 0},
			want2: nil,
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

			req := &walletMarginSymbolRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec(WalletMarginSymbolRequest{Symbol: "9433", Exchange: StockExchangeToushou})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}
