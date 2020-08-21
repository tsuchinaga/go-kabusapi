package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_NewTokenRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 bool
		want *tokenRequester
	}{
		{name: "本番用URLが取れる",
			arg1: true,
			want: &tokenRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/token"}}},
		{name: "検証用URLが取れる",
			arg1: false,
			want: &tokenRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/token"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := NewTokenRequester(test.arg1)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.arg1, got)
			}
		})
	}
}

func Test_tokenRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *TokenResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"ResultCode": 0, "Token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}`,
			want1:  &TokenResponse{ResultCode: 0, Token: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
			want2:  nil,
		},
		{name: "異常レスポンスをパースして返せる",
			status: http.StatusUnauthorized,
			body:   `{"Code": 4001001,"Message": "内部エラー"}`,
			want1:  nil,
			want2: ErrorResponse{
				StatusCode: http.StatusUnauthorized,
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

			req := &tokenRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec(TokenRequest{APIPassword: "xxxxxx"})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}
