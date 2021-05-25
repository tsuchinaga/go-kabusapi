package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_restClient_SendOrderFuture(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *SendOrderFutureResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"Result": 0, "OrderId": "20200529A01N06848002"}`,
			want1:  &SendOrderFutureResponse{Result: 0, OrderID: "20200529A01N06848002"},
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
			mux.HandleFunc("/sendorder/future", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			req := &restClient{url: ts.URL}
			got1, got2 := req.SendOrderFuture("", SendOrderFutureRequest{})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_SendOrderFutureRequest_toJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		req  SendOrderFutureRequest
		want string
	}{
		{name: "エントリーならsendOrderFutureEntryRequestにして出す",
			req: SendOrderFutureRequest{
				Password:           "password",
				Symbol:             "165110019",
				Exchange:           FutureExchangeEvening,
				TradeType:          TradeTypeEntry,
				TimeInForce:        TimeInForceFAK,
				Side:               SideBuy,
				Qty:                1,
				ClosePositionOrder: 0,
				ClosePositions:     nil,
				FrontOrderType:     FutureFrontOrderTypeMarket,
				Price:              0,
				ExpireDay:          YmdNUMToday,
				ReverseLimitOrder: &FutureReverseLimitOrder{
					TriggerPrice:      25000,
					UnderOver:         UnderOverUnder,
					AfterHitOrderType: FutureAfterHitOrderTypeMarket,
					AfterHitPrice:     0,
				},
			},
			want: `{"Password":"password","Symbol":"165110019","Exchange":24,"TradeType":1,"TimeInForce":2,"Side":"2","Qty":1,"FrontOrderType":120,"Price":0,"ExpireDay":0,"ReverseLimitOrder":{"TriggerPrice":25000,"UnderOver":1,"AfterHitOrderType":1,"AfterHitPrice":0}}`,
		},
		{name: "エグジットで返済建玉指定があれば決済順序は出さない",
			req: SendOrderFutureRequest{
				Password:           "password",
				Symbol:             "165110019",
				Exchange:           FutureExchangeEvening,
				TradeType:          TradeTypeExit,
				TimeInForce:        TimeInForceFAK,
				Side:               SideSell,
				Qty:                1,
				ClosePositionOrder: ClosePositionOrderUnspecified,
				ClosePositions:     []ClosePosition{{HoldID: "20200903E01N04773904", Qty: 1}},
				FrontOrderType:     FutureFrontOrderTypeMarket,
				Price:              0,
				ExpireDay:          YmdNUMToday,
				ReverseLimitOrder: &FutureReverseLimitOrder{
					TriggerPrice:      25000,
					UnderOver:         UnderOverOver,
					AfterHitOrderType: FutureAfterHitOrderTypeLimit,
					AfterHitPrice:     25000,
				},
			},
			want: `{"Password":"password","Symbol":"165110019","Exchange":24,"TradeType":2,"TimeInForce":2,"Side":"1","Qty":1,"ClosePositions":[{"HoldID":"20200903E01N04773904","Qty":1}],"FrontOrderType":120,"Price":0,"ExpireDay":0,"ReverseLimitOrder":{"TriggerPrice":25000,"UnderOver":2,"AfterHitOrderType":2,"AfterHitPrice":25000}}`,
		},
		{name: "エグジットで返済建玉指定がなければ返済建玉指定は出さない",
			req: SendOrderFutureRequest{
				Password:           "password",
				Symbol:             "165110019",
				Exchange:           FutureExchangeEvening,
				TradeType:          TradeTypeExit,
				TimeInForce:        TimeInForceFAK,
				Side:               SideBuy,
				Qty:                1,
				ClosePositionOrder: ClosePositionOrderDateAscProfitDesc,
				FrontOrderType:     FutureFrontOrderTypeMarket,
				Price:              0,
				ExpireDay:          YmdNUMToday,
			},
			want: `{"Password":"password","Symbol":"165110019","Exchange":24,"TradeType":2,"TimeInForce":2,"Side":"2","Qty":1,"ClosePositionOrder":0,"FrontOrderType":120,"Price":0,"ExpireDay":0,"ReverseLimitOrder":null}`,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := test.req.toJSON()
			if test.want != string(got) || err != nil {
				t.Errorf("%s error\nwant: %s\ngot: %s, %v\n", t.Name(), test.want, string(got), err)
			}
		})
	}
}
