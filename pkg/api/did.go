package api

import (
	"encoding/json"
)

// RegisterDID 注册DID文档
func (c *Client) RegisterDID(req *RegisterDIDRequest) (*CommonResponse, error) {
	path := "/api/sys/v1/did/register"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp CommonResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryDID 查询DID文档
func (c *Client) QueryDID(req *QueryDIDRequest) (*QueryDIDResponse, error) {
	path := "/api/sys/v1/did/search"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp QueryDIDResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDID 更新DID文档
func (c *Client) UpdateDID(req *UpdateDIDRequest) (*UpdateDIDResponse, error) {
	path := "/api/sys/v1/did/update"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp UpdateDIDResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
