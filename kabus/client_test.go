package kabus

import (
	"reflect"
	"testing"
)

func Test_getClient(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg  bool
		want string
	}{
		{name: "本番用のホストが取れる", arg: true, want: "localhost:18080"},
		{name: "検証用のホストが取れる", arg: false, want: "localhost:18081"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := host(test.arg)
			if test.want != got {
				t.Errorf("%s error\nwant: %s\ngot: %s\n", t.Name(), test.want, got)
			}
		})
	}
}

func Test_baseURL(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		isProd bool
		want   string
	}{
		{name: "本番用URLが取れる", isProd: true, want: "http://localhost:18080/kabusapi/"},
		{name: "検証用URLが取れる", isProd: false, want: "http://localhost:18081/kabusapi/"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := baseURL(test.isProd)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), test.want, got)
			}
		})
	}
}
