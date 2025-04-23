package queue

import (
	"context"
	"sync"

	"github.com/zeromicro/go-zero/core/logx"
)

// ConsumerManager 消费者管理器
type ConsumerManager struct {
	consumers []Consumer
	queue     QueueManager
}

// Consumer 消费者接口
type Consumer interface {
	Start(ctx context.Context) error
}

// NewConsumerManager 创建消费者管理器
func NewConsumerManager(queue QueueManager) *ConsumerManager {
	return &ConsumerManager{
		queue: queue,
	}
}

// Register 注册消费者
func (m *ConsumerManager) Register(consumer Consumer) {
	logger := logx.WithContext(context.Background())
	logger.Infof("注册消费者: %T", consumer)
	m.consumers = append(m.consumers, consumer)
}

// StartAll 启动所有消费者
func (m *ConsumerManager) StartAll(ctx context.Context) error {
	logger := logx.WithContext(ctx)
	logger.Infof("开始启动所有消费者，数量: %d", len(m.consumers))

	var wg sync.WaitGroup
	errChan := make(chan error, len(m.consumers))

	// 先注册所有消费者
	for _, consumer := range m.consumers {
		wg.Add(1)
		go func(c Consumer) {
			defer wg.Done()
			logger.Infof("注册消费者: %T", c)
			if err := c.Start(ctx); err != nil {
				logger.Errorf("消费者注册失败: %T, 错误: %v", c, err)
				errChan <- err
			} else {
				logger.Infof("消费者注册成功: %T", c)
			}
		}(consumer)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	// 启动消息分发
	logger.Infof("开始启动消息分发")
	if err := m.queue.Start(ctx); err != nil {
		logger.Errorf("消息分发启动失败: %v", err)
		return err
	}

	logger.Infof("所有消费者启动完成")
	return nil
}
