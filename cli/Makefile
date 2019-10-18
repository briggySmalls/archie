SHELL := /bin/bash

define src
	$(shell find $(1) -type f -name '*.go' -not -path "./vendor/*")
endef

BIN_DIR := bin

CLI := archie
CLI_MAIN := main.go
CLI_SRC := $(call src,cmd) $(CLI_MAIN)

WASM := $(BIN_DIR)/archie-lib.wasm
WASM_MAIN := ./wasm/main.go
WASM_SRC := $(call src,wasm)

$(CLI): $(CLI_SRC) pack
	go build $(LDFLAGS) -o $(CLI) $(CLI_MAIN)

pack: $(WASM) runtime
	packr

$(WASM): export GOOS=js
$(WASM): export GOARCH=wasm
$(WASM): $(WASM_SRC)
	go build $(LDFLAGS) -o $(WASM) $(WASM_SRC)

runtime:
	cp $(shell go env GOROOT)"/misc/wasm/wasm_exec.js" $(BIN_DIR)

clean:
	rm -rf $(BIN_DIR)

fmt:
	@gofmt -l -w $(CLI_SRC) $(WASM_SRC)

simplify:
	@gofmt -s -l -w $(CLI_SRC) $(WASM_SRC)

check:
	@echo "[Formatting]"
	@test -z $(shell gofmt -l $(MAIN) | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@echo "[Vetting]"
	@for d in $$(go list ./... | grep -v /vendor/); do go vet $${d}; done
	@echo "[Linting]"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done

test:
	@go test $(TEST_FLAGS) ./...

coverage: TEST_FLAGS+= -coverprofile=$(COVERAGE_RESULTS)
coverage: test
	@go tool cover -html=coverage.out

print-%: ; @echo $* = $($*)
