package statistics

import (
	"go-mq/domain/statistics"
)

func POTODOGetStatistics(res *Statistics) *statistics.Statistics {
	return &statistics.Statistics{
		Id:                   res.Id,
		UserId:               res.UserId,
		TotalPatients:        res.TotalPatients,
		ActivePatients:       res.ActivePatients,
		TodayAppointments:    res.TodayAppointments,
		UpcomingAppointments: res.UpcomingAppointments,
		HealthRecords:        res.HealthRecords,
		RecentOperations:     res.RecentOperations,
		Date:                 res.Date,
		IsDeleted:            res.IsDeleted,
		CreatedAt:            res.CreatedAt,
		UpdatedAt:            res.UpdatedAt,
	}
}

func DOTOPOStatistics(res *statistics.Statistics) *Statistics {
	return &Statistics{
		Id:                   res.Id,
		UserId:               res.UserId,
		TotalPatients:        res.TotalPatients,
		ActivePatients:       res.ActivePatients,
		TodayAppointments:    res.TodayAppointments,
		UpcomingAppointments: res.UpcomingAppointments,
		HealthRecords:        res.HealthRecords,
		RecentOperations:     res.RecentOperations,
		Date:                 res.Date,
		CreatedAt:            res.CreatedAt,
		UpdatedAt:            res.UpdatedAt,
	}
}

func POTODOGetStatisticsList(res []Statistics) []*statistics.Statistics {
	resp := make([]*statistics.Statistics, 0, len(res))
	for _, v := range res {
		resp = append(resp, POTODOGetStatistics(&v))
	}
	return resp
}
