package data

import golangjwt "github.com/golang-jwt/jwt/v5"

// SecurityConfig 获取加解密的配置
type SecurityConfig struct {
	JWTKey string `yaml:"jwt_key"`
}

func (c *SecurityConfig) GetConfigName() string {
	return "security"
}

type CustomClaims struct {
	UserID int64 `json:"user_id"`
	golangjwt.RegisteredClaims
}
