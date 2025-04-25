package queue

// MessageType 消息类型
type MessageType string

const (
	// HealthRecordSaved 健康记录保存消息
	HealthRecordSaved MessageType = "health.record.saved"
)
