package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_NewSendOrderRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *sendOrderRequester
	}{
		{name: "本番用URLが取れる",
			arg1: "token1", arg2: true,
			want: &sendOrderRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/sendorder", token: "token1"}}},
		{name: "検証用URLが取れる",
			arg1: "token2", arg2: false,
			want: &sendOrderRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/sendorder", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := NewSendOrderRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_sendOrderRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *SendOrderResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"Result": 0, "OrderId": "20200529A01N06848002"}`,
			want1:  &SendOrderResponse{Result: 0, OrderID: "20200529A01N06848002"},
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

			req := &sendOrderRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec(SendOrderRequest{})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_SendOrderRequest_toJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		req  SendOrderRequest
		want string
	}{
		{name: "現物ならそのまま出す",
			req: SendOrderRequest{
				Password:           "password",
				Symbol:             "1320",
				Exchange:           ExchangeToushou,
				SecurityType:       SecurityTypeKabu,
				Side:               SideBuy,
				CashMargin:         CashMarginCash,
				MarginTradeType:    MarginTradeTypeSystem,
				DelivType:          DelivTypeCash,
				FundType:           FundTypeTransferMargin,
				AccountType:        AccountTypeGeneral,
				Qty:                1.0,
				ClosePositionOrder: ClosePositionOrderUnspecified,
				ClosePositions:     []ClosePosition{},
				Price:              0,
				ExpireDay:          YmdNUM{Time: time.Date(2020, 8, 24, 0, 0, 0, 0, time.Local)},
				FrontOrderType:     FrontOrderTypeMarket},
			want: `{"Password":"password","Symbol":"1320","Exchange":1,"SecurityType":1,"Side":"2","CashMargin":1,"MarginTradeType":1,"DelivType":2,"FundType":"AA","AccountType":2,"Qty":1,"ClosePositionOrder":-1,"ClosePositions":[],"Price":0,"ExpireDay":20200824,"FrontOrderType":10}`,
		},
		{name: "信用新規で返済建玉指定があれば決済順序は出さない",
			req: SendOrderRequest{
				Password:           "password",
				Symbol:             "1320",
				Exchange:           ExchangeToushou,
				SecurityType:       SecurityTypeKabu,
				Side:               SideBuy,
				CashMargin:         CashMarginMarginEntry,
				MarginTradeType:    MarginTradeTypeSystem,
				DelivType:          DelivTypeCash,
				FundType:           FundTypeTransferMargin,
				AccountType:        AccountTypeGeneral,
				Qty:                1.0,
				ClosePositionOrder: ClosePositionOrderUnspecified,
				ClosePositions:     []ClosePosition{{HoldID: "position-id", Qty: 10}},
				Price:              0,
				ExpireDay:          YmdNUM{Time: time.Date(2020, 8, 24, 0, 0, 0, 0, time.Local)},
				FrontOrderType:     FrontOrderTypeMarket},
			want: `{"Password":"password","Symbol":"1320","Exchange":1,"SecurityType":1,"Side":"2","CashMargin":2,"MarginTradeType":1,"DelivType":2,"FundType":"AA","AccountType":2,"Qty":1,"ClosePositions":[{"HoldID":"position-id","Qty":10}],"Price":0,"ExpireDay":20200824,"FrontOrderType":10}`,
		},
		{name: "信用返済で返済建玉指定がなければ返済建玉指定は出さない",
			req: SendOrderRequest{
				Password:           "password",
				Symbol:             "1320",
				Exchange:           ExchangeToushou,
				SecurityType:       SecurityTypeKabu,
				Side:               SideBuy,
				CashMargin:         CashMarginMarginExit,
				MarginTradeType:    MarginTradeTypeSystem,
				DelivType:          DelivTypeCash,
				FundType:           FundTypeTransferMargin,
				AccountType:        AccountTypeGeneral,
				Qty:                1.0,
				ClosePositionOrder: ClosePositionOrderDateAscProfitDesc,
				ClosePositions:     []ClosePosition{},
				Price:              0,
				ExpireDay:          YmdNUM{Time: time.Date(2020, 8, 24, 0, 0, 0, 0, time.Local)},
				FrontOrderType:     FrontOrderTypeMarket},
			want: `{"Password":"password","Symbol":"1320","Exchange":1,"SecurityType":1,"Side":"2","CashMargin":3,"MarginTradeType":1,"DelivType":2,"FundType":"AA","AccountType":2,"Qty":1,"ClosePositionOrder":0,"Price":0,"ExpireDay":20200824,"FrontOrderType":10}`,
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
