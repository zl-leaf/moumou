package param

type SecurityConfig struct {
	JWTKey string `yaml:"jwt_key"`
}

func (c *SecurityConfig) GetConfigName() string {
	return "security"
}
