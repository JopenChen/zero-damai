package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	// Mysql 数据库配置
	Mysql struct {
		Datasource string
	}
}
