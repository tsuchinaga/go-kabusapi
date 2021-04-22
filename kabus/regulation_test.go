package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_restClient_Regulation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *RegulationResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   regulationBody200,
			want1: &RegulationResponse{
				Symbol: "5614",
				RegulationsInfo: []RegulationsInfo{
					{
						Exchange:      RegulationExchangeToushou,
						Product:       RegulationProductReceipt,
						Side:          RegulationSideBuy,
						Reason:        "品受停止（貸借申込停止銘柄（日証金規制））",
						LimitStartDay: YmdHmString{Time: time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local)},
						LimitEndDay:   YmdHmString{Time: time.Date(2999, 12, 31, 0, 0, 0, 0, time.Local)},
						Level:         RegulationLevelError,
					}, {
						Exchange:      RegulationExchangeUnspecified,
						Product:       RegulationProductCash,
						Side:          RegulationSideBuy,
						Reason:        "その他（代用不適格銘柄）",
						LimitStartDay: YmdHmString{Time: time.Date(2021, 1, 27, 0, 0, 0, 0, time.Local)},
						LimitEndDay:   YmdHmString{Time: time.Date(2021, 2, 17, 0, 0, 0, 0, time.Local)},
						Level:         RegulationLevelError,
					},
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
			mux.HandleFunc("/regulations/5614@1", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			req := &restClient{url: ts.URL}
			got1, got2 := req.Regulation("", RegulationRequest{Symbol: "5614", Exchange: StockExchangeToushou})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

const regulationBody200 = `{
  "Symbol": "5614",
  "RegulationsInfo": [
    {
      "Exchange": 1,
      "Product": 8,
      "Side": 2,
      "Reason": "品受停止（貸借申込停止銘柄（日証金規制））",
      "LimitStartDay": "2020/10/01 00:00",
      "LimitEndDay": "2999/12/31 00:00",
      "Level": 2
    },
    {
      "Exchange": 0,
      "Product": 1,
      "Side": 2,
      "Reason": "その他（代用不適格銘柄）",
      "LimitStartDay": "2021/01/27 00:00",
      "LimitEndDay": "2021/02/17 00:00",
      "Level": 2
    }
  ]
}`
