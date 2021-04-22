package kabus

// Exchange - 市場
type Exchange int

const (
	// 株式
	ExchangeUnspecified Exchange = 0 // 指定なし
	ExchangeToushou     Exchange = 1 // 東証
	ExchangeMeishou     Exchange = 3 // 名証
	ExchangeFukushou    Exchange = 5 // 福証
	ExchangeSatsushou   Exchange = 6 // 札証

	// 先物
	ExchangeAll     Exchange = 2  // 日通し
	ExchangeDaytime Exchange = 23 // 日中
	ExchangeEvening Exchange = 24 // 夜間
)

// OrderExchange - 注文可能市場
type OrderExchange int

const (
	OrderExchangeUnspecified OrderExchange = 0  // 指定なし
	OrderExchangeToushou     OrderExchange = 1  // 東証
	OrderExchangeMeishou     OrderExchange = 3  // 名証
	OrderExchangeFukushou    OrderExchange = 5  // 福証
	OrderExchangeSatsushou   OrderExchange = 6  // 札証
	OrderExchangeSOR         OrderExchange = 9  // SOR
	OrderExchangeAll         OrderExchange = 2  // 日通し
	OrderExchangeDaytime     OrderExchange = 23 // 日中
	OrderExchangeEvening     OrderExchange = 24 // 夜間
)

// StockExchange - 株式市場
type StockExchange int

const (
	StockExchangeUnspecified StockExchange = 0 // 指定なし
	StockExchangeToushou     StockExchange = 1 // 東証
	StockExchangeMeishou     StockExchange = 3 // 名証
	StockExchangeFukushou    StockExchange = 5 // 福証
	StockExchangeSatsushou   StockExchange = 6 // 札証
)

// FutureExchange - 先物市場
type FutureExchange int

const (
	FutureExchangeUnspecified FutureExchange = 0  // 指定なし
	FutureExchangeAll         FutureExchange = 2  // 日通し
	FutureExchangeDaytime     FutureExchange = 23 // 日中
	FutureExchangeEvening     FutureExchange = 24 // 夜間
)

// OptionExchange - オプション市場
type OptionExchange int

const (
	OptionExchangeUnspecified OptionExchange = 0  // 指定なし
	OptionExchangeAll         OptionExchange = 2  // 日通し
	OptionExchangeDaytime     OptionExchange = 23 // 日中
	OptionExchangeEvening     OptionExchange = 24 // 夜間
)

// SecurityType - 商品種別
type SecurityType int

const (
	SecurityTypeUnspecified SecurityType = 0   // 指定なし
	SecurityTypeStock       SecurityType = 1   // 株式
	SecurityTypeNK225       SecurityType = 101 // 日経225先物
	SecurityTypeNK225Mini   SecurityType = 901 // 日経225mini先物
	SecurityTypeJPX400      SecurityType = 121 // JPX日経インデックス400先物
	SecurityTypeTOPIX       SecurityType = 107 // TOPIX先物
	SecurityTypeTOPIXMini   SecurityType = 907 // ミニTOPIX先物
	SecurityTypeMothers     SecurityType = 154 // 東証マザーズ指数先物
	SecurityTypeREIT        SecurityType = 155 // 東証REIT指数先物
	SecurityTypeDOW         SecurityType = 144 // NYダウ先物
	SecurityTypeVI          SecurityType = 145 // 日経平均VI先物
	SecurityTypeCORE30      SecurityType = 171 // TOPIX Core30先物
)

// Side - 売買区分
type Side string

const (
	SideUnspecified Side = ""  // 指定なし
	SideSell        Side = "1" // 売
	SideBuy         Side = "2" // 買
)

// CashMargin - 現物信用区分
type CashMargin int

const (
	CashMarginUnspecified CashMargin = 0 // 指定なし
	CashMarginCash        CashMargin = 1 // 現物
	CashMarginMarginEntry CashMargin = 2 // 信用新規
	CashMarginMarginExit  CashMargin = 3 // 信用返済
)

// MarginTradeType - 信用取引区分
type MarginTradeType int

const (
	MarginTradeTypeUnspecified  MarginTradeType = 0 // 指定なし
	MarginTradeTypeSystem       MarginTradeType = 1 // 制度信用
	MarginTradeTypeGeneralLong  MarginTradeType = 2 // 一般信用
	MarginTradeTypeGeneralShort MarginTradeType = 3 // 一般信用(売短)
)

// DelivType - 受渡区分
type DelivType int

const (
	DelivTypeUnspecified DelivType = 0 // 指定なし
	DelivTypeAuto        DelivType = 1 // 自動振替
	DelivTypeCash        DelivType = 2 // お預り金
)

// FundType - 資産区分
type FundType string

const (
	FundTypeUnspecified     FundType = "  " // 指定なし
	FundTypeProtected       FundType = "02" // 保護
	FundTypeTransferMargin  FundType = "AA" // 信用代用
	FundTypeTransferDeposit FundType = "BB" // 証拠金代用
	FundTypeMarginTrading   FundType = "11" // 信用取引
)

// AccountType - 口座種別
type AccountType int

const (
	AccountTypeUnspecified AccountType = 0  // 指定なし
	AccountTypeGeneral     AccountType = 2  // 一般
	AccountTypeSpecific    AccountType = 4  // 特定
	AccountTypeCorporation AccountType = 12 // 法人
)

// ClosePositionOrder - 決済順序
type ClosePositionOrder int

const (
	ClosePositionOrderUnspecified        ClosePositionOrder = -1 // 指定なし
	ClosePositionOrderDateAscProfitDesc  ClosePositionOrder = 0  // 日付（古い順）、損益（高い順）
	ClosePositionOrderDateAscProfitAsc   ClosePositionOrder = 1  // 日付（古い順）、損益（低い順）
	ClosePositionOrderDateDescProfitDesc ClosePositionOrder = 2  // 日付（新しい順）、損益（高い順）
	ClosePositionOrderDateDescProfitAc   ClosePositionOrder = 3  // 日付（新しい順）、損益（低い順）
	ClosePositionOrderProfitDescDateAsc  ClosePositionOrder = 4  // 損益（高い順）、日付（古い順）
	ClosePositionOrderProfitDescDateDesc ClosePositionOrder = 5  // 損益（高い順）、日付（新しい順）
	ClosePositionOrderProfitAscDateAsc   ClosePositionOrder = 6  // 損益（低い順）、日付（古い順）
	ClosePositionOrderProfitAscDateDesc  ClosePositionOrder = 7  // 損益（低い順）、日付（古い順）
)

// StockFrontOrderType - 株式執行条件
type StockFrontOrderType int

const (
	StockFrontOrderTypeUnspecified StockFrontOrderType = 0  // 指定なし
	StockFrontOrderTypeMarket      StockFrontOrderType = 10 // 成行
	StockFrontOrderTypeMOOM        StockFrontOrderType = 13 // 寄成（前場） Market on Open Morning
	StockFrontOrderTypeMOOA        StockFrontOrderType = 14 // 寄成（後場） Market on Open Afternoon
	StockFrontOrderTypeMOCM        StockFrontOrderType = 15 // 引成（前場） Market on Close Morning
	StockFrontOrderTypeMOCA        StockFrontOrderType = 16 // 引成（後場） Market on Close Afternoon
	StockFrontOrderTypeIOCMarket   StockFrontOrderType = 17 // IOC成行
	StockFrontOrderTypeLimit       StockFrontOrderType = 20 // 指値
	StockFrontOrderTypeLOOM        StockFrontOrderType = 21 // 寄指（前場） Limit on Open Morning
	StockFrontOrderTypeLOOA        StockFrontOrderType = 22 // 寄指（後場） Limit on Open Afternoon
	StockFrontOrderTypeLOCM        StockFrontOrderType = 23 // 引指（前場） Limit on Close Morning
	StockFrontOrderTypeLOCA        StockFrontOrderType = 24 // 引指（後場） Limit on Close Afternoon
	StockFrontOrderTypeFunariM     StockFrontOrderType = 25 // 不成（前場） Funari Morning
	StockFrontOrderTypeFunariA     StockFrontOrderType = 26 // 不成（後場） Funari Afternoon
	StockFrontOrderTypeIOCLimit    StockFrontOrderType = 27 // IOC指値
)

// CurrentPriceChangeStatus - 現値前値比較
type CurrentPriceChangeStatus string

const (
	CurrentPriceChangeStatusUnspecified          CurrentPriceChangeStatus = ""     // 指定なし
	CurrentPriceChangeStatusNoEffect             CurrentPriceChangeStatus = "0000" // 変化なし
	CurrentPriceChangeStatusNoChange             CurrentPriceChangeStatus = "0056" // 変わらず
	CurrentPriceChangeStatusUp                   CurrentPriceChangeStatus = "0057" // UP
	CurrentPriceChangeStatusDown                 CurrentPriceChangeStatus = "0058" // DOWN
	CurrentPriceChangeStatusOpenPriceAfterBreak  CurrentPriceChangeStatus = "0059" // 中断板寄り後の初値
	CurrentPriceChangeStatusTradingSessionClose  CurrentPriceChangeStatus = "0060" // ザラバ引け
	CurrentPriceChangeStatusClose                CurrentPriceChangeStatus = "0061" // 板寄り引け
	CurrentPriceChangeStatusBreakClose           CurrentPriceChangeStatus = "0062" // 中断引け
	CurrentPriceChangeStatusDownClose            CurrentPriceChangeStatus = "0063" // ダウン引け
	CurrentPriceChangeStatusTarnOverClose        CurrentPriceChangeStatus = "0064" // 逆転終値
	CurrentPriceChangeStatusSpecialQuoteClose    CurrentPriceChangeStatus = "0066" // 特別気配引け
	CurrentPriceChangeStatusReservationClose     CurrentPriceChangeStatus = "0067" // 一時留保引け
	CurrentPriceChangeStatusStopClose            CurrentPriceChangeStatus = "0068" // 売買停止引け
	CurrentPriceChangeCircuitBreakerClose        CurrentPriceChangeStatus = "0069" // サーキットブレーカ引け
	CurrentPriceChangeDynamicCircuitBreakerClose CurrentPriceChangeStatus = "0431" // ダイナミックサーキットブレーカ引け
)

// CurrentPriceStatus - 現値ステータス
type CurrentPriceStatus int

const (
	CurrentPriceStatusUnspecified                CurrentPriceStatus = 0  // 指定なし
	CurrentPriceStatusCurrentPrice               CurrentPriceStatus = 1  // 現値
	CurrentPriceStatusNoContinuousTicks          CurrentPriceStatus = 2  // 不連続歩み
	CurrentPriceStatusItayose                    CurrentPriceStatus = 3  // 板寄せ
	CurrentPriceStatusSystemError                CurrentPriceStatus = 4  // システム障害
	CurrentPriceStatusPause                      CurrentPriceStatus = 5  // 中断
	CurrentPriceStatusStopTrading                CurrentPriceStatus = 6  // 売買停止
	CurrentPriceStatusRestart                    CurrentPriceStatus = 7  // 売買停止・システム停止解除
	CurrentPriceStatusClosePrice                 CurrentPriceStatus = 8  // 終値
	CurrentPriceStatusSystemStop                 CurrentPriceStatus = 9  // システム停止
	CurrentPriceStatusRoughQuote                 CurrentPriceStatus = 10 // 概算値
	CurrentPriceStatusReference                  CurrentPriceStatus = 11 // 参考値
	CurrentPriceStatusInCircuitBreak             CurrentPriceStatus = 12 // サーキットブレイク実施中
	CurrentPriceStatusRestoration                CurrentPriceStatus = 13 // システム障害解除
	CurrentPriceStatusReleaseCircuitBreak        CurrentPriceStatus = 14 // システム障害解除
	CurrentPriceStatusReleasePause               CurrentPriceStatus = 15 // 中断解除
	CurrentPriceStatusInReservation              CurrentPriceStatus = 16 // 一時留保中
	CurrentPriceStatusReleaseReservation         CurrentPriceStatus = 17 // 一時留保解除
	CurrentPriceStatusFileError                  CurrentPriceStatus = 18 // ファイル障害
	CurrentPriceStatusReleaseFileError           CurrentPriceStatus = 19 // ファイル障害解除
	CurrentPriceStatusSpreadStrategy             CurrentPriceStatus = 20 // Spread/Strategy
	CurrentPriceStatusInDynamicCircuitBreak      CurrentPriceStatus = 21 // ダイナミックサーキットブレイク発動
	CurrentPriceStatusReleaseDynamicCircuitBreak CurrentPriceStatus = 22 // ダイナミックサーキットブレイク解除
	CurrentPriceStatusContractedInItayose        CurrentPriceStatus = 23 // 板寄せ約定
)

// BidAskSign - 最良気配フラグ
type BidAskSign string

const (
	BidAskSignUnspecified            BidAskSign = ""     // 指定なし
	BidAskSignNoEffect               BidAskSign = "0000" // 事象なし
	BidAskSignGeneral                BidAskSign = "0101" // 一般気配
	BidAskSignSpecial                BidAskSign = "0102" // 特別気配
	BidAskSignAttention              BidAskSign = "0103" // 注意気配
	BidAskSignBeforeOpen             BidAskSign = "0107" // 寄前気配
	BidAskSignSpecialBeforeStop      BidAskSign = "0108" // 停止前特別気配
	BidAskSignAfterClose             BidAskSign = "0109" // 引け後気配
	BidAskSignNotExistsContractPoint BidAskSign = "0116" // 寄前気配約定成立ポイントなし
	BidAskSignExistsContractPoint    BidAskSign = "0117" // 寄前気配約定成立ポイントあり
	BidAskSignContinuous             BidAskSign = "0118" // 連続約定気配
	BidAskSignContinuousBeforeStop   BidAskSign = "0119" // 停止前の連続約定気配
	BidAskSignMoving                 BidAskSign = "0120" // 買い上がり売り下がり中
)

// PriceRangeGroup - 呼び値グループ
type PriceRangeGroup string

const (
	PriceRangeGroupUnspecified PriceRangeGroup = ""      // 指定なし
	PriceRangeGroup10000       PriceRangeGroup = "10000" // グループ 10000
	PriceRangeGroup10003       PriceRangeGroup = "10003" // グループ 10003
	PriceRangeGroup10118       PriceRangeGroup = "10118" // グループ 10118
	PriceRangeGroup10119       PriceRangeGroup = "10119" // グループ 10119
	PriceRangeGroup10318       PriceRangeGroup = "10318" // グループ 10318
	PriceRangeGroup10706       PriceRangeGroup = "10706" // グループ 10706
	PriceRangeGroup10718       PriceRangeGroup = "10718" // グループ 10718
	PriceRangeGroup12122       PriceRangeGroup = "12122" // グループ 12122
	PriceRangeGroup14473       PriceRangeGroup = "14473" // グループ 14473
	PriceRangeGroup14515       PriceRangeGroup = "14515" // グループ 14515
	PriceRangeGroup15411       PriceRangeGroup = "15411" // グループ 15411
	PriceRangeGroup15569       PriceRangeGroup = "15569" // グループ 15569
	PriceRangeGroup17163       PriceRangeGroup = "17163" // グループ 17163
)

// Product - 商品
type Product int

const (
	ProductAll    Product = 0 // 全て
	ProductCash   Product = 1 // 現物
	ProductMargin Product = 2 // 信用
	ProductFuture Product = 3 // 先物
	ProductOption Product = 4 // OP
)

// State - 状態
type State int

const (
	StateUnspecified State = 0 // 指定なし
	StateWait        State = 1 // 待機（発注待機）
	StateProcessing  State = 2 // 処理中（発注送信中）
	StateProcessed   State = 3 // 処理済（発注済・訂正済）
	StateInCancel    State = 4 // 訂正取消送信中
	StateDone        State = 5 // 終了（発注エラー・取消済・全約定・失効・期限切れ）
)

// OrderState - 注文状態
type OrderState int

const (
	OrderStateUnspecified OrderState = 0 // 指定なし
	OrderStateWait        OrderState = 1 // 待機（発注待機）
	OrderStateProcessing  OrderState = 2 // 処理中（発注送信中）
	OrderStateProcessed   OrderState = 3 // 処理済（発注済・訂正済）
	OrderStateInCancel    OrderState = 4 // 訂正取消送信中
	OrderStateDone        OrderState = 5 // 終了（発注エラー・取消済・全約定・失効・期限切れ）
)

// OrderDetailState - 注文詳細状態
type OrderDetailState int

const (
	OrderDetailStateUnspecified OrderDetailState = 0 // 指定なし
	OrderDetailStateWait        OrderDetailState = 1 // 待機（発注待機）
	OrderDetailStateProcessing  OrderDetailState = 2 // 処理中（発注送信中・訂正送信中・取消送信中）
	OrderDetailStateProcessed   OrderDetailState = 3 // 処理済（発注済・訂正済・取消済・全約定・期限切れ）
	OrderDetailStateError       OrderDetailState = 4 // エラー
	OrderDetailStateDeleted     OrderDetailState = 5 // 削除済み
)

// OrdType - 執行条件
type OrdType int

const (
	OrdTypeUnspecified   OrdType = 0 // 指定なし
	OrdTypeInTrading     OrdType = 1 // ザラバ
	OrdTypeOpen          OrdType = 2 // 寄り
	OrdTypeClose         OrdType = 3 // 引け
	OrdTypeNoContracted  OrdType = 4 // 不成
	OrdTypeMarketToLimit OrdType = 5 // 対当指値
	OrdTypeIOC           OrdType = 6 // IOC
)

// RecType - 明細種別
type RecType int

const (
	RecTypeUnspecified RecType = 0 // 指定なし
	RecTypeReceived    RecType = 1 // 受付
	RecTypeCarried     RecType = 2 // 繰越
	RecTypeExpired     RecType = 3 // 期限切れ
	RecTypeOrdered     RecType = 4 // 発注
	RecTypeModified    RecType = 5 // 訂正
	RecTypeCanceled    RecType = 6 // 取消
	RecTypeRevocation  RecType = 7 // 失効
	RecTypeContracted  RecType = 8 // 約定
)

// TradeType - 取引区分
type TradeType int

const (
	TradeTypeUnspecified TradeType = 0 // 指定なし
	TradeTypeEntry       TradeType = 1 // 新規
	TradeTypeExit        TradeType = 2 // 返済
)

// TimeInForce - 有効期間条件
type TimeInForce int

const (
	TimeInForceUnspecified TimeInForce = 0 // 指定なし
	TimeInForceFAS         TimeInForce = 1 // FAS
	TimeInForceFAK         TimeInForce = 2 // FAK
	TimeInForceFOK         TimeInForce = 3 // FOK
)

// FutureFrontOrderType - 先物執行条件
type FutureFrontOrderType int

const (
	FutureFrontOrderTypeUnspecified FutureFrontOrderType = 0   // 指定なし
	FutureFrontOrderTypeMarket      FutureFrontOrderType = 120 // 成行
	FutureFrontOrderTypeMarketClose FutureFrontOrderType = 18  // 引成（派生）
	FutureFrontOrderTypeLimit       FutureFrontOrderType = 20  // 指値
	FutureFrontOrderTypeLimitClose  FutureFrontOrderType = 28  // 引指（派生）
)

// OptionFrontOrderType - 先物執行条件
type OptionFrontOrderType int

const (
	OptionFrontOrderTypeUnspecified OptionFrontOrderType = 0   // 指定なし
	OptionFrontOrderTypeMarket      OptionFrontOrderType = 120 // 成行
	OptionFrontOrderTypeMarketClose OptionFrontOrderType = 18  // 引成（派生）
	OptionFrontOrderTypeLimit       OptionFrontOrderType = 20  // 指値
	OptionFrontOrderTypeLimitClose  OptionFrontOrderType = 28  // 引指（派生）
)

// FutureCode - 先物コード
type FutureCode string

const (
	FutureCodeUnspecified FutureCode = ""          // 指定なし
	FutureCodeNK225       FutureCode = "NK225"     // NK225
	FutureCodeNK225Mini   FutureCode = "NK225mini" // NK225mini
	FutureCodeTOPIX       FutureCode = "TOPIX"     // TOPIX
	FutureCodeTOPIXMini   FutureCode = "TOPIXmini" // TOPIXmini
	FutureCodeMOTHERS     FutureCode = "MOTHERS"   // MOTHERS
	FutureCodeJPX400      FutureCode = "JPX400"    // JPX400
	FutureCodeDOW         FutureCode = "DOW"       // DOW
	FutureCodeVI          FutureCode = "VI"        // VI
	FutureCodeCore30      FutureCode = "Core30"    // Core30
	FutureCodeREIT        FutureCode = "REIT"      // REIT
)

// PutOrCall - コール or プット
type PutOrCall string

const (
	PutOrCallUnspecified PutOrCall = "" // 指定なし
	PutOrCallPut         PutOrCall = "P"
	PutOrCallCall        PutOrCall = "C"
)

// PutOrCallNum - コール or プット
type PutOrCallNum int

const (
	PutOrCallNumUnspecified PutOrCallNum = 0 // 指定なし
	PutOrCallNumPut         PutOrCallNum = 1
	PutOrCallNumCall        PutOrCallNum = 2
)

// IsGetOrderDetail - 注文詳細を取得するか
type IsGetOrderDetail string

const (
	IsGetOrderDetailUnspecified IsGetOrderDetail = ""      // 指定なし
	IsGetOrderDetailTrue        IsGetOrderDetail = "true"  // 注文詳細を出力する
	IsGetOrderDetailFalse       IsGetOrderDetail = "false" // 注文詳細を出力しない
)

// Underlyer - 原資産コード
type Underlyer string

const (
	UnderlyerUnspecified  Underlyer = ""
	UnderlyerNK225        Underlyer = "NK225"        // 日経225
	UnderlyerNK300        Underlyer = "NK300"        // 日経225
	UnderlyerMOTHERS      Underlyer = "MOTHERS"      // 東証マザーズ
	UnderlyerJPX400       Underlyer = "JPX400"       // JPX日経400
	UnderlyerTOPIX        Underlyer = "TOPIX"        // TOPIX
	UnderlyerNKVI         Underlyer = "NKVI"         // 日経平均VI
	UnderlyerDJIA         Underlyer = "DJIA"         // NYダウ
	UnderlyerTSEREITINDEX Underlyer = "TSEREITINDEX" // 東証REIT指数
	UnderlyerTOPIXCORE30  Underlyer = "TOPIXCORE30"  // TOPIX Core30
)

// RankingType - ランキング種別
type RankingType string

const (
	RankingTypeUnspecified                 RankingType = ""   // 指定なし
	RankingTypePriceIncreaseRate           RankingType = "1"  // 値上がり率
	RankingTypePriceDecreaseRate           RankingType = "2"  // 値下がり率
	RankingTypeVolume                      RankingType = "3"  // 売買高上位
	RankingTypeValue                       RankingType = "4"  // 売買代金
	RankingTypeTickCount                   RankingType = "5"  // TICK回数
	RankingTypeVolumeRapidIncrease         RankingType = "6"  // 売買高急増
	RankingTypeValueRapidIncrease          RankingType = "7"  // 売買代金急増
	RankingTypeMarginSellBalanceIncrease   RankingType = "8"  // 信用売残増
	RankingTypeMarginSellBalanceDecrease   RankingType = "9"  // 信用売残減
	RankingTypeMarginBuyBalanceIncrease    RankingType = "10" // 信用買残増
	RankingTypeMarginBuyBalanceDecrease    RankingType = "11" // 信用買残減
	RankingTypeMarginHighMagnification     RankingType = "12" // 信用高倍率
	RankingTypeMarginLowMagnification      RankingType = "13" // 信用低倍率
	RankingTypePriceIncreaseRateByCategory RankingType = "14" // 業種別値上がり率
	RankingTypePriceDecreaseRateByCategory RankingType = "15" // 業種別値下がり率
)

// ExchangeDivision - 市場
type ExchangeDivision string

const (
	ExchangeDivisionUnspecified ExchangeDivision = ""    // 指定なし
	ExchangeDivisionALL         ExchangeDivision = "ALL" // 全市場（デフォルト）
	ExchangeDivisionToushou     ExchangeDivision = "T"   // 東証全体
	ExchangeDivisionToushou1    ExchangeDivision = "T1"  // 東証一部
	ExchangeDivisionToushou2    ExchangeDivision = "T2"  // 東証二部
	ExchangeDivisionMothers     ExchangeDivision = "TM"  // マザーズ
	ExchangeDivisionJASDAQ      ExchangeDivision = "TQ"  // JASDAQ
	ExchangeDivisionMeishou     ExchangeDivision = "M"   // 名証
	ExchangeDivisionFukushou    ExchangeDivision = "FK"  // 福証
	ExchangeDivisionSatsushou   ExchangeDivision = "S"   // 札証
)

// RankingTrend - ランキングのトレンド
type RankingTrend string

const (
	RankingTrendUnspecified   RankingTrend = ""  // 指定なし
	RankingTrendNoData        RankingTrend = "0" // 対象データ無し
	RankingTrendRiseOver20    RankingTrend = "1" // 過去10営業日より20位以上上昇
	RankingTrendRise          RankingTrend = "2" // 過去10営業日より1～19位上昇
	RankingTrendUnchanged     RankingTrend = "3" // 過去10営業日と変わらず
	RankingTrendDescent       RankingTrend = "4" // 過去10営業日より1～19位下落
	RankingTrendDescentOver20 RankingTrend = "5" // 過去10営業日より20位以上下落
)

// GetSymbolInfo - 銘柄情報の追加情報出力フラグ
type GetSymbolInfo string

const (
	GetSymbolInfoUnspecified = ""      // 指定なし
	GetSymbolInfoTrue        = "true"  // 追加情報を出力する
	GetSymbolInfoFalse       = "false" // 追加情報を出力しない
)

// GetPositionInfo - 残高照会の追加情報出力フラグ
type GetPositionInfo string

const (
	GetPositionInfoUnspecified = ""      // 指定なし
	GetPositionInfoTrue        = "true"  // 追加情報を出力する
	GetPositionInfoFalse       = "false" // 追加情報を出力しない
)

// ExchangeSymbol - 通貨
type ExchangeSymbol string

const (
	ExchangeSymbolUnspecified ExchangeSymbol = ""       // 指定なし
	ExchangeSymbolUSDJPY      ExchangeSymbol = "usdjpy" //USD/JPY
	ExchangeSymbolEURJPY      ExchangeSymbol = "eurjpy" // EUR/JPY
	ExchangeSymbolGBPJPY      ExchangeSymbol = "gbpjpy" // GBP/JPY
	ExchangeSymbolAUDJPY      ExchangeSymbol = "audjpy" // AUD/JPY
	ExchangeSymbolCHFJPY      ExchangeSymbol = "chfjpy" // CHF/JPY
	ExchangeSymbolCADJPY      ExchangeSymbol = "cadjpy" // CAD/JPY
	ExchangeSymbolNZDJPY      ExchangeSymbol = "nzdjpy" // NZD/JPY
	ExchangeSymbolZARJPY      ExchangeSymbol = "zarjpy" // ZAR/JPY
	ExchangeSymbolEURUSD      ExchangeSymbol = "eurusd" // EUR/USD
	ExchangeSymbolGBPUSD      ExchangeSymbol = "gbpusd" // GBP/USD
	ExchangeSymbolAUDUSD      ExchangeSymbol = "audusd" // AUD/USD
)

// ExchangeSymbolDetail - 通貨内容
type ExchangeSymbolDetail string

const (
	ExchangeSymbolDetailUnspecified ExchangeSymbolDetail = ""        // 指定なし
	ExchangeSymbolDetailUSDJPY      ExchangeSymbolDetail = "USD/JPY" // USD/JPY
	ExchangeSymbolDetailEURJPY      ExchangeSymbolDetail = "EUR/JPY" // EUR/JPY
	ExchangeSymbolDetailGBPJPY      ExchangeSymbolDetail = "GBP/JPY" // GBP/JPY
	ExchangeSymbolDetailAUDJPY      ExchangeSymbolDetail = "AUD/JPY" // AUD/JPY
	ExchangeSymbolDetailCHFJPY      ExchangeSymbolDetail = "CHF/JPY" // CHF/JPY
	ExchangeSymbolDetailCADJPY      ExchangeSymbolDetail = "CAD/JPY" // CAD/JPY
	ExchangeSymbolDetailNZDJPY      ExchangeSymbolDetail = "NZD/JPY" // NZD/JPY
	ExchangeSymbolDetailZARJPY      ExchangeSymbolDetail = "ZAR/JPY" // ZAR/JPY
	ExchangeSymbolDetailEURUSD      ExchangeSymbolDetail = "EUR/USD" // EUR/USD
	ExchangeSymbolDetailGBPUSD      ExchangeSymbolDetail = "GBP/USD" // GBP/USD
	ExchangeSymbolDetailAUDUSD      ExchangeSymbolDetail = "AUD/USD" // AUD/USD
)
