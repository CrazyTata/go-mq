CREATE TABLE IF NOT EXISTS `statistics` (
    `id` INT UNSIGNED AUTO_INCREMENT COMMENT '统计ID',  
    `user_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '用户ID', 
    `total_patients` int NOT NULL DEFAULT 0 COMMENT '总患者数',
    `active_patients` int NOT NULL DEFAULT 0 COMMENT '活跃患者数',
    `today_appointments` int NOT NULL DEFAULT 0 COMMENT '今日预约数',
    `upcoming_appointments` int NOT NULL DEFAULT 0 COMMENT '未来预约数',
    `health_records` int NOT NULL DEFAULT 0 COMMENT '健康记录总数',
    `recent_operations` int NOT NULL DEFAULT 0 COMMENT '最近操作记录数',
    `date` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '日期',
    `is_deleted` TINYINT NOT NULL DEFAULT 0 COMMENT '是否删除',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_date` (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='统计表';