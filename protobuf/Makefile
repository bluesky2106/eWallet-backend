build:
	protoc -I. --go_out=plugins=grpc:. base.proto
	protoc -I. --go_out=plugins=grpc:. user.proto
	protoc -I. --go_out=plugins=grpc:. product.proto
	protoc -I. --go_out=plugins=grpc:. email.proto