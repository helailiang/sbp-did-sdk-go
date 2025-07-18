
**SBP-DID OpenAPI接口设计说明文档**

<a name="heading_0"></a>1. **设计说明**

`  `OpenAPI面向运营人员及SDK提供了多个开放的API接口，主要涉及的是DID权限、DID文档、发证方以及VC相关的API接口。包括DID项目所对应的合约访问权限上链、DID文档的查询、注册以及更新，发证方的查询、注册、更新以及启用或禁用，VC模板的查询、注册、更新以及启用或禁用，VC存证的查询和保存，VC的核验、吊销、吊销状态查询以及启用或禁用等API。

<a name="heading_1"></a>2. **接口设计**

<a name="heading_2"></a>2.1 **内部接口**

<a name="heading_3"></a>2.1.1 **项目管理**

<a name="heading_4"></a>2.1.1.1 **获取Token**

增加token有效期

|**接口地址**|/api/sys/v1/getToken|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>获取Token，处理逻辑如下：</p><p>- 验证签名值是否正确</p><p>- 验证项目状态是否启用</p><p>- 验证通过返回Token</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|签名|signature|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|`      `1|项目编号|projectNo|String|Y||
|`      `2|用户名|clientName|String|Y||
|`     `3|是否是项目owner|isProjectOwner|Boolean|N|true=是，false=不是，默认为false|
|**响应参数**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|Y||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": {},<br>`    `"message": ""<br>}||||||

<a name="heading_5"></a>2.1.1.2 **~~更新Token（弃用）~~**

|**~~接口地址~~**|~~/api/sys/v1/updateToken~~|||||
| :-: | :- | :- | :- | :- | :- |
|**~~接口描述~~**|<p>~~更新Token，处理逻辑如下：~~</p><p>- ~~验证签名值是否正确~~</p><p>- ~~验证Token是否在启用中~~</p><p>- ~~验证项目状态是否启用~~</p><p>- ~~验证token和用户名是否匹配~~</p><p>- ~~验证通过返回新的Token~~</p>|**~~调用方式~~** |~~POST~~|||
|**~~请求参数~~**||||||
|**~~Header~~**||||||
|**~~序号~~**|**~~字段名~~**|**~~字段~~**|**~~类型~~**|**~~必填~~**|**~~备注~~**|
|~~1~~|~~签名~~|~~signature~~|~~String~~|~~Y~~||
|~~2~~|~~令牌~~|~~token~~|~~String~~|~~Y~~||
|**~~Body~~**||||||
|**~~序号~~**|**~~字段名~~**|**~~字段~~**|**~~类型~~**|**~~必填~~**|**~~备注~~**|
|`      `~~1~~|~~项目编号~~|~~projectNo~~|~~String~~|~~Y~~||
|`      `~~2~~|~~用户名~~|~~clientName~~|~~String~~|~~Y~~||
|**~~响应参数~~**||||||
|**~~序号~~**|**~~字段名~~**|**~~字段~~**|**~~类型~~**|**~~必填~~**|**~~备注~~**|
|~~1~~|~~新的令牌~~|~~token~~|~~String~~|~~Y~~||
|**~~响应示例~~**||||||
|~~{~~<br>`    `~~"code": "0",~~<br>`    `~~"data": {},~~<br>`    `~~"message": ""~~<br>~~}~~||||||

<a name="heading_6"></a>2.1.1.3 **项目启用/停用**

|**接口地址**|/api/sys/v1/project/enable|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>项目启用，处理逻辑如下：</p><p>- 验证签名值是否正确</p><p>- 验证Token是否在启用中</p><p>- 验证项目状态是否已启用/已禁用</p><p>- 验证通过调用合约修改项目为已启用/已禁用</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|签名|signature|String|Y||
|~~2~~|~~令牌~~|~~token~~|~~String~~|~~Y~~||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|项目状态|status|Int|Y|1=启用 2=停用|
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_7"></a>2.1.1.4 **更改项目可见性**

|**接口地址**|/api/sys/v1/project/visibility/update|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>更改项目为公开状态，处理逻辑如下：</p><p>- 验证签名值是否正确</p><p>- ~~验证Token是否在启用中~~</p><p>- 验证项目状态是否启用</p><p>- ~~验证是否项目owner发起更改~~</p><p>- 验证通过调用合约修改项目为公开状态</p><p>&emsp;更改项目为私有状态，处理逻辑如下：</p><p>- 验证签名值是否正确</p><p>- ~~验证Token是否在启用中~~</p><p>- 验证项目状态是否启用</p><p>- ~~验证是否项目owner发起更改~~</p><p>- 验证通过调用合约修改项目为私有状态</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|签名|signature|String|Y||
|2|令牌|token|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|项目可见性|visibility|Int|Y|1=公开 2=私有|
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_8"></a>2.1.1.5 **更改项目method**

|**接口地址**|/api/sys/v1/project/method/update|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>更改项目method，处理逻辑如下：</p><p>- 验证签名值是否正确</p><p>- ~~验证Token是否在启用中~~</p><p>- 验证项目状态是否为启用状态</p><p>- ~~验证是否项目owner发起更改~~</p><p>- 调用合约修改Method</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|签名|signature|String|Y||
|~~2~~|~~令牌~~|~~token~~|~~String~~|~~Y~~||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|Method名称|method|String|Y||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_9"></a>2.1.1.6 **更改Issuer、VC模版审核配置**

|**接口地址**|/api/sys/v1/project/review|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>更改Issuer、VC模版审核配置，处理逻辑如下：</p><p>- 验证签名值是否正确</p><p>- ~~验证Token是否在启用中~~</p><p>- 验证项目状态是否为启用状态</p><p>- ~~验证是否项目owner发起更改~~</p><p>- 调用合约修改发证方审核、VC模板审核配置</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|签名|signature|String|Y||
|2|~~令牌~~|~~token~~|~~String~~|~~Y~~||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|发证方审核|issuerReview|Boolean|Y|true=开启审核，false=关闭审核|
|3|VC模板审核|vcTemplateReview|Boolean|Y|true=开启审核，false=关闭审核|
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_10"></a>2.1.1.7 **更改项目管理员**

|**接口地址**|/api/sys/v1/project/owner/update|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>更改项目管理员，处理逻辑如下：</p><p>- 验证签名值是否正确</p><p>- ~~验证Token是否在启用中~~</p><p>- 验证项目状态是否启用或者停用</p><p>- ~~验证是否项目owner发起更改~~</p><p>- 验证通过调用合约更改项目owner</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|签名|signature|String|Y||
|2|~~令牌~~|~~token~~|~~String~~|~~Y~~||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|用户名|clientName|String|Y||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_11"></a>2.1.1.8 **批量更改函数选择器**

|**接口地址**|/api/sys/v1/project/selector/update|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>批量授权函数选择器，处理逻辑如下：</p><p>- 验证签名值是否正确</p><p>- 验证Token是否在启用中</p><p>- 验证项目状态</p><p>- ~~验证是否项目owner发起更改~~</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|签名|signature|String|Y||
|2|~~令牌~~|~~token~~|~~String~~|~~Y~~||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|用户名|clientName|String|N||
|3|链账户|account|String|N||
|4|函数选择器列表|accountSelectors|List|Y||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_12"></a>2.1.1.9 **更改项目成员状态**

|**接口地址**|/api/sys/v1/project/member/update|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>批量授权函数选择器，处理逻辑如下：</p><p>- 验证签名值是否正确</p><p>- ~~验证Token是否在启用中~~</p><p>- 验证项目状态</p><p>- 验证项目成员状态</p><p>- ~~验证是否项目owner发起更改~~</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|签名|signature|String|Y||
|2|~~令牌~~|~~token~~|~~String~~|~~Y~~||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|用户名|clientName|String|N||
|3|链账户|account|String|N||
|5|状态|status|Int|Y|1=启用 2=禁用|
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_13"></a>2.1.1.10 **发证方、VC模板审核结果回调**

|**接口地址**|/api/sys/v1/project/review/callback|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>审核结果回调，处理逻辑如下：</p><p>- 验证签名值是否正确</p><p>- 查询申请记录是否存在</p><p>- 判断审核状态是否为待审核</p><p>- 更新数据库表审核状态及意见</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|1|签名|signature|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|审核单号|auditNo|String|Y||
|3|审核结果|auditResult|Boolean|Y|true=审核通过 false=审核未通过|
|4|`            `审核意见|auditComments|String|Y||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_14"></a>2.2 **外部接口**

<a name="heading_15"></a>2.2.1 **DID**

<a name="heading_16"></a>2.2.1.1 **注册DID**

|**接口地址**|/api/sys/v1/did/register|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>注册DID，处理逻辑如下：</p><p>- 验证Token是否在启用中（公开项目不需要）</p><p>- 验证DID项目是否在启用中</p><p>- 验证用户是否有向对应项目写入DID的权限</p><p>- 验证DID标识符中的DID Methods是否符合项目配置</p><p>- 验证上链交易签名模式</p><p>- 验证DID 文档的格式是否符合SBP-DID平台的统一要求</p><p>&emsp;验证通过后将DID文档写入DID合约。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|DID文档|didDocument|Text|Y||
|3|业务数据签名值|signature|String|Y||
|4|交易数据签名值|txSignature|String|N||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_17"></a>2.2.1.2  **更新DID文档**

|**接口地址**|/api/sys/v1/did/update|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>更新DID文档，处理逻辑如下：</p><p>- 验证Token是否在启用中（公开项目不需要）</p><p>- 验证DID项目是否在启用中</p><p>- 验证用户是否有向对应项目写入DID的权限</p><p>- 验证DID标识符中的DID Methods是否符合项目配置</p><p>- 验证DID标识是否属于对应的项目</p><p>- 验证DID 文档的格式是否符合SBP-DI平台的统一要求</p><p>&emsp;验证通过后将新的DID文档更新到区块链。 </p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|DID文档|didDocument|Text|Y||
|3|公钥索引|index||||
|4|业务数据签名值|signature|String|Y||
|5|交易数据签名值|txSignature|String|N||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_18"></a>2.2.1.3 **查询DID文档**

|**接口地址**|/api/sys/v1/did/search|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>查询DID文档，处理逻辑如下：</p><p>- 公开项目，用户查询DID文档需要提供项目编号、DID标识；</p><p>- 验证DID项目是否在启用中</p><p>- 验证项目中是否存在对应的DID</p><p>- 私有项目，用户查询DID文档需要提供Token、项目编号、DID标识；</p><p>- 验证Token是否在启用中（公开项目不需要）</p><p>- 验证DID项目是否在启用中</p><p>- 验证用户是否有对应项目查询DID文档的权限</p><p>- 验证项目中是否存在对应的DID</p><p>验证通过后返回将对应的DID文档原文。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|N||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|DID标识符|did|String|Y||
|2|项目编号|projectNo|String|Y||
|**响应参数**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|DID文档|didDocument|Text|Y||
|**响应示例**||||||
|<p>{<br>`    `"code": "0",<br>`    `"data": {</p><p>`        `\"didDocument\": "{\"@context\":\"https://w3id.org/did/v1\",\"id\":\"did:eproof:19916083176f4b3cb2b0daae9725c079\",\"version\":\"v1\",\"created\":\"2023-06-03T10:19:24Z\",\"updated\":\"2023-06-03T10:19:24Z\",\"verificationMethod\":[{\"id\":\"did:eproof:19916083176f4b3cb2b0daae9725c079#key-1\",\"type\":\"SHA256withRSA\",\"controller\":\"did:eproof:19916083176f4b3cb2b0daae9725c079\",\"publicKeyMultibase\":\"zH3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV\"}],\"authentication\":[\"#key-1\"],\"proof\":{\"type\":\"DataIntegrityProof\",\"created\":\"2025-04-27T17:58:33Z\",\"verificationMethod\":\"did:corpid:8adf89f92b354d73acdfa91e619204ee#key-1\",\"cryptosuite\":\"ecdsa-rdfc-2019\",\"proofPurpose\":\"assertionMethod\",\"proofValue\":\"z2LeuoNi3yR1b6c3fkRsEvXJ5ex8X4RdutyK7L6HAo2bJQwr21w85Y5KWy3DptXR8ke52Assqik6wKTy9DKqkEZ2r\"}}"<br>        <br>`    `},<br>`    `"message": ""<br>}</p>||||||

<a name="heading_19"></a>2.2.2 **发证方**

<a name="heading_20"></a>2.2.2.1 **注册发证方**

|**接口地址**|/api/sys/v1/issuer/register|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>注册发证方，处理逻辑如下：</p><p>- 验证Token是否在启用中（公开项目不需要）</p><p>- 验证DID项目是否在启用中</p><p>- 验证用户是否有向对应项目注册发证方的权限</p><p>- 验证发证方名称在项目内是否还没有被注册（项目内已注册或注册待审核）</p><p>- 验签DID签名值</p><p>&emsp;验证通过后：</p><p>- **如果未启用审核**，则将发证方DID标识、发证方名称写入到Issuer合约。</p><p>- **开启了发证方注册审核，**注册请求将发送至项目Owner，经过项目Owner审核通过后，将发证方DID标识、发证方名称写入到Issuer合约。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|DID标识符|issuerDid|String|Y||
|2|发证方名称|issuerName|String|` `Y||
|3|项目编号|projectNo|String|Y||
|5|联系人|contactPerson|String|Y||
|6|联系电话|contactNumber|String|Y||
|7|联系邮箱|contactEmail|String|Y||
|8|业务描述|businessScenario|String|Y||
|9|业务数据签名值|signature|String|Y||
|10|交易数据签名值|txSignature|String|N||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_21"></a>2.2.2.2  **更新发证方**

|**接口地址**|/api/sys/v1/issuer/update|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>更新发证方，处理逻辑如下：</p><p></p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|DID标识符|issuerDid|String|Y||
|2|发证方名称|issuerName|String|` `Y||
|3|业务数据签名值|signature|String|Y||
|4|交易数据签名值|txSignature|String|N||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_22"></a>2.2.2.3  **查询发证方**

|**接口地址**|/api/sys/v1/issuer/search|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>查询发证方，处理逻辑如下：</p><p>- 公开项目</p><p>- 验证DID项目是否在启用中</p><p>- 验证项目中是否存在对应的发证方DID</p><p>- 私有项目：需提供Token、DID标识、项目编号</p><p>- 验证Token是否在启用中</p><p>- 验证用户是否有对应项目查询发证方的权限</p><p>- 验证项目中是否存在对应的发证方DID</p><p>验证通过后返回将对应的DID和发证方名称。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|N||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|DID标识符|issuerDid|String|Y||
|**响应参数**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|DID标识符|issuerDid|String|Y||
|2|发证方名称|issuerName|String|` `Y||
|**响应示例**||||||
|<p>{<br>`    `"code": "0",<br>`    `"data": {</p><p>`    `"issuerDid": "",</p><p>`    `"issuerName": ""</p><p>},<br>"message": ""</p><p>}</p>||||||

<a name="heading_23"></a>2.2.2.4 **启用/禁用发证方**

|**接口地址**|/api/sys/v1/issuer/status/update|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>启用发证方，处理逻辑如下：</p><p>- 验证Token是否在启用中</p><p>- 验证DID项目是否在启用中</p><p>- 验证用户是否为项目Owner</p><p>- 验证项目中是否存在对应的发证方DID</p><p>&emsp;验证通过后将更改链上的发证方状态。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|DID标识符|issuerDid|String|Y||
|2|项目编号|projectNo|String|Y||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|<p>{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""</p><p>}</p>||||||

<a name="heading_24"></a>2.2.3 **VC**

<a name="heading_25"></a>2.2.3.1  **注册VC模板**

|**接口地址**|/api/sys/v1/vc/register|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>注册VC模板，处理逻辑如下：</p><p>- 验证Token是否在启用中（公开项目不需要）</p><p>- 验证DID项目是否在启用中</p><p>- 验证对应项目内是否存在对应的发证方</p><p>- 验证用户是否有向对应项目注册VC模板的权限</p><p>- 验证VC模板的ID是否还没有被注册</p><p>- 通过发证方DID获取公钥验证签名值</p><p>- 验证交易签名模式</p><p>&emsp;**验证通过后将**</p><p>1. **开启VC模板注册审核，**注册请求将发送至项目Owner，经过审核后将发证方的VC模版内容写入Issuer合约。</p><p>2. **未开启VC模板注册审核，**则直接将发证方的VC模版内容写入Issuer合约。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|DID标识符|issuerDid|String|Y||
|2|项目编号|projectNo|String|Y||
|3|模板数据|vcTemplate|` `VCTemplate|Y||
|4|业务数据签名值|signature|String|Y||
|5|交易数据签名值|txSignature|String|N||
|**VCTemplate**||||||
|1|VC模板ID|templateId|String|Y||
|2|VC模板名称|templateName|String|Y||
|3|发证方DID标识符|issuerDid|String|Y||
|4|回调地址|issuanceEndpoint|String|Y||
|5|VC模板描述|templateDescription|String|Y||
|6|属性字段定义|registrationFields|Array|Y||
|7|属性字段定义|credentialSubjects|Array|Y||
|8|VC模板签名|proof|Proof|Y||
|**registrationFields属性结构**||||||
|1|属性名称|fieldName|String|Y||
|2|属性描述|description|String|Y||
|3|是否必填项|mandatory|Boolean|Y||
|**credentialSubjects属性结构**||||||
|1|属性名称|subjectName|String|||
|3|属性描述|description|String|Y||
|**Proof**||||||
|1|创建时间|created|String|Y||
|2|proof类型|type|String|Y||
|3|验证方法|verificationMethod|String|Y||
|4|加密套件|cryptosuite|String|Y||
|5|证明目的|proofPurpose|String|Y||
|6|签名值|proofValue|String|Y||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_26"></a>2.2.3.2 **更新VC模板**

|**接口地址**|/api/sys/v1/vc/update|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>更新VC模板，处理逻辑如下：</p><p>- 验证Token是否在启用中（公开项目不需要）</p><p>- 验证DID项目是否在启用中</p><p>- 验证项目中是否存在对应的VC模板ID</p><p>- 验证用户是否有向对应项目更新VC模板的权限</p><p>- 通过发证方DID获取公钥验证签名值</p><p>- 验证交易签名模式</p><p>&emsp;**验证通过后将**</p><p>1. **开启VC模板注册审核，**更新请求将发送至项目Owner，经过审核后再将发证方的新的VC模版内容写入Issuer合约。</p><p>2. **未开启VC模板注册审核，**则直接将发证方的新的VC模版内容写入Issuer合约。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|DID标识符|issuerDid|String|Y||
|2|项目编号|projectNo|String|Y||
|3|模板数据|vcTemplate|` `VCTemplate|Y||
|4|业务数据签名值|signature|String|Y||
|5|交易数据签名值|txSignature|String|N||
|**VCTemplate**||||||
|1|VC模板ID|templateId|String|Y||
|2|VC模板名称|templateName|String|Y||
|3|发证方DID标识符|issuerDid|String|Y||
|4|回调地址|issuanceEndpoint|String|Y||
|5|VC模板描述|templateDescription|String|Y||
|6|属性字段定义|registrationFields|Array|Y||
|7|属性字段定义|credentialSubjects|Array|Y||
|8|VC模板签名|proof|Proof|Y||
|**registrationFields属性结构**||||||
|1|属性名称|fieldName|String|Y||
|2|属性描述|description|String|Y||
|3|是否必填项|mandatory|Boolean|Y||
|**credentialSubjects属性结构**||||||
|1|属性名称|subjectName|String|||
|3|属性描述|description|String|Y||
|3|属性描述|description|String|Y||
|**Proof**||||||
|1|创建时间|created|String|Y||
|2|proof类型|type|String|Y||
|3|验证方法|verificationMethod|String|Y||
|4|加密套件|cryptosuite|String|Y||
|5|证明目的|proofPurpose|String|Y||
|6|签名值|proofValue|String|Y||
<a name="heading_27"></a>2.2.3.3 **启用/禁用VC模板**

|**接口地址**|/api/sys/v1/vc/status/update|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>启用VC模板，处理逻辑如下：</p><p>- 验证Token是否在启用中（公开项目不需要）</p><p>- 验证DID项目是否在启用中</p><p>- 验证项目中是否存在对应的VC模板ID</p><p>- 验证用户是否有向对应项目注册VC模板的权限（有注册权限才可以执行Enable/Disable操作）</p><p>- 验证VC Tempalte的当前状态和当前操作是否不一致</p><p>- 通过发证方DID获取公钥验证签名值</p><p>- 验证交易签名模式</p><p>&emsp;**验证通过后将**更新VC模版的状态。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|DID标识符|issuerDid|String|Y||
|2|项目编号|projectNo|String|Y||
|3|VC模板ID|vcTemplateId|String|Y||
|4|业务数据签名值|signature|String|Y||
|5|交易数据签名值|txSignature|String|N||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""<br>}||||||

<a name="heading_28"></a>2.2.3.4 **查询VC模板**

|**接口地址**|/api/sys/v1/vc/search|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>查询VC模板，处理逻辑如下：</p><p>- 公开项目：用户需要提供项目编号、VC 模板ID；</p><p>- 验证DID项目是否在启用中</p><p>- 验证VC模板ID是否存在（项目内）</p><p>- 私有项目：用户需要提供Token、项目编号、VC 模板ID；</p><p>- 验证Token是否在启用中</p><p>- 验证DID项目是否在启用中</p><p>- 验证VC模板ID是否存在（项目内）</p><p>- 验证用户是否拥有查询VC Template的项目权限</p><p>验证通过后，返回VC模板原文数据。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|N||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|2|项目编号|projectNo|String|Y||
|3|VC模板ID|vcTemplateId|String|Y||
|**响应参数**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|VC模板数据|vcTemplate|String|Y||
|**响应示例**||||||
|<p>{<br>`    `"code": "0",<br>`    `"data": {</p><p>`        `\"vcTemplate\": "{\"templateId\":\"441gb4gq5j7tozuf04c3efxmccrltkmx\",\"templateName\":\"NUS e-Diploma\",\"templateDescription\":\"National University e-Diploma\",\"issuanceEndpoint\":\"https://xxxx/oai/sys/v1/util/receiveVcApplyCallback\",\"issuerDid\":\"did:corpid:0252ab3abedd408390475ed4133f576e\",\"credentialSubjects\":[{\"subjectName\":\"Holder DID\",\"description\":\"Holder's sgDID identifier\"},{\"subjectName\":\"Faculty\",\"description\":\"The faculty of the student.\"}],\"registrationFields\":[{\"fieldName\":\"StudentNo.\",\"description\":\"Student ID\",\"mandatory\":true},{\"fieldName\":\"Photo\",\"description\":\"The photo of the student.\",\"mandatory\":true}],\"proof\":{\"type\":\"DataIntegrityProof\",\"verificationMethod\":\"did:corpid:8adf89f92b354d73acdfa91e619204ee#key-1\",\"cryptosuite\":\"ecdsa-rdfc-2019\",\"created\":\"2021-11-13T18:19:39Z\",\"proofPurpose\":\"assertionMethod\",\"proofValue\":\"26466a7b1e31f04a5a6de8c81665229e364db03a8cfefd11cfb031bc6f7b3853ea479448b1d71686562e2da4edc2eb7897be221e14fbe98c768a4dd2224763d7\"}}"<br>`    `},<br>`    `"message": ""<br>}</p>||||||

<a name="heading_29"></a>2.2.3.5 **VC存证**

|**接口地址**|/api/sys/v1/vc/evidence|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>VC存证，处理逻辑如下：</p><p>- 验证Token是否在启用中（公开项目不需要）</p><p>- 验证DID项目是否在启用中</p><p>- 验证项目中发证方和VC ID还没有在项目中存证</p><p>- 验证用户是否拥有VC存证的项目权限</p><p>- 通过发证方DID获取公钥验证发证方DID签名值</p><p>&emsp;验证通过后，将VC ID、VC hash、发证方DID存储至区块链。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|N||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|VC ID|vcId|String|Y||
|2|项目编号|projectNo|String|Y||
|3|VC哈希|vcHash|String|Y||
|4|发证方DID标识符|issuerDid|String|Y||
|5|业务数据签名值|signature|String|Y||
|6|交易数据签名值|txSignature|String|N||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|<p>{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""</p><p>}</p>||||||

<a name="heading_30"></a>2.2.3.6 **查询VC存证**

|**接口地址**|/api/sys/v1/vc/evidence/search|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>查询VC存证，处理逻辑如下：</p><p>- 公开项目：项目编号、VC ID、发证方DID</p><p>- 验证DID项目是否在启用中</p><p>- 私有项目：Token、项目编号、VC ID、发证方DID</p><p>- 验证Token是否在启用中</p><p>- 验证DID项目是否在启用中</p><p>- 验证用户是否拥有查询VC存证的项目权限</p><p>验证通过后，返回VC存证记录。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|N||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|VC ID|vcId|String|Y||
|2|项目编号|projectNo|String|Y||
|3|DID标识符|issuerDid|String|Y||
|**响应参数**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|VC ID|vcId|String|Y||
|2|VC哈希|vcHash|String|Y||
|3|DID标识符|issuerDid|String|Y||
|4|交易哈希|txHash|String|Y||
|<p>{<br>`    `"code": "0",<br>`    `"data": {</p><p>`        `"vcId": "",</p><p>`  `"vcHash": "",</p><p>`  `"issuerDid": "",</p><p>`  `"txHash": ""</p><p>},<br>"message": ""</p><p>}</p>||||||

<a name="heading_31"></a>2.2.3.7 **吊销VC**

|**接口地址**|/api/sys/v1/vc/revoke|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>吊销VC，处理逻辑如下：</p><p>- 验证Token是否在启用中（公开项目不需要）</p><p>- 验证DID项目是否在启用中</p><p>- 验证在该项目的存证数据中，是否存在对应的VC ID和发证方</p><p>- 验证用户是否拥有吊销VC的项目权限</p><p>- 验证VC是否已被吊销</p><p>- 通过发证方DID获取公钥验证发证方DID签名值</p><p>&emsp;验证通过后，将VC hash存证的记录关联的状态更新为吊销。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|Y||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|VC原文|vc|String|Y||
|**响应参数**||||||
|无||||||
|**响应示例**||||||
|<p>{<br>`    `"code": "0",<br>`    `"data": null,<br>`    `"message": ""</p><p>}</p>||||||

<a name="heading_32"></a>2.2.3.8 **查询VC吊销状态**

|**接口地址**|/api/sys/v1/vc/status/search|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>查询VC吊销，处理逻辑如下：</p><p>- 公开项目：项目编号、VC ID、发证方DID标识符</p><p>- 验证DID项目是否在启用中</p><p>- 私有项目：Token、项目编号、VC ID、发证方DID标识符</p><p>- 验证DID项目是否在启用中</p><p>- 验证Token是否在启用中（公开项目不需要）</p><p>- 验证用户是否拥有查询VC吊销状态的项目权限</p><p>验证通过后，到链上查询对应的VC ID的吊销状态，如果是吊销状态则返回True；如果没被吊销返回False。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|N||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|VC ID|vcId|String|Y||
|2|项目编号|projectNo|String|Y||
|3|DID标识符|issuerDid|String|Y||
|**响应参数**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|VC吊销状态|revokeStatus|boolean|Y||
|<p>{<br>`    `"code": "0",<br>`    `"data": {</p><p>`        `"revokeStatus": true</p><p>},<br>"message": ""</p><p>}</p>||||||

<a name="heading_33"></a>2.2.3.9 **核验VC**

|**接口地址**|/api/sys/v1/vc/verify|||||
| :-: | :- | :- | :- | :- | :- |
|**接口描述**|<p>核验VC，处理逻辑如下：</p><p>- 验证token是否被冻结</p><p>- 验证DID项目否被停用</p><p>- 验证用户是否加入了对应项目，或者是项目Owner</p><p>&emsp;验证通过后，通过到链上查询对应项目的数据，对VC进行以下核验：</p><p>1. 验证VC的Proof中的发证方签名</p><p>2. 验证VC是否跟链上存证记录的VC hash一致</p><p>3. 验证VC是否没有吊销记录</p><p>&emsp;全部核验通过后返回验证通过，任意一项不通过则返回核验失败。</p>|**调用方式** |POST|||
|**请求参数**||||||
|**Header**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|令牌|token|String|N||
|**Body**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|项目编号|projectNo|String|Y||
|2|VC原文|vc|String|Y||
|**响应参数**||||||
|**序号**|**字段名**|**字段**|**类型**|**必填**|**备注**|
|1|VC核验状态|verificationStatus|boolean|Y||
|<p>{<br>`    `"code": "0",<br>`    `"data": {</p><p>`        `"verificationStatus": true</p><p>},<br>"message": ""</p><p>}</p>||||||


