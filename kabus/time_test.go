package kabus

import (
	"reflect"
	"testing"
	"time"
)

func Test_YmdTHms_MarshalJSON(t *testing.T) {
	t.Parallel()
	tt := YmdTHms{time.Date(2020, 8, 21, 17, 54, 22, 0, time.Local)}
	want := []byte(`"2020-08-21T17:54:22"`)
	got, err := tt.MarshalJSON()
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Errorf("%s error\nwant: %s, %+v\ngot: %s\n", t.Name(), want, err, got)
	}
}

func Test_YmdTHms_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		src  []byte
		want YmdTHms
	}{
		{name: "正常系のパース", src: []byte(`"2020-08-21T17:54:22"`), want: YmdTHms{time.Date(2020, 8, 21, 17, 54, 22, 0, time.Local)}},
		{name: "nullはゼロ値にする", src: []byte(`null`), want: YmdTHms{}},
		{name: "空文字はゼロ値にする", src: []byte(`""`), want: YmdTHms{}},
		{name: "nilはゼロ値にする", src: nil, want: YmdTHms{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := YmdTHms{}
			err := got.UnmarshalJSON(test.src)
			if !reflect.DeepEqual(test.want, got) || err != nil {
				t.Errorf("%s error\nwant: %s, %v\ngot: %s\n", t.Name(), test.want, err, got)
			}
		})
	}
}

func Test_YmdTHmsSSS_MarshalJSON(t *testing.T) {
	t.Parallel()
	tt := YmdTHmsSSS{time.Date(2020, 8, 21, 17, 54, 22, 123456000, time.Local)}
	want := []byte(`"2020-08-21T17:54:22.123456"`)
	got, err := tt.MarshalJSON()
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Errorf("%s error\nwant: %s, %+v\ngot: %s\n", t.Name(), want, err, got)
	}
}

func Test_YmdTHmsSSS_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		src  []byte
		want YmdTHmsSSS
	}{
		{name: "正常系のパース", src: []byte(`"2020-08-21T17:54:22.123456"`), want: YmdTHmsSSS{time.Date(2020, 8, 21, 17, 54, 22, 123456000, time.Local)}},
		{name: "タイムゾーンは無視してパースする", src: []byte(`"2020-08-21T17:54:22.123456+09:00"`), want: YmdTHmsSSS{time.Date(2020, 8, 21, 17, 54, 22, 123456000, time.Local)}},
		{name: "nullはゼロ値にする", src: []byte(`null`), want: YmdTHmsSSS{}},
		{name: "空文字はゼロ値にする", src: []byte(`""`), want: YmdTHmsSSS{}},
		{name: "nilはゼロ値にする", src: nil, want: YmdTHmsSSS{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := YmdTHmsSSS{}
			err := got.UnmarshalJSON(test.src)
			if !reflect.DeepEqual(test.want, got) || err != nil {
				t.Errorf("%s error\nwant: %s, %v\ngot: %s\n", t.Name(), test.want, err, got)
			}
		})
	}
}

func Test_YmdNUM_MarshalJSON(t *testing.T) {
	t.Parallel()
	tt := YmdNUM{time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local)}
	want := []byte(`20200821`)
	got, err := tt.MarshalJSON()
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Errorf("%s error\nwant: %s, %+v\ngot: %s\n", t.Name(), want, err, got)
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
