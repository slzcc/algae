SHELL := /bin/bash

compile:
	@glide up
	@rm -rf go.mod
	@go mod init algae
	@go mod vendor
	@go build

compile_mod:
	@export GO111MODULE=on
	@export GOPROXY=http://gopkg.ops.zhangyue-ops.com/repository/go-proxy
	@go mod download

test:
