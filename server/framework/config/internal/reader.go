package internal

import (
	"github.com/moumou/server/framework/env"
	"go.uber.org/config"
	"sync"
)

var (
	ConfigProvider config.Provider
	once           sync.Once
)

func InitProvider(configDir string) error {
	var err error
	once.Do(func() {
		filePath := buildConfigFilePath(configDir, env.GetEnv())
		ConfigProvider, err = config.NewYAML(config.File(filePath))
	})
	return err
}

func buildConfigFilePath(configDir, env string) string {
	fileName := "config_" + env + ".yml"

	if configDir[len(configDir)-1] != '/' {
		return configDir + "/" + fileName
	}
	return configDir + fileName
}
