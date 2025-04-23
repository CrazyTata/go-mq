package health_records

import (
	"go-mq/domain/health"
)

func POTODOGetHealthRecords(res *HealthRecords) *health.HealthRecords {
	return &health.HealthRecords{
		Id:          res.Id,
		PatientId:   res.PatientId,
		PatientName: res.PatientName,
		Date:        res.Date,
		RecordType:  res.RecordType,
		Diagnosis:   res.Diagnosis,
		Treatment:   res.Treatment,
		Notes:       res.Notes,
		VitalSigns:  res.VitalSigns,
		Medications: res.Medications,
		Attachments: res.Attachments,
		UserId:      res.UserId,
		IsDeleted:   res.IsDeleted,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
}

func DOTOPOHealthRecords(res *health.HealthRecords) *HealthRecords {
	return &HealthRecords{
		Id:          res.Id,
		PatientId:   res.PatientId,
		PatientName: res.PatientName,
		Date:        res.Date,
		RecordType:  res.RecordType,
		Diagnosis:   res.Diagnosis,
		Treatment:   res.Treatment,
		Notes:       res.Notes,
		VitalSigns:  res.VitalSigns,
		Medications: res.Medications,
		Attachments: res.Attachments,
		UserId:      res.UserId,
		IsDeleted:   res.IsDeleted,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
}

func POTODOGetHealthRecordsList(res []HealthRecords) []*health.HealthRecords {
	resp := make([]*health.HealthRecords, 0, len(res))
	for _, v := range res {
		resp = append(resp, POTODOGetHealthRecords(&v))
	}
	return resp
}
