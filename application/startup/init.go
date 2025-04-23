package startup

import (
	"context"
	"go-mq/infrastructure/provider"
	"go-mq/infrastructure/svc"
)

func Init(ctx *svc.ServiceContext) {
	initServices(ctx)
	initConsumers(ctx)
}

// initServices 初始化所有服务
func initServices(ctx *svc.ServiceContext) {
	// 1. 初始化仓储层
	// subjectRepo := subjectmodel.NewSubjectModel(ctx.GetDB(), ctx.GetCache())

	// // 2. 初始化服务层
	// subjectService := service.NewSubjectService(subjectRepo)

	// // 3. 初始化外观层
	// apiKeyFacade := facade.NewApiKeyFacade(subjectService)
	// apiKeyFacade.Init() // 注册事件监听
	// ctx.ApiKeyAuth = middleware.NewApiKeyAuthMiddleware(apiKeyFacade.GetKeyManager()).Handle
	return
}

func initConsumers(svcCtx *svc.ServiceContext) {

	manager := provider.InitializeConsumerManager(svcCtx)

	// 注册健康记录消费者
	healthConsumer := provider.InitializeHealthConsumer(svcCtx)
	manager.Register(healthConsumer)

	// TODO: 注册其他消费者
	// 例如：用户消费者、预约消费者等

	manager.StartAll(context.Background())
	return
}
