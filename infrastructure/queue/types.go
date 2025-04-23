package queue

// MessageType 消息类型
type MessageType string

const (
	// UserCreated 用户创建消息
	UserCreated MessageType = "user.created"
	// UserUpdated 用户更新消息
	UserUpdated MessageType = "user.updated"
	// UserDeleted 用户删除消息
	UserDeleted MessageType = "user.deleted"
	// AppointmentCreated 预约创建消息
	AppointmentCreated MessageType = "appointment.created"
	// AppointmentUpdated 预约更新消息
	AppointmentUpdated MessageType = "appointment.updated"
	// AppointmentCancelled 预约取消消息
	AppointmentCancelled MessageType = "appointment.cancelled"
	// HealthRecordSaved 健康记录保存消息
	HealthRecordSaved MessageType = "health.record.saved"
	// HealthRecordDeleted 健康记录删除消息
	HealthRecordDeleted MessageType = "health.record.deleted"
	// OperationCreated 手术创建消息
	OperationCreated MessageType = "operation.created"
	// OperationUpdated 手术更新消息
	OperationUpdated MessageType = "operation.updated"
)
