package kabus

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// httpClient - HTTPクライアント
type httpClient struct {
	url   string // URL
	token string // リクエストトークン
}

// get - GETリクエスト
func (c *httpClient) get(ctx context.Context, pathParam string, queryParam string) (int, []byte, error) {
	u, err := url.Parse(c.url)
	if err != nil {
		return 0, nil, err
	}
	if pathParam != "" {
		u.Path += "/" + pathParam
	}
	if queryParam != "" {
		u.RawQuery = queryParam
	}

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return 0, nil, err
	}
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

// post - POSTリクエスト
func (c *httpClient) post(ctx context.Context, request []byte) (int, []byte, error) {
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
func (c *httpClient) put(ctx context.Context, request []byte) (int, []byte, error) {
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

// parseResponse - レスポンスをパースする
func parseResponse(code int, body []byte, v interface{}) error {
	if code == http.StatusOK {
		if err := json.Unmarshal(body, v); err != nil {
			return err
		}
		return nil
	} else {
		var errRes ErrorResponse
		if err := json.Unmarshal(body, &errRes); err != nil {
			return err
		}
		errRes.StatusCode = code
		errRes.Body = string(body)
		return errRes
	}
}

// createURL - リクエスト先のURLを生成する
func createURL(path string, isProd bool) string {
	return "http://" + getHost(isProd) + strings.ReplaceAll("/kabusapi/"+path, "//", "/")
}

// wsClient - WSクライアント
type wsClient struct {
	url    string // URL
	isProd bool   // 本番用か
}

// createWS - リクエスト先のWS URLを生成する
func createWS(isProd bool) string {
	return "ws://" + getHost(isProd) + "/kabusapi/websocket"
}

// getHost - 本番か検証のホストを返す
func getHost(isProd bool) string {
	if isProd {
		return "localhost:18080"
	}
	return "localhost:18081"
}
