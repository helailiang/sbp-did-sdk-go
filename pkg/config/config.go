package config

import (
	"errors"
	"fmt"
	"strings"
)

// Config 定义SDK的配置结构
type Config struct {
	// 华为云配置
	HuaweiCloudEndpoint  string `json:"huawei_cloud_endpoint" yaml:"huawei_cloud_endpoint"`
	HuaweiCloudAccessKey string `json:"huawei_cloud_access_key" yaml:"huawei_cloud_access_key"`
	HuaweiCloudSecretKey string `json:"huawei_cloud_secret_key" yaml:"huawei_cloud_secret_key"`
	HuaweiCloudRegion    string `json:"huawei_cloud_region" yaml:"huawei_cloud_region"`
	
	// DCI配置（可选）
	DCIAccessKey string `json:"dci_access_key" yaml:"dci_access_key"`
	DCISecretKey string `json:"dci_secret_key" yaml:"dci_secret_key"`
	DCIAccessAddress string `json:"dci_access_address" yaml:"dci_access_address"`
	
	// OpenAPI配置
	OpenAPIEndpoint string `json:"openapi_endpoint" yaml:"openapi_endpoint"`
	ProjectID       string `json:"project_id" yaml:"project_id"`
	ProjectVisibility string `json:"project_visibility" yaml:"project_visibility"` // public 或 private
	Token           string `json:"token" yaml:"token"` // 私有项目需要提供Token
	
	// 算法配置
	DefaultAlgorithm string `json:"default_algorithm" yaml:"default_algorithm"` // ECDSA, RSA, SM2
	DefaultHashAlgorithm string `json:"default_hash_algorithm" yaml:"default_hash_algorithm"` // SHA256, SM3
	
	// 日志配置
	LogLevel string `json:"log_level" yaml:"log_level"` // debug, info, warn, error
	LogFile  string `json:"log_file" yaml:"log_file"`
}

// NewConfig 创建新的配置实例
func NewConfig() *Config {
	return &Config{
		DefaultAlgorithm:     "ECDSA",
		DefaultHashAlgorithm: "SHA256",
		ProjectVisibility:    "public",
		LogLevel:            "info",
	}
}

// Validate 验证配置的有效性
func (c *Config) Validate() error {
	var errors []string
	
	// 验证华为云配置
	if c.HuaweiCloudEndpoint == "" {
		errors = append(errors, "HuaweiCloudEndpoint is required")
	}
	if c.HuaweiCloudAccessKey == "" {
		errors = append(errors, "HuaweiCloudAccessKey is required")
	}
	if c.HuaweiCloudSecretKey == "" {
		errors = append(errors, "HuaweiCloudSecretKey is required")
	}
	
	// 验证OpenAPI配置
	if c.OpenAPIEndpoint == "" {
		errors = append(errors, "OpenAPIEndpoint is required")
	}
	if c.ProjectID == "" {
		errors = append(errors, "ProjectID is required")
	}
	
	// 验证项目可见性
	if c.ProjectVisibility != "public" && c.ProjectVisibility != "private" {
		errors = append(errors, "ProjectVisibility must be 'public' or 'private'")
	}
	
	// 如果是私有项目，需要Token
	if c.ProjectVisibility == "private" && c.Token == "" {
		errors = append(errors, "Token is required for private projects")
	}
	
	// 验证算法配置
	if !isValidAlgorithm(c.DefaultAlgorithm) {
		errors = append(errors, fmt.Sprintf("DefaultAlgorithm must be one of: ECDSA, RSA, SM2, got: %s", c.DefaultAlgorithm))
	}
	
	if !isValidHashAlgorithm(c.DefaultHashAlgorithm) {
		errors = append(errors, fmt.Sprintf("DefaultHashAlgorithm must be one of: SHA256, SM3, got: %s", c.DefaultHashAlgorithm))
	}
	
	if len(errors) > 0 {
		return fmt.Errorf("configuration validation failed: %s", strings.Join(errors, "; "))
	}
	
	return nil
}

// isValidAlgorithm 检查算法是否有效
func isValidAlgorithm(algorithm string) bool {
	validAlgorithms := []string{"ECDSA", "RSA", "SM2"}
	for _, valid := range validAlgorithms {
		if algorithm == valid {
			return true
		}
	}
	return false
}

// isValidHashAlgorithm 检查哈希算法是否有效
func isValidHashAlgorithm(algorithm string) bool {
	validAlgorithms := []string{"SHA256", "SM3"}
	for _, valid := range validAlgorithms {
		if algorithm == valid {
			return true
		}
	}
	return false
}

// IsPrivateProject 检查是否为私有项目
func (c *Config) IsPrivateProject() bool {
	return c.ProjectVisibility == "private"
}

// GetToken 获取Token，私有项目必须提供
func (c *Config) GetToken() (string, error) {
	if c.IsPrivateProject() && c.Token == "" {
		return "", errors.New("token is required for private projects")
	}
	return c.Token, nil
}

// GetHuaweiCloudConfig 获取华为云配置
func (c *Config) GetHuaweiCloudConfig() map[string]string {
	return map[string]string{
		"endpoint":  c.HuaweiCloudEndpoint,
		"accessKey": c.HuaweiCloudAccessKey,
		"secretKey": c.HuaweiCloudSecretKey,
		"region":    c.HuaweiCloudRegion,
	}
}

// GetOpenAPIConfig 获取OpenAPI配置
func (c *Config) GetOpenAPIConfig() map[string]string {
	return map[string]string{
		"endpoint": c.OpenAPIEndpoint,
		"projectID": c.ProjectID,
		"visibility": c.ProjectVisibility,
		"token": c.Token,
	}
} 