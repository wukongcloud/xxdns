GOCMD			:=$(shell which go)
GOBUILD			:=$(GOCMD) build

IMPORT_PATH		:=github.com/wukongcloud/xxdns/utils
BUILD_TIME		:=$(shell date "+%F %T")
COMMIT_ID       :=$(shell git rev-parse HEAD)
GO_VERSION      :=$(shell $(GOCMD) version)
VERSION			:=$(shell git describe --tags)
BUILD_USER		:=$(shell whoami)

FLAG			:="-w -s -X '${IMPORT_PATH}.BuildTime=${BUILD_TIME}' -X '${IMPORT_PATH}.CommitID=${COMMIT_ID}' -X '${IMPORT_PATH}.GoVersion=${GO_VERSION}'  -X '${IMPORT_PATH}.Version=${VERSION}' -X '${IMPORT_PATH}.BuildUser=${BUILD_USER}'"

BINARY_DIR=bin
BINARY_NAME:=xxdns

# linux
build-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags $(FLAG) -o $(BINARY_DIR)/$(BINARY_NAME)-linux

#mac
build-darwin:
	CGO_ENABLED=0 GOOS=darwin $(GOBUILD) -ldflags $(FLAG) -o $(BINARY_DIR)/$(BINARY_NAME)-darwin

# windows
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags $(FLAG) -o $(BINARY_DIR)/$(BINARY_NAME)-win.exe
# 全平台
build-all:
	make build-linux
	make build-darwin
	make build-win
	./upx $(BINARY_DIR)/$(BINARY_NAME)-*build.sh