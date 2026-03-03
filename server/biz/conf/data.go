package conf

import "github.com/moumou/server/pkgs/database"

type Data struct {
	ServerConfig   ServerConfig      `json:"server"`
	SecurityConfig SecurityConfig    `json:"security"`
	DBConfig       database.DbConfig `json:"database"`
	SystemConfig   SystemConfig      `json:"system"`
}

// SystemConfig 系统初始化接口认证配置（如 /system/initialize）
type SystemConfig struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

// ServerConfig 服务配置
type ServerConfig struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Timeout int    `yml:"timeout"`
}

// SecurityConfig 获取加解密的配置
type SecurityConfig struct {
	JWTKey string `json:"jwt_key"`
	Exp    int    `json:"exp"`
}
