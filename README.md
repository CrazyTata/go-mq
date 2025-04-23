# go-mq

基于 DDD (Domain-Driven Design) 架构的 Go 消息队列项目。

## 项目简介

go-mq 是一个基于 Go 语言开发的消息队列项目，主要用于医疗健康领域的数据处理和统计。项目采用领域驱动设计（DDD）架构，实现了健康记录的生产和统计数据的消费功能。

## 核心特性

- 🏥 健康记录管理
- 📊 医疗数据统计
- 🔄 消息队列处理
- 📈 实时数据更新
- 🔒 数据安全性保障

## 技术栈

### 核心框架
- Go 1.22+
- go-zero v1.6.0 (微服务框架)
- JWT 认证
- Redis 缓存

### 数据库
- MySQL (数据持久化)
- Redis (缓存和消息队列)

### CI/CD
- GitHub Actions

## 项目结构

```
.
├── cmd/                    # 应用程序入口
│   └── main.go            # 主程序入口
├── domain/                # 领域层
│   ├── health/           # 健康记录领域
│   │   ├── health.go     # 健康记录模型
│   │   └── repository.go # 健康记录仓储接口
│   ├── statistics/       # 统计领域
│   │   ├── statistics.go # 统计模型
│   │   └── repository.go # 统计仓储接口
│   └── vars.go           # 领域变量
├── application/          # 应用层
│   ├── dto/             # 数据传输对象
│   ├── service/         # 应用服务
│   └── usecase/         # 用例实现
├── infrastructure/       # 基础设施层
│   ├── persistence/     # 持久化实现
│   ├── config/         # 配置管理
│   ├── intergration/   # 集成服务
│   ├── provider/       # 提供者
│   ├── sql/           # 数据表DDL
│   ├── startup/        # 服务启动初始化
│   └── svc/            # 服务上下文
├── interfaces/          # 接口层
│   ├── api/            # API 接口
│   │   └── handler/    # HTTP 处理器
│   └── middleware/     # 中间件
├── common/             # 公共组件
│   ├── errors/        # 错误定义
│   ├── utils/         # 工具函数
│   └── constants/     # 常量定义
├── etc/               # 配置文件
├── pkg/               # 可重用的包
├── scripts/           # 脚本文件
├── test/              # 测试文件
├── vendor/            # 依赖包
├── go.mod             # Go 模块文件
├── go.sum             # Go 依赖版本锁定
├── Makefile           # 构建脚本
├── Dockerfile         # Docker 构建文件
└── docker-compose.yml # Docker 编排文件
```

## 功能特性

### 健康记录管理
- 健康记录创建
- 健康记录更新
- 健康记录删除
- 健康记录查询

### 医疗数据统计
- 患者总数统计
- 活跃患者统计
- 预约数据统计
- 健康记录统计
- 操作记录统计

### 消息队列处理
- 健康记录生产
- 统计数据消费
- 实时数据更新
- 数据一致性保证

## 快速开始

### 安装

```bash
go get github.com/your-username/go-mq
```

### 配置

```yaml
# etc/config.yaml
server:
  port: 8888
  host: 0.0.0.0

database:
  host: localhost
  port: 3306
  username: root
  password: root
  dbname: go_mq

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
```

### 使用示例

#### 创建健康记录

```bash
curl -X POST http://localhost:8888/v1/health/save \
  -H "Content-Type: application/json" \
  -d '{
    "patient_id": 1,
    "patient_name": "张三",
    "date": "2024-04-23",
    "record_type": "门诊",
    "diagnosis": "感冒",
    "treatment": "休息",
    "notes": "多喝水",
    "vital_signs": "正常",
    "medications": "感冒药",
    "attachments": ""
  }'
```

响应示例：
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 1
  }
}
```

#### 获取健康记录列表

```bash
curl -X GET "http://localhost:8888/v1/health/list?page=1&page_size=10"
```

响应示例：
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "patient_id": 1,
        "patient_name": "张三",
        "date": "2024-04-23",
        "record_type": "门诊",
        "diagnosis": "感冒",
        "treatment": "休息",
        "notes": "多喝水",
        "vital_signs": "正常",
        "medications": "感冒药",
        "attachments": "",
        "user_id": "user123",
        "created_at": "2024-04-23 10:00:00",
        "updated_at": "2024-04-23 10:00:00"
      }
    ],
    "total": 1,
    "page": 1,
    "page_size": 10
  }
}
```

## 开发指南

1. 遵循 DDD 原则进行开发
2. 使用 Go 1.22 或更高版本
3. 实现适当的错误处理和日志记录
4. 编写单元测试和集成测试
5. 遵循代码规范

## 构建和运行

```bash
# 安装依赖
go mod download

# 运行测试
make test

# 构建项目
make build

# 运行服务
make run
```

## 贡献指南

1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

MIT License
