package kabus

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
)

// client - HTTPクライアント
type client struct {
	url   string // URL
	token string // リクエストトークン
}

// post - POSTリクエスト
func (c *client) post(ctx context.Context, request []byte) (int, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", c.url, bytes.NewReader(request))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if c.token != "" {
		req.Header.Set("X-API-KEY", c.token)
	}

	// リクエスト送信
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, b, nil
}

// put - PUTリクエスト
func (c *client) put(ctx context.Context, request []byte) (int, []byte, error) {
	req, err := http.NewRequestWithContext(ctx, "PUT", c.url, bytes.NewReader(request))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if c.token != "" {
		req.Header.Set("X-API-KEY", c.token)
	}

	// リクエスト送信
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	return res.StatusCode, b, nil
}
