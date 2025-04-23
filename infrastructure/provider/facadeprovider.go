package provider

import (
	"go-mq/application/service"
	"go-mq/application/service/facade"

	"github.com/google/wire"
)

// FacedeServiceProviderSet 门面服务层依赖提供者集合
var FacedeServiceProviderSet = wire.NewSet(
	ProviderUserFacade,
)

// ProviderUserFacade 提供用户门面服务实现
func ProviderUserFacade(healthService *service.HealthService, statisticsService *service.StatisticsService) *facade.UserFacade {
	return facade.NewUserFacade(healthService, statisticsService)
}
