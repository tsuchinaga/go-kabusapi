package kabus

import "time"

// YmdHms - YYYY-MM-DDTHH:MM:SSフォーマット TODO いいかんじの名前に変える
type YmdTHms struct {
	time.Time
}

func (t *YmdTHms) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.Format("2006-01-02T15:04:05") + `"`), nil
}

func (t *YmdTHms) UnmarshalJSON(b []byte) error {
	tt, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
	if err != nil {
		return err
	}
	*t = YmdTHms{time.Date(tt.Year(), tt.Month(), tt.Day(), tt.Hour(), tt.Minute(), tt.Second(), tt.Nanosecond(), time.Local)}
	return nil
}
