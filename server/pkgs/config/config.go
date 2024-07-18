package config

import (
	"errors"
	"github.com/moumou/server/pkgs/config/internal"
)

func Init(configDir string) error {
	return internal.InitProvider(configDir)
}

func GetConfig(cnfName string, cnf any) error {
	return internal.ConfigProvider.Get(cnfName).Populate(cnf)
}

func GetConfigs(cnfNames []string, cnfs []any) error {
	if len(cnfNames) != len(cnfs) {
		return errors.New("配置名称和配置长度不一致")
	}
	for i := 0; i < len(cnfNames); i++ {
		if err := GetConfig(cnfNames[i], cnfs[i]); err != nil {
			return err
		}
	}
	return nil
}
