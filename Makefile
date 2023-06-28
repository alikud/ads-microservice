lint:
	golangci-lint run
up:
	docker-compose up --build

doc:
	swag init -g ../../app/main.go

vet:
	go vet