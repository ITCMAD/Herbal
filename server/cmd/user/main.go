package main

import (
	"Herbal/server/cmd/user/config"
	"Herbal/server/cmd/user/initialize"
	"Herbal/server/shared/kitex_gen/user/userservice"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"net"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()

	r, info := initialize.InitRegistry()
	db := initialize.InitDB()
	rdb := initialize.InitRedis()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	svr := userservice.NewServer(&UserServiceImpl{
		MysqlManager: db,
		RedisManager: rdb,
	},
		server.WithServiceAddr(utils.NewNetAddr("tcp", net.JoinHostPort(config.GlobalServerConfig.Host, config.GlobalServerConfig.Port))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)
	err := svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
