.PHONY: all
all: help

APP=goland
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
VERSION?=1.0

BUILD?=$(shell git describe --tags --always --dirty)
DELVE:=$(shell command -v dlv 2> /dev/null)
GOLINT:=$(shell command -v golangci-lint 2> /dev/null)
GOSEC:=$(shell command -v gosec 2> /dev/null)
GLICE:=$(shell command -v glice 2> /dev/null)
GOJUNITCOVER:=$(shell command -v go-junit-report 2> /dev/null)
GOMOCKGEN:=$(shell command -v mockgen 2> /dev/null)
SWAGGER:=$(shell command -v swagger 2> /dev/null)
RICHGO=$(shell command -v richgo 2> /dev/null)

BIN_DIR=bin
APP_EXECUTABLE=./$(BIN_DIR)/$(APP)
REPORTS_DIR=reports
PERF_REPORTS_DIR=perf-reports

PROTOCOL?=https
PORT?=8000
HOSTNAME?=localhost
APP_URL=$(PROTOCOL)://$(HOSTNAME):$(PORT)/

ifeq ($(RICHGO),)
	GOBIN=go
else
	GOBIN=richgo
endif

setup: ## Setup necessary dependencies and folder structure
ifndef GOLINT
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.16.0
endif
ifndef GOCOVER
	$(GOBIN) get golang.org/x/tools/cmd/cover
endif
ifndef GOJUNITCOVER
	$(GOBIN) get github.com/jstemmer/go-junit-report
endif
ifndef GOSEC
	curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b $(GOPATH)/bin 1.3.0
endif
ifndef GLICE
	$(GOBIN) get github.com/ribice/glice
	$(GOBIN) install github.com/ribice/glice
endif
ifndef GOMOCKGEN
	$(GOBIN) get github.com/golang/mock/gomock
	$(GOBIN) install github.com/golang/mock/mockgen
endif
ifndef SWAGGER
	curl -sfL https://github.com/go-swagger/go-swagger/releases/download/v0.19.0/swagger_linux_amd64 -o $(GOPATH)/bin/swagger
	chmod +x $(GOPATH)/bin/swagger
endif
ifndef DELVE
	CGO_ENABLED=1 $(GOBIN) get github.com/go-delve/delve/cmd/dlv
endif

help:
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean:
	$(GOBIN) clean -r -cache -testcache
	rm -rf $(APP_EXECUTABLE) $(REPORTS_DIR)/* $(PERF_REPORTS_DIR) /generated_mocks *.out *.log

run: compile ## Build and start app locally (outside docker)
	GIN_MODE=debug PORT=$(PORT) $(APP_EXECUTABLE)

build: fmt test analyze compile

fmt: ## Run the code formatter
	$(GOBIN) fmt $(ALL_PACKAGES)

test: generate-docs generate-mocks ## Run tests
	mkdir -p $(REPORTS_DIR)
	GIN_MODE=test $(GOBIN) test $(ALL_PACKAGES) -v -coverprofile ./$(REPORTS_DIR)/coverage

compile: generate-docs ## Build the app
	$(GOBIN) build -i -o $(APP_EXECUTABLE)


generate-mocks: ## Generate mocks to be used only for unit testing
	rm -rf ./generated_mocks
	mkdir -p ./generated_mocks
	# Application mocks
	mockgen -source=./controllers/hello.go -destination=./generated_mocks/hello.go -package=generated_mocks
	mockgen -source=./services/hello.go -destination=./generated_mocks/hello.go -package=generated_mocks

analyze: generate-docs lint gosec

generate-docs:  ## Generates the static files to be embedded into the application + swagger.json
	mkdir -p bin 2> /dev/null
	swagger generate spec -b ./ -o ./bin/swagger.json --scan-models

lint: ## Run the code linter
	golangci-lint run

gosec:
	mkdir -p $(REPORTS_DIR)
	echo 'gosec -fmt=text -out=$(REPORTS_DIR)/gosec-report.txt ./...'


