package kabus

import (
	"reflect"
	"testing"
)

func Test_NewWSRequester(t *testing.T) {
	t.Parallel()

	want := &wsRequester{wsClient: wsClient{url: "ws://localhost:18080/kabusapi/websocket", isProd: true}}
	got := NewWSRequester(true, nil)
	got.onNext = nil // テストのためにnil化

	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s error\nwant: %+v\ngot: %+v\n", t.Name(), want, got)
	}
}
