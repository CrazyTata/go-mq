package queue

import "context"

// Message 消息接口
type Message struct {
	ID   string
	Type MessageType
	Body []byte
}

// QueueManager 队列管理接口
type QueueManager interface {
	// Publish 发布消息
	Publish(ctx context.Context, message *Message) error
	// Consume 消费消息
	Consume(ctx context.Context, queueType []MessageType, handler func(message *Message) error) error
	// Start 启动消息分发
	Start(ctx context.Context) error
	// Close 关闭连接
	Close() error
}
