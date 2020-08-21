package kabus

import (
	"reflect"
	"testing"
	"time"
)

func Test_YmdTHms_MarshalJSON(t *testing.T) {
	tt := YmdTHms{time.Date(2020, 8, 21, 17, 54, 22, 0, time.Local)}
	want := []byte(`"2020-08-21T17:54:22"`)
	got, err := tt.MarshalJSON()
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Errorf("%s error\nwant: %s, %+v\ngot: %s\n", t.Name(), want, err, got)
	}
}

func Test_YmdTHms_UnmarshalJSON(t *testing.T) {
	want := YmdTHms{time.Date(2020, 8, 21, 17, 54, 22, 0, time.Local)}

	got := YmdTHms{}
	err := got.UnmarshalJSON([]byte(`"2020-08-21T17:54:22"`))
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Errorf("%s error\nwant: %s, %v\ngot: %s\n", t.Name(), want, err, got)
	}
}
