package facade

import (
	"context"
	"go-mq/application/service"
	"go-mq/domain/statistics"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFacade struct {
	healthService     *service.HealthService
	statisticsService *service.StatisticsService
}

func NewUserFacade(healthService *service.HealthService, statisticsService *service.StatisticsService) *UserFacade {
	return &UserFacade{
		healthService:     healthService,
		statisticsService: statisticsService,
	}
}

func (s *UserFacade) UpdateHealthStatistics(userId string) error {
	ctx := context.Background()
	logger := logx.WithContext(ctx)

	healthRecords, err := s.healthService.CountByUserID(ctx, userId)
	if err != nil {
		logger.Errorf("UpdateHealthStatistics  healthService.CountByUserID error: %v", err)
		return err
	}
	statisticsDo, err := s.statisticsService.GetByDateAndUserId(ctx, userId, time.Now().Format("2006-01-02"))
	if err != nil {
		logger.Errorf("UpdateHealthStatistics  statisticsService.GetByDateAndUserId error: %v", err)
		return err
	}
	if statisticsDo == nil {
		statisticsDo = statistics.Create(userId, 0, 0, 0, 0, healthRecords, 0, time.Now().Format("2006-01-02"))
		_, err = s.statisticsService.Create(ctx, statisticsDo)
		if err != nil {
			logger.Errorf("UpdateUserStatistics error: %v", err)
			return err
		}
	} else {
		statisticsDo.Update(0, 0, 0, 0, healthRecords, 0)
		if err = s.statisticsService.Update(ctx, statisticsDo); err != nil {
			logger.Errorf("UpdateUserStatistics error: %v", err)
			return err
		}
	}
	return nil
}
