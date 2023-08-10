package consts

const (
	ApiConfigPath  = "./server/cmd/api/config.yaml"
	UserConfigPath = "./server/cmd/user/config.yaml"

	HlogFilePath = "./tmp/hlog/logs/"
	KlogFilePath = "./tmp/klog/logs/"

	UserSnowflakeNode = iota
	NacosSnowflakeNode

	MySqlDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	RedisProfileClientDB = 0
)
