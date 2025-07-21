package api

import (
	"encoding/json"
)

// RegisterIssuer 注册发证方
func (c *Client) RegisterIssuer(req *RegisterIssuerRequest) (*CommonResponse, error) {
	path := "/api/sys/v1/issuer/register"
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

// QueryIssuer 查询发证方
func (c *Client) QueryIssuer(req *QueryIssuerRequest) (*QueryIssuerResponse, error) {
	path := "/api/sys/v1/issuer/search"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp QueryIssuerResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateIssuer 更新发证方
func (c *Client) UpdateIssuer(req *UpdateIssuerRequest) (*UpdateIssuerResponse, error) {
	path := "/api/sys/v1/issuer/update"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp UpdateIssuerResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IssuerStatus 启用/禁用发证方
func (c *Client) IssuerStatus(req *IssuerStatusRequest) (*IssuerStatusResponse, error) {
	path := "/api/sys/v1/issuer/status/update"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp IssuerStatusResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
 