# HELP
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: run application in cmd/goscawa
.PHONY: run
run:
	go run cmd/goscawa/main.go 


## build: build application cmd/goscawa into bin/
.PHONY: build
build:
	go build -ldflags='-s' -o=./bin/server ./cmd/goscawa

## test: run tests
.PHONY: test
test:
	go test pkg/*