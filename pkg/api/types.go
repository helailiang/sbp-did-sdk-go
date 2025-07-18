package api

// 通用响应结构
// code: "0" 表示成功
// data: 具体数据
// message: 错误信息

type CommonResponse struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// ========== DID ========== //
type RegisterDIDRequest struct {
	ProjectNo   string `json:"projectNo"`
	DIDDocument string `json:"didDocument"` // JSON字符串
	Signature   string `json:"signature"`
	TxSignature string `json:"txSignature,omitempty"`
}

type QueryDIDRequest struct {
	DID       string `json:"did"`
	ProjectNo string `json:"projectNo"`
}

type QueryDIDResponse struct {
	Code string `json:"code"`
	Data struct {
		DIDDocument string `json:"didDocument"`
	} `json:"data"`
	Message string `json:"message"`
}

// ========== Issuer ========== //
type Issuer struct {
	IssuerDid  string `json:"issuerDid"`
	IssuerName string `json:"issuerName"`
}
type RegisterIssuerRequest struct {
	Issuer
	IssuerDid        string `json:"issuerDid"`
	IssuerName       string `json:"issuerName"`
	ProjectNo        string `json:"projectNo"`
	ContactPerson    string `json:"contactPerson"`
	ContactNumber    string `json:"contactNumber"`
	ContactEmail     string `json:"contactEmail"`
	BusinessScenario string `json:"businessScenario"`
	Signature        string `json:"signature"`
	TxSignature      string `json:"txSignature,omitempty"`
}

type QueryIssuerRequest struct {
	IssuerDid string `json:"issuerDid"`
}

type QueryIssuerResponse struct {
	Code    string `json:"code"`
	Data    Issuer `json:"data"`
	Message string `json:"message"`
}

// ========== VC模板 ========== //
type Proof struct {
	Created            string `json:"created"`
	ProofType          string `json:"type"`
	VerificationMethod string `json:"verificationMethod"`
	Cryptosuite        string `json:"cryptosuite"`
	ProofPurpose       string `json:"proofPurpose"`
	ProofValue         string `json:"proofValue"`
}

type RegistrationField struct {
	FieldName   string `json:"fieldName"`
	Description string `json:"description"`
	Mandatory   bool   `json:"mandatory"`
}

type CredentialSubject struct {
	SubjectName string `json:"subjectName"`
	Description string `json:"description"`
}

type VCTemplate struct {
	TemplateId          string              `json:"templateId"`
	TemplateName        string              `json:"templateName"`
	TemplateDescription string              `json:"templateDescription"`
	IssuanceEndpoint    string              `json:"issuanceEndpoint"`
	IssuerDid           string              `json:"issuerDid"`
	RegistrationFields  []RegistrationField `json:"registrationFields"`
	CredentialSubjects  []CredentialSubject `json:"credentialSubjects"`
	Proof               Proof               `json:"proof"`
}

type RegisterVCTemplateRequest struct {
	IssuerDid   string     `json:"issuerDid"`
	ProjectNo   string     `json:"projectNo"`
	VCTemplate  VCTemplate `json:"vcTemplate"`
	Signature   string     `json:"signature"`
	TxSignature string     `json:"txSignature,omitempty"`
}

type QueryVCTemplateRequest struct {
	ProjectNo    string `json:"projectNo"`
	VCTemplateId string `json:"vcTemplateId"`
}

type QueryVCTemplateResponse struct {
	Code string `json:"code"`
	Data struct {
		VCTemplate string `json:"vcTemplate"`
	} `json:"data"`
	Message string `json:"message"`
}

// ========== VC ========== //
// W3C标准VC结构体
// https://www.w3.org/TR/vc-data-model/
type VerifiableCredential struct {
	Context           []string               `json:"@context"`
	ID                string                 `json:"id"`
	Type              []string               `json:"type"`
	Issuer            string                 `json:"issuer"`
	IssuanceDate      string                 `json:"issuanceDate"`
	ExpirationDate    string                 `json:"expirationDate,omitempty"`
	CredentialSubject map[string]interface{} `json:"credentialSubject"`
	CredentialSchema  interface{}            `json:"credentialSchema,omitempty"`
	Proof             interface{}            `json:"proof,omitempty"`
}

type IssueVCRequest struct {
	ProjectNo    string               `json:"projectNo"`
	VCTemplateId string               `json:"vcTemplateId"`
	VC           VerifiableCredential `json:"vc"`
	Signature    string               `json:"signature"`
	TxSignature  string               `json:"txSignature,omitempty"`
}

type QueryVCEvidenceRequest struct {
	VcId      string `json:"vcId"`
	ProjectNo string `json:"projectNo"`
	IssuerDid string `json:"issuerDid"`
}

type QueryVCEvidenceResponse struct {
	Code string `json:"code"`
	Data struct {
		VcId      string `json:"vcId"`
		VcHash    string `json:"vcHash"`
		IssuerDid string `json:"issuerDid"`
		TxHash    string `json:"txHash"`
	} `json:"data"`
	Message string `json:"message"`
}

// ========== DID ========== //
type UpdateDIDRequest struct {
	ProjectNo   string `json:"projectNo"`
	DIDDocument string `json:"didDocument"`
	Index       int    `json:"index,omitempty"`
	Signature   string `json:"signature"`
	TxSignature string `json:"txSignature,omitempty"`
}

type UpdateDIDResponse = CommonResponse

// ========== Issuer ========== //
type UpdateIssuerRequest struct {
	IssuerDid   string `json:"issuerDid"`
	IssuerName  string `json:"issuerName"`
	Signature   string `json:"signature"`
	TxSignature string `json:"txSignature,omitempty"`
}

type UpdateIssuerResponse = CommonResponse

type IssuerStatusRequest struct {
	IssuerDid string `json:"issuerDid"`
	ProjectNo string `json:"projectNo"`
}

type IssuerStatusResponse = CommonResponse

// ========== VC模板 ========== //
type UpdateVCTemplateRequest struct {
	IssuerDid   string     `json:"issuerDid"`
	ProjectNo   string     `json:"projectNo"`
	VCTemplate  VCTemplate `json:"vcTemplate"`
	Signature   string     `json:"signature"`
	TxSignature string     `json:"txSignature,omitempty"`
}

type UpdateVCTemplateResponse = CommonResponse

type VCTemplateStatusRequest struct {
	IssuerDid    string `json:"issuerDid"`
	ProjectNo    string `json:"projectNo"`
	VCTemplateId string `json:"vcTemplateId"`
	Signature    string `json:"signature"`
	TxSignature  string `json:"txSignature,omitempty"`
}

type VCTemplateStatusResponse = CommonResponse

// ========== VC ========== //
type IssueVCResponse = CommonResponse

type VCEvidenceRequest struct {
	VcId        string `json:"vcId"`
	ProjectNo   string `json:"projectNo"`
	VcHash      string `json:"vcHash"`
	IssuerDid   string `json:"issuerDid"`
	Signature   string `json:"signature"`
	TxSignature string `json:"txSignature,omitempty"`
}

type VCEvidenceResponse = CommonResponse

type VCEvidenceSearchRequest struct {
	VcId      string `json:"vcId"`
	ProjectNo string `json:"projectNo"`
	IssuerDid string `json:"issuerDid"`
}

type VCEvidenceSearchResponse = QueryVCEvidenceResponse

type VCRevokeRequest struct {
	ProjectNo string `json:"projectNo"`
	VC        string `json:"vc"`
}

type VCRevokeResponse = CommonResponse

type VCRevokeStatusRequest struct {
	VcId      string `json:"vcId"`
	ProjectNo string `json:"projectNo"`
	IssuerDid string `json:"issuerDid"`
}

type VCRevokeStatusResponse struct {
	Code string `json:"code"`
	Data struct {
		RevokeStatus bool `json:"revokeStatus"`
	} `json:"data"`
	Message string `json:"message"`
}

type VCVerifyRequest struct {
	ProjectNo string `json:"projectNo"`
	VC        string `json:"vc"`
}

type VCVerifyResponse struct {
	Code string `json:"code"`
	Data struct {
		VerificationStatus bool `json:"verificationStatus"`
	} `json:"data"`
	Message string `json:"message"`
}

// ========== 项目管理 ========== //
type GetTokenRequest struct {
	ProjectNo      string `json:"projectNo"`
	ClientName     string `json:"clientName"`
	IsProjectOwner bool   `json:"isProjectOwner,omitempty"`
}

type GetTokenResponse struct {
	Code string `json:"code"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Message string `json:"message"`
}

type ProjectEnableRequest struct {
	ProjectNo string `json:"projectNo"`
	Status    int    `json:"status"`
}

type ProjectEnableResponse = CommonResponse

// ... 其他项目管理相关结构体可继续补充 ...
