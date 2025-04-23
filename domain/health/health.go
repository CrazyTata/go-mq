package health

import (
	"database/sql"
	"time"
)

// Health 健康记录领域模型
type HealthRecords struct {
	Id          int64
	PatientId   int64          // 患者ID
	PatientName string         // 患者姓名
	Date        sql.NullTime   // 日期
	RecordType  string         // 记录类型
	Diagnosis   sql.NullString // 诊断
	Treatment   sql.NullString // 治疗
	Notes       sql.NullString // 备注
	VitalSigns  sql.NullString // 生命体征
	Medications sql.NullString // 药物
	Attachments sql.NullString // 附件
	UserId      string         // 用户ID
	IsDeleted   int64          // 是否删除
	CreatedAt   time.Time      // 创建时间
	UpdatedAt   time.Time      // 更新时间
}

func Create(patientId int64, patientName string, date time.Time, recordType string, diagnosis string, treatment string, notes string, vitalSigns string, medications string, attachments string, userId string) *HealthRecords {
	return &HealthRecords{
		PatientId:   patientId,
		PatientName: patientName,
		Date:        sql.NullTime{Time: date, Valid: !date.IsZero()},
		RecordType:  recordType,
		Diagnosis:   sql.NullString{String: diagnosis, Valid: diagnosis != ""},
		Treatment:   sql.NullString{String: treatment, Valid: treatment != ""},
		Notes:       sql.NullString{String: notes, Valid: notes != ""},
		VitalSigns:  sql.NullString{String: vitalSigns, Valid: vitalSigns != ""},
		Medications: sql.NullString{String: medications, Valid: medications != ""},
		Attachments: sql.NullString{String: attachments, Valid: attachments != ""},
		UserId:      userId,
	}
}

func (h *HealthRecords) Update(patientName string, date time.Time, recordType string, diagnosis string, treatment string, notes string, vitalSigns string, medications string, attachments string, patientId int64) {
	h.PatientId = patientId
	h.PatientName = patientName
	h.Date = sql.NullTime{Time: date, Valid: !date.IsZero()}
	h.RecordType = recordType
	h.Diagnosis = sql.NullString{String: diagnosis, Valid: diagnosis != ""}
	h.Treatment = sql.NullString{String: treatment, Valid: treatment != ""}
	h.Notes = sql.NullString{String: notes, Valid: notes != ""}
	h.VitalSigns = sql.NullString{String: vitalSigns, Valid: vitalSigns != ""}
	h.Medications = sql.NullString{String: medications, Valid: medications != ""}
	h.Attachments = sql.NullString{String: attachments, Valid: attachments != ""}
	h.Attachments = sql.NullString{String: attachments, Valid: attachments != ""}
}

func (h *HealthRecords) MarkAsDeleted() {
	h.IsDeleted = 1
}
