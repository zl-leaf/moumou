package env

import "os"

const (
	PROD = "prod"
	DEV  = "dev"
)

func GetEnv() string {
	env := os.Getenv("ENV")
	if env == PROD {
		return PROD
	}
	return DEV
}
