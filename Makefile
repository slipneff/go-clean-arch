run-dev:
	go run cmd/service/main.go -env-mode=development -config-path=config.yaml
seed:
	go run cmd/seeder/main.go -env-mode=development -config-path=config.yaml

.PHONY: clean
generate:
		protoc --go_out=internal/generated/ --go_opt=paths=source_relative --go-grpc_out=internal/generated/ --go-grpc_opt=paths=source_relative proto/user/user.proto