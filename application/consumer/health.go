package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"go-mq/application/service/facade"
	"go-mq/domain/health"
	"go-mq/infrastructure/queue"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthConsumer struct {
	queue      queue.QueueManager `wire:"ProviderGoQueueManager"` // 队列管理器
	userFacade *facade.UserFacade
}

func NewHealthConsumer(queue queue.QueueManager, userFacade *facade.UserFacade) *HealthConsumer {
	return &HealthConsumer{queue: queue, userFacade: userFacade}
}

func (c *HealthConsumer) Start(ctx context.Context) error {
	logger := logx.WithContext(ctx)
	logger.Infof("HealthConsumer 开始启动")
	defer logger.Infof("HealthConsumer 启动完成")

	return c.queue.Consume(ctx, []queue.MessageType{queue.HealthRecordSaved}, c.handleMessage)
}

func (c *HealthConsumer) handleMessage(message *queue.Message) error {
	ctx := context.Background()
	logger := logx.WithContext(ctx)
	logger.Infof("开始处理健康记录消息: %v", message)

	if message == nil {
		logger.Errorf("消息为空")
		return nil
	}

	var healthRecord health.HealthRecords
	if err := json.Unmarshal(message.Body, &healthRecord); err != nil {
		logger.Errorf("解析健康记录失败: %v, 错误: %v", message, err)
		return fmt.Errorf("failed to unmarshal health record: %v", err)
	}

	logger.Infof("成功解析健康记录: %v", healthRecord)

	if err := c.userFacade.UpdateHealthStatistics(healthRecord.UserId); err != nil {
		logger.Errorf("更新健康统计失败: %v, 错误: %v", healthRecord.UserId, err)
		return err
	}

	logger.Infof("健康统计更新成功: %v", healthRecord.UserId)
	return nil
}
