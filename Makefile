.PHONY: run 
# 构建
build:
	@echo "拉取代码.."
	@git pull
	@echo "构建镜像..."
	@docker-compose build go-mq
	@echo "停止容器..."
	@docker-compose stop go-mq
	@echo "启动容器..."
	docker-compose up -d go-mq

model:
	#@goctl model mysql ddl -src infrastructure/sql/user.sql -dir infrastructure/persistence/model/user -c
	@goctl model mysql ddl -src infrastructure/sql/patients.sql -dir infrastructure/persistence/model/patients -c
	@goctl model mysql ddl -src infrastructure/sql/appointments.sql -dir infrastructure/persistence/model/appointments -c
	@goctl model mysql ddl -src infrastructure/sql/health_records.sql -dir infrastructure/persistence/model/health_records -c
	@goctl model mysql ddl -src infrastructure/sql/operation_records.sql -dir infrastructure/persistence/model/operation_records -c
	@goctl model mysql ddl -src infrastructure/sql/statistics.sql -dir infrastructure/persistence/model/statistics -c
	