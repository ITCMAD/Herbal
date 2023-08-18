package initialize

import (
	"Herbal/server/cmd/user/config"
	"Herbal/server/shared/consts"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	hlog.Info("Init rdb")
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GlobalServerConfig.RedisInfo.Host, config.GlobalServerConfig.RedisInfo.Port),
		Password: config.GlobalServerConfig.RedisInfo.Password,
		DB:       consts.RedisProfileClientDB,
	})

}
