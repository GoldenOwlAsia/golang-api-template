run:
	go run main.go wire_gen.go
start:
	go run main.go wire_gen.go
watch:
	air -d
build:
	go build -o api .
tidy:
	go mod download && go mod tidy
test:
	go test ./... -v -cover
lint:
	golangci-lint run -v