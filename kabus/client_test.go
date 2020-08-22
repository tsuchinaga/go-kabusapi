package kabus

import "testing"

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
			got := getHost(test.arg)
			if test.want != got {
				t.Errorf("%s error\nwant: %s\ngot: %s\n", t.Name(), test.want, got)
			}
		})
	}
}
