package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_restClient_SendOrderStock(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *SendOrderStockResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   `{"Result": 0, "OrderId": "20200529A01N06848002"}`,
			want1:  &SendOrderStockResponse{Result: 0, OrderID: "20200529A01N06848002"},
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
			mux.HandleFunc("/sendorder", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			req := &restClient{url: ts.URL}
			got1, got2 := req.SendOrderStock("", SendOrderStockRequest{})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

func Test_SendOrderStockRequest_toJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		req  SendOrderStockRequest
		want string
	}{
		{name: "現物ならそのまま出す",
			req: SendOrderStockRequest{
				Password:           "password",
				Symbol:             "1320",
				Exchange:           StockExchangeToushou,
				SecurityType:       SecurityTypeStock,
				Side:               SideBuy,
				CashMargin:         CashMarginCash,
				MarginTradeType:    MarginTradeTypeSystem,
				DelivType:          DelivTypeCash,
				FundType:           FundTypeTransferMargin,
				AccountType:        AccountTypeGeneral,
				Qty:                1.0,
				ClosePositionOrder: ClosePositionOrderUnspecified,
				ClosePositions:     []ClosePosition{},
				FrontOrderType:     StockFrontOrderTypeMarket,
				Price:              0,
				ExpireDay:          YmdNUM{Time: time.Date(2020, 8, 24, 0, 0, 0, 0, time.Local)},
				ReverseLimitOrder: &StockReverseLimitOrder{
					TriggerSec:        TriggerSecOrderSymbol,
					TriggerPrice:      1000,
					UnderOver:         UnderOverUnder,
					AfterHitOrderType: StockAfterHitOrderTypeMarket,
					AfterHitPrice:     0,
				}},
			want: `{"Password":"password","Symbol":"1320","Exchange":1,"SecurityType":1,"Side":"2","CashMargin":1,"MarginTradeType":1,"DelivType":2,"FundType":"AA","AccountType":2,"Qty":1,"ClosePositionOrder":-1,"ClosePositions":[],"FrontOrderType":10,"Price":0,"ExpireDay":20200824,"ReverseLimitOrder":{"TriggerSec":1,"TriggerPrice":1000,"UnderOver":1,"AfterHitOrderType":1,"AfterHitPrice":0}}`,
		},
		{name: "信用新規で返済建玉指定があれば決済順序は出さない",
			req: SendOrderStockRequest{
				Password:           "password",
				Symbol:             "1320",
				Exchange:           StockExchangeToushou,
				SecurityType:       SecurityTypeStock,
				Side:               SideBuy,
				CashMargin:         CashMarginMarginEntry,
				MarginTradeType:    MarginTradeTypeSystem,
				DelivType:          DelivTypeCash,
				FundType:           FundTypeTransferMargin,
				AccountType:        AccountTypeGeneral,
				Qty:                1.0,
				ClosePositionOrder: ClosePositionOrderUnspecified,
				ClosePositions:     []ClosePosition{{HoldID: "position-id", Qty: 10}},
				FrontOrderType:     StockFrontOrderTypeMarket,
				Price:              0,
				ExpireDay:          YmdNUM{Time: time.Date(2020, 8, 24, 0, 0, 0, 0, time.Local)},
				ReverseLimitOrder: &StockReverseLimitOrder{
					TriggerSec:        TriggerSecOrderN225,
					TriggerPrice:      25000,
					UnderOver:         UnderOverOver,
					AfterHitOrderType: StockAfterHitOrderTypeLimit,
					AfterHitPrice:     25000,
				}},
			want: `{"Password":"password","Symbol":"1320","Exchange":1,"SecurityType":1,"Side":"2","CashMargin":2,"MarginTradeType":1,"DelivType":2,"FundType":"AA","AccountType":2,"Qty":1,"ClosePositions":[{"HoldID":"position-id","Qty":10}],"FrontOrderType":10,"Price":0,"ExpireDay":20200824,"ReverseLimitOrder":{"TriggerSec":2,"TriggerPrice":25000,"UnderOver":2,"AfterHitOrderType":2,"AfterHitPrice":25000}}`,
		},
		{name: "信用返済で返済建玉指定がなければ返済建玉指定は出さない",
			req: SendOrderStockRequest{
				Password:           "password",
				Symbol:             "1320",
				Exchange:           StockExchangeToushou,
				SecurityType:       SecurityTypeStock,
				Side:               SideBuy,
				CashMargin:         CashMarginMarginExit,
				MarginTradeType:    MarginTradeTypeSystem,
				DelivType:          DelivTypeCash,
				FundType:           FundTypeTransferMargin,
				AccountType:        AccountTypeGeneral,
				Qty:                1.0,
				ClosePositionOrder: ClosePositionOrderDateAscProfitDesc,
				ClosePositions:     []ClosePosition{},
				FrontOrderType:     StockFrontOrderTypeMarket,
				Price:              0,
				ExpireDay:          YmdNUM{Time: time.Date(2020, 8, 24, 0, 0, 0, 0, time.Local)}},
			want: `{"Password":"password","Symbol":"1320","Exchange":1,"SecurityType":1,"Side":"2","CashMargin":3,"MarginTradeType":1,"DelivType":2,"FundType":"AA","AccountType":2,"Qty":1,"ClosePositionOrder":0,"FrontOrderType":10,"Price":0,"ExpireDay":20200824,"ReverseLimitOrder":null}`,
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
