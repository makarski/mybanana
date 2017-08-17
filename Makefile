.PHONY: all build build-app run run-app deps version config test logs stop

GO ?= go
COMPOSE ?= docker-compose
DOCKER ?= docker
ECHO ?= echo

BIN_DIR = $(GOPATH)/bin
BIN_NAME = $(BIN_DIR)/makarski
REVISION_NAME = $(BIN_DIR)/REVISION
IMPORT_PATH = github.com/makarski/mybanana
CNT_NAME_APP = mybanana-api
CNT_NAME_DB = mybanana-db

all: build

deps:
	$(GO) get -u github.com/go-chi/chi
	$(GO) get -u github.com/go-gorp/gorp
	$(GO) get -u github.com/kelseyhightower/envconfig

build:
	CNT_NAME_APP=$(CNT_NAME_APP) CNT_NAME_DB=$(CNT_NAME_DB) $(COMPOSE) build

test: deps
test:
	$(GO) vet ./...
	$(GO) fmt ./...
	$(GO) test ./... -v

build-app: test version config
	cd $(GOPATH)/src/$(IMPORT_PATH)
	BIN_DIR=$(BIN_DIR) $(GO) build -o $(BIN_NAME)

run:
	CNT_NAME_APP=$(CNT_NAME_APP) CNT_NAME_DB=$(CNT_NAME_DB) $(COMPOSE) up -d

logs:
	$(DOCKER) logs $(CNT_NAME_APP) --follow

stop:
	$(DOCKER) stop $(CNT_NAME_APP) $(CNT_NAME_DB)

run-app:
	BIN_DIR=$(BIN_DIR) $(BIN_NAME)

config:
# 	todo

version:
	# $(ECHO)	`git log -n 1 --pretty=format:"%H"` > $(REVISION_NAME)
