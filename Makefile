# SPDX-License-Identifier: LGPL-3.0-only

MAKEFLAGS += --no-builtin-rules --no-builtin-variables
.SUFFIXES:

.PHONY: all
all: build test

.PHONY: test
test: build
	go test ./...

.PHONY: build
build: go.mod
	go build ./...

.PHONY: clean
clean:
	go clean -testcache

