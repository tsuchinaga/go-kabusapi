package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_NewUnregisterRequester(t *testing.T) {
	t.Parallel()

	want := &unregisterRequester{client{token: "token", url: "http://localhost:18080/kabusapi/unregister"}}
	got := NewUnregisterRequester("token")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_unregisterRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *UnregisterResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"RegistList": [{"Symbol": "9433","Exchange": 1}]}`,
			want1:  &UnregisterResponse{RegistList: []RegisteredSymbol{{Symbol: "9433", Exchange: ExchangeToushou}}},
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

			req := &unregisterRequester{client{url: ts.URL}}
			got1, got2 := req.Exec(UnregisterRequest{Symbols: []UnregistSymbol{{Symbol: "9433", Exchange: ExchangeToushou}}})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}
