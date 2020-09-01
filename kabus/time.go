package kabus

import (
	"time"
)

var (
	YmdNUMToday = YmdNUM{isToday: true} // 当日
)

// NewYmdNUM - time.Timeを与えてYmdNUMを生成する
func NewYmdNUM(t time.Time) YmdNUM {
	return YmdNUM{Time: t}
}

// YmdNUM - YYYYMMDDフォーマット(数値) TODO いいかんじの名前に変える
type YmdNUM struct {
	time.Time
	isToday bool
}

func (t YmdNUM) MarshalJSON() ([]byte, error) {
	if t.isToday {
		return []byte("0"), nil
	}
	// 8桁を保つために最低でも10000101にする
	if t.IsZero() || t.Year() < 1000 {
		return []byte("10000101"), nil
	}
	return []byte(t.Time.Format("20060102")), nil
}

func (t *YmdNUM) UnmarshalJSON(b []byte) error {
	if b == nil || string(b) == `""` || string(b) == "null" {
		return nil
	}
	tt, err := time.Parse(`20060102`, string(b))
	if err != nil {
		return err
	}
	*t = YmdNUM{Time: time.Date(tt.Year(), tt.Month(), tt.Day(), tt.Hour(), tt.Minute(), tt.Second(), tt.Nanosecond(), time.Local)}
	return nil
}
