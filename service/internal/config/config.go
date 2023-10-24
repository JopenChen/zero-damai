package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	// Mysql 数据库配置
	Mysql struct {
		Datasource string
	}

	// JWT 配置
	Jwt struct {
		Secret string // 秘钥
		Expire int64  // 过期时间，单位：s
	}
}
