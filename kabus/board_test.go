package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_NewBoardRequester(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg1 string
		arg2 bool
		want *boardRequester
	}{
		{name: "本番用URLが取れる",
			arg1: "token1", arg2: true,
			want: &boardRequester{client{url: "http://localhost:18080/kabusapi/board", token: "token1"}}},
		{name: "検証用URLが取れる",
			arg1: "token2", arg2: false,
			want: &boardRequester{client{url: "http://localhost:18081/kabusapi/board", token: "token2"}}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := NewBoardRequester(test.arg1, test.arg2)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_boardRequester_Exec(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		want1  *BoardResponse
		want2  error
	}{
		{name: "正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   boardBody200,
			want1: &BoardResponse{
				Symbol:                   "5401",
				SymbolName:               "新日鐵住金",
				Exchange:                 ExchangeToushou,
				ExchangeName:             "東証１部",
				CurrentPrice:             2408,
				CurrentPriceTime:         time.Date(2020, 7, 22, 15, 0, 0, 0, time.Local),
				CurrentPriceChangeStatus: CurrentPriceChangeStatusDown,
				CurrentPriceStatus:       CurrentPriceStatusCurrentPrice,
				CalcPrice:                343.7,
				PreviousClose:            1048,
				PreviousCloseTime:        YmdTHms{Time: time.Date(2020, 7, 21, 0, 0, 0, 0, time.Local)},
				ChangePreviousClose:      1360,
				ChangePreviousClosePer:   129.77,
				OpeningPrice:             2380,
				OpeningPriceTime:         time.Date(2020, 7, 22, 9, 0, 0, 0, time.Local),
				HighPrice:                2418,
				HighPriceTime:            time.Date(2020, 7, 22, 13, 25, 47, 0, time.Local),
				LowPrice:                 2370,
				LowPriceTime:             time.Date(2020, 7, 22, 10, 0, 4, 0, time.Local),
				TradingVolume:            4571500,
				TradingVolumeTime:        time.Date(2020, 7, 22, 15, 0, 0, 0, time.Local),
				VWAP:                     2394.4262,
				TradingValue:             10946119350,
				BidQty:                   100,
				BidPrice:                 2408.5,
				BidTime:                  time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local),
				BidSign:                  BidAskSignGeneral,
				MarketOrderSellQty:       0,
				Sell1:                    FirstBoardSign{Time: time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local), Sign: BidAskSignGeneral, Price: 2408.5, Qty: 100},
				Sell2:                    BoardSign{Price: 2409, Qty: 800},
				Sell3:                    BoardSign{Price: 2409.5, Qty: 2100},
				Sell4:                    BoardSign{Price: 2410, Qty: 800},
				Sell5:                    BoardSign{Price: 2410.5, Qty: 500},
				Sell6:                    BoardSign{Price: 2411, Qty: 8400},
				Sell7:                    BoardSign{Price: 2411.5, Qty: 1200},
				Sell8:                    BoardSign{Price: 2412, Qty: 27200},
				Sell9:                    BoardSign{Price: 2412.5, Qty: 400},
				Sell10:                   BoardSign{Price: 2413, Qty: 16400},
				AskQty:                   200,
				AskPrice:                 2407.5,
				AskTime:                  time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local),
				AskSign:                  BidAskSignGeneral,
				MarketOrderBuyQty:        0,
				Buy1:                     FirstBoardSign{Time: time.Date(2020, 7, 22, 14, 59, 59, 0, time.Local), Sign: BidAskSignGeneral, Price: 2407.5, Qty: 200},
				Buy2:                     BoardSign{Price: 2407, Qty: 400},
				Buy3:                     BoardSign{Price: 2406.5, Qty: 1000},
				Buy4:                     BoardSign{Price: 2406, Qty: 5800},
				Buy5:                     BoardSign{Price: 2405.5, Qty: 7500},
				Buy6:                     BoardSign{Price: 2405, Qty: 2200},
				Buy7:                     BoardSign{Price: 2404.5, Qty: 16700},
				Buy8:                     BoardSign{Price: 2404, Qty: 30100},
				Buy9:                     BoardSign{Price: 2403.5, Qty: 1300},
				Buy10:                    BoardSign{Price: 2403, Qty: 3000},
				OverSellQty:              974900,
				UnderBuyQty:              756000,
				TotalMarketValue:         3266254659361.4,
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

			req := &boardRequester{client{url: ts.URL}}
			got1, got2 := req.Exec(BoardRequest{Symbol: "5401", Exchange: ExchangeToushou})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

const boardBody200 = `{
  "Symbol": "5401",
  "SymbolName": "新日鐵住金",
  "Exchange": 1,
  "ExchangeName": "東証１部",
  "CurrentPrice": 2408,
  "CurrentPriceTime": "2020-07-22T15:00:00+09:00",
  "CurrentPriceChangeStatus": "0058",
  "CurrentPriceStatus": 1,
  "CalcPrice": 343.7,
  "PreviousClose": 1048,
  "PreviousCloseTime": "2020-07-21T00:00:00",
  "ChangePreviousClose": 1360,
  "ChangePreviousClosePer": 129.77,
  "OpeningPrice": 2380,
  "OpeningPriceTime": "2020-07-22T09:00:00+09:00",
  "HighPrice": 2418,
  "HighPriceTime": "2020-07-22T13:25:47+09:00",
  "LowPrice": 2370,
  "LowPriceTime": "2020-07-22T10:00:04+09:00",
  "TradingVolume": 4571500,
  "TradingVolumeTime": "2020-07-22T15:00:00+09:00",
  "VWAP": 2394.4262,
  "TradingValue": 10946119350,
  "BidQty": 100,
  "BidPrice": 2408.5,
  "BidTime": "2020-07-22T14:59:59+09:00",
  "BidSign": "0101",
  "MarketOrderSellQty": 0,
  "Sell1": {
    "Time": "2020-07-22T14:59:59+09:00",
    "Sign": "0101",
    "Price": 2408.5,
    "Qty": 100
  },
  "Sell2": {
    "Price": 2409,
    "Qty": 800
  },
  "Sell3": {
    "Price": 2409.5,
    "Qty": 2100
  },
  "Sell4": {
    "Price": 2410,
    "Qty": 800
  },
  "Sell5": {
    "Price": 2410.5,
    "Qty": 500
  },
  "Sell6": {
    "Price": 2411,
    "Qty": 8400
  },
  "Sell7": {
    "Price": 2411.5,
    "Qty": 1200
  },
  "Sell8": {
    "Price": 2412,
    "Qty": 27200
  },
  "Sell9": {
    "Price": 2412.5,
    "Qty": 400
  },
  "Sell10": {
    "Price": 2413,
    "Qty": 16400
  },
  "AskQty": 200,
  "AskPrice": 2407.5,
  "AskTime": "2020-07-22T14:59:59+09:00",
  "AskSign": "0101",
  "MarketOrderBuyQty": 0,
  "Buy1": {
    "Time": "2020-07-22T14:59:59+09:00",
    "Sign": "0101",
    "Price": 2407.5,
    "Qty": 200
  },
  "Buy2": {
    "Price": 2407,
    "Qty": 400
  },
  "Buy3": {
    "Price": 2406.5,
    "Qty": 1000
  },
  "Buy4": {
    "Price": 2406,
    "Qty": 5800
  },
  "Buy5": {
    "Price": 2405.5,
    "Qty": 7500
  },
  "Buy6": {
    "Price": 2405,
    "Qty": 2200
  },
  "Buy7": {
    "Price": 2404.5,
    "Qty": 16700
  },
  "Buy8": {
    "Price": 2404,
    "Qty": 30100
  },
  "Buy9": {
    "Price": 2403.5,
    "Qty": 1300
  },
  "Buy10": {
    "Price": 2403,
    "Qty": 3000
  },
  "OverSellQty": 974900,
  "UnderBuyQty": 756000,
  "TotalMarketValue": 3266254659361.4
}`
