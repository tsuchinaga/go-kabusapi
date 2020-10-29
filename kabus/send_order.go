package kabus

// ClosePosition - 返済建玉
type ClosePosition struct {
	HoldID string `json:"HoldID"` // 返済建玉ID
	Qty    int    `json:"Qty"`    // 返済建玉数量
}
