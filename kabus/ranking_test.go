package kabus

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_restClient_Ranking(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status int
		body   string
		arg1   RankingType
		arg2   ExchangeDivision
		want1  *RankingResponse
		want2  error
	}{
		{name: "Type=1の正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   rankingBodyType1,
			arg1:   RankingTypePriceIncreaseRate,
			arg2:   ExchangeDivisionALL,
			want1: &RankingResponse{Type: RankingTypePriceIncreaseRate, ExchangeDivision: ExchangeDivisionToushou, PriceRanking: []PriceRanking{
				{No: 1, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "1689", SymbolName: "ガスETF/ETF(C)", CurrentPrice: 2, ChangeRatio: 1, ChangePercentage: 100, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, TradingVolume: 5722.4, Turnover: 10.4136, ExchangeName: "東証ETF/ETN", CategoryName: "その他"},
				{No: 2, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, ChangePercentage: 54.65, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, TradingVolume: 3117.5, Turnover: 3194.7121, ExchangeName: "東証JQS", CategoryName: "電気機器"},
				{No: 3, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "8143", SymbolName: "ラピーヌ", CurrentPrice: 430, ChangeRatio: 80, ChangePercentage: 22.85, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, TradingVolume: 627.8, Turnover: 262.624, ExchangeName: "東証２部", CategoryName: "繊維製品"},
				{No: 4, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "4549", SymbolName: "栄研化", CurrentPrice: 2435, ChangeRatio: 398, ChangePercentage: 19.53, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, TradingVolume: 2077.7, Turnover: 4900.3547, ExchangeName: "東証１部", CategoryName: "医薬品"},
				{No: 5, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "4575", SymbolName: "CANBAS", CurrentPrice: 600, ChangeRatio: 97, ChangePercentage: 19.28, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, TradingVolume: 579.6, Turnover: 315.8453, ExchangeName: "東証ﾏｻﾞｰｽﾞ", CategoryName: "医薬品"},
			}},
			want2: nil,
		},
		{name: "Type=5の正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   rankingBodyType5,
			arg1:   RankingTypeTickCount,
			arg2:   ExchangeDivisionALL,
			want1: &RankingResponse{Type: RankingTypeTickCount, ExchangeDivision: ExchangeDivisionToushou,
				TickRanking: []TickRanking{
					{No: 1, Trend: RankingTrendRiseOver20, AverageRanking: 22, Symbol: "2929", SymbolName: "ﾌｧｰﾏﾌｰｽﾞ", CurrentPrice: 2748, ChangeRatio: 99, TickCount: 40579, UpCount: 12722, DownCount: 12798, ChangePercentage: 3.73, TradingVolume: 16086.8, Turnover: 43810.0498, ExchangeName: "東証２部", CategoryName: "食料品"},
					{No: 2, Trend: RankingTrendUnchanged, AverageRanking: 2, Symbol: "9984", SymbolName: "ｿﾌﾄﾊﾞﾝｸG", CurrentPrice: 8285, ChangeRatio: -309, TickCount: 32219, UpCount: 8655, DownCount: 8562, ChangePercentage: -3.59, TradingVolume: 16688.8, Turnover: 138143.1773, ExchangeName: "東証１部", CategoryName: "情報・通信業"},
					{No: 3, Trend: RankingTrendRiseOver20, AverageRanking: 31, Symbol: "7751", SymbolName: "キヤノン", CurrentPrice: 2476.5, ChangeRatio: -1.5, TickCount: 20875, UpCount: 5702, DownCount: 5642, ChangePercentage: -0.06, TradingVolume: 13403.1, Turnover: 33139.97, ExchangeName: "東証１部", CategoryName: "電気機器"},
					{No: 4, Trend: RankingTrendRiseOver20, AverageRanking: 41, Symbol: "2413", SymbolName: "ｴﾑｽﾘｰ", CurrentPrice: 9125, ChangeRatio: -422, TickCount: 19492, UpCount: 6690, DownCount: 6963, ChangePercentage: -4.42, TradingVolume: 11325.1, Turnover: 103804.0663, ExchangeName: "東証１部", CategoryName: "サービス業"},
					{No: 5, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "4552", SymbolName: "JCRﾌｧｰﾏ", CurrentPrice: 3115, ChangeRatio: 373, TickCount: 18514, UpCount: 5040, DownCount: 4847, ChangePercentage: 13.6, TradingVolume: 7019.8, Turnover: 21201.682, ExchangeName: "東証１部", CategoryName: "医薬品"},
				}},
			want2: nil,
		},
		{name: "Type=6の正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   rankingBodyType6,
			arg1:   RankingTypeVolumeRapidIncrease,
			arg2:   ExchangeDivisionALL,
			want1: &RankingResponse{Type: RankingTypeVolumeRapidIncrease, ExchangeDivision: ExchangeDivisionToushou,
				VolumeRapidRanking: []VolumeRapidRanking{
					{No: 1, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "1490", SymbolName: "上場ﾍﾞｰﾀ/ETF", CurrentPrice: 7750, ChangeRatio: 40, RapidTradePercentage: 49900, TradingVolume: 1, CurrentPriceTime: HmString{time.Date(0, 1, 1, 13, 20, 0, 0, time.Local)}, ChangePercentage: 0.51, ExchangeName: "東証ETF/ETN", CategoryName: "その他"},
					{No: 2, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, RapidTradePercentage: 28189.47, TradingVolume: 3117.5, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 54.65, ExchangeName: "東証JQS", CategoryName: "電気機器"},
					{No: 3, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "5010", SymbolName: "日精蝋", CurrentPrice: 194, ChangeRatio: 7, RapidTradePercentage: 11951.4, TradingVolume: 1453.4, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 3.74, ExchangeName: "東証２部", CategoryName: "石油・石炭製品"},
					{No: 4, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "8143", SymbolName: "ラピーヌ", CurrentPrice: 430, ChangeRatio: 80, RapidTradePercentage: 9648.44, TradingVolume: 627.8, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 22.85, ExchangeName: "東証２部", CategoryName: "繊維製品"},
					{No: 5, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "8184", SymbolName: "島忠", CurrentPrice: 5470, ChangeRatio: -10, RapidTradePercentage: 2520.35, TradingVolume: 1382.5, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: -0.18, ExchangeName: "東証監理", CategoryName: "小売業"},
				}},
			want2: nil,
		},
		{name: "Type=7の正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   rankingBodyType7,
			arg1:   RankingTypeValueRapidIncrease,
			arg2:   ExchangeDivisionALL,
			want1: &RankingResponse{Type: RankingTypeValueRapidIncrease, ExchangeDivision: ExchangeDivisionToushou,
				ValueRapidRanking: []ValueRapidRanking{
					{No: 1, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "6907", SymbolName: "ｼﾞｵﾏﾃｯｸ", CurrentPrice: 1013, ChangeRatio: 358, RapidPaymentPercentage: 55381.47, Turnover: 3194.7121, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 54.65, ExchangeName: "東証JQS", CategoryName: "電気機器"},
					{No: 2, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "1490", SymbolName: "上場ﾍﾞｰﾀ/ETF", CurrentPrice: 7750, ChangeRatio: 40, RapidPaymentPercentage: 50159.4, Turnover: 7.75, CurrentPriceTime: HmString{time.Date(0, 1, 1, 13, 20, 0, 0, time.Local)}, ChangePercentage: 0.51, ExchangeName: "東証ETF/ETN", CategoryName: "その他"},
					{No: 3, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "5010", SymbolName: "日精蝋", CurrentPrice: 194, ChangeRatio: 7, RapidPaymentPercentage: 14014.72, Turnover: 308.8866, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 3.74, ExchangeName: "東証２部", CategoryName: "石油・石炭製品"},
					{No: 4, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "8143", SymbolName: "ラピーヌ", CurrentPrice: 430, ChangeRatio: 80, RapidPaymentPercentage: 11547.74, Turnover: 262.624, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 22.85, ExchangeName: "東証２部", CategoryName: "繊維製品"},
					{No: 5, Trend: RankingTrendRiseOver20, AverageRanking: 999, Symbol: "2332", SymbolName: "クエスト", CurrentPrice: 1405, ChangeRatio: 132, RapidPaymentPercentage: 2788.22, Turnover: 180.1434, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 10.36, ExchangeName: "東証JQS", CategoryName: "情報・通信業"},
				}},
			want2: nil,
		},
		{name: "Type=12の正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   rankingBodyType12,
			arg1:   RankingTypeMarginHighMagnification,
			arg2:   ExchangeDivisionALL,
			want1: &RankingResponse{Type: RankingTypeMarginHighMagnification, ExchangeDivision: ExchangeDivisionToushou,
				MarginRanking: []MarginRanking{
					{No: 1, Symbol: "3150", SymbolName: "グリムス", Ratio: 14467, SellRapidPaymentPercentage: 0.1, SellLastWeekRatio: -0.5, BuyRapidPaymentPercentage: 1446.7, BuyLastWeekRatio: 139.7, ExchangeName: "東証１部", CategoryName: "卸売業"},
					{No: 2, Symbol: "6955", SymbolName: "ＦＤＫ", Ratio: 10536.5, SellRapidPaymentPercentage: 0.2, SellLastWeekRatio: -0.8, BuyRapidPaymentPercentage: 2107.3, BuyLastWeekRatio: 121.6, ExchangeName: "東証２部", CategoryName: "電気機器"},
					{No: 3, Symbol: "4592", SymbolName: "ｻﾝﾊﾞｲｵ", Ratio: 10392.5, SellRapidPaymentPercentage: 0.2, SellLastWeekRatio: 0.1, BuyRapidPaymentPercentage: 2078.5, BuyLastWeekRatio: -82.4, ExchangeName: "東証ﾏｻﾞｰｽﾞ", CategoryName: "医薬品"},
					{No: 4, Symbol: "2354", SymbolName: "YEDIGIT", Ratio: 9970, SellRapidPaymentPercentage: 0.1, SellLastWeekRatio: -15.3, BuyRapidPaymentPercentage: 997, BuyLastWeekRatio: -152.1, ExchangeName: "東証２部", CategoryName: "情報・通信業"},
					{No: 5, Symbol: "3909", SymbolName: "ｼｮｰｹｰｽ", Ratio: 8452, SellRapidPaymentPercentage: 0.1, SellLastWeekRatio: -0.7, BuyRapidPaymentPercentage: 845.2, BuyLastWeekRatio: 5.9, ExchangeName: "東証１部", CategoryName: "情報・通信業"},
				}},
			want2: nil,
		},
		{name: "Type=14の正常レスポンスをパースして返せる",
			status: http.StatusOK,
			body:   rankingBodyType14,
			arg1:   RankingTypePriceIncreaseRateByCategory,
			arg2:   ExchangeDivisionALL,
			want1: &RankingResponse{Type: RankingTypePriceIncreaseRateByCategory, ExchangeDivision: ExchangeDivisionUnspecified,
				CategoryPriceRanking: []CategoryPriceRanking{
					{No: 1, Trend: RankingTrendRise, AverageRanking: 18, Category: "343", CategoryName: "IS 空運", CurrentPrice: 170.97, ChangeRatio: 6.72, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 4.09},
					{No: 2, Trend: RankingTrendRise, AverageRanking: 16, Category: "341", CategoryName: "IS 陸運", CurrentPrice: 1895.49, ChangeRatio: 15.41, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 0.82},
					{No: 3, Trend: RankingTrendRise, AverageRanking: 22, Category: "348", CategoryName: "IS 銀行", CurrentPrice: 123.22, ChangeRatio: 0.21, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 0.17},
					{No: 4, Trend: RankingTrendRiseOver20, AverageRanking: 25, Category: "342", CategoryName: "IS 海運", CurrentPrice: 300.77, ChangeRatio: 0.33, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 0.11},
					{No: 4, Trend: RankingTrendRise, AverageRanking: 13, Category: "347", CategoryName: "IS 小売", CurrentPrice: 1390.16, ChangeRatio: 1.56, CurrentPriceTime: HmString{time.Date(0, 1, 1, 15, 0, 0, 0, time.Local)}, ChangePercentage: 0.11},
				}},
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
			mux.HandleFunc("/ranking", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.status)
				_, _ = w.Write([]byte(test.body))
			})
			ts := httptest.NewServer(mux)
			defer ts.Close()

			req := &restClient{url: ts.URL}
			got1, got2 := req.Ranking("", RankingRequest{Type: test.arg1, ExchangeDivision: test.arg2})
			if !reflect.DeepEqual(test.want1, got1) || !reflect.DeepEqual(test.want2, got2) {
				t.Errorf("%s error\nwant: %+v, %v\ngot: %+v, %v\n", t.Name(), test.want1, test.want2, got1, got2)
			}
		})
	}
}

var rankingBodyType1 = `{"Type":"1","ExchangeDivision":"T","Ranking":[{"No":1,"Trend":"1","AverageRanking":999.0,"Symbol":"1689","SymbolName":"ガスETF/ETF(C)","CurrentPrice":2.0000,"ChangeRatio":1.00000,"ChangePercentage":100.00,"CurrentPriceTime":"15:00","TradingVolume":5722.4000,"Turnover":10.4136,"ExchangeName":"東証ETF/ETN","CategoryName":"その他"},{"No":2,"Trend":"1","AverageRanking":999.0,"Symbol":"6907","SymbolName":"ｼﾞｵﾏﾃｯｸ","CurrentPrice":1013.0000,"ChangeRatio":358.00000,"ChangePercentage":54.65,"CurrentPriceTime":"15:00","TradingVolume":3117.5000,"Turnover":3194.7121,"ExchangeName":"東証JQS","CategoryName":"電気機器"},{"No":3,"Trend":"1","AverageRanking":999.0,"Symbol":"8143","SymbolName":"ラピーヌ","CurrentPrice":430.0000,"ChangeRatio":80.00000,"ChangePercentage":22.85,"CurrentPriceTime":"15:00","TradingVolume":627.8000,"Turnover":262.6240,"ExchangeName":"東証２部","CategoryName":"繊維製品"},{"No":4,"Trend":"1","AverageRanking":999.0,"Symbol":"4549","SymbolName":"栄研化","CurrentPrice":2435.0000,"ChangeRatio":398.00000,"ChangePercentage":19.53,"CurrentPriceTime":"15:00","TradingVolume":2077.7000,"Turnover":4900.3547,"ExchangeName":"東証１部","CategoryName":"医薬品"},{"No":5,"Trend":"1","AverageRanking":999.0,"Symbol":"4575","SymbolName":"CANBAS","CurrentPrice":600.0000,"ChangeRatio":97.00000,"ChangePercentage":19.28,"CurrentPriceTime":"15:00","TradingVolume":579.6000,"Turnover":315.8453,"ExchangeName":"東証ﾏｻﾞｰｽﾞ","CategoryName":"医薬品"}]}`
var rankingBodyType5 = `{"Type":"5","ExchangeDivision":"T","Ranking":[{"No":1,"Trend":"1","AverageRanking":22.0,"Symbol":"2929","SymbolName":"ﾌｧｰﾏﾌｰｽﾞ","CurrentPrice":2748.0000,"ChangeRatio":99.00000,"TickCount":40579,"UpCount":12722,"DownCount":12798,"ChangePercentage":3.73,"TradingVolume":16086.8000,"Turnover":43810.0498,"ExchangeName":"東証２部","CategoryName":"食料品"},{"No":2,"Trend":"3","AverageRanking":2.0,"Symbol":"9984","SymbolName":"ｿﾌﾄﾊﾞﾝｸG","CurrentPrice":8285.0000,"ChangeRatio":-309.00000,"TickCount":32219,"UpCount":8655,"DownCount":8562,"ChangePercentage":-3.59,"TradingVolume":16688.8000,"Turnover":138143.1773,"ExchangeName":"東証１部","CategoryName":"情報・通信業"},{"No":3,"Trend":"1","AverageRanking":31.0,"Symbol":"7751","SymbolName":"キヤノン","CurrentPrice":2476.5000,"ChangeRatio":-1.50000,"TickCount":20875,"UpCount":5702,"DownCount":5642,"ChangePercentage":-0.06,"TradingVolume":13403.1000,"Turnover":33139.9700,"ExchangeName":"東証１部","CategoryName":"電気機器"},{"No":4,"Trend":"1","AverageRanking":41.0,"Symbol":"2413","SymbolName":"ｴﾑｽﾘｰ","CurrentPrice":9125.0000,"ChangeRatio":-422.00000,"TickCount":19492,"UpCount":6690,"DownCount":6963,"ChangePercentage":-4.42,"TradingVolume":11325.1000,"Turnover":103804.0663,"ExchangeName":"東証１部","CategoryName":"サービス業"},{"No":5,"Trend":"1","AverageRanking":999.0,"Symbol":"4552","SymbolName":"JCRﾌｧｰﾏ","CurrentPrice":3115.0000,"ChangeRatio":373.00000,"TickCount":18514,"UpCount":5040,"DownCount":4847,"ChangePercentage":13.60,"TradingVolume":7019.8000,"Turnover":21201.6820,"ExchangeName":"東証１部","CategoryName":"医薬品"}]}`
var rankingBodyType6 = `{"Type":"6","ExchangeDivision":"T","Ranking":[{"No":1,"Trend":"1","AverageRanking":999.0,"Symbol":"1490","SymbolName":"上場ﾍﾞｰﾀ/ETF","CurrentPrice":7750.0000,"ChangeRatio":40.00000,"RapidTradePercentage":49900.00,"TradingVolume":1.0000,"CurrentPriceTime":"13:20","ChangePercentage":0.51,"ExchangeName":"東証ETF/ETN","CategoryName":"その他"},{"No":2,"Trend":"1","AverageRanking":999.0,"Symbol":"6907","SymbolName":"ｼﾞｵﾏﾃｯｸ","CurrentPrice":1013.0000,"ChangeRatio":358.00000,"RapidTradePercentage":28189.47,"TradingVolume":3117.5000,"CurrentPriceTime":"15:00","ChangePercentage":54.65,"ExchangeName":"東証JQS","CategoryName":"電気機器"},{"No":3,"Trend":"1","AverageRanking":999.0,"Symbol":"5010","SymbolName":"日精蝋","CurrentPrice":194.0000,"ChangeRatio":7.00000,"RapidTradePercentage":11951.40,"TradingVolume":1453.4000,"CurrentPriceTime":"15:00","ChangePercentage":3.74,"ExchangeName":"東証２部","CategoryName":"石油・石炭製品"},{"No":4,"Trend":"1","AverageRanking":999.0,"Symbol":"8143","SymbolName":"ラピーヌ","CurrentPrice":430.0000,"ChangeRatio":80.00000,"RapidTradePercentage":9648.44,"TradingVolume":627.8000,"CurrentPriceTime":"15:00","ChangePercentage":22.85,"ExchangeName":"東証２部","CategoryName":"繊維製品"},{"No":5,"Trend":"1","AverageRanking":999.0,"Symbol":"8184","SymbolName":"島忠","CurrentPrice":5470.0000,"ChangeRatio":-10.00000,"RapidTradePercentage":2520.35,"TradingVolume":1382.5000,"CurrentPriceTime":"15:00","ChangePercentage":-0.18,"ExchangeName":"東証監理","CategoryName":"小売業"}]}`
var rankingBodyType7 = `{"Type":"7","ExchangeDivision":"T","Ranking":[{"No":1,"Trend":"1","AverageRanking":999.0,"Symbol":"6907","SymbolName":"ｼﾞｵﾏﾃｯｸ","CurrentPrice":1013.0000,"ChangeRatio":358.00000,"RapidPaymentPercentage":55381.47,"Turnover":3194.7121,"CurrentPriceTime":"15:00","ChangePercentage":54.65,"ExchangeName":"東証JQS","CategoryName":"電気機器"},{"No":2,"Trend":"1","AverageRanking":999.0,"Symbol":"1490","SymbolName":"上場ﾍﾞｰﾀ/ETF","CurrentPrice":7750.0000,"ChangeRatio":40.00000,"RapidPaymentPercentage":50159.40,"Turnover":7.7500,"CurrentPriceTime":"13:20","ChangePercentage":0.51,"ExchangeName":"東証ETF/ETN","CategoryName":"その他"},{"No":3,"Trend":"1","AverageRanking":999.0,"Symbol":"5010","SymbolName":"日精蝋","CurrentPrice":194.0000,"ChangeRatio":7.00000,"RapidPaymentPercentage":14014.72,"Turnover":308.8866,"CurrentPriceTime":"15:00","ChangePercentage":3.74,"ExchangeName":"東証２部","CategoryName":"石油・石炭製品"},{"No":4,"Trend":"1","AverageRanking":999.0,"Symbol":"8143","SymbolName":"ラピーヌ","CurrentPrice":430.0000,"ChangeRatio":80.00000,"RapidPaymentPercentage":11547.74,"Turnover":262.6240,"CurrentPriceTime":"15:00","ChangePercentage":22.85,"ExchangeName":"東証２部","CategoryName":"繊維製品"},{"No":5,"Trend":"1","AverageRanking":999.0,"Symbol":"2332","SymbolName":"クエスト","CurrentPrice":1405.0000,"ChangeRatio":132.00000,"RapidPaymentPercentage":2788.22,"Turnover":180.1434,"CurrentPriceTime":"15:00","ChangePercentage":10.36,"ExchangeName":"東証JQS","CategoryName":"情報・通信業"}]}`
var rankingBodyType12 = `{"Type":"12","ExchangeDivision":"T","Ranking":[{"No":1,"Symbol":"3150","SymbolName":"グリムス","Ratio":14467.0000,"SellRapidPaymentPercentage":0.1000,"SellLastWeekRatio":-0.5000,"BuyRapidPaymentPercentage":1446.7000,"BuyLastWeekRatio":139.7000,"ExchangeName":"東証１部","CategoryName":"卸売業"},{"No":2,"Symbol":"6955","SymbolName":"ＦＤＫ","Ratio":10536.5000,"SellRapidPaymentPercentage":0.2000,"SellLastWeekRatio":-0.8000,"BuyRapidPaymentPercentage":2107.3000,"BuyLastWeekRatio":121.6000,"ExchangeName":"東証２部","CategoryName":"電気機器"},{"No":3,"Symbol":"4592","SymbolName":"ｻﾝﾊﾞｲｵ","Ratio":10392.5000,"SellRapidPaymentPercentage":0.2000,"SellLastWeekRatio":0.1000,"BuyRapidPaymentPercentage":2078.5000,"BuyLastWeekRatio":-82.4000,"ExchangeName":"東証ﾏｻﾞｰｽﾞ","CategoryName":"医薬品"},{"No":4,"Symbol":"2354","SymbolName":"YEDIGIT","Ratio":9970.0000,"SellRapidPaymentPercentage":0.1000,"SellLastWeekRatio":-15.3000,"BuyRapidPaymentPercentage":997.0000,"BuyLastWeekRatio":-152.1000,"ExchangeName":"東証２部","CategoryName":"情報・通信業"},{"No":5,"Symbol":"3909","SymbolName":"ｼｮｰｹｰｽ","Ratio":8452.0000,"SellRapidPaymentPercentage":0.1000,"SellLastWeekRatio":-0.7000,"BuyRapidPaymentPercentage":845.2000,"BuyLastWeekRatio":5.9000,"ExchangeName":"東証１部","CategoryName":"情報・通信業"}]}`
var rankingBodyType14 = `{"Type":"14","ExchangeDivision":null,"Ranking":[{"No":1,"Trend":"2","AverageRanking":18.0,"Category":"343","CategoryName":"IS 空運","CurrentPrice":170.9700,"ChangeRatio":6.72000,"CurrentPriceTime":"15:00","ChangePercentage":4.09},{"No":2,"Trend":"2","AverageRanking":16.0,"Category":"341","CategoryName":"IS 陸運","CurrentPrice":1895.4900,"ChangeRatio":15.41000,"CurrentPriceTime":"15:00","ChangePercentage":0.82},{"No":3,"Trend":"2","AverageRanking":22.0,"Category":"348","CategoryName":"IS 銀行","CurrentPrice":123.2200,"ChangeRatio":0.21000,"CurrentPriceTime":"15:00","ChangePercentage":0.17},{"No":4,"Trend":"1","AverageRanking":25.0,"Category":"342","CategoryName":"IS 海運","CurrentPrice":300.7700,"ChangeRatio":0.33000,"CurrentPriceTime":"15:00","ChangePercentage":0.11},{"No":4,"Trend":"2","AverageRanking":13.0,"Category":"347","CategoryName":"IS 小売","CurrentPrice":1390.1600,"ChangeRatio":1.56000,"CurrentPriceTime":"15:00","ChangePercentage":0.11}]}`
