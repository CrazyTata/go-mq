//go:build wireinject
// +build wireinject

package provider

import (
	"go-mq/application/consumer"
	"go-mq/application/service"
	"go-mq/application/service/facade"
	"go-mq/infrastructure/queue"
	"go-mq/infrastructure/svc"

	"github.com/google/wire"
)

// userProviderSet 定义用户服务相关的依赖提供者
var providerSet = wire.NewSet(

	// 基础设施层
	RepositoryProviderSet,

	// 应用服务层
	ServiceProviderSet,

	// 门面服务层
	FacedeServiceProviderSet,
)

// InitializeUserFacade 初始化统计门面服务
func InitializeUserFacade(svcCtx *svc.ServiceContext) *facade.UserFacade {
	wire.Build(providerSet)
	return nil // 这个返回值会被wire自动生成的代码替换
}

// InitializeHealthService 初始化健康服务
func InitializeHealthService(svcCtx *svc.ServiceContext) *service.HealthService {
	wire.Build(providerSet)
	return nil // 这个返回值会被wire自动生成的代码替换
}

// InitializeStatisticsService 初始化统计服务
func InitializeStatisticsService(svcCtx *svc.ServiceContext) *service.StatisticsService {
	wire.Build(providerSet)
	return nil // 这个返回值会被wire自动生成的代码替换
}

// InitializeHealthConsumer 初始化健康消费者
func InitializeHealthConsumer(svcCtx *svc.ServiceContext) *consumer.HealthConsumer {
	wire.Build(providerSet)
	return nil // 这个返回值会被wire自动生成的代码替换
}

// InitializeConsumerManager 初始化消费者管理器
func InitializeConsumerManager(svcCtx *svc.ServiceContext) *queue.ConsumerManager {
	wire.Build(providerSet)
	return nil // 这个返回值会被wire自动生成的代码替换
}
