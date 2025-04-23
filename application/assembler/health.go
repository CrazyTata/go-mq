package assembler

import (
	"go-mq/application/dto"
	domain "go-mq/domain/health"
	healthModel "go-mq/infrastructure/persistence/model/health_records"
	"time"

	"github.com/samber/lo"
)

func DOTOPOHealthRecord(do *domain.HealthRecords) *healthModel.HealthRecords {
	return &healthModel.HealthRecords{
		Id:          do.Id,
		PatientId:   do.PatientId,
		PatientName: do.PatientName,
		Date:        do.Date,
		RecordType:  do.RecordType,
		Diagnosis:   do.Diagnosis,
		Treatment:   do.Treatment,
		Notes:       do.Notes,
		VitalSigns:  do.VitalSigns,
		Medications: do.Medications,
		Attachments: do.Attachments,
		UserId:      do.UserId,
		IsDeleted:   do.IsDeleted,
		CreatedAt:   do.CreatedAt,
		UpdatedAt:   do.UpdatedAt,
	}
}

func POTODOHealthRecord(po *healthModel.HealthRecords) *domain.HealthRecords {
	return &domain.HealthRecords{
		Id:          po.Id,
		PatientId:   po.PatientId,
		PatientName: po.PatientName,
		Date:        po.Date,
		RecordType:  po.RecordType,
		Diagnosis:   po.Diagnosis,
		Treatment:   po.Treatment,
		Notes:       po.Notes,
		VitalSigns:  po.VitalSigns,
		Medications: po.Medications,
		Attachments: po.Attachments,
		UserId:      po.UserId,
		IsDeleted:   po.IsDeleted,
		CreatedAt:   po.CreatedAt,
		UpdatedAt:   po.UpdatedAt,
	}
}

func DOTODTOHealth(do *domain.HealthRecords) *dto.HealthResponse {
	if do == nil {
		return nil
	}
	date := ""
	if do.Date.Valid {
		date = do.Date.Time.Format(time.DateOnly)
	}

	return &dto.HealthResponse{
		Id:          do.Id,
		PatientId:   do.PatientId,
		PatientName: do.PatientName,
		Date:        date,
		RecordType:  do.RecordType,
		Diagnosis:   do.Diagnosis.String,
		Treatment:   do.Treatment.String,
		Notes:       do.Notes.String,
		VitalSigns:  do.VitalSigns.String,
		Medications: do.Medications.String,
		Attachments: do.Attachments.String,
		UserId:      do.UserId,
		CreatedAt:   do.CreatedAt.Format(time.DateTime),
		UpdatedAt:   do.UpdatedAt.Format(time.DateTime),
	}
}

func DOTODTOHealthList(do []*domain.HealthRecords, total int64) *dto.HealthListResponse {
	return &dto.HealthListResponse{
		Total: total,
		List: lo.Map(do, func(do *domain.HealthRecords, _ int) *dto.HealthResponse {
			return DOTODTOHealth(do)
		}),
	}
}
