package api

import (
	"encoding/json"
)

// QueryVCEvidence 查询VC存证
func (c *Client) QueryVCEvidence(req *QueryVCEvidenceRequest) (*QueryVCEvidenceResponse, error) {
	path := "/api/sys/v1/vc/evidence/search"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp QueryVCEvidenceResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// IssueVC 签发VC
func (c *Client) IssueVC(req *IssueVCRequest) (*IssueVCResponse, error) {
	path := "/api/sys/v1/vc/issue"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp IssueVCResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// VCEvidence VC存证
func (c *Client) VCEvidence(req *VCEvidenceRequest) (*VCEvidenceResponse, error) {
	path := "/api/sys/v1/vc/evidence"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp VCEvidenceResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// VCRevoke 吊销VC
func (c *Client) VCRevoke(req *VCRevokeRequest) (*VCRevokeResponse, error) {
	path := "/api/sys/v1/vc/revoke"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp VCRevokeResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// VCRevokeStatus 查询VC吊销状态
func (c *Client) VCRevokeStatus(req *VCRevokeStatusRequest) (*VCRevokeStatusResponse, error) {
	path := "/api/sys/v1/vc/status/search"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp VCRevokeStatusResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// VCVerify 核验VC
func (c *Client) VCVerify(req *VCVerifyRequest) (*VCVerifyResponse, error) {
	path := "/api/sys/v1/vc/verify"
	headers := map[string]string{}
	if c.Token != "" {
		headers["token"] = c.Token
	}
	respBytes, _, err := c.Post(path, req, headers)
	if err != nil {
		return nil, err
	}
	var resp VCVerifyResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
 