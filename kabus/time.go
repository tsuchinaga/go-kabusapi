package kabus

import (
	"strings"
	"time"
)

// YmdHms - YYYY-MM-DDTHH:MM:SSフォーマット TODO いいかんじの名前に変える
type YmdTHms struct {
	time.Time
}

func (t YmdTHms) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.Format("2006-01-02T15:04:05") + `"`), nil
}

func (t *YmdTHms) UnmarshalJSON(b []byte) error {
	if b == nil || string(b) == `""` || string(b) == "null" {
		return nil
	}
	tt, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
	if err != nil {
		return err
	}
	*t = YmdTHms{time.Date(tt.Year(), tt.Month(), tt.Day(), tt.Hour(), tt.Minute(), tt.Second(), tt.Nanosecond(), time.Local)}
	return nil
}

// YmdTHmsSSS - YYYY-MM-DDTHH:MM:SS.SSSSSSフォーマット TODO いいかんじの名前に変える
type YmdTHmsSSS struct {
	time.Time
}

func (t YmdTHmsSSS) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.Format("2006-01-02T15:04:05.999999") + `"`), nil
}

func (t *YmdTHmsSSS) UnmarshalJSON(b []byte) error {
	if b == nil || string(b) == `""` || string(b) == "null" {
		return nil
	}
	tt, err := time.Parse(`"2006-01-02T15:04:05.999999"`, strings.ReplaceAll(string(b), "+09:00", ""))
	if err != nil {
		return err
	}
	*t = YmdTHmsSSS{time.Date(tt.Year(), tt.Month(), tt.Day(), tt.Hour(), tt.Minute(), tt.Second(), tt.Nanosecond(), time.Local)}
	return nil
}

// YmdNUM - YYYYMMDDフォーマット(数値) TODO いいかんじの名前に変える
type YmdNUM struct {
	time.Time
}

func (t YmdNUM) MarshalJSON() ([]byte, error) {
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
	*t = YmdNUM{time.Date(tt.Year(), tt.Month(), tt.Day(), tt.Hour(), tt.Minute(), tt.Second(), tt.Nanosecond(), time.Local)}
	return nil
}
