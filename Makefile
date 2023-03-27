# SPDX-License-Identifier: AGPL-3.0-or-later

MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

.PHONY: all
all: build test

.PHONY: test
test: build
	go test -v -tags netgo ./...

.PHONY: build
build: go.mod
	go build -tags netgo ./...

.PHONY: clean
clean:
	go clean -testcache

