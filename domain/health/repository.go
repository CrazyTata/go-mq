package health

import (
	"context"
)

// HealthRepository 健康记录仓储接口
type HealthRepository interface {
	// Create 创建健康记录
	Create(ctx context.Context, patient *HealthRecords) (int64, error)
	// GetByID 根据ID获取健康记录
	GetByID(ctx context.Context, id int64) (*HealthRecords, error)
	// GetByUserID 根据用户ID获取健康记录列表
	GetByUserID(ctx context.Context, userID string, page, pageSize int64) ([]*HealthRecords, int64, error)
	// Update 更新健康记录信息
	UpdateHealth(ctx context.Context, patient *HealthRecords) error
	// Search 搜索健康记录
	GetList(ctx context.Context, userID string, keyword string, status string, order, patientId, page, pageSize int64) ([]*HealthRecords, int64, error)
	// CountByUserID 根据用户ID获取健康记录数量
	CountByUserID(ctx context.Context, userID string) (int64, error)
}
