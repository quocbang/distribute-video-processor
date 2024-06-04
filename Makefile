.PHONY: proto run db

proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/quality/quality.proto

run: 
	go run cmd/http/main.go

db:
	docker-compose --env-file=.env  -f docker/docker-compose.yml down && docker-compose --env-file=.env -f docker/docker-compose.yml up --build -d;