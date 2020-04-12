#!make

test: test-errors test-config test-mysql test-redis test-rabbitmq
	# go test -timeout 300s -cover -a -v ./...

test-errors:
	cd ./errors && \
	go test -timeout 300s -cover -a -v

test-config:
	cd ./config && \
	go test -timeout 300s -cover -a -v

test-mysql:
	cd ./libs/mysql && \
	go test -timeout 300s -cover -a -v

test-redis:
	cd ./libs/redis && \
	go test -timeout 300s -cover -a -v

test-rabbitmq:
	cd ./libs/rabbitmq && \
	go test -timeout 300s -cover -a -v

integration-test:
	cd ./integration_test && \
	go test -timeout 300s -cover -a -v ./...

start-gateway:
	cd ./gateway && \
	export service=gateway && \
	go run main.go

start-cache:
	cd ./entry_cache && \
	export service=entry-cache && \
	go run main.go

start-store:
	cd ./entry_store && \
	export service=entry-store && \
	go run main.go

start-bo-store:
	cd ./bo_entry_store && \
	export service=bo-entry-store && \
	go run main.go

start-bo:
	cd ./bo_controller && \
	export service=bo-controller && \
	go run main.go