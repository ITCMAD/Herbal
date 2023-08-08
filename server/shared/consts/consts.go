package consts

const (
	UserConfigPath = "./server/cmd/user/config.yaml"
	KlogFilePath   = "./tmp/klog/logs/"

	UserSnowflakeNode = iota
	NacosSnowflakeNode

	MySqlDSN = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	RedisProfileClientDB = 0
)
