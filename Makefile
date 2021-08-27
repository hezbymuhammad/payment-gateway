build:
	go build -o bin/main main.go

coverage:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

dep:
	go mod download

lint:
	golangci-lint run ./...

pretty:
	gofmt -s -w .

run:
	go run main.go

test:
	go test -v -race ./...

tidy:
	go mod tidy

vendor:
	go mod vendor
