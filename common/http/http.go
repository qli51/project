package http

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"bytes"

	"github.com/tal-tech/go-zero/core/logx"
)

type HTTPClient struct {
	header http.Header
	client *http.Client
	Logger logx.Logger
	ctx    context.Context
}

func NewHttpClient(ctx context.Context, header http.Header) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
		header: header,
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}

func (client *HTTPClient) BuildUrl(agent, endpoint, urlStr string, params map[string]interface{}) string {
	if params == nil {
		return fmt.Sprintf("http://%s:%s%s", agent, endpoint, urlStr)
	}

	query := url.Values{}

	for key, param := range params {
		// 考虑传入值为切片的情况
		if values, ok := param.([]string); ok {
			for _, v := range values {
				query.Add(key, v)
			}
		} else {
			query.Add(key, fmt.Sprintf("%s", param))
		}
	}

	return fmt.Sprintf("http://%s:%s%s?%s", agent, endpoint, urlStr, query.Encode())
}

func (client *HTTPClient) Request(url, method string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		client.Logger.Errorf("create new request failed: %s", err)
		return nil, err
	}

	//设置请求头
	if client.header != nil {
		req.Header = client.header
	}

	return client.client.Do(req)
}
