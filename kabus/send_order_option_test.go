package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_NewSendOrderOptionRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *sendOrderOptionRequester
	}{
		{name: "本番用URLが取れる",
			arg1: "token1", arg2: true,
			want: &sendOrderOptionRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/sendorder/option", token: "token1"}}},
		{name: "検証用URLが取れる",
			arg1: "token2", arg2: false,
			want: &sendOrderOptionRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/sendorder/option", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := NewSendOrderOptionRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_sendOrderOptionRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *SendOrderOptionResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"Result": 0, "OrderId": "20200529A01N06848002"}`,
			want1:  &SendOrderOptionResponse{Result: 0, OrderID: "20200529A01N06848002"},
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

			req := &sendOrderOptionRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec(SendOrderOptionRequest{})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_SendOrderOptionRequest_toJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		req  SendOrderOptionRequest
		want string
	}{
		{name: "エントリーならsendOrderOptionEntryRequestにして出す",
			req: SendOrderOptionRequest{
				Password:           "password",
				Symbol:             "135113218",
				Exchange:           OptionExchangeEvening,
				TradeType:          TradeTypeEntry,
				TimeInForce:        TimeInForceFAK,
				Side:               SideBuy,
				Qty:                1,
				ClosePositionOrder: 0,
				ClosePositions:     nil,
				FrontOrderType:     OptionFrontOrderTypeMarket,
				Price:              0,
				ExpireDay:          YmdNUMToday,
			},
			want: `{"Password":"password","Symbol":"135113218","Exchange":24,"TradeType":1,"TimeInForce":2,"Side":"2","Qty":1,"FrontOrderType":120,"Price":0,"ExpireDay":0}`,
		},
		{name: "エグジットで返済建玉指定があれば決済順序は出さない",
			req: SendOrderOptionRequest{
				Password:           "password",
				Symbol:             "135113218",
				Exchange:           OptionExchangeEvening,
				TradeType:          TradeTypeExit,
				TimeInForce:        TimeInForceFAK,
				Side:               SideSell,
				Qty:                1,
				ClosePositionOrder: ClosePositionOrderUnspecified,
				ClosePositions:     []ClosePosition{{HoldID: "20200903E01N04773904", Qty: 1}},
				FrontOrderType:     OptionFrontOrderTypeMarket,
				Price:              0,
				ExpireDay:          YmdNUMToday,
			},
			want: `{"Password":"password","Symbol":"135113218","Exchange":24,"TradeType":2,"TimeInForce":2,"Side":"1","Qty":1,"ClosePositions":[{"HoldID":"20200903E01N04773904","Qty":1}],"FrontOrderType":120,"Price":0,"ExpireDay":0}`,
		},
		{name: "エグジットで返済建玉指定がなければ返済建玉指定は出さない",
			req: SendOrderOptionRequest{
				Password:           "password",
				Symbol:             "135113218",
				Exchange:           OptionExchangeEvening,
				TradeType:          TradeTypeExit,
				TimeInForce:        TimeInForceFAK,
				Side:               SideBuy,
				Qty:                1,
				ClosePositionOrder: ClosePositionOrderDateAscProfitDesc,
				FrontOrderType:     OptionFrontOrderTypeMarket,
				Price:              0,
				ExpireDay:          YmdNUMToday,
			},
			want: `{"Password":"password","Symbol":"135113218","Exchange":24,"TradeType":2,"TimeInForce":2,"Side":"2","Qty":1,"ClosePositionOrder":0,"FrontOrderType":120,"Price":0,"ExpireDay":0}`,
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
