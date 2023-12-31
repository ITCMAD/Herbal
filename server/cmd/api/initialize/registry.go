package initialize

import (
	"Herbal/server/cmd/api/config"
	"Herbal/server/shared/consts"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"

	"net"
)

func InitRegistry() (registry.Registry, *registry.Info) {
	v := viper.New()
	v.SetConfigFile(consts.ApiConfigPath)
	if err := v.ReadInConfig(); err != nil {
		hlog.Fatalf("read viper config failed: %s", err)
	}
	if err := v.Unmarshal(&config.GlobalNacosConfig); err != nil {
		hlog.Fatalf("unmarshal err failed: %s", err)
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(config.GlobalNacosConfig.Host, config.GlobalNacosConfig.Port),
	}

	cc := constant.ClientConfig{
		NamespaceId:         config.GlobalNacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "info",
		Username:            config.GlobalNacosConfig.User,
		Password:            config.GlobalNacosConfig.Password,
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		hlog.Errorf("create registry err: %s", err.Error())
	}

	r := nacos.NewNacosRegistry(cli, nacos.WithRegistryGroup(config.GlobalNacosConfig.Group))

	sf, err := snowflake.NewNode(consts.NacosSnowflakeNode)
	if err != nil {
		hlog.Fatalf("generate nacos service name failed:%s", err)
	}
	info := &registry.Info{
		ServiceName: config.GlobalServerConfig.Name,
		Addr:        utils.NewNetAddr("tcp", net.JoinHostPort(config.GlobalServerConfig.Host, config.GlobalServerConfig.Port)),
		Weight:      10,
		Tags:        map[string]string{"ID": sf.Generate().Base36()},
	}
	return r, info
}
