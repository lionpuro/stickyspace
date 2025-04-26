build:
	@go build -o tmp/run .

run: build
	@./tmp/run

dev:
	@air -c .air.toml

fmt:
	@gofmt -l -s -w .
