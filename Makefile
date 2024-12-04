VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
HTTPS_GIT := https://github.com/scalarorg/scalar-core.git

WASM := true
WASMVM_VERSION := v2.1.3

MAX_WASM_SIZE := $(shell echo "$$((3 * 1024 * 1024))")
IBC_WASM_HOOKS := false
# Export env var to go build so Cosmos SDK can see it
export CGO_ENABLED := 1

$(info $$WASM is [${WASM}])
$(info $$IBC_WASM_HOOKS is [${IBC_WASM_HOOKS}])
$(info $$MAX_WASM_SIZE is [${MAX_WASM_SIZE}])
$(info $$CGO_ENABLED is [${CGO_ENABLED}])

ifndef $(WASM_CAPABILITIES)
# Wasm capabilities: https://github.com/CosmWasm/cosmwasm/blob/main/docs/CAPABILITIES-BUILT-IN.md
WASM_CAPABILITIES := "iterator,staking,stargate,cosmwasm_1_1,cosmwasm_1_2,cosmwasm_1_3"
else
WASM_CAPABILITIES := ""
endif

ifeq ($(MUSLC), true)
STATIC_LINK_FLAGS := -linkmode=external -extldflags '-Wl,-z,muldefs -static'
BUILD_TAGS := ledger,muslc
else
STATIC_LINK_FLAGS := ""
BUILD_TAGS := ledger
endif

ARCH := x86_64
ifeq ($(shell uname -m), arm64)
ARCH := aarch64
endif

ifndef $(VERSION)
VERSION := v0.0.1
endif

DENOM := scal

GO_MOD_PATH := github.com/scalarorg/scalar-core

ldflags = "-X github.com/cosmos/cosmos-sdk/version.Name=scalar \
	-X github.com/cosmos/cosmos-sdk/version.AppName=scalard \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(BUILD_TAGS)" \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X $(GO_MOD_PATH)/x/scalarnet/exported.NativeAsset=$(DENOM) \
	-X $(GO_MOD_PATH)/app.WasmEnabled=$(WASM) \
	-X $(GO_MOD_PATH)/app.IBCWasmHooksEnabled=$(IBC_WASM_HOOKS) \
	-X $(GO_MOD_PATH)/app.WasmCapabilities=$(WASM_CAPABILITIES) \
	-X $(GO_MOD_PATH)/app.MaxWasmSize=${MAX_WASM_SIZE} \
	-w -s ${STATIC_LINK_FLAGS}"

BUILD_FLAGS := -tags $(BUILD_TAGS) -ldflags $(ldflags) -trimpath

BIN_NAME=scalard
BIN_PATH=./bin/${BIN_NAME}

# Build the project with release flags
.PHONY: build
build: go.sum
	@go build -o ${BIN_PATH} -mod=readonly $(BUILD_FLAGS) ./cmd/${BIN_NAME}

.PHONY: run
run:
	@HOME=$(PWD) ./entrypoint.sh

.PHONY: start
start: build
	@make run

.PHONY: dev-init
dev-init:
	./scripts/mock-init.sh

.PHONY: dev
dev:
	@HOME_DIR=$(PWD) ./scripts/entrypoint.debug.sh

.PHONY: dbg
dbg: build
	make dev


# Build a release image
.PHONY: docker-image
docker-image:
	@DOCKER_BUILDKIT=1 docker build \
		--build-arg WASM="${WASM}" \
		--build-arg WASMVM_VERSION="${WASMVM_VERSION}" \
		--build-arg IBC_WASM_HOOKS="${IBC_WASM_HOOKS}" \
		--build-arg ARCH="${ARCH}" \
		-t scalarorg/scalar-core .

docker-run:
	@docker run -it scalarorg/scalar-core 

proto-all: proto-update-deps proto-format proto-lint proto-gen

PROTO_GEN_IMAGE := scalar/proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@if ! docker images $(PROTO_GEN_IMAGE) | grep -q $(PROTO_GEN_IMAGE); then \
		DOCKER_BUILDKIT=1 docker build -t $(PROTO_GEN_IMAGE) -f ./Dockerfile.protocgen .; \
	fi

	@$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(PROTO_GEN_IMAGE) sh ./scripts/protocgen.sh
	@echo "Generating Protobuf Swagger endpoint"
	@$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(PROTO_GEN_IMAGE) sh ./scripts/protoc-swagger-gen.sh
	@statik -src=./client/docs/static -dest=./client/docs -f -m

proto-format:
	@echo "Formatting Protobuf files"
	@$(DOCKER) run --rm -v $(CURDIR):/workspace \
	--workdir /workspace tendermintdev/docker-build-proto \
	$( find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; )

proto-lint:
	@echo "Linting Protobuf files"
	@$(DOCKER_BUF) lint

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=main

proto-update-deps:
	@echo "Updating Protobuf deps"
	@$(DOCKER_BUF) mod update

proto-clean:
	find x -type f \( -name "*.pb.go" -o -name "*.pb.gw.go" \) -delete

.PHONY: proto-all proto-gen proto-gen-any proto-format proto-lint proto-check-breaking proto-update-deps proto-clean


# Install all generate prerequisites
.Phony: prereqs
prereqs:
	@which mdformat &>/dev/null || ( \
		echo "Installing mdformat in a virtual environment..." && \
		python3 -m venv .venv && \
		. .venv/bin/activate && \
		pip install mdformat )
	@which protoc &>/dev/null || echo "Please install protoc for grpc (https://grpc.io/docs/languages/go/quickstart/)"
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/matryer/moq@latest
	go install github.com/rakyll/statik@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0


###############################
######## 	Commands	#######
###############################

KEYRING_BACKEND := test
HOME := ./.scalar
CHAIN_ID := demo
FROM := alice

cfgwtx:
	@if [ -z "$(ARGS)" ]; then \
		echo "ARGS is required"; \
		exit 1; \
	fi
	@$(BIN_PATH) tx btc confirm-gateway-txs $(ARGS) --from $(FROM) --keyring-backend $(KEYRING_BACKEND) --home $(HOME) --chain-id $(CHAIN_ID)


