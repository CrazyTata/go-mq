package queue

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ QueueManager = &GoMQ{}

// GoMQ 基于 Go channel 的消息队列实现
type GoMQ struct {
	queues      map[MessageType]chan *Message
	subscribers map[MessageType][]chan *Message
	mu          sync.RWMutex
	closed      bool
}

var (
	instance *GoMQ
	once     sync.Once
)

// NewGoMQ 创建一个新的 GoMQ 实例
func NewGoMQ() *GoMQ {
	once.Do(func() {
		instance = &GoMQ{
			queues:      make(map[MessageType]chan *Message),
			subscribers: make(map[MessageType][]chan *Message),
		}
	})
	return instance
}

// Publish 发布消息到队列
func (q *GoMQ) Publish(ctx context.Context, message *Message) error {
	logger := logx.WithContext(ctx)
	logger.Infof("开始发布消息: %v", message)
	if q.closed {
		return errors.New("queue is closed")
	}

	q.mu.RLock()
	queue, exists := q.queues[message.Type]
	q.mu.RUnlock()

	if !exists {
		logger.Errorf("未注册该类型的消费者，无法发布消息: %v", message.Type)
		return errors.New("no consumer for message type")
	}
	logger.Infof("准备发送消息到队列: %v", message)
	// 将消息发送到队列
	select {
	case queue <- message:
		logger.Infof("消息成功发送到队列: %v", message)
		return nil
	case <-ctx.Done():
		logger.Errorf("发布消息超时: %v", ctx.Err())
		return ctx.Err()
	default:
		logger.Errorf("队列已满，无法发送消息: %v", message)
		return errors.New("queue is full")
	}
}

// Consume 注册消息消费者
func (q *GoMQ) Consume(ctx context.Context, queueTypes []MessageType, handler func(message *Message) error) error {
	logger := logx.WithContext(ctx)
	logger.Infof("注册消费者订阅消息类型: %v", queueTypes)
	if q.closed {
		return errors.New("queue is closed")
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	// 为每个消息类型创建订阅者通道
	subscriber := make(chan *Message, 1000)
	for _, queueType := range queueTypes {
		// 确保队列存在
		if _, exists := q.queues[queueType]; !exists {
			logger.Infof("创建新的队列通道: %v", queueType)
			q.queues[queueType] = make(chan *Message, 1000)
		}
		q.subscribers[queueType] = append(q.subscribers[queueType], subscriber)
	}
	logger.Infof("当前订阅者: %v", q.subscribers)

	// 启动消费者 goroutine
	go func() {
		logger.Infof("消费者 goroutine 启动")
		defer logger.Infof("消费者 goroutine 停止")

		for {
			select {
			case <-ctx.Done():
				logger.Infof("消费者 goroutine 收到取消信号")
				return
			case msg := <-subscriber:
				logger.Infof("消费者收到消息: %v", msg)
				if err := handler(msg); err != nil {
					logger.Errorf("处理消息失败: %v, 错误: %v", msg, err)
					continue
				}
				logger.Infof("消息处理成功: %v", msg)
			}
		}
	}()

	return nil
}

// Start 启动消息分发
func (q *GoMQ) Start(ctx context.Context) error {
	logger := logx.WithContext(ctx)
	logger.Infof("开始启动消息分发")

	q.mu.RLock()
	queueTypes := make([]MessageType, 0, len(q.queues))
	for queueType := range q.queues {
		queueTypes = append(queueTypes, queueType)
	}
	q.mu.RUnlock()

	// 为每个消息类型启动一个分发 goroutine
	for _, queueType := range queueTypes {
		queue := q.queues[queueType]
		go func(qt MessageType, ch chan *Message) {
			logger.Infof("分发 goroutine 启动: %v", qt)
			defer logger.Infof("分发 goroutine 停止: %v", qt)

			// 使用 time.After 替代 Timer，避免内存泄漏
			idleDuration := time.Millisecond * 100
			idleTimer := time.After(idleDuration)

			for {
				select {
				case <-ctx.Done():
					logger.Infof("分发 goroutine 收到取消信号: %v", qt)
					return
				case msg, ok := <-ch:
					if !ok {
						logger.Infof("队列通道已关闭: %v", qt)
						return
					}

					// 重置空闲定时器
					idleTimer = time.After(idleDuration)

					logger.Infof("从队列 %v 获取到消息: %v", qt, msg)

					q.mu.RLock()
					subscribers := q.subscribers[msg.Type]
					q.mu.RUnlock()

					if len(subscribers) == 0 {
						logger.Infof("没有订阅者订阅消息类型: %v", msg.Type)
						continue
					}

					logger.Infof("准备将消息分发给 %d 个订阅者", len(subscribers))
					// 发送消息给所有订阅者
					for i, sub := range subscribers {
						select {
						case sub <- msg:
							logger.Infof("消息成功发送给订阅者 %d: %v", i, msg)
						case <-ctx.Done():
							logger.Infof("分发过程中收到取消信号")
							return
						}
					}
				case <-idleTimer:
					// 没有消息时，重置空闲定时器
					idleTimer = time.After(idleDuration)
				}
			}
		}(queueType, queue)
	}

	return nil
}

// Close 关闭队列
func (q *GoMQ) Close() error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return nil
	}

	q.closed = true

	// 关闭所有队列通道
	for _, queue := range q.queues {
		close(queue)
	}

	// 关闭所有订阅者通道
	for _, subscribers := range q.subscribers {
		for _, sub := range subscribers {
			close(sub)
		}
	}

	return nil
}
