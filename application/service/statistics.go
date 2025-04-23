package service

import (
	"context"
	"go-mq/domain/statistics"
)

type StatisticsService struct {
	statisticsRepo statistics.StatisticsRepository
}

func NewStatisticsService(statisticsRepo statistics.StatisticsRepository) *StatisticsService {
	return &StatisticsService{
		statisticsRepo: statisticsRepo,
	}
}

func (s *StatisticsService) GetByDateAndUserId(ctx context.Context, userId, date string) (*statistics.Statistics, error) {
	return s.statisticsRepo.GetStatistics(ctx, userId, date)
}

func (s *StatisticsService) Create(ctx context.Context, statistics *statistics.Statistics) (int64, error) {
	return s.statisticsRepo.InsertStatistics(ctx, statistics)
}

func (s *StatisticsService) Update(ctx context.Context, statistics *statistics.Statistics) error {
	return s.statisticsRepo.UpdateStatistics(ctx, statistics)
}
