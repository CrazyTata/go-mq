package provider

import (
	"go-mq/domain/health"
	"go-mq/domain/statistics"
	healthModel "go-mq/infrastructure/persistence/model/health_records"
	statisticsModel "go-mq/infrastructure/persistence/model/statistics"
	"go-mq/infrastructure/svc"

	"github.com/google/wire"
)

// RepositoryProviderSet 仓储层依赖提供者集合
var RepositoryProviderSet = wire.NewSet(
	ProviderHealthRepo,
	ProviderStatisticsRepo,
)

// ProviderHealthRepo 提供健康记录仓储实现
func ProviderHealthRepo(svcCtx *svc.ServiceContext) health.HealthRepository {
	return healthModel.NewHealthRecordsModel(svcCtx.GetDB(), svcCtx.GetCache())
}

// ProviderStatisticsRepo 提供统计仓储实现
func ProviderStatisticsRepo(svcCtx *svc.ServiceContext) statistics.StatisticsRepository {
	return statisticsModel.NewStatisticsModel(svcCtx.GetDB(), svcCtx.GetCache())
}
