package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

func GetConsulConfig(addr, prefix string) (config.Config, error) {
	source := consul.NewSource(
		// 配置中心地址
		consul.WithAddress(addr),
		// 配置中心前缀，默认值为 /micro/config
		consul.WithPrefix(prefix),
		// 获取配置时不需要使用前缀
		consul.StripPrefix(true),
	)
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}
	return config, config.Load(source)
}
