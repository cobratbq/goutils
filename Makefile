# SPDX-License-Identifier: LGPL-3.0-only

MAKEFLAGS += --no-builtin-rules --no-builtin-variables
.SUFFIXES:

.PHONY: all
all: build test

.PHONY: test
test: build
	go test -tags netgo ./...

.PHONY: build
build: go.mod
	go build -tags netgo ./...

.PHONY: clean
clean:
	go clean -testcache

