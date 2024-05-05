package config

import "github.com/moumou/server/framework/config/internal"

func Init(configDir string) error {
	return internal.InitProvider(configDir)
}

func GetConfig(cnfName string, cnf any) error {
	return internal.ConfigProvider.Get(cnfName).Populate(cnf)
}
