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
	TokenAnalysis struct {
		Url     string // 令牌分析URL
		Android struct {
			Appid       string // Android应用ID
			Version     string // 版本号
			StrictCheck string // 严格检查
			APPSecret   string // 应用密钥
		}
		Ios struct {
			Appid       string // iOS应用ID
			Version     string // 版本号
			StrictCheck string // 严格检查
			APPSecret   string // 应用密钥
		}
	}
	MSG struct {
		SMSTemplate string // 短信模板
		Url         string // 短信URL
		ApiKey      string // API密钥
		Account     string // 账号
	}

	Qiniu struct {
		AccessKey string
		SecretKey string
		Bucket    string
		Domain    string
		Region    string
	}

	Email struct {
		Host        string
		Port        int
		Username    string
		Password    string
		From        string
		FrontendURL string
	}

	DqConf dq.DqConf
}
