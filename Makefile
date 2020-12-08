SHELL := /bin/bash

compile:
	@export GO111MODULE=on
	@export GOPROXY=http://gopkg.ops.zhangyue-ops.com/repository/go-proxy
	@go mod download

test:
