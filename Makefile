.DEFAULT_GOAL := build 
.PHONY : hello   build   
docker : 
	@echo "Starting all docker containers"
	@echo "=============================="
	@echo "Starting  consul service"
	docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
	@echo "Starting mySQL database"
	docker run --name movieexample_db -e MYSQL_ROOT_ PASSWORD=password -e MYSQL_DATABASE=movieexample -p 3306:3306 -d mysql:latest
	@echo "Starting Jaegar for tracing "
	docker run -d --name jaeger  -e COLLECTOR_OTLP_ENABLED=true   -p 6831:6831/udp   -p 6832:6832/udp   -p 5778:5778  -p 16686:16686  -p 4317:4317   -p 4318:4318   -p 14250:14250   -p 14268:14268   -p 14269:14269   -p 9411:9411   jaegertracing/all-in-one:1.37
	@echo "Starting Apache Kafka for ingestion "
	docker run -p 9092:9092 --name kafka --link zookeeper:zookeeper \
	-e KAFKA_BROKER_ID=1 \
	-e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
	-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092,PLAINTEXT_HOST://localhost:9092 \
	-e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT \
	-e KAFKA_INTER_BROKER_LISTENER_NAME=PLAINTEXT \
	-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
	-e KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR=1 \
	-e KAFKA_TRANSACTION_STATE_LOG_MIN_ISR=1 \
	confluentinc/cp-kafka:latest
build: docker
	@echo "alll external service running succesfully "

