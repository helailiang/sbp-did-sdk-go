package main

import (
	"fmt"
	"log"

	"github.com/sbp-did/sbp-did-sdk-go/pkg/config"
	"github.com/sbp-did/sbp-did-sdk-go/pkg/crypto"
	"github.com/sbp-did/sbp-did-sdk-go/pkg/did"
	"github.com/sbp-did/sbp-did-sdk-go/pkg/utils"
)

func main() {
	fmt.Println("SBP DID SDK Go - 演示程序")
	fmt.Println("==========================")

	// 创建配置
	cfg := createConfig()

	// 演示密钥生成 (SDK-001)
	fmt.Println("\n1. 生成密钥对 (SDK-001)")
	demoKeyGeneration(cfg)

	// 演示DID标识符计算 (SDK-002)
	fmt.Println("\n2. 计算DID标识符 (SDK-002)")
	demoDIDCalculation(cfg)

	// 演示DID文档组装 (SDK-003)
	fmt.Println("\n3. 组装DID文档 (SDK-003)")
	demoDIDDocumentAssembly(cfg)

	// 演示哈希计算 (SDK-017)
	fmt.Println("\n4. 哈希计算 (SDK-017)")
	demoHashCalculation()

	// 演示加密解密 (SDK-018, SDK-019)
	fmt.Println("\n5. 加密解密 (SDK-018, SDK-019)")
	demoEncryptionDecryption(cfg)

	// 演示签名验证 (SDK-020, SDK-021)
	fmt.Println("\n6. 签名验证 (SDK-020, SDK-021)")
	demoSignatureVerification(cfg)

	// 演示VC模板ID生成 (SDK-011)
	fmt.Println("\n7. 生成VC模板ID (SDK-011)")
	demoVCTemplateIDGeneration()

	fmt.Println("\n演示完成！")
}

func createConfig() *config.Config {
	cfg := config.NewConfig()

	// 设置华为云配置
	cfg.HuaweiCloudEndpoint = "https://your-huawei-cloud-endpoint"
	cfg.HuaweiCloudAccessKey = "your-access-key"
	cfg.HuaweiCloudSecretKey = "your-secret-key"
	cfg.HuaweiCloudRegion = "cn-north-4"

	// 设置OpenAPI配置
	cfg.OpenAPIEndpoint = "https://your-openapi-endpoint"
	cfg.ProjectID = "your-project-id"
	cfg.ProjectVisibility = "public"
	cfg.Token = ""

	// 设置算法配置
	cfg.DefaultAlgorithm = "ECDSA"
	cfg.DefaultHashAlgorithm = "SHA256"

	return cfg
}

func demoKeyGeneration(cfg *config.Config) {
	fmt.Println("生成ECDSA密钥对...")

	keyPair, err := crypto.GenerateKeyPair(cfg, "ECDSA", "demo-key")
	if err != nil {
		log.Printf("密钥生成失败: %v", err)
		return
	}

	fmt.Printf("密钥对生成成功:\n")
	fmt.Printf("  算法: %s\n", keyPair.Algorithm)
	fmt.Printf("  密钥名称: %s\n", keyPair.KeyName)

	// 获取公钥十六进制
	publicKeyHex, err := keyPair.GetPublicKeyHex()
	if err != nil {
		log.Printf("获取公钥失败: %v", err)
		return
	}
	fmt.Printf("  公钥: %s\n", publicKeyHex[:64]+"...")

	// 获取私钥十六进制
	privateKeyHex, err := keyPair.GetPrivateKeyHex()
	if err != nil {
		log.Printf("获取私钥失败: %v", err)
		return
	}
	fmt.Printf("  私钥: %s\n", privateKeyHex[:64]+"...")
}

func demoDIDCalculation(cfg *config.Config) {
	fmt.Println("计算DID标识符...")

	// 生成密钥对用于演示
	keyPair, err := crypto.GenerateKeyPair(cfg, "ECDSA", "did-demo-key")
	if err != nil {
		log.Printf("密钥生成失败: %v", err)
		return
	}

	// 计算DID标识符
	didIdentifier, err := did.CalculateDIDIdentifier(keyPair, "did:sbp:")
	if err != nil {
		log.Printf("DID标识符计算失败: %v", err)
		return
	}

	fmt.Printf("DID标识符: %s\n", didIdentifier)

	// 验证DID标识符
	if err := did.ValidateDIDIdentifier(didIdentifier); err != nil {
		log.Printf("DID标识符验证失败: %v", err)
		return
	}
	fmt.Println("DID标识符验证通过")
}

func demoDIDDocumentAssembly(cfg *config.Config) {
	fmt.Println("组装DID文档...")

	// 生成密钥对
	keyPair, err := crypto.GenerateKeyPair(cfg, "ECDSA", "doc-demo-key")
	if err != nil {
		log.Printf("密钥生成失败: %v", err)
		return
	}

	// 计算DID标识符
	didIdentifier, err := did.CalculateDIDIdentifier(keyPair, "did:sbp:")
	if err != nil {
		log.Printf("DID标识符计算失败: %v", err)
		return
	}

	// 业务属性
	businessAttributes := map[string]interface{}{
		"name":         "张三",
		"email":        "zhangsan@example.com",
		"organization": "示例组织",
	}

	// 组装DID文档
	doc, err := did.AssembleDIDDocument(cfg, keyPair, "ECDSA", didIdentifier, businessAttributes)
	if err != nil {
		log.Printf("DID文档组装失败: %v", err)
		return
	}

	fmt.Printf("DID文档组装成功:\n")
	fmt.Printf("  ID: %s\n", doc.ID)
	fmt.Printf("  创建时间: %s\n", doc.Created)
	fmt.Printf("  验证方法数量: %d\n", len(doc.VerificationMethod))
	fmt.Printf("  业务属性数量: %d\n", len(doc.CustomFields))

	// 转换为JSON
	jsonData, err := doc.ToJSON()
	if err != nil {
		log.Printf("DID文档JSON转换失败: %v", err)
		return
	}

	fmt.Printf("  JSON长度: %d 字节\n", len(jsonData))
}

func demoHashCalculation() {
	fmt.Println("计算哈希值...")

	testData := "Hello, SBP DID SDK!"

	// 使用SHA256计算哈希
	sha256Hash, err := utils.CalculateHashFromString(testData, utils.SHA256)
	if err != nil {
		log.Printf("SHA256哈希计算失败: %v", err)
		return
	}

	fmt.Printf("原始数据: %s\n", testData)
	fmt.Printf("SHA256哈希: %s\n", sha256Hash)

	// 使用SM3计算哈希
	sm3Hash, err := utils.CalculateHashFromString(testData, utils.SM3)
	if err != nil {
		log.Printf("SM3哈希计算失败: %v", err)
		return
	}

	fmt.Printf("SM3哈希: %s\n", sm3Hash)
}

func demoEncryptionDecryption(cfg *config.Config) {
	fmt.Println("演示加密解密...")

	// 生成RSA密钥对
	keyPair, err := crypto.GenerateKeyPair(cfg, "RSA", "encryption-demo-key")
	if err != nil {
		log.Printf("密钥生成失败: %v", err)
		return
	}

	plainText := "这是需要加密的敏感数据"

	// 加密
	encryptionResult, err := crypto.Encrypt(cfg, keyPair, []byte(plainText), "RSA")
	if err != nil {
		log.Printf("加密失败: %v", err)
		return
	}

	fmt.Printf("原始数据: %s\n", plainText)
	fmt.Printf("加密结果: %s\n", encryptionResult.EncryptedData[:64]+"...")

	// 解密
	encryptedData, err := crypto.DecryptFromHex(cfg, keyPair, encryptionResult.EncryptedData, "RSA")
	if err != nil {
		log.Printf("解密失败: %v", err)
		return
	}

	fmt.Printf("解密结果: %s\n", encryptedData.DecryptedData)
}

func demoSignatureVerification(cfg *config.Config) {
	fmt.Println("演示签名验证...")

	// 生成ECDSA密钥对
	keyPair, err := crypto.GenerateKeyPair(cfg, "ECDSA", "signature-demo-key")
	if err != nil {
		log.Printf("密钥生成失败: %v", err)
		return
	}

	dataToSign := "这是需要签名的数据"

	// 签名
	signatureResult, err := crypto.Sign(cfg, keyPair, []byte(dataToSign), "ECDSA")
	if err != nil {
		log.Printf("签名失败: %v", err)
		return
	}

	fmt.Printf("原始数据: %s\n", dataToSign)
	fmt.Printf("签名结果: %s\n", signatureResult.Signature[:64]+"...")

	// 验证签名
	verificationResult, err := crypto.VerifySignatureFromHex(cfg, keyPair, dataToSign, signatureResult.Signature, "ECDSA")
	if err != nil {
		log.Printf("签名验证失败: %v", err)
		return
	}

	fmt.Printf("验证结果: %t\n", verificationResult.Valid)
	if !verificationResult.Valid {
		fmt.Printf("验证消息: %s\n", verificationResult.Message)
	}
}

func demoVCTemplateIDGeneration() {
	fmt.Println("生成VC模板ID...")

	// 生成多个VC模板ID
	for i := 0; i < 3; i++ {
		templateID := utils.GenerateVCTemplateID()
		fmt.Printf("VC模板ID %d: %s\n", i+1, templateID)

		// 验证UUID格式
		if utils.ValidateUUID(templateID) {
			fmt.Printf("  UUID格式验证通过\n")
		} else {
			fmt.Printf("  UUID格式验证失败\n")
		}
	}
}
