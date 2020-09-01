package kabus

import (
	"reflect"
	"testing"
	"time"
)

func Test_YmdNUM_MarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		time YmdNUM
		want []byte
	}{
		{name: "正常な日付をパースできる", time: YmdNUM{time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local)}, want: []byte(`20200821`)},
		{name: "time.Timeのゼロ値は10000101にしておく", time: YmdNUM{time.Time{}}, want: []byte("10000101")},
		{name: "1000年以前は10000101にしておく", time: YmdNUM{time.Date(999, 12, 31, 0, 0, 0, 0, time.Local)}, want: []byte("10000101")},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := test.time.MarshalJSON()
			if !reflect.DeepEqual(test.want, got) || err != nil {
				t.Errorf("%s error\nwant: %s, %+v\ngot: %s\n", t.Name(), test.want, err, got)
			}
		})
	}
}

func Test_YmdNUM_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		src  []byte
		want YmdNUM
	}{
		{name: "正常系のパース", src: []byte(`20200821`), want: YmdNUM{time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local)}},
		{name: "nullはゼロ値にする", src: []byte(`null`), want: YmdNUM{}},
		{name: "空文字はゼロ値にする", src: []byte(`""`), want: YmdNUM{}},
		{name: "nilはゼロ値にする", src: nil, want: YmdNUM{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := YmdNUM{}
			err := got.UnmarshalJSON(test.src)
			if !reflect.DeepEqual(test.want, got) || err != nil {
				t.Errorf("%s error\nwant: %s, %v\ngot: %s\n", t.Name(), test.want, err, got)
			}
		})
	}
}
