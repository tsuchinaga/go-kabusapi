package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_NewRegisterRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *registerRequester
	}{
		{name: "本番用URLが取れる",
			arg1: "token1", arg2: true,
			want: &registerRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/register", token: "token1"}}},
		{name: "検証用URLが取れる",
			arg1: "token2", arg2: false,
			want: &registerRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/register", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := NewRegisterRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_registerRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *RegisterResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"RegistList": [{"Symbol": "9433","Exchange": 1}]}`,
			want1:  &RegisterResponse{RegisterList: []RegisteredSymbol{{Symbol: "9433", Exchange: StockExchangeToushou}}},
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

			req := &registerRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec(RegisterRequest{Symbols: []RegisterSymbol{{Symbol: "9433", Exchange: ExchangeToushou}}})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}
