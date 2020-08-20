package kabus

import "fmt"

// ErrorResponse - kabusapiからのエラーレスポンス
type ErrorResponse struct {
	StatusCode int    `json:"-"`       // HTTPステータス
	Body       string `json:"-"`       // エラーBODY
	Code       int    `json:"Code"`    // エラーコード
	Message    string `json:"Message"` // エラーメッセージ
}

func (r ErrorResponse) Error() string {
	return fmt.Sprintf(`{StatusCode:%d Body:%s Code:%d Message %s}`, r.StatusCode, r.Body, r.Code, r.Message)
}
