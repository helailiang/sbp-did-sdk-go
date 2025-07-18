package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client OpenAPI通用RESTful客户端
// 支持Token鉴权
// 用于封装所有OpenAPI的RESTful接口调用

type Client struct {
	BaseURL    string
	Token      string // 私有项目需要
	ProjectNO  string // 私有项目需要
	HTTPClient *http.Client
}

// NewClient 创建OpenAPI客户端
func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL:    baseURL,
		Token:      token,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// doRequest 执行HTTP请求
func (c *Client) doRequest(method, path string, body interface{}, headers map[string]string) ([]byte, int, error) {
	url := c.BaseURL + path
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}

	// 默认JSON
	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return respBody, resp.StatusCode, fmt.Errorf("http error: %d, %s", resp.StatusCode, string(respBody))
	}

	return respBody, resp.StatusCode, nil
}

// Get GET请求
func (c *Client) Get(path string, headers map[string]string) ([]byte, int, error) {
	return c.doRequest("GET", path, nil, headers)
}

// Post POST请求
func (c *Client) Post(path string, body interface{}, headers map[string]string) ([]byte, int, error) {
	return c.doRequest("POST", path, body, headers)
}

// Put PUT请求
func (c *Client) Put(path string, body interface{}, headers map[string]string) ([]byte, int, error) {
	return c.doRequest("PUT", path, body, headers)
}

// Delete DELETE请求
func (c *Client) Delete(path string, headers map[string]string) ([]byte, int, error) {
	return c.doRequest("DELETE", path, nil, headers)
}
