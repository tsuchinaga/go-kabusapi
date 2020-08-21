package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_NewSymbolRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *symbolRequester
	}{
		{name: "本番用URLが取れる",
			arg1: "token1", arg2: true,
			want: &symbolRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/symbol", token: "token1"}}},
		{name: "検証用URLが取れる",
			arg1: "token2", arg2: false,
			want: &symbolRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/symbol", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := NewSymbolRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_symbolRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *SymbolResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   symbolBody200,
			want1: &SymbolResponse{
				Symbol:             "9433",
				SymbolName:         "ＫＤＤＩ",
				DisplayName:        "ＫＤＤＩ",
				Exchange:           ExchangeToushou,
				ExchangeName:       "東証１部",
				BisCategory:        "5250",
				TotalMarketValue:   7654484465100,
				TotalStocks:        4484,
				TradingUnit:        100,
				FiscalYearEndBasic: 20210331,
				PriceRangeGroup:    PriceRangeGroup10003,
				KCMarginBuy:        true,
				KCMarginSell:       true,
				MarginBuy:          true,
				MarginSell:         true,
				UpperLimit:         4041,
				LowerLimit:         2641,
			},
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

			req := &symbolRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec(SymbolRequest{Symbol: "9433", Exchange: ExchangeToushou})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

const symbolBody200 = `{
  "Symbol": "9433",
  "SymbolName": "ＫＤＤＩ",
  "DisplayName": "ＫＤＤＩ",
  "Exchange": 1,
  "ExchangeName": "東証１部",
  "BisCategory": "5250",
  "TotalMarketValue": 7654484465100,
  "TotalStocks": 4484,
  "TradingUnit": 100,
  "FiscalYearEndBasic": 20210331,
  "PriceRangeGroup": "10003",
  "KCMarginBuy": true,
  "KCMarginSell": true,
  "MarginBuy": true,
  "MarginSell": true,
  "UpperLimit": 4041,
  "LowerLimit": 2641
}`
