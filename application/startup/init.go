package startup

import (
	"context"
	"go-mq/infrastructure/provider"
	"go-mq/infrastructure/svc"
)

func Init(ctx *svc.ServiceContext) {
	initConsumers(ctx)
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
