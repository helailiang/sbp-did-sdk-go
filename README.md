# SBP DID SDK Go

## 项目概述

SBP DID SDK Go 是一个基于Go语言开发的去中心化身份（Decentralized Identifier, DID）软件开发工具包。该SDK提供了完整的DID生命周期管理功能，包括密钥生成、DID文档创建、凭证模板管理、可验证凭证（VC）签发与验证、可验证表示（VP）生成与验证等核心功能。

### 已实现功能

✅ **SDK-001**: 生成公钥和私钥对 (ECDSA, RSA, SM2)  
✅ **SDK-002**: 计算DID标识符  
✅ **SDK-003**: 组装DID文档  
✅ **SDK-004**: DID文档证明（OpenAPI接口封装）  
✅ **SDK-005**: 查询DID文档（OpenAPI接口封装）  
✅ **SDK-006**: 注册发证方（OpenAPI接口封装）  
✅ **SDK-007**: 更新发证方（OpenAPI接口封装）  
✅ **SDK-008**: 查询发证方（OpenAPI接口封装）  
✅ **SDK-009**: 组装VC模板  
✅ **SDK-010**: 证明VC模板（OpenAPI接口封装）  
✅ **SDK-011**: 生成VC模板ID (UUIDv4)  
✅ **SDK-012**: 查询VC模板（OpenAPI接口封装）  
✅ **SDK-013**: 签发VC（OpenAPI接口封装）  
✅ **SDK-014**: 验证VC（OpenAPI接口封装）  
✅ **SDK-015**: 生成VP  
✅ **SDK-016**: 验证VP  
✅ **SDK-017**: 计算哈希值 (SHA256, SM3)  
✅ **SDK-018**: 加密数据  
✅ **SDK-019**: 解密数据  
✅ **SDK-020**: 签名数据  
✅ **SDK-021**: 验证签名  
✅ **SDK-022**: 查询VC哈希（OpenAPI接口封装）  
✅ **SDK-023**: 查询VC状态（OpenAPI接口封装）

> **说明：** 所有链上相关操作均已通过OpenAPI接口进行RESTful封装，SDK方法与官方文档100%一致。

### OpenAPI接口封装
- 所有DID、Issuer、VC模板、VC等链上操作均通过`pkg/api`目录下的RESTful客户端进行封装。
- 支持Token鉴权，所有接口均严格使用POST，token放header，参数放body，响应结构体与官方文档一致。
- 具体接口包括：DID注册/更新/查询，Issuer注册/更新/查询/状态变更，VC模板注册/更新/查询/状态变更，VC签发/存证/查询/吊销/核验/状态查询等。

## 核心功能

### 1. 密钥管理 (SDK-001)
- **功能**: 生成公钥和私钥对
- **支持算法**: ECDSA、RSA、SM2
- **特点**: 支持华为云HSM硬件安全模块
- **前置条件**: 需要配置DCI接入地址和账号

### 2. DID标识符计算 (SDK-002)
- **功能**: 基于公钥计算DID标识符
- **支持算法**: ECDSA、RSA、SM2
- **输入**: 公钥、DID Method
- **输出**: 完整的DID标识符

### 3. DID文档组装与上链 (SDK-003, SDK-004, SDK-005)
- **组装DID文档**: 输入公钥、算法、DID标识符、业务属性字段值，输出未签名DID文档
- **DID文档证明/注册/查询**: 通过OpenAPI接口完成Proof签名、上链、查询等操作
- **接口方法**：`RegisterDID`, `QueryDID`, `UpdateDID`

### 4. 发证方管理 (SDK-006, SDK-007, SDK-008)
- **注册/更新/查询发证方**: 通过OpenAPI接口完成发证方信息的注册、更新、查询、状态变更
- **接口方法**：`RegisterIssuer`, `UpdateIssuer`, `QueryIssuer`, `IssuerStatus`

### 5. VC模板管理 (SDK-009, SDK-010, SDK-011, SDK-012)
- **组装VC模板**: 创建VC模板内容
- **注册/查询/更新/状态变更VC模板**: 通过OpenAPI接口完成VC模板的注册、查询、更新、状态变更
- **接口方法**：`RegisterVCTemplate`, `QueryVCTemplate`, `UpdateVCTemplate`, `VCTemplateStatus`

### 6. 可验证凭证管理 (SDK-013, SDK-014, SDK-022, SDK-023)
- **签发/存证/查询/吊销/核验VC**: 通过OpenAPI接口完成VC的签发、存证、查询、吊销、核验、状态查询
- **接口方法**：`IssueVC`, `VCEvidence`, `QueryVCEvidence`, `VCRevoke`, `VCRevokeStatus`, `VCVerify`

### 7. 可验证表示管理 (SDK-015, SDK-016)
- **生成/验证VP**: 组装和验证可验证表示

### 8. 加密解密功能 (SDK-018, SDK-019)
- **加密**: 使用公钥加密数据
- **解密**: 使用私钥解密数据
- **支持算法**: ECDSA、RSA、SM2

### 9. 签名验证功能 (SDK-020, SDK-021)
- **签名**: 使用私钥对数据进行签名
- **验证签名**: 使用公钥验证签名
- **支持算法**: ECDSA、RSA、SM2

### 10. 哈希计算 (SDK-017)
- **功能**: 对数据进行哈希计算
- **支持算法**: SHA256、SM3

## 项目架构

```
sbp-did-sdk-go/
├── README.md                 # 项目说明文档
├── go.mod                    # Go模块定义
├── go.sum                    # 依赖校验文件
├── cmd/                      # 命令行工具
│   └── main.go              # 主程序入口
├── pkg/                      # 核心包
│   ├── config/              # 配置管理
│   ├── crypto/              # 加密算法实现
│   ├── did/                 # DID相关功能
│   ├── vc/                  # 可验证凭证功能
│   ├── vp/                  # 可验证表示功能
│   ├── issuer/              # 发证方管理
│   ├── template/            # 模板管理
│   ├── api/                 # OpenAPI接口封装（链上操作核心）
│   └── utils/               # 工具函数
├── examples/                 # 使用示例
├── tests/                    # 测试文件
└── docs/                     # 文档
```

## 安装和使用

### 环境要求
- Go 1.19+
- 华为云账号和配置
- DCI接入配置（可选）

### 安装
```bash
go get github.com/your-org/sbp-did-sdk-go
```

### 基本使用示例

```go
package main

import (
    "fmt"
    "log"
    "github.com/your-org/sbp-did-sdk-go/pkg/config"
    "github.com/your-org/sbp-did-sdk-go/pkg/crypto"
    "github.com/your-org/sbp-did-sdk-go/pkg/did"
)

func main() {
    // 初始化配置
    cfg := &config.Config{
        HuaweiCloudEndpoint: "https://your-huawei-cloud-endpoint",
        HuaweiCloudAccessKey: "your-access-key",
        HuaweiCloudSecretKey: "your-secret-key",
        OpenAPIEndpoint: "https://your-openapi-endpoint",
        ProjectID: "your-project-id",
        ProjectVisibility: "public", // or "private"
        Token: "", // 私有项目需要提供Token
    }
    // 生成密钥对
    keyPair, err := crypto.GenerateKeyPair(cfg, "ECDSA", "my-key")
    if err != nil {
        log.Fatal(err)
    }
    // 计算DID标识符
    didIdentifier, err := did.CalculateDIDIdentifier(keyPair.PublicKey, "did:sbp:")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Generated DID: %s\n", didIdentifier)
}
```

### OpenAPI接口调用示例

```go
import (
    "fmt"
    "github.com/your-org/sbp-did-sdk-go/pkg/api"
)

func main() {
    // 1. 初始化API客户端
    client := api.NewClient("https://your-openapi-endpoint", "your-token")
    projectNo := "your-project-id"

    // 2. 注册DID文档
    didReq := &api.RegisterDIDRequest{
        ProjectNo:   projectNo,
        DIDDocument: "{...}", // DID文档JSON字符串
        Signature:   "签名值",
    }
    didResp, err := client.RegisterDID(didReq)
    if err != nil {
        fmt.Println("注册DID失败:", err)
    } else {
        fmt.Println("注册DID响应:", didResp)
    }

    // 3. 查询DID文档
    queryReq := &api.QueryDIDRequest{
        DID:       "did:sbp:xxx",
        ProjectNo: projectNo,
    }
    queryResp, err := client.QueryDID(queryReq)
    if err != nil {
        fmt.Println("查询DID失败:", err)
    } else {
        fmt.Println("查询DID响应:", queryResp)
    }

    // 4. 注册发证方
    issuerReq := &api.RegisterIssuerRequest{
        IssuerDid:        "did:sbp:issuer001",
        IssuerName:       "示例发证方",
        ProjectNo:        projectNo,
        ContactPerson:    "张三",
        ContactNumber:    "123456789",
        ContactEmail:     "issuer@example.com",
        BusinessScenario: "示例业务",
        Signature:        "签名值",
    }
    issuerResp, err := client.RegisterIssuer(issuerReq)
    if err != nil {
        fmt.Println("注册发证方失败:", err)
    } else {
        fmt.Println("注册发证方响应:", issuerResp)
    }

    // 5. 签发VC
    vcReq := &api.IssueVCRequest{
        // 按VC结构体补充字段
    }
    vcResp, err := client.IssueVC(vcReq)
    if err != nil {
        fmt.Println("签发VC失败:", err)
    } else {
        fmt.Println("签发VC响应:", vcResp)
    }
}
```

### 主要API方法与参数说明

- `RegisterDID(req *RegisterDIDRequest) (*CommonResponse, error)`
- `QueryDID(req *QueryDIDRequest) (*QueryDIDResponse, error)`
- `UpdateDID(req *UpdateDIDRequest) (*UpdateDIDResponse, error)`
- `RegisterIssuer(req *RegisterIssuerRequest) (*CommonResponse, error)`
- `QueryIssuer(req *QueryIssuerRequest) (*QueryIssuerResponse, error)`
- `UpdateIssuer(req *UpdateIssuerRequest) (*UpdateIssuerResponse, error)`
- `IssuerStatus(req *IssuerStatusRequest) (*IssuerStatusResponse, error)`
- `RegisterVCTemplate(req *RegisterVCTemplateRequest) (*CommonResponse, error)`
- `QueryVCTemplate(req *QueryVCTemplateRequest) (*QueryVCTemplateResponse, error)`
- `UpdateVCTemplate(req *UpdateVCTemplateRequest) (*UpdateVCTemplateResponse, error)`
- `VCTemplateStatus(req *VCTemplateStatusRequest) (*VCTemplateStatusResponse, error)`
- `IssueVC(req *IssueVCRequest) (*IssueVCResponse, error)`
- `VCEvidence(req *VCEvidenceRequest) (*VCEvidenceResponse, error)`
- `QueryVCEvidence(req *QueryVCEvidenceRequest) (*QueryVCEvidenceResponse, error)`
- `VCRevoke(req *VCRevokeRequest) (*VCRevokeResponse, error)`
- `VCRevokeStatus(req *VCRevokeStatusRequest) (*VCRevokeStatusResponse, error)`
- `VCVerify(req *VCVerifyRequest) (*VCVerifyResponse, error)`

详细参数结构体定义请见 [`pkg/api/types.go`](pkg/api/types.go)。

## 配置说明

### 华为云配置
- `HuaweiCloudEndpoint`: 华为云服务端点
- `HuaweiCloudAccessKey`: 访问密钥
- `HuaweiCloudSecretKey`: 秘密密钥

### OpenAPI配置
- `OpenAPIEndpoint`: OpenAPI服务端点
- `ProjectID`: 项目编号
- `ProjectVisibility`: 项目可见性（public/private）
- `Token`: 私有项目的访问令牌

## 安全特性

1. **HSM支持**: 支持华为云硬件安全模块
2. **多种算法**: 支持ECDSA、RSA、SM2等加密算法
3. **签名验证**: 完整的数字签名和验证机制
4. **权限控制**: 基于Token的访问控制

## 开发指南

### 添加新功能
1. 在相应的包中创建新文件
2. 实现功能逻辑
3. 添加单元测试
4. 更新文档

### 运行测试
```bash
go test ./...
```

### 构建
```bash
go build -o sbp-did-sdk cmd/main.go
```

## 版本变更与改进

- **2024-06**：完成全部OpenAPI接口的RESTful封装，所有链上操作均可通过SDK方法直接调用，结构体和方法与官方文档100%一致。
- 后续如有新接口、用法或改进建议，欢迎随时反馈与补充。

## 许可证

[许可证信息]

## 联系方式

[联系信息] 