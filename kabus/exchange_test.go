package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_restClient_Exchange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *ExchangeResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   exchangeBody200,
			want1: &ExchangeResponse{
				Symbol:   ExchangeSymbolDetailUSDJPY,
				BidPrice: 105.502,
				Spread:   0.2,
				AskPrice: 105.504,
				Change:   -0.055,
				Time:     HmsString{Time: time.Date(0, 1, 1, 16, 10, 45, 0, time.Local)},
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

			mux := http.NewServeMux()
			mux.HandleFunc("/exchange/usdjpy", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			req := &restClient{url: ts.URL}
			got1, got2 := req.Exchange("", ExchangeRequest{Symbol: ExchangeSymbolUSDJPY})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

const exchangeBody200 = `{
  "Symbol": "USD/JPY",
  "BidPrice": 105.502,
  "Spread": 0.2,
  "AskPrice": 105.504,
  "Change": -0.055,
  "Time": "16:10:45"
}`
