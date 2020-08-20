package kabus

import "time"

// SendOrderRequest - 注文発注のリクエストパラメータ
type SendOrderRequest struct {
	Password           string             `json:"Password"`           // 注文パスワード
	Symbol             string             `json:"Symbol"`             // 銘柄コード
	Exchange           Exchange           `json:"Exchange"`           // 市場コード
	SecurityType       SecurityType       `json:"SecurityType"`       // 商品種別
	Side               Side               `json:"Side"`               // 売買区分
	CashMargin         CashMargin         `json:"CashMargin"`         // 現物信用区分
	MarginTradeType    MarginTradeType    `json:"MarginTradeType"`    // 信用取引区分 ※信用取引の場合必須
	DelivType          DelivType          `json:"DelivType"`          // 受渡区分 ※株式取引の場合必須で、現物売・信用返済では「指定なし」を指定する
	FundType           FundType           `json:"FundType"`           // 資産区分 ※株式取引の場合必須で、現物売・信用返済では「指定なし」を指定する
	AccountType        AccountType        `json:"AccountType"`        // 口座種別
	Qty                int                `json:"Qty"`                // 注文数量
	ClosePositionOrder ClosePositionOrder `json:"ClosePositionOrder"` // 決済順序 ※信用取引の場合必須
	ClosePositions     []ClosePosition    `json:"ClosePositions"`     // 返済建玉指定
	Price              int                `json:"Price"`              // 注文価格
	ExpireDay          time.Time          `json:"ExpireDay"`          // 注文有効期限（年月日）
	FrontOrderType     FrontOrderType     `json:"FrontOrderType"`     // 執行条件
}

// ClosePosition - 返済建玉
type ClosePosition struct {
	HoldID string `json:"HoldID"` // 返済建玉ID
	Qty    int    `json:"Qty"`    // 返済建玉数量
}

// SendOrderResponse - 注文発注のレスポンス
type SendOrderResponse struct {
	Result  int    `json:"Result"`  // 結果コード 0が成功、それ以外はエラー
	OrderID string `json:"OrderId"` // 受付注文番号
}
