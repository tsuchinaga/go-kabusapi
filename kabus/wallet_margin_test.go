package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_restClient_WalletMargin(t *testing.T) {
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

			mux := http.NewServeMux()
			mux.HandleFunc("/wallet/margin", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			req := &restClient{url: ts.URL}
			got1, got2 := req.WalletMargin("")
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_restClient_WalletMarginSymbol(t *testing.T) {
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

			mux := http.NewServeMux()
			mux.HandleFunc("/wallet/margin/9433@1", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			req := &restClient{url: ts.URL}
			got1, got2 := req.WalletMarginSymbol("", WalletMarginSymbolRequest{Symbol: "9433", Exchange: StockExchangeToushou})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}
