package statistics

import (
	"context"
	"database/sql"
	"go-mq/domain/statistics"
	"go-mq/infrastructure/persistence"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ statistics.StatisticsRepository = (*customStatisticsModel)(nil)

type (
	// StatisticsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStatisticsModel.
	StatisticsModel interface {
		statisticsModel
	}

	customStatisticsModel struct {
		*defaultStatisticsModel
	}
)

// NewStatisticsModel returns a model for the database table.
func NewStatisticsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) statistics.StatisticsRepository {
	return &customStatisticsModel{
		defaultStatisticsModel: newStatisticsModel(conn, c, opts...),
	}
}

func (m *customStatisticsModel) InsertStatistics(ctx context.Context, statistics *statistics.Statistics) (int64, error) {
	po := DOTOPOStatistics(statistics)
	res, err := m.Insert(ctx, po)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (m *customStatisticsModel) UpdateStatistics(ctx context.Context, statistics *statistics.Statistics) error {
	po := DOTOPOStatistics(statistics)
	return m.Update(ctx, po)
}

func (m *customStatisticsModel) GetStatistics(ctx context.Context, userId, date string) (*statistics.Statistics, error) {
	rowBuilder := persistence.RowBuilder(m.table, statisticsRows)
	query, args, err := rowBuilder.Where("user_id = ?", userId).Where("date = ?", date).ToSql()
	if err != nil {
		return nil, err
	}
	var res Statistics
	err = m.QueryRowNoCacheCtx(ctx, &res, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return POTODOGetStatistics(&res), nil
}
