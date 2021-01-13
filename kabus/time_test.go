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
		{name: "正常な日付をパースできる", time: YmdNUM{Time: time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local)}, want: []byte(`20200821`)},
		{name: "time.Timeのゼロ値は10000101にしておく", time: YmdNUM{Time: time.Time{}}, want: []byte("10000101")},
		{name: "1000年以前は10000101にしておく", time: YmdNUM{Time: time.Date(999, 12, 31, 0, 0, 0, 0, time.Local)}, want: []byte("10000101")},
		{name: "当日指定は0にする", time: YmdNUM{Time: time.Date(999, 12, 31, 0, 0, 0, 0, time.Local), isToday: true}, want: []byte("0")},
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
		{name: "正常系のパース", src: []byte(`20200821`), want: YmdNUM{Time: time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local)}},
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

func Test_YmdNUMToday(t *testing.T) {
	t.Parallel()
	want := YmdNUM{isToday: true}
	got := YmdNUMToday
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_NewYmdNUM(t *testing.T) {
	t.Parallel()
	arg := time.Date(2020, 9, 1, 14, 22, 47, 0, time.Local)
	want := YmdNUM{Time: arg}
	got := NewYmdNUM(arg)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_YmNUM_MarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		time YmNUM
		want []byte
	}{
		{name: "正常な日付をパースできる", time: YmNUM{Time: time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local)}, want: []byte(`202008`)},
		{name: "time.Timeのゼロ値は100001にしておく", time: YmNUM{Time: time.Time{}}, want: []byte("100001")},
		{name: "1000年以前は100001にしておく", time: YmNUM{Time: time.Date(999, 12, 31, 0, 0, 0, 0, time.Local)}, want: []byte("100001")},
		{name: "当月指定は0にする", time: YmNUM{Time: time.Time{}, isThisMonth: true}, want: []byte("0")},
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

func Test_YmNUM_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		src  []byte
		want YmNUM
	}{
		{name: "正常系のパース", src: []byte(`202008`), want: YmNUM{Time: time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local)}},
		{name: "nullはゼロ値にする", src: []byte(`null`), want: YmNUM{}},
		{name: "空文字はゼロ値にする", src: []byte(`""`), want: YmNUM{}},
		{name: "nilはゼロ値にする", src: nil, want: YmNUM{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := YmNUM{}
			err := got.UnmarshalJSON(test.src)
			if !reflect.DeepEqual(test.want, got) || err != nil {
				t.Errorf("%s error\nwant: %v, %v\ngot: %v\n", t.Name(), test.want, err, got)
			}
		})
	}
}

func Test_YmNUMToday(t *testing.T) {
	t.Parallel()
	want := YmNUM{isThisMonth: true}
	got := YmNUMToday
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_NewYmNUM(t *testing.T) {
	t.Parallel()
	arg := time.Date(2020, 9, 1, 14, 22, 47, 0, time.Local)
	want := YmNUM{Time: arg}
	got := NewYmNUM(arg)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}

func Test_YmNUM_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		time YmNUM
		want string
	}{
		{name: "正常な日付をパースできる", time: YmNUM{Time: time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local)}, want: `202008`},
		{name: "time.Timeのゼロ値は100001にしておく", time: YmNUM{Time: time.Time{}}, want: "100001"},
		{name: "1000年以前は100001にしておく", time: YmNUM{Time: time.Date(999, 12, 31, 0, 0, 0, 0, time.Local)}, want: "100001"},
		{name: "当月指定は0にする", time: YmNUM{Time: time.Time{}, isThisMonth: true}, want: "0"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := test.time.String()
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %s\ngot: %s\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_YmString_MarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		time YmString
		want []byte
	}{
		{name: "正常な日付をパースできる", time: YmString{Time: time.Date(2020, 8, 21, 0, 0, 0, 0, time.Local)}, want: []byte(`2020/08`)},
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

func Test_YmString_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		src  []byte
		want YmString
	}{
		{name: "正常系のパース", src: []byte(`"2020/08"`), want: YmString{Time: time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local)}},
		{name: "nullはゼロ値にする", src: []byte(`null`), want: YmString{}},
		{name: "空文字はゼロ値にする", src: []byte(`""`), want: YmString{}},
		{name: "nilはゼロ値にする", src: nil, want: YmString{}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := YmString{}
			err := got.UnmarshalJSON(test.src)
			if !reflect.DeepEqual(test.want, got) || err != nil {
				t.Errorf("%s error\nwant: %v, %v\ngot: %v\n", t.Name(), test.want, err, got)
			}
		})
	}
}
