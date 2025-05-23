package config

import (
	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/rest"
)

// Config 应用配置
type Config struct {
	rest.RestConf // REST服务配置

	DB struct {
		DataSource string // 数据库连接字符串
	}

	Domain string // 回调基础URL

	Redis struct {
		Host string // Redis主机
		Pass string // Redis密码
		Type string // Redis类型
		Tls  bool   // Redis是否启用TLS
	}
	// jwt 配置
	FrontendAuth struct {
		AccessSecret string `json:",optional,default=13safhasfuawefc0f0"`
		AccessExpire int64  `json:",optional,default=25920000"`
	}

	DqConf dq.DqConf
}
