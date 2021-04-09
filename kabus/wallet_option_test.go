package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_restClient_WalletOption(t *testing.T) {
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

			mux := http.NewServeMux()
			mux.HandleFunc("/wallet/option", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			req := &restClient{url: ts.URL}
			got1, got2 := req.WalletOption("")
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_restClient_WalletOptionSymbol(t *testing.T) {
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

			mux := http.NewServeMux()
			mux.HandleFunc("/wallet/option/145124818@2", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			req := &restClient{url: ts.URL}
			got1, got2 := req.WalletOptionSymbol("", WalletOptionSymbolRequest{Symbol: "145124818", Exchange: OptionExchangeAll})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}
