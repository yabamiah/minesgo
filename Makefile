GOCMD=go
BINARY_NAME=minesgo
VERSION?=0.0.0

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all build vendor

build:
	GOARCH=amd64 GOOS=darwin $(GOCMD) build -o build/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux $(GOCMD) build -o build/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows $(GOCMD) build -o build/${BINARY_NAME}-windows main.go

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm -rf ${BINARY_NAME}-darwin
	rm -rf ${BINARY_NAME}-linux
	rm -rf ${BINARY_NAME}-windows

help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)