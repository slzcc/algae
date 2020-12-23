SHELL := /bin/bash

compile:
	@go mod download
	@go build

test:
