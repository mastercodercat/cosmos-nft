PACKAGES=$(shell go list ./... | grep -v '/simulation')
BINDIR ?= $(GOPATH)/bin
# VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
# COMMIT := $(shell git log -1 --format='%H')

SIMAPP := ./app

# TODO: Update the ldflags with the app, client & server names
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=GreenApp \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=greend \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=greencli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

# Uncomment when you have some tests
#test:
# 	@go test -mod=readonly $(PACKAGES)
#.PHONY: test


test-sim-nondeterminism:
	@echo "Running non-determinism test..."
	@go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=100 -BlockSize=200 -Commit=true -Period=0 -v -timeout 24h

test-sim-custom-genesis-fast:
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.gaiad/config/genesis.json will be used."
	@go test -mod=readonly $(SIMAPP) -run TestFullGaiaSimulation -Genesis=${HOME}/.gaiad/config/genesis.json \
		-Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h

test-sim-import-export: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) 50 5 TestAppImportExport

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/greend
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/greencli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify


# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

###
# Find OS and Go environment
# GO contains the Go binary
# FS contains the OS file separator
###
ifeq ($(OS),Windows_NT)
  GO := $(shell where go.exe 2> NUL)
  FS := "\\"
else
  GO := $(shell command -v go 2> /dev/null)
  FS := "/"
endif

ifeq ($(GO),)
  $(error could not find go. Is it in PATH? $(GO))
endif

tools: runsim

GOPATH ?= $(shell $(GO) env GOPATH)

TOOLS_DESTDIR  ?= $(GOPATH)/bin
RUNSIM 			= $(TOOLS_DESTDIR)/runsim

runsim: $(RUNSIM)

$(RUNSIM):
	@echo "Installing runsim..."
	@(cd /tmp && go get github.com/cosmos/tools/cmd/runsim@v1.0.0)

tools-clean:
	rm -f $(RUNSIM)
	rm -f tools-stamp
