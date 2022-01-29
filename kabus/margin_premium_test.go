package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_restClient_MarginPremium(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *MarginPremiumResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   marginPremiumBody200,
			want1: &MarginPremiumResponse{
				Symbol: "9433",
				GeneralMargin: MarginPremiumDetail{
					MarginPremiumType:  MarginPremiumTypeUnspecified,
					MarginPremium:      0,
					UpperMarginPremium: 0,
					LowerMarginPremium: 0,
					TickMarginPremium:  0,
				},
				DayTrade: MarginPremiumDetail{
					MarginPremiumType:  MarginPremiumTypeAuction,
					MarginPremium:      0.55,
					UpperMarginPremium: 1,
					LowerMarginPremium: 0.3,
					TickMarginPremium:  0.01,
				},
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
			mux.HandleFunc("/margin/marginpremium/9433", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			req := &restClient{url: ts.URL}
			got1, got2 := req.MarginPremium("", MarginPremiumRequest{Symbol: "9433"})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

const marginPremiumBody200 = `{
  "Symbol": "9433",
  "GeneralMargin": {
    "MarginPremiumType": null,
    "MarginPremium": null,
    "UpperMarginPremium": null,
    "LowerMarginPremium": null,
    "TickMarginPremium": null
  },
  "DayTrade": {
    "MarginPremiumType": 2,
    "MarginPremium": 0.55,
    "UpperMarginPremium": 1,
    "LowerMarginPremium": 0.3,
    "TickMarginPremium": 0.01
  }
}`
