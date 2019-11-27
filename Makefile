SHELL := /bin/bash

# The name of the executable (default is current directory name)
TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

# These will be provided to the target
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Testing flags
TEST_FLAGS=-v
COVERAGE_RESULTS=coverage.out

.PHONY: fmt simplify vet lint run test coverage

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

vet:
	@for d in $$(go list ./... | grep -v /vendor/); do go vet $${d}; done

lint:
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done

test:
	@for d in $$(go list ./... | grep -v /vendor/); do go test $(TEST_FLAGS) $${d}; done

coverage: TEST_FLAGS+= -coverprofile=$(COVERAGE_RESULTS)
coverage: test
	@go tool cover -html=coverage.out

run: install
	@$(TARGET)

print-%  : ; @echo $* = $($*)
