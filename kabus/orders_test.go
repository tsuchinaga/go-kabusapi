package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_NewOrdersRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *ordersRequester
	}{
		{name: "本番用URLが取れる",
			arg1: "token1", arg2: true,
			want: &ordersRequester{httpClient: httpClient{url: "http://localhost:18080/kabusapi/orders", token: "token1"}}},
		{name: "検証用URLが取れる",
			arg1: "token2", arg2: false,
			want: &ordersRequester{httpClient: httpClient{url: "http://localhost:18081/kabusapi/orders", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := NewOrdersRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_ordersRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *OrdersResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   ordersBody200,
			want1: &OrdersResponse{
				{
					ID:              "20200715A02N04738436",
					State:           StateDone,
					OrderState:      OrderStateDone,
					OrdType:         OrdTypeInTrading,
					RecvTime:        time.Date(2020, 7, 16, 18, 0, 51, 763683000, time.Local),
					Symbol:          "8306",
					SymbolName:      "三菱ＵＦＪフィナンシャル・グループ",
					Exchange:        StockExchangeToushou,
					ExchangeName:    "東証１部",
					Price:           704.5,
					OrderQty:        1500,
					CumQty:          1500,
					Side:            SideSell,
					CashMargin:      CashMarginMarginEntry,
					AccountType:     AccountTypeSpecific,
					DelivType:       DelivTypeCash,
					ExpireDay:       NewYmdNUM(time.Date(2020, 7, 2, 0, 0, 0, 0, time.Local)),
					MarginTradeType: MarginTradeTypeSystem,
					Details: []OrderDetail{
						{
							SeqNum:        1,
							ID:            "20200715A02N04738436",
							RecType:       RecTypeReceived,
							ExchangeID:    "00000000-0000-0000-0000-00000000",
							State:         StateProcessed,
							TransactTime:  time.Date(2020, 7, 16, 18, 0, 51, 763683000, time.Local),
							OrdType:       OrdTypeInTrading,
							Price:         704.5,
							Qty:           1500,
							ExecutionID:   "",
							ExecutionDay:  time.Date(2020, 7, 2, 18, 2, 0, 0, time.Local),
							DelivDay:      NewYmdNUM(time.Date(2020, 7, 6, 0, 0, 0, 0, time.Local)),
							Commission:    0,
							CommissionTax: 0,
						},
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

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			}))
			defer ts.Close()

			req := &ordersRequester{httpClient{url: ts.URL}}
			got1, got2 := req.Exec(OrdersRequest{Product: ProductAll})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

const ordersBody200 = `[
  {
    "ID": "20200715A02N04738436",
    "State": 5,
    "OrderState": 5,
    "OrdType": 1,
    "RecvTime": "2020-07-16T18:00:51.763683+09:00",
    "Symbol": "8306",
    "SymbolName": "三菱ＵＦＪフィナンシャル・グループ",
    "Exchange": 1,
    "ExchangeName": "東証１部",
    "Price": 704.5,
    "OrderQty": 1500,
    "CumQty": 1500,
    "Side": "1",
    "CashMargin": 2,
    "AccountType": 4,
    "DelivType": 2,
    "ExpireDay": 20200702,
    "MarginTradeType": 1,
    "Details": [
      {
        "SeqNum": 1,
        "ID": "20200715A02N04738436",
        "RecType": 1,
        "ExchangeID": "00000000-0000-0000-0000-00000000",
        "State": 3,
        "TransactTime": "2020-07-16T18:00:51.763683+09:00",
        "OrdType": 1,
        "Price": 704.5,
        "Qty": 1500,
        "ExecutionID": "",
        "ExecutionDay": "2020-07-02T18:02:00+09:00",
        "DelivDay": 20200706,
        "Commission": 0,
        "CommissionTax": 0
      }
    ]
  }
]`

func Test_OrdersRequest_toQuery(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		request OrdersRequest
		want    string
	}{
		{name: "初期値ではproductだけが出る", request: OrdersRequest{}, want: "product=0"},
		{name: "productを指定したら任意のパラメータが出る", request: OrdersRequest{Product: ProductMargin}, want: "product=2"},
		{name: "IDを指定したらidが出る", request: OrdersRequest{ID: "20200715A02N04738436"}, want: "product=0&id=20200715A02N04738436"},
		{name: "UpdateTimeを指定したらupdtimeが出る", request: OrdersRequest{UpdateTime: time.Date(2020, 12, 17, 20, 31, 9, 0, time.Local)}, want: "product=0&updtime=20201217203109"},
		{name: "IsGetOrderDetailを指定したらdetailsが出る", request: OrdersRequest{IsGetOrderDetail: IsGetOrderDetailFalse}, want: "product=0&details=false"},
		{name: "Symbolを指定したらsymbolが出る", request: OrdersRequest{Symbol: "8306"}, want: "product=0&symbol=8306"},
		{name: "Stateを指定したらstateが出る", request: OrdersRequest{State: OrderStateProcessed}, want: "product=0&state=3"},
		{name: "Sideを指定したらsideが出る", request: OrdersRequest{Side: SideBuy}, want: "product=0&side=2"},
		{name: "CashMarginを指定したらcashmarginが出る", request: OrdersRequest{CashMargin: CashMarginMarginEntry}, want: "product=0&cashmargin=2"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.request.toQuery()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
