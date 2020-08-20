package kabus

import "testing"

func Test_ErrorResponse_Error(t *testing.T) {
	t.Parallel()

	want := `{StatusCode:401 Body:{"Code":4001013,"Message":"トークン取得失敗"} Code:4001013 Message トークン取得失敗}`
	got := ErrorResponse{
		StatusCode: 401,
		Body:       `{"Code":4001013,"Message":"トークン取得失敗"}`,
		Code:       4001013,
		Message:    "トークン取得失敗",
	}.Error()

	if want != got {
		t.Errorf("%s error\nwant: %s\ngot: %s\n", t.Name(), want, got)
	}
}
