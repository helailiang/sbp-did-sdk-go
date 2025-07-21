package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/helailiang/sbp-did-sdk-go/pkg/api"
	"github.com/helailiang/sbp-did-sdk-go/pkg/config"
	"github.com/helailiang/sbp-did-sdk-go/pkg/crypto"
	"github.com/helailiang/sbp-did-sdk-go/pkg/did"
	"github.com/helailiang/sbp-did-sdk-go/pkg/utils"
)

// BasicUsageExample 基本使用示例
func BasicUsageExample() {
	fmt.Println("SBP DID SDK Go - 基本使用示例")
	fmt.Println("================================")

	// 1. 创建配置
	cfg := createExampleConfig()

	// 2. 生成密钥对
	fmt.Println("\n步骤1: 生成密钥对")
	keyPair, err := crypto.GenerateKeyPair(cfg, "ECDSA", "example-key")
	if err != nil {
		log.Fatalf("密钥生成失败: %v", err)
	}
	fmt.Printf("✓ 密钥对生成成功 (算法: %s)\n", keyPair.Algorithm)

	// 3. 计算DID标识符
	fmt.Println("\n步骤2: 计算DID标识符")
	didIdentifier, err := did.CalculateDIDIdentifier(keyPair, "did:sbp:")
	if err != nil {
		log.Fatalf("DID标识符计算失败: %v", err)
	}
	fmt.Printf("✓ DID标识符: %s\n", didIdentifier)

	// 4. 组装DID文档
	fmt.Println("\n步骤3: 组装DID文档")
	businessAttributes := map[string]interface{}{
		"name":         "示例用户",
		"email":        "user@example.com",
		"organization": "示例组织",
		"role":         "开发者",
	}

	doc, err := did.AssembleDIDDocument(cfg, keyPair, "ECDSA", didIdentifier, businessAttributes)
	if err != nil {
		log.Fatalf("DID文档组装失败: %v", err)
	}
	fmt.Printf("✓ DID文档组装成功\n")
	docJ, _ := json.Marshal(doc)
	fmt.Printf("  - ID: %s\n", string(docJ))
	fmt.Printf("  - 验证方法数量: %d\n", len(doc.VerificationMethod))
	fmt.Printf("  - 业务属性数量: %d\n", len(doc.CustomFields))

	// 5. 计算哈希
	fmt.Println("\n步骤4: 计算数据哈希")
	testData := "这是需要哈希计算的数据"
	hash, err := utils.CalculateHashFromString(testData, utils.SHA256)
	if err != nil {
		log.Fatalf("哈希计算失败: %v", err)
	}
	fmt.Printf("✓ 数据哈希计算成功\n")
	fmt.Printf("  - 原始数据: %s\n", testData)
	fmt.Printf("  - SHA256哈希: %s\n", hash)

	// 6. 签名验证
	fmt.Println("\n步骤5: 数字签名验证")
	dataToSign := "这是需要签名的数据"
	signature, err := crypto.Sign(cfg, keyPair, []byte(dataToSign), "ECDSA")
	if err != nil {
		log.Fatalf("签名失败: %v", err)
	}
	fmt.Printf("✓ 数据签名成功\n")
	fmt.Printf("  - 原始数据: %s\n", dataToSign)
	fmt.Printf("  - 签名: %s...\n", signature.Signature[:32])

	// 验证签名
	verification, err := crypto.VerifySignatureFromHex(cfg, keyPair, dataToSign, signature.Signature, "ECDSA")
	if err != nil {
		log.Fatalf("签名验证失败: %v", err)
	}
	fmt.Printf("✓ 签名验证结果: %t\n", verification.Valid)

	// 7. 生成VC模板ID
	fmt.Println("\n步骤6: 生成VC模板ID")
	templateID := utils.GenerateVCTemplateID()
	fmt.Printf("✓ VC模板ID生成成功: %s\n", templateID)

	fmt.Println("\n✓ 所有基本功能演示完成！")
}

// createExampleConfig 创建示例配置
func createExampleConfig() *config.Config {
	cfg := config.NewConfig()

	// 设置华为云配置（示例值）
	cfg.HuaweiCloudEndpoint = "https://your-huawei-cloud-endpoint"
	cfg.HuaweiCloudAccessKey = "your-access-key"
	cfg.HuaweiCloudSecretKey = "your-secret-key"
	cfg.HuaweiCloudRegion = "cn-north-4"

	// 设置OpenAPI配置（示例值）
	cfg.OpenAPIEndpoint = "https://your-openapi-endpoint"
	cfg.ProjectID = "your-project-id"
	cfg.ProjectVisibility = "public"
	cfg.Token = ""

	// 设置算法配置
	cfg.DefaultAlgorithm = "ECDSA"
	cfg.DefaultHashAlgorithm = "SHA256"

	return cfg
}

// AdvancedUsageExample 高级使用示例
func AdvancedUsageExample() {
	fmt.Println("\nSBP DID SDK Go - 高级使用示例")
	fmt.Println("================================")

	cfg := createExampleConfig()

	// 1. 多种算法密钥生成
	fmt.Println("\n1. 多种算法密钥生成")
	algorithms := []string{"ECDSA", "RSA", "SM2"}

	for _, algo := range algorithms {
		keyPair, err := crypto.GenerateKeyPair(cfg, algo, fmt.Sprintf("%s-key", algo))
		if err != nil {
			log.Printf("生成%s密钥失败: %v", algo, err)
			continue
		}
		fmt.Printf("✓ %s密钥对生成成功\n", algo)

		// 获取公钥
		publicKeyHex, err := keyPair.GetPublicKeyHex()
		if err != nil {
			log.Printf("获取%s公钥失败: %v", algo, err)
			continue
		}
		fmt.Printf("  - 公钥长度: %d 字节\n", len(publicKeyHex)/2)
	}

	// 2. 多种哈希算法
	fmt.Println("\n2. 多种哈希算法")
	testData := "测试数据"
	hashAlgorithms := []utils.HashAlgorithm{utils.SHA256, utils.SM3}

	for _, hashAlgo := range hashAlgorithms {
		hash, err := utils.CalculateHashFromString(testData, hashAlgo)
		if err != nil {
			log.Printf("计算%s哈希失败: %v", hashAlgo, err)
			continue
		}
		fmt.Printf("✓ %s哈希: %s\n", hashAlgo, hash)
	}

	// 3. 加密解密演示
	fmt.Println("\n3. 加密解密演示")
	keyPair, err := crypto.GenerateKeyPair(cfg, "RSA", "encryption-key")
	if err != nil {
		log.Fatalf("生成RSA密钥失败: %v", err)
	}

	plainText := "敏感数据需要加密"
	encryptionResult, err := crypto.Encrypt(cfg, keyPair, []byte(plainText), "RSA")
	if err != nil {
		log.Fatalf("加密失败: %v", err)
	}
	fmt.Printf("✓ 数据加密成功\n")
	fmt.Printf("  - 原始数据: %s\n", plainText)
	fmt.Printf("  - 加密数据: %s...\n", encryptionResult.EncryptedData[:32])

	// 解密
	decryptionResult, err := crypto.DecryptFromHex(cfg, keyPair, encryptionResult.EncryptedData, "RSA")
	if err != nil {
		log.Fatalf("解密失败: %v", err)
	}
	fmt.Printf("✓ 数据解密成功\n")
	fmt.Printf("  - 解密结果: %s\n", decryptionResult.DecryptedData)

	fmt.Println("\n✓ 高级功能演示完成！")
}

func OpenAPIUsageExample() {
	fmt.Println("\nSBP DID SDK Go - OpenAPI接口调用示例")
	fmt.Println("====================================")

	// 1. 初始化API客户端
	openAPIEndpoint := "https://your-openapi-endpoint" // 替换为实际OpenAPI地址
	token := "your-token"                              // 私有项目需要Token，公开项目可为空
	projectID := "your-project-id"
	client := api.NewClient(openAPIEndpoint, token)

	//Bclient := api.NewClient(openAPIEndpoint, token)
	//Bclient.IssuerStatus()

	// 2. 注册DID文档
	fmt.Println("\n1. 注册DID文档（上链）")
	didReq := &api.RegisterDIDRequest{
		ProjectNo:   projectID,
		DIDDocument: "{...}", // DID文档JSON字符串
		Signature:   "签名值",
	}
	resp, err := client.RegisterDID(didReq)
	if err != nil {
		fmt.Printf("注册DID文档失败: %v\n", err)
	} else {
		fmt.Printf("注册DID文档响应: %+v\n", resp)
	}

	// 3. 查询DID文档
	fmt.Println("\n2. 查询DID文档")
	queryReq := &api.QueryDIDRequest{
		DID:       "did:sbp:example123",
		ProjectNo: projectID,
	}
	queryResp, err := client.QueryDID(queryReq)
	if err != nil {
		fmt.Printf("查询DID文档失败: %v\n", err)
	} else {
		fmt.Printf("查询DID文档响应: %+v\n", queryResp)
	}

	// 4. 注册发证方
	fmt.Println("\n3. 注册发证方")
	issuerReq := &api.RegisterIssuerRequest{
		Issuer: api.Issuer{
			IssuerDid:  "issuer-001",
			IssuerName: "示例发证方",
		},
		ProjectNo:        projectID,
		ContactPerson:    "张三",
		ContactNumber:    "123456789",
		ContactEmail:     "issuer@example.com",
		BusinessScenario: "示例业务",
		Signature:        "签名值",
	}
	issuerResp, err := client.RegisterIssuer(issuerReq)
	if err != nil {
		fmt.Printf("注册发证方失败: %v\n", err)
	} else {
		fmt.Printf("注册发证方响应: %+v\n", issuerResp)
	}

	// 5. 签发VC
	fmt.Println("\n4. 签发VC")
	vcReq := &api.IssueVCRequest{
		// 按VC结构体补充字段
	}
	vcResp, err := client.IssueVC(vcReq)
	if err != nil {
		fmt.Printf("签发VC失败: %v\n", err)
	} else {
		fmt.Printf("签发VC响应: %+v\n", vcResp)
	}
}

// RegisterIssuerExample 演示如何注册发证方（OpenAPI接口封装）
func RegisterIssuerExample() {
	fmt.Println("\nSBP DID SDK Go - 注册发证方示例")
	fmt.Println("================================")

	// 1. 初始化API客户端
	openAPIEndpoint := "https://your-openapi-endpoint" // 替换为实际OpenAPI地址
	token := "your-token"                              // 私有项目需要Token，公开项目可为空
	projectNo := "your-project-id"
	client := api.NewClient(openAPIEndpoint, token)

	// 2. 构造注册发证方请求
	issuerDid := "did:sbp:issuer001"
	issuerName := "示例发证方"
	contactPerson := "张三"
	contactNumber := "123456789"
	contactEmail := "issuer@example.com"
	businessScenario := "示例业务"

	// 3. 生成签名（实际业务中应使用私钥对关键信息签名，这里仅为演示）
	signature := "签名值" // TODO: 替换为真实签名

	req := &api.RegisterIssuerRequest{

		IssuerDid:        issuerDid,
		IssuerName:       issuerName,
		ProjectNo:        projectNo,
		ContactPerson:    contactPerson,
		ContactNumber:    contactNumber,
		ContactEmail:     contactEmail,
		BusinessScenario: businessScenario,
		Signature:        signature,
		// TxSignature:   "", // 如有链上交易签名可补充
	}

	// 4. 调用注册接口
	resp, err := client.RegisterIssuer(req)
	if err != nil {
		fmt.Println("注册发证方失败:", err)
	} else if resp.Code != "0" {
		fmt.Println("注册发证方失败，错误信息:", resp.Message)
	} else {
		fmt.Println("注册发证方成功，返回数据:", resp.Data)
	}
}

// SignAndIssueVCExample 演示签发VC的完整业务流程
func SignAndIssueVCExample() {
	fmt.Println("\nSBP DID SDK Go - 签发VC完整流程示例")
	fmt.Println("====================================")

	// 1. 初始化API客户端
	openAPIEndpoint := "https://your-openapi-endpoint" // 替换为实际OpenAPI地址
	token := "your-token"                              // 私有项目需要Token，公开项目可为空
	projectNo := "your-project-id"
	client := api.NewClient(openAPIEndpoint, token)

	// 2. 获取VC模板信息
	vcTemplateId := "template-001" // 替换为实际模板ID
	queryTemplateReq := &api.QueryVCTemplateRequest{
		ProjectNo:    projectNo,
		VCTemplateId: vcTemplateId,
	}
	templateResp, err := client.QueryVCTemplate(queryTemplateReq)
	if err != nil {
		fmt.Println("获取VC模板失败:", err)
		return
	}
	if templateResp.Code != "0" {
		fmt.Println("获取VC模板失败，错误信息:", templateResp.Message)
		return
	}
	// 反序列化模板内容
	var template api.VCTemplate
	err = json.Unmarshal([]byte(templateResp.Data.VCTemplate), &template)
	if err != nil {
		fmt.Println("解析VC模板失败:", err)
		return
	}

	// 3. 校验Credential Subjects
	credentialSubject := map[string]interface{}{
		"name": "张三",
		"age":  18,
		// ...补充其他字段...
	}
	missingFields := []string{}
	for _, field := range template.RegistrationFields {
		if field.Mandatory {
			if _, ok := credentialSubject[field.FieldName]; !ok {
				missingFields = append(missingFields, field.FieldName)
			}
		}
	}
	if len(missingFields) > 0 {
		fmt.Println("Credential Subject 缺少必填字段:", missingFields)
		return
	}

	// 4. 组装VC
	vc := api.VerifiableCredential{
		Context:           []string{"https://www.w3.org/2018/credentials/v1"},
		ID:                "vc-001", // 可用UUID生成
		Type:              []string{"VerifiableCredential"},
		Issuer:            template.IssuerDid,
		IssuanceDate:      "2024-06-01T00:00:00Z", // 替换为当前时间
		CredentialSubject: credentialSubject,
	}

	// 5. 计算VC哈希
	vcBytes, err := json.Marshal(vc)
	if err != nil {
		fmt.Println("VC序列化失败:", err)
		return
	}
	hash, err := utils.CalculateHash(vcBytes, utils.SHA256)
	if err != nil {
		fmt.Println("VC哈希计算失败:", err)
		return
	}
	fmt.Println("VC哈希:", hash)

	// 6. VC哈希存证
	vcEvidenceReq := &api.VCEvidenceRequest{
		VcId:      vc.ID,
		ProjectNo: projectNo,
		VcHash:    hash,
		IssuerDid: template.IssuerDid,
		Signature: "签名值", // TODO: 替换为真实签名
	}
	vcEvidenceResp, err := client.VCEvidence(vcEvidenceReq)
	if err != nil {
		fmt.Println("VC哈希存证失败:", err)
		return
	}
	if vcEvidenceResp.Code != "0" {
		fmt.Println("VC哈希存证失败，错误信息:", vcEvidenceResp.Message)
		return
	}
	fmt.Println("VC哈希存证成功")

	// 7. 调用IssueVC接口签发VC
	issueVCReq := &api.IssueVCRequest{
		ProjectNo:    projectNo,
		VCTemplateId: vcTemplateId,
		VC:           vc,
		Signature:    "签名值", // TODO: 替换为真实签名
	}
	issueVCResp, err := client.IssueVC(issueVCReq)
	if err != nil {
		fmt.Println("签发VC失败:", err)
		return
	}
	if issueVCResp.Code != "0" {
		fmt.Println("签发VC失败，错误信息:", issueVCResp.Message)
		return
	}
	fmt.Println("签发VC成功，返回数据:", issueVCResp.Data)
}

// MultiKeyDIDDocumentExample 演示如何组装包含Multiple Keys的DID文档
func MultiKeyDIDDocumentExample() {
	fmt.Println("\nSBP DID SDK Go - 多密钥DID文档组装示例")
	fmt.Println("====================================")
	did := "did:example:123456789abcdefghi"

	// 假设有两个密钥，一个ECDSA，一个Ed25519
	vm1 := did.VerificationMethod{
		ID:              did + "#key-1",
		Type:            "EcdsaSecp256k1VerificationKey2019",
		Controller:      did,
		PublicKeyBase58: "2bVtQw1...ecdsaBase58...",
	}
	vm2 := did.VerificationMethod{
		ID:              did + "#key-2",
		Type:            "Ed25519VerificationKey2018",
		Controller:      did,
		PublicKeyBase58: "3fGkQw1...ed25519Base58...",
	}

	didDoc := did.AssembleMultiKeyDIDDocument(
		did,
		[]did.VerificationMethod{vm1, vm2},
		[]string{vm1.ID, vm2.ID},
		[]string{vm1.ID},
	)

	docBytes, _ := json.MarshalIndent(didDoc, "", "  ")
	fmt.Println(string(docBytes))
}

func main() {
	// 运行基本使用示例
	BasicUsageExample()

	// 运行高级使用示例
	AdvancedUsageExample()
	OpenAPIUsageExample()
}
 