package nacos

import (
	"log"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// NewNacosConf 初始化nacos客户端服务
func NewNacosConf(conf *conf.Dbs, logger log.Logger) vo.NacosClientParam {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.Ip, conf.Nacos.Port),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         conf.Nacos.NamespaceId,
		TimeoutMs:           conf.Nacos.TimeoutMs,
		NotLoadCacheAtStart: conf.Nacos.NotLoadCacheAtStart,
		LogDir:              conf.Nacos.LogDir,
		CacheDir:            conf.Nacos.CacheDir,
		LogLevel:            conf.Nacos.LogLevel,
	}

	return vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	}
}

func NewDiscovery(param vo.NacosClientParam) registry.Discovery {
	client, err := clients.NewNamingClient(param)
	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}

// NewRegistrar 服务注册业务注入
func NewRegistrar(param vo.NacosClientParam) registry.Registrar {
	client, err := clients.NewNamingClient(param)
	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}
