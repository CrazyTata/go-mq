package provider

import (
	"go-mq/application/consumer"
	"go-mq/application/service"
	"go-mq/application/service/facade"
	"go-mq/domain/health"
	"go-mq/domain/statistics"
	"go-mq/infrastructure/queue"
	"go-mq/infrastructure/svc"

	"github.com/google/wire"
)

// ServiceProviderSet 服务层依赖提供者集合
var ServiceProviderSet = wire.NewSet(
	ProviderHealthService,
	ProviderStatisticsService,
	ProviderQueueManager,
	ProviderHealthConsumer,
	ProviderConsumerManager,
)

// ProviderHealthService 提供健康服务实现
func ProviderHealthService(healthRepo health.HealthRepository, queueManager queue.QueueManager) *service.HealthService {
	return service.NewHealthService(healthRepo, queueManager)
}

// ProviderStatisticsService 提供统计服务实现
func ProviderStatisticsService(statisticsRepo statistics.StatisticsRepository) *service.StatisticsService {
	return service.NewStatisticsService(statisticsRepo)
}

// ProviderGoQueueManager 提供 Go 队列管理器实现
func ProviderQueueManager(svcCtx *svc.ServiceContext) queue.QueueManager {
	return queue.NewGoMQ()
}

// ProviderHealthConsumer 提供健康消费者实现
func ProviderHealthConsumer(queueManager queue.QueueManager, userFacade *facade.UserFacade) *consumer.HealthConsumer {
	return consumer.NewHealthConsumer(queueManager, userFacade)
}

// ProviderConsumerManager 提供消费者管理器实现
func ProviderConsumerManager(queueManager queue.QueueManager) *queue.ConsumerManager {
	return queue.NewConsumerManager(queueManager)
}
