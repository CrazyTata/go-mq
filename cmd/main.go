package main

import (
	"flag"
	"fmt"
	"go-mq/application/startup"
	"go-mq/common/redis"
	"go-mq/infrastructure/config"
	"go-mq/infrastructure/svc"
	"go-mq/interfaces/api"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/config.yaml", "配置文件路径")

func main() {

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	redis.Init(c.Redis.Host, c.Redis.Pass)
	defer redis.Close()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 配置 CORS
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 处理预检请求
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next(w, r)
		}
	})

	api.RegisterHandlers(server, ctx)
	startup.Init(ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
