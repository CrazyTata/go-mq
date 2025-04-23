package api

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"go-mq/infrastructure/svc"
	"go-mq/interfaces/api/handler/health"
)

// RegisterHandlers 注册HTTP处理器
func RegisterHandlers(server *rest.Server, svc *svc.ServiceContext) {

	// 健康记录
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v1/health/save",
				Handler: health.SaveHandler(svc),
			},
		},
	)
}
