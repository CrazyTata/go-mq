package queue

import (
	"context"
	"encoding/json"
	"go-mq/infrastructure/svc"
	"sync"
	"time"

	"github.com/samber/lo"
	"github.com/zeromicro/go-queue/dq"
	"github.com/zeromicro/go-zero/core/logx"
)

var _ QueueManager = &DQQueue{}

// DQConfig DQ队列配置
type DQConfig struct {
	// 重试配置
	MaxRetries     int
	RetryInterval  time.Duration
	ReconnectDelay time.Duration
	MaxReconnects  int
	// 消息延迟时间
	DelayTime time.Duration
}

// DefaultDQConfig 返回默认 DQ 配置
func DefaultDQConfig(svcCtx *svc.ServiceContext) *DQConfig {
	return &DQConfig{
		MaxRetries:     3,
		RetryInterval:  time.Second * 5,
		ReconnectDelay: time.Second * 5,
		MaxReconnects:  5,
		DelayTime:      time.Second * 5,
	}
}

// DQQueue 实现
type DQQueue struct {
	config     *DQConfig
	producer   dq.Producer
	consumer   dq.Consumer
	mu         sync.RWMutex
	isClosed   bool
	handlers   map[MessageType]func(message *Message) error
	queueTypes []MessageType
}

// NewDQQueue 创建新的 DQQueue 实例
func NewDQQueue(svcCtx *svc.ServiceContext) *DQQueue {
	config := DefaultDQConfig(svcCtx)

	// 创建生产者
	producer := dq.NewProducer(svcCtx.GetConfig().DqConf.Beanstalks)

	// 创建消费者
	consumer := dq.NewConsumer(svcCtx.GetConfig().DqConf)

	return &DQQueue{
		config:     config,
		producer:   producer,
		consumer:   consumer,
		handlers:   make(map[MessageType]func(message *Message) error),
		queueTypes: make([]MessageType, 0),
	}
}

// Publish 发布消息
func (q *DQQueue) Publish(ctx context.Context, message *Message) error {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.isClosed {
		return ErrQueueClosed
	}

	msgBody, err := json.Marshal(message)
	if err != nil {
		return err
	}

	for i := 0; i < q.config.MaxRetries; i++ {
		_, err := q.producer.Delay(msgBody, q.config.DelayTime)
		if err == nil {
			return nil
		}
		time.Sleep(q.config.RetryInterval)
	}

	return ErrPublishFailed
}

// Consume 注册消息消费者
func (q *DQQueue) Consume(ctx context.Context, queueTypes []MessageType, handler func(message *Message) error) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isClosed {
		return ErrQueueClosed
	}

	// 注册消息处理器
	for _, queueType := range queueTypes {
		q.handlers[queueType] = handler
		if !lo.Contains(q.queueTypes, queueType) {
			q.queueTypes = append(q.queueTypes, queueType)
		}
	}

	return nil
}

// Start 启动消息分发
func (q *DQQueue) Start(ctx context.Context) error {
	q.mu.RLock()
	if q.isClosed {
		q.mu.RUnlock()
		return ErrQueueClosed
	}
	queueTypes := make([]MessageType, len(q.queueTypes))
	copy(queueTypes, q.queueTypes)
	q.mu.RUnlock()

	// 使用单个 goroutine 来处理消费
	go func() {
		for {
			select {
			case <-ctx.Done():
				logx.Info("Context cancelled, stopping consumer")
				return
			default:
				if q.isClosed {
					logx.Info("Queue is closed, stopping consumer")
					return
				}

				logx.Info("Starting consumer...")

				// 直接调用消费方法
				q.consumer.Consume(func(body []byte) {
					// 解析消息
					var msg Message
					if err := json.Unmarshal(body, &msg); err != nil {
						logx.Errorf("Failed to unmarshal message: %v", err)
						return
					}

					// 检查消息类型是否在订阅列表中
					if !lo.Contains(queueTypes, msg.Type) {
						logx.Infof("Message type %v not in queue types %v", msg.Type, queueTypes)
						return
					}

					// 获取对应的处理器
					q.mu.RLock()
					handler, exists := q.handlers[msg.Type]
					q.mu.RUnlock()

					if !exists {
						logx.Errorf("No handler found for message type: %v", msg.Type)
						return
					}

					// 处理消息
					if err := handler(&msg); err != nil {
						logx.Errorf("Failed to handle message: %v", err)
					} else {
						logx.Info("Message handled successfully")
					}
				})

				// 如果发生错误，等待一段时间后重试
				logx.Info("Consumer finished, waiting for retry...")
				time.Sleep(q.config.RetryInterval)
			}
		}
	}()

	return nil
}

// Close 关闭连接
func (q *DQQueue) Close() error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isClosed {
		return nil
	}

	q.isClosed = true

	// 关闭生产者
	if q.producer != nil {
		q.producer.Close()
	}

	return nil
}

// 错误定义
var (
	ErrQueueClosed   = NewError("queue is closed")
	ErrPublishFailed = NewError("failed to publish message")
)

// Error 自定义错误类型
type Error struct {
	message string
}

// NewError 创建新的错误
func NewError(message string) error {
	return &Error{message: message}
}

func (e *Error) Error() string {
	return e.message
}
