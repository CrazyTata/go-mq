package statistics

import (
	"context"
)

type StatisticsRepository interface {
	InsertStatistics(ctx context.Context, statistics *Statistics) (int64, error)
	UpdateStatistics(ctx context.Context, statistics *Statistics) error
	GetStatistics(ctx context.Context, userId, date string) (*Statistics, error)
}
