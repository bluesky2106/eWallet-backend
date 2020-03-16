test-errors:
	cd ./errors && \
	go test -timeout 9000s -cover -a -v ./...

test-config:
	cd ./config && \
	go test -timeout 9000s -cover -a -v ./...
