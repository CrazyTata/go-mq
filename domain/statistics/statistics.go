package statistics

import "time"

type Statistics struct {
	Id                   int64     `db:"id"`                    // 统计ID
	UserId               string    `db:"user_id"`               // 用户ID
	TotalPatients        int64     `db:"total_patients"`        // 总患者数
	ActivePatients       int64     `db:"active_patients"`       // 活跃患者数
	TodayAppointments    int64     `db:"today_appointments"`    // 今日预约数
	UpcomingAppointments int64     `db:"upcoming_appointments"` // 未来预约数
	HealthRecords        int64     `db:"health_records"`        // 健康记录总数
	RecentOperations     int64     `db:"recent_operations"`     // 最近操作记录数
	Date                 string    `db:"date"`                  // 日期
	IsDeleted            int64     `db:"is_deleted"`            // 是否删除
	CreatedAt            time.Time `db:"created_at"`            // 创建时间
	UpdatedAt            time.Time `db:"updated_at"`            // 更新时间
}

func Create(userId string, totalPatients, activePatients, todayAppointments, upcomingAppointments, healthRecords, recentOperations int64, date string) *Statistics {
	return &Statistics{
		UserId:               userId,
		TotalPatients:        totalPatients,
		ActivePatients:       activePatients,
		TodayAppointments:    todayAppointments,
		UpcomingAppointments: upcomingAppointments,
		HealthRecords:        healthRecords,
		RecentOperations:     recentOperations,
		Date:                 date,
	}
}

func (s *Statistics) Update(totalPatients, activePatients, todayAppointments, upcomingAppointments, healthRecords, recentOperations int64) {
	if 0 != totalPatients {
		s.TotalPatients = totalPatients
	}
	if 0 != activePatients {
		s.ActivePatients = activePatients
	}
	if 0 != todayAppointments {
		s.TodayAppointments = todayAppointments
	}
	if 0 != upcomingAppointments {
		s.UpcomingAppointments = upcomingAppointments
	}
	if 0 != healthRecords {
		s.HealthRecords = healthRecords
	}
	if 0 != recentOperations {
		s.RecentOperations = recentOperations
	}
}
