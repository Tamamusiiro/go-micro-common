package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	PWD  string `json:"pwd"`
	DB   string `json:"db"`
	Port string `json:"port"`
}

func GetMysqlConfig(config config.Config, path ...string) MysqlConfig {
	mysqlConfig := MysqlConfig{}
	config.Get(path...).Scan(&mysqlConfig)
	return mysqlConfig
}
