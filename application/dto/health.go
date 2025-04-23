package dto

// HealthRequest 创建健康记录请求
type HealthRequest struct {
	Id          int64  `json:"id,optional"`
	PatientID   int64  `json:"patient_id"`
	PatientName string `json:"patient_name"`
	Date        string `json:"date"`
	RecordType  string `json:"record_type"`
	Diagnosis   string `json:"diagnosis"`
	Treatment   string `json:"treatment"`
	Notes       string `json:"notes"`
	VitalSigns  string `json:"vital_signs"`
	Medications string `json:"medications"`
	Attachments string `json:"attachments"`
}

// CreateHealthResponse 创建健康记录响应
type CreateHealthResponse struct {
	ID int64 `json:"id"`
}

// HealthResponse 健康记录响应
type HealthResponse struct {
	Id          int64  `json:"id"`
	PatientId   int64  `json:"patient_id"`   // 患者ID
	PatientName string `json:"patient_name"` // 患者姓名
	Date        string `json:"date"`         // 日期
	RecordType  string `json:"record_type"`  // 记录类型
	Diagnosis   string `json:"diagnosis"`    // 诊断
	Treatment   string `json:"treatment"`    // 治疗
	Notes       string `json:"notes"`        // 备注
	VitalSigns  string `json:"vital_signs"`  // 生命体征
	Medications string `json:"medications"`  // 药物
	Attachments string `json:"attachments"`  // 附件
	UserId      string `json:"user_id"`      // 用户ID
	CreatedAt   string `json:"created_at"`   // 创建时间
	UpdatedAt   string `json:"updated_at"`
}

// HealthListResponse 健康记录列表响应
type HealthListResponse struct {
	List     []*HealthResponse `json:"list"`
	Total    int64             `json:"total"`
	Page     int64             `json:"page"`
	PageSize int64             `json:"page_size"`
}

// GetHealthListRequest 获取健康记录列表请求
type GetHealthListRequest struct {
	Page      int64  `form:"page"`
	PageSize  int64  `form:"page_size"`
	PatientId int64  `form:"patient_id,optional"`
	Search    string `form:"search,optional"`
	Status    string `form:"status,optional"`
	Order     int64  `form:"order,optional"` // 1: 降序, 0: 升序
}
