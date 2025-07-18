package api

import (
	"encoding/json"
)

// RegisterVCTemplate 注册VC模板
func (c *Client) RegisterVCTemplate(req *RegisterVCTemplateRequest) (*CommonResponse, error) {
	path := "/api/sys/v1/vc/register"
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

// QueryVCTemplate 查询VC模板
func (c *Client) QueryVCTemplate(req *QueryVCTemplateRequest) (*QueryVCTemplateResponse, error) {
	path := "/api/sys/v1/vc/search"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp QueryVCTemplateResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
