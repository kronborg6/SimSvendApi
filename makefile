build:
	@go build -o bin/SimSvend

run: build
	@./bin/SimSvend

test:
	@go test -v ./...