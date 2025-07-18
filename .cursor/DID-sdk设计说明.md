1. #### SDK-001 Generate Public and Private Key Pair

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-001                                                      |
| Function Name                | Generate Public and Private Key Pair                         |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | This method can return a pair of public and private keys of Elliptic Curve Digital Signature Algorithm ("ECDSA"), Rivest, Shamir, Adleman ("RSA" and ShangMi2 ("SM2") types of public and private keys. |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：用户需要配置DCI接入地址和账号。功能说明：支持通过SDK直接调用DCI的HSM创建密钥检查逻辑：检查是否配置DCI账户是可用。检查密钥名称是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限控制                                                   |

1. #### SDK-002 Calculate DID Identifier

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-002                                                      |
| Function Name                | Calculate DID Identifier                                     |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | The method requires passing a public key of ECDSA, RSA and SM2 and DID Method. The method will internally hash the public key and return the full DID identifier. |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：用户需要配置华为云接入地址和账号。功能说明：支持通过SDK生成DID标识符检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查DID Method是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限控制                                                   |

1. #### SDK-003 Assemble DID Document

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-003                                                      |
| Function Name                | Assemble DID Document                                        |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | This method requires the public key, algorithm, DID identifier, and business attribute field value to be written into the DID Document, and returns the unsigned DID Document. |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：用户需要配置华为云接入地址和账号。功能说明：支持通过SDK对DID文档进行组装。检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查DID标识符是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限控制                                                   |

1. #### SDK-004 Proof DID Document

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-004                                                      |
| Function Name                | Proof DID Document                                           |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 入参需要包含项目可见性（私有项目需要Token）、项目编号，那么SDK内部方法执行完成后自动调用API。该方法是对组装的DID文档进行Proof签名（支持HSM），并组装带有proof属性的完整DID文档，然后调用API将DID文档上链。 |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址及华为云接入地址和账号。功能说明：支持通过SDK调用OpenAPI进行DID文档的注册和更新。检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查项目编号是否为空值。检查项目可见性是私有还是公开，如项目为私有，则需要检查Token是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-005 Query DID Document

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-005                                                      |
| Function Name                | Query DID Document                                           |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 调用时提供DID标识、项目编号、项目可见性，SDK调用API查询DID文档。 |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址。检查逻辑：检查项目编号是否为空值。检查项目可见性是私有还是公开，如项目为私有，则需要检查Token是否为空值。检查DID标识符是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-006 Register Issuer

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-006                                                      |
| Function Name                | Register Issuer                                              |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 调用时提供发证方标识符、发证方名称、发证方资料、DID私钥（支持HSM）、项目编号及项目可见性，SDK将组装调用参数并使用私钥签名，然后调用API的注册发证方方法。 |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址及华为云接入地址和账号。检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查项目编号是否为空值。检查项目可见性是私有还是公开，如项目为私有，则需要检查Token是否为空值。检查待更新发证方数据信息是否为空值，包括发证方名称、发证方资料。检查发证方标识符是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-007 Update Issuer

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-007                                                      |
| Function Name                | Update Issuer                                                |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 调用时提供发证方标识符、发证方名称、发证方资料、DID私钥（支持HSM）、项目编号及项目可见性，SDK将组装调用参数并使用私钥签名，然后调用API的更新发证方方法。 |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址及华为云接入地址和账号。检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查项目编号是否为空值。检查项目可见性是私有还是公开，如项目为私有，则需要检查Token是否为空值。检查发证方标识符是否为空值。检查待更新发证方数据信息是否为空值，包括发证方名称、发证方资料。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-008 Query Issuer

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-008                                                      |
| Function Name                | Query Issuer                                                 |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 调用时提供项目编号、发证方标识符，SDK调用API的查询发证方方法，返回发证方查询结果。 |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址。检查逻辑：检查项目编号是否为空值。检查项目可见性是私有还是公开，如项目为私有，则需要检查Token是否为空值。检查发证方标识符是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-009 Assemble VC Template

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-009                                                      |
| Function Name                | Assemble VC Template                                         |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | This method requires the Issuer DID identifier, custom VC business attribute field constraints, template title, and template description and it returns the VC template content without a signature. |
| Mode                         | Online                                                       |
| Business Rules               | 检查逻辑：检查密钥名称是否为空值。检查发证方标识符是否为空值。检查待注册的VC模板数据是为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限验证                                                   |

1. #### SDK-010 Proof VC Template

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-010                                                      |
| Function Name                | Proof VC Template                                            |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 入参如包含了Token（公开项目不需要）、API地址、项目ID，那么SDK内部方法执行完成后自动调用API。该方法是对组装的VC模版进行Proof签名（支持HSM），并组装带有proof属性的VC模版内容，然后调用API将VC模版上链。 |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址及华为云接入地址和账号。检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查项目编号是否为空值。检查项目可见性是私有还是公开，如项目为私有，则需要检查Token是否为空值。检查发证方标识符是否为空值。检查待注册或待更新的VC模板数据是为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-011 Generate VC Template ID

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-011                                                      |
| Function Name                | Generate VC Template ID                                      |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 调用该方法将随机生成一个 UUIDv4 格式的字符串，作为凭证模版ID |
| Mode                         | Online                                                       |
| Business Rules               |                                                              |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限验证                                                   |

1. #### SDK-012 Query VC Template

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-012                                                      |
| Function Name                | Query VC Template                                            |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 调用时提供项目编号、项目可见性、发证方标识符及VC模板ID，SDK将然后调用API查询VC模板方法，返回查询结果。 |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址。检查逻辑：检查项目编号是否为空值。检查项目可见性是私有还是公开，如项目为私有，则需要检查Token是否为空值。检查发证方标识符是否为空值。检查VC模板ID是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-013 Issue VC

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-013                                                      |
| Function Name                | Issue VC                                                     |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 该方法需要发证方提供项目编号、项目可见性、发证方标识符、VC模板标识、Credential Subjects以及过期时间，通过封装相关方法自动组装VC并使用用户提供的私钥进行签名。同时，还可以在签发后调用API进行VC Hash的存证。 |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址及华为云接入地址和账号。检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查项目编号是否为空值。检查项目可见性是私有还是公开，如项目为私有，则需要检查Token是否为空值。检查发证方标识符是否为空值。检查需要生成VC对应的Credential Subjects是否为空值。检查VC模板ID是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-014 Verify VC

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-014                                                      |
| Function Name                | Verify VC                                                    |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 该方法需要验证方提供项目编号、项目可见性、发证方标识符、VC原文及VC ID，通过调用API查询验证VC所需数据，然后对VC的Proof、存证以及吊销状态进行核验。 |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：用户需要配置OpenAPI地址。检查逻辑：检查项目编号是否为空值。检查项目可见性是私有还是公开，如项目为私有，则需要检查Token是否为空值。检查发证方标识符是否为空值。检查VC原文数据是否为空值。检查VC ID是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-015 Generate VP

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-015                                                      |
| Function Name                | Generate VP                                                  |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         |                                                              |
| Mode                         | Online                                                       |
| Business Rules               |                                                              |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限验证                                                   |

1. #### SDK-016 Verify VP

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-016                                                      |
| Function Name                | Verify VP                                                    |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 通过SDK对VP中的发证方签名和持证人的签名进行签名。            |
| Mode                         | Online                                                       |
| Business Rules               | 通过SDK对VP中的发证方签名和持证人的签名进行签名。            |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-017 Calculate Hash

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-017                                                      |
| Function Name                | Calculate Hash                                               |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | This method requires VC original text. The method will perform hash calculation on the VC original text according to the SHA256 or SM3 algorithm and return the hash data. |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址及华为云接入地址和账号。功能说明：支持通过SDK对明文数据进行哈希计算。检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查待生成哈希的明文数据是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限验证                                                   |

1. #### SDK-018 Encryption

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-018                                                      |
| Function Name                | Encryption                                                   |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | This method requires the public key of ECDSA, RSA and SM2 and the plain text data to be encrypted. The method will use the public key to encrypt the plain text data according to the algorithm type and return the encrypted data. |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：用户需要配置华为云接入地址和账号。功能说明：支持通过SDK直接调用华为云的HSM进行加密检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查待加密的明文数据是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限验证                                                   |

1. #### SDK-019 Decryption

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-019                                                      |
| Function Name                | Decryption                                                   |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | This method requires the private key of ECDSA, RSA and SM2 and encrypted data. The method will use the private key to decrypt the encrypted data according to the algorithm type and return the plain text data. |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：用户需要配置华为云接入地址和账号。功能说明：支持通过SDK直接调用华为云的HSM进行解密检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查待密钥的密文数据是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限验证                                                   |

1. #### SDK-020 Signature

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-020                                                      |
| Function Name                | Signature                                                    |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | This method requires the private key of ECDSA, RSA and SM2 and the data to be signed. The method will use the private key to sign the data according to the algorithm type and return the signature value. |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：用户需要配置华为云接入地址和账号。功能说明：支持通过SDK直接调用华为云的HSM进行签名检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查待签名的明文数据是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限验证                                                   |

1. #### SDK-021 Verify Signature

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-021                                                      |
| Function Name                | Verify Signature                                             |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | This method requires the public key of ECDSA, RSA and SM2, the data of signature and the signature value. It will use a public key to verify the signature result data according to the algorithm type and return the verification result |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：用户需要配置华为云接入地址和账号。功能说明：支持通过SDK直接调用华为云的HSM进行验签检查逻辑：检查是否配置华为云账户是可用。检查密钥名称是否为空值。检查待验签的明文数据是否为空值。检查签名后的签名数据是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | 无权限验证                                                   |

1. #### SDK-022 Query VC Hash

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-022                                                      |
| Function Name                | Query VC Hash                                                |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 支持通过SDK调用OpenAPI进行VC存证哈希查询。                   |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址。功能说明：支持通过SDK调用OpenAPI进行VC存证哈希查询。检查逻辑：检查项目可见性是公开或私有，如项目可见性为私有则需要检验Token是否为空值。检查DID项目编号是否为空值。检查VC ID是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |

1. #### SDK-023 Query VC Status

| **Item**                     | **Description**                                              |
| ---------------------------- | ------------------------------------------------------------ |
| Function ID                  | SDK-023                                                      |
| Function Name                | Query VC Status                                              |
| Category                     | SDK                                                          |
| Related Requirements         | F-REQ-AS-006                                                 |
| Function Description         | 支持通过SDK调用OpenAPI进行VC状态查询。                       |
| Mode                         | Online                                                       |
| Business Rules               | 前置条件：需要配置OpenAPI地址。功能说明：支持通过SDK调用OpenAPI进行VC状态查询。检查逻辑：检查项目可见性是公开或私有，如项目可见性为私有则需要检验Token是否为空值。检查DID项目编号是否为空值。检查VC ID是否为空值。检查发证方DID是否为空值。 |
| User Input Screens and Forms | As some layout design changes may occur, the image below may differ from the final one. <<Screens Cap>> |
| Security Requirements        | SDK无权限验证，访问API的权限验证参考相关API设计              |