package initialize

import (
	"Herbal/server/cmd/user/config"
	"Herbal/server/shared/consts"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"strings"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.UserConfigPath)
	if err := v.ReadInConfig(); err != nil {
		klog.Fatalf("read viper config failed: %s", err)
	}
	if err := v.Unmarshal(&config.GlobalNacosConfig); err != nil {
		klog.Fatalf("unmarshal err failed: %s", err)
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
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		klog.Fatalf("create config client failed: %s", err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: config.GlobalNacosConfig.DataId,
		Group:  config.GlobalNacosConfig.Group,
	})
	if err != nil {
		klog.Fatalf("get config failed: %s", err.Error())
	}
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(strings.NewReader(content))
	if err != nil {
		klog.Fatal(fmt.Errorf("failed to read config: %s", err))
	}

	err = viper.Unmarshal(&config.GlobalServerConfig)
	if err != nil {
		klog.Fatalf("nacos config failed: %s", err)
	}
}
