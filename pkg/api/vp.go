package api

import (
	"fmt"
)

// GenerateVP 生成VP
func (c *Client) GenerateVP(projectID string, vpReq interface{}) ([]byte, error) {
	path := fmt.Sprintf("/v1/vp/%s/generate", projectID)
	resp, _, err := c.Post(path, vpReq, nil)
	return resp, err
}

// VerifyVP 验证VP
func (c *Client) VerifyVP(projectID string, verifyReq interface{}) ([]byte, error) {
	path := fmt.Sprintf("/v1/vp/%s/verify", projectID)
	resp, _, err := c.Post(path, verifyReq, nil)
	return resp, err
}
