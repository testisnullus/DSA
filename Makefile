.PHONY: build
build:
	@go build -o ./bin/dsa

.PHONY: run
run: build
	@./bin/dsa