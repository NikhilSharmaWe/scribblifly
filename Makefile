build:
	@go build -o bin/scribblifly

run: build
	@./bin/gobank
