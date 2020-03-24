#!make

test:
	go test -timeout 9000s -cover -a -v ./...

test-errors:
	cd ./errors && \
	go test -timeout 9000s -cover -a -v

test-config:
	cd ./config && \
	go test -timeout 9000s -cover -a -v

test-mysql:
	cd ./libs/mysql && \
	go test -timeout 9000s -cover -a -v

test-redis:
	cd ./libs/redis && \
	go test -timeout 9000s -cover -a -v

start-gateway:
	cd ./gateway && \
	export service=gateway && \
	go run main.go

start-cache:
	cd ./entry_cache && \
	export service=entry-cache && \
	go run main.go