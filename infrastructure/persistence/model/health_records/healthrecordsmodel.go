package health_records

import (
	"context"
	"database/sql"
	"go-mq/domain/health"
	"go-mq/infrastructure/persistence"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ health.HealthRepository = (*customHealthRecordsModel)(nil)

type (
	// HealthRecordsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHealthRecordsModel.
	HealthRecordsModel interface {
		healthRecordsModel
	}

	customHealthRecordsModel struct {
		*defaultHealthRecordsModel
	}
)

// NewHealthRecordsModel returns a model for the database table.
func NewHealthRecordsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) health.HealthRepository {
	return &customHealthRecordsModel{
		defaultHealthRecordsModel: newHealthRecordsModel(conn, c, opts...),
	}
}

// Create 创建健康记录
func (m *customHealthRecordsModel) Create(ctx context.Context, patient *health.HealthRecords) (int64, error) {
	po := DOTOPOHealthRecords(patient)
	res, err := m.Insert(ctx, po)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// GetByID 根据ID获取健康记录
func (m *customHealthRecordsModel) GetByID(ctx context.Context, id int64) (*health.HealthRecords, error) {
	rowBuilder := persistence.RowBuilder(m.table, healthRecordsRows)
	query, values, err := rowBuilder.Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return nil, err
	}
	var res HealthRecords
	err = m.QueryRowNoCacheCtx(ctx, &res, query, values...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return POTODOGetHealthRecords(&res), nil
}

// GetByUserID 根据用户ID获取健康记录列表
func (m *customHealthRecordsModel) GetByUserID(ctx context.Context, userID string, page, pageSize int64) ([]*health.HealthRecords, int64, error) {
	rowBuilder := persistence.RowBuilder(m.table, healthRecordsRows)
	countBuilder := persistence.CountBuilder(m.table)
	if userID != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"user_id": userID})
		countBuilder = countBuilder.Where(squirrel.Eq{"user_id": userID})
	}
	var count int64
	query, values, err := countBuilder.ToSql()
	if err != nil {
		return nil, 0, err
	}
	err = m.QueryRowNoCacheCtx(ctx, &count, query, values...)
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return nil, 0, nil
	}
	query1, values1, err := rowBuilder.Limit(uint64(pageSize)).Offset(uint64((page - 1) * pageSize)).ToSql()
	if err != nil {
		return nil, 0, err
	}
	var res []HealthRecords
	err = m.QueryRowsNoCacheCtx(ctx, &res, query1, values1...)
	if err != nil {
		return nil, 0, err
	}
	return POTODOGetHealthRecordsList(res), count, nil
}

// Update 更新健康记录信息
func (m *customHealthRecordsModel) UpdateHealth(ctx context.Context, patient *health.HealthRecords) error {
	po := DOTOPOHealthRecords(patient)
	return m.Update(ctx, po)
}

// Search 搜索健康记录
func (m *customHealthRecordsModel) GetList(ctx context.Context, userID string, keyword string, status string, order int64, patientId int64, page, pageSize int64) ([]*health.HealthRecords, int64, error) {
	rowBuilder := persistence.RowBuilder(m.table, healthRecordsRows)
	countBuilder := persistence.CountBuilder(m.table)
	if userID != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"user_id": userID})
		countBuilder = countBuilder.Where(squirrel.Eq{"user_id": userID})
	}
	if keyword != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"patient_name": keyword})
		countBuilder = countBuilder.Where(squirrel.Eq{"patient_name": keyword})
	}
	if status != "" {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"record_type": status})
		countBuilder = countBuilder.Where(squirrel.Eq{"record_type": status})
	}
	if patientId != 0 {
		rowBuilder = rowBuilder.Where(squirrel.Eq{"patient_id": patientId})
		countBuilder = countBuilder.Where(squirrel.Eq{"patient_id": patientId})
	}
	if order == 1 {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy("id ASC")
	}
	var count int64
	query, values, err := countBuilder.ToSql()
	err = m.QueryRowNoCacheCtx(ctx, &count, query, values...)
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return nil, 0, nil
	}
	query1, values1, err := rowBuilder.OrderBy("id DESC").Limit(uint64(pageSize)).Offset(uint64((page - 1) * pageSize)).ToSql()
	if err != nil {
		return nil, 0, err
	}
	var res []HealthRecords
	err = m.QueryRowsNoCacheCtx(ctx, &res, query1, values1...)
	if err != nil {
		return nil, 0, err
	}
	return POTODOGetHealthRecordsList(res), count, nil
}

// CountByUserID 根据用户ID获取健康记录数量
func (m *customHealthRecordsModel) CountByUserID(ctx context.Context, userID string) (int64, error) {
	countBuilder := persistence.CountBuilder(m.table)
	if userID != "" {
		countBuilder = countBuilder.Where(squirrel.Eq{"user_id": userID})
	}
	var count int64
	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}
	err = m.QueryRowNoCacheCtx(ctx, &count, query, values...)
	if err != nil {
		return 0, err
	}
	return count, nil
}
