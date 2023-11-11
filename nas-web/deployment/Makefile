BUILD_PATH := $(shell dirname $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST)))))
BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

all: build

build: nas_web

nas_web:
	${BUILD_ENV} go build  -ldflags "-s -w"  -mod=mod -o ${BUILD_PATH}/bin/nas_web ${BUILD_PATH}/cmd/nas_web.go
