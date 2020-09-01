package kabus

import (
	"io"
	"time"

	"golang.org/x/net/websocket"
)

// PriceMessage - 時価情報 ※時価情報・板情報と完全に一致している
type PriceMessage struct {
	Symbol                   string                   `json:"Symbol"`                   // 銘柄コード
	SymbolName               string                   `json:"SymbolName"`               // 銘柄名
	Exchange                 Exchange                 `json:"Exchange"`                 // 市場コード
	ExchangeName             string                   `json:"ExchangeName"`             // 市場名称
	CurrentPrice             float64                  `json:"CurrentPrice"`             // 現値
	CurrentPriceTime         time.Time                `json:"CurrentPriceTime"`         // 現値時刻
	CurrentPriceChangeStatus CurrentPriceChangeStatus `json:"CurrentPriceChangeStatus"` // 現値前値比較
	CurrentPriceStatus       CurrentPriceStatus       `json:"CurrentPriceStatus"`       // 現値ステータス
	CalcPrice                float64                  `json:"CalcPrice"`                // 計算用現値
	PreviousClose            float64                  `json:"PreviousClose"`            // 前日終値
	PreviousCloseTime        time.Time                `json:"PreviousCloseTime"`        // 前日終値日付
	ChangePreviousClose      float64                  `json:"ChangePreviousClose"`      // 前日比
	ChangePreviousClosePer   float64                  `json:"ChangePreviousClosePer"`   // 騰落率
	OpeningPrice             float64                  `json:"OpeningPrice"`             // 始値
	OpeningPriceTime         time.Time                `json:"OpeningPriceTime"`         // 始値時刻
	HighPrice                float64                  `json:"HighPrice"`                // 高値
	HighPriceTime            time.Time                `json:"HighPriceTime"`            // 高値時刻
	LowPrice                 float64                  `json:"LowPrice"`                 // 安値
	LowPriceTime             time.Time                `json:"LowPriceTime"`             // 安値時刻
	TradingVolume            float64                  `json:"TradingVolume"`            // 売買高
	TradingVolumeTime        time.Time                `json:"TradingVolumeTime"`        // 売買高時刻
	VWAP                     float64                  `json:"VWAP"`                     // 売買高加重平均価格（VWAP）
	TradingValue             float64                  `json:"TradingValue"`             // 売買代金
	BidQty                   float64                  `json:"BidQty"`                   // 最良売気配数量
	BidPrice                 float64                  `json:"BidPrice"`                 // 最良売気配値段
	BidTime                  time.Time                `json:"BidTime"`                  // 最良売気配時刻
	BidSign                  BidAskSign               `json:"BidSign"`                  // 最良売気配フラグ
	MarketOrderSellQty       float64                  `json:"MarketOrderSellQty"`       // 売成行数量
	Sell1                    FirstBoardSign           `json:"Sell1"`                    // 売気配数量1本目
	Sell2                    BoardSign                `json:"Sell2"`                    // 売気配数量2本目
	Sell3                    BoardSign                `json:"Sell3"`                    // 売気配数量3本目
	Sell4                    BoardSign                `json:"Sell4"`                    // 売気配数量4本目
	Sell5                    BoardSign                `json:"Sell5"`                    // 売気配数量5本目
	Sell6                    BoardSign                `json:"Sell6"`                    // 売気配数量6本目
	Sell7                    BoardSign                `json:"Sell7"`                    // 売気配数量7本目
	Sell8                    BoardSign                `json:"Sell8"`                    // 売気配数量8本目
	Sell9                    BoardSign                `json:"Sell9"`                    // 売気配数量9本目
	Sell10                   BoardSign                `json:"Sell10"`                   // 売気配数量10本目
	AskQty                   float64                  `json:"AskQty"`                   // 最良買気配数量
	AskPrice                 float64                  `json:"AskPrice"`                 // 最良買気配値段
	AskTime                  time.Time                `json:"AskTime"`                  // 最良買気配時刻
	AskSign                  BidAskSign               `json:"AskSign"`                  // 最良買気配フラグ
	MarketOrderBuyQty        float64                  `json:"MarketOrderBuyQty"`        // 買成行数量
	Buy1                     FirstBoardSign           `json:"Buy1"`                     // 買気配数量1本目
	Buy2                     BoardSign                `json:"Buy2"`                     // 買気配数量2本目
	Buy3                     BoardSign                `json:"Buy3"`                     // 買気配数量3本目
	Buy4                     BoardSign                `json:"Buy4"`                     // 買気配数量4本目
	Buy5                     BoardSign                `json:"Buy5"`                     // 買気配数量5本目
	Buy6                     BoardSign                `json:"Buy6"`                     // 買気配数量6本目
	Buy7                     BoardSign                `json:"Buy7"`                     // 買気配数量7本目
	Buy8                     BoardSign                `json:"Buy8"`                     // 買気配数量8本目
	Buy9                     BoardSign                `json:"Buy9"`                     // 買気配数量9本目
	Buy10                    BoardSign                `json:"Buy10"`                    // 買気配数量10本目
	OverSellQty              float64                  `json:"OverSellQty"`              // OVER気配数量
	UnderBuyQty              float64                  `json:"UnderBuyQty"`              // UNDER気配数量
	TotalMarketValue         float64                  `json:"TotalMarketValue"`         // 時価総額
}

// NewWSRequester - 時価PUSH配信リクエスタの生成
func NewWSRequester(isProd bool, onNext func(PriceMessage) error) *wsRequester {
	r := &wsRequester{
		wsClient: wsClient{url: createWS(isProd), isProd: isProd},
		onNext:   onNext,
	}
	if r.onNext == nil {
		r.onNext = func(PriceMessage) error { return nil }
	}
	return r
}

// wsRequester - 時価PUSH配信リクエスタ
type wsRequester struct {
	wsClient
	ws       *websocket.Conn
	onNext   func(PriceMessage) error // メッセージ受領時に叩かれる処理
	isClosed bool
}

// Open - web socketを開く
// 受け取ったメッセージはonNext関数に渡す
func (r *wsRequester) Open() error {
	var err error
	r.ws, err = websocket.Dial(r.url, "", "http://"+getHost(r.isProd))
	if err != nil {
		return err
	}
	defer r.ws.Close()

	for {
		var msg PriceMessage
		if err := websocket.JSON.Receive(r.ws, &msg); err != nil && err != io.EOF && !r.isClosed {
			// エラーがあって、EOFではなくて、自ら閉じたわけでなければエラーを返す
			return err
		}

		// 自ら接続を閉じたのであれば終了する
		if r.isClosed {
			return nil
		}

		if err := r.onNext(msg); err != nil {
			return err
		}
	}
}

// Close - web socketを閉じる
func (r *wsRequester) Close() error {
	r.isClosed = true
	return r.ws.Close()
}
