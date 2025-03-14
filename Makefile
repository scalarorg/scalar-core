# Include .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif


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

SCALAR_BIN_PATH ?= bin/scalard
SCALAR_BIN_NAME ?= scalard
SCALAR_HOME_DIR ?= .scalar/scalar
SCALAR_CHAIN_ID ?= scalar-testnet-1
SCALAR_KEYRING_BACKEND ?= test
LOCAL_LIB_PATH ?= $(shell pwd)/lib
ifndef USER_ID
USER_ID ?= $(shell id -u)
endif
ifndef GROUP_ID
GROUP_ID ?= $(shell id -g)
endif
export CGO_LDFLAGS := ${CGO_LDFLAGS} -lbitcoin_vault_ffi  -L${LOCAL_LIB_PATH}

$(info ⛳️ Makefile Environment Variables ⛳️)
$(info $$USER_ID is [${USER_ID}])
$(info $$GROUP_ID is [${GROUP_ID}])
$(info $$WASM is [${WASM}])
$(info $$IBC_WASM_HOOKS is [${IBC_WASM_HOOKS}])
$(info $$MAX_WASM_SIZE is [${MAX_WASM_SIZE}])
$(info $$CGO_ENABLED is [${CGO_ENABLED}])
$(info $$CGO_LDFLAGS is [${CGO_LDFLAGS}])
$(info $$LOCAL_LIB_PATH is [${LOCAL_LIB_PATH}])

$(info $$SCALAR_BIN_NAME is [${SCALAR_BIN_NAME}])
$(info $$SCALAR_BIN_PATH is [${SCALAR_BIN_PATH}])
$(info $$SCALAR_HOME_DIR is [${SCALAR_HOME_DIR}])
$(info $$SCALAR_CHAIN_ID is [${SCALAR_CHAIN_ID}])
$(info $$SCALAR_KEYRING_BACKEND is [${SCALAR_KEYRING_BACKEND}])

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
IMAGE_TAG := ${COMMIT}
else
IMAGE_TAG := ${VERSION}
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

BUILD_FLAGS := -tags $(BUILD_TAGS) -ldflags $(ldflags) -trimpath -buildvcs=false


# Build the project with release flags
.PHONY: build
build: go.sum
	@go build -o ./bin/scalard -mod=readonly $(BUILD_FLAGS) ./cmd/scalard

# Build the project with release flags in a docker container
.PHONY: docker-build
docker-build: go.sum
	@go build -o ./bin/docker/scalard -mod=readonly $(BUILD_FLAGS) ./cmd/scalard

.PHONY: run
run:
	@HOME=$(PWD) ./entrypoint.sh

.PHONY: start
start: build
	@make run

.PHONY: dev-init
dev-init:
	@./scripts/dev-init.sh

.PHONY: init
init:
	echo "🚒 deprecated"

.PHONY: dev
# Usage: make dev SCALAR_HOME_DIR=.scalar/node1/scalard
dev:
	@if [ -z "$(N)" ]; then \
		SCALAR_HOME_DIR=${SCALAR_HOME_DIR} ./scripts/entrypoint.debug.sh; \
	else \
		echo "Running node ${N}"; \
		export SCALAR_HOME_DIR=./.scalar/node${N}/scalard; \
		./scripts/entrypoint.debug.sh; \
	fi

.PHONY: dbg
dbg: build
	make dev


# Build a release image
.PHONY: docker-image-test
docker-image-test:
	@docker build \
		--build-arg WASM="${WASM}" \
		--build-arg WASMVM_VERSION="${WASMVM_VERSION}" \
		--build-arg IBC_WASM_HOOKS="${IBC_WASM_HOOKS}" \
		--build-arg ARCH="${ARCH}" \
		--build-arg USER_ID="${USER_ID}" \
		--build-arg GROUP_ID="${GROUP_ID}" \
		-t scalarorg/scalar-core:${IMAGE_TAG} .

# Build a release image
.PHONY: docker-image
docker-image:
	@docker build \
		--build-arg WASM="${WASM}" \
		--build-arg WASMVM_VERSION="${WASMVM_VERSION}" \
		--build-arg IBC_WASM_HOOKS="${IBC_WASM_HOOKS}" \
		--build-arg ARCH="${ARCH}" \
		--build-arg USER_ID="${USER_ID}" \
		--build-arg GROUP_ID="${GROUP_ID}" \
		-t scalarorg/scalar-core .

docker-run:
	@DOCKER_BUILDKIT=1 docker run -it scalarorg/scalar-core 

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
	LOOKING_DIR="x utils"; \
	find $$LOOKING_DIR -type f \( -name "*.pb.go" -o -name "*.pb.gw.go" \) -delete

.PHONY: proto-all proto-gen proto-gen-any proto-format proto-lint proto-check-breaking proto-update-deps proto-clean


.PHONY: generate
generate: prereqs docs generate-mocks

.PHONY: generate-mocks
generate-mocks:
	go generate -x ./...

.PHONY: docs
docs:
	@echo "Removing old clidocs"

	@if find docs/cli -name "*.md"  | grep -q .; then \
		rm docs/cli/*.md; \
	fi

	@echo "Generating new cli docs"
	@go run  $(BUILD_FLAGS) cmd/scalard/main.go --docs docs/cli
	@# ensure docs are canonically formatted
	@mdformat docs/cli/*


# Install all generate prerequisites
.Phony: prereqs
prereqs:
	@which mdformat &>/dev/null || ( \
		echo "Installing mdformat in a virtual environment..." && \
		python3 -m venv .venv && \
		. .venv/bin/activate && \
		pip3 install mdformat && \
		sudo ln -sf $(PWD)/.venv/bin/mdformat /usr/local/bin/mdformat )
	@which protoc &>/dev/null || echo "Please install protoc for grpc (https://grpc.io/docs/languages/go/quickstart/)"
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/tools/cmd/stringer@latest
	go install github.com/matryer/moq@latest
	go install github.com/rakyll/statik@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0


###############################
######## 	Commands	#######
##############################


######
# Usage: SCALAR_HOME_DIR=.scalar/node1/scalard make cfst WALLET=node1 ARGS="bitcoin-testnet4 07b50c84f889e2f1315da875fc91734e2bac8d0153ff9a98d9da14caa4fc7d57"
######
.PHONY: cfs cfs2 cfs3 cfd2 cfd3 cfd4
cfs:
	@if [ -z "$(ARGS)" ]; then \
		echo "ARGS is required"; \
		exit 1; \
	fi

	$(SCALAR_BIN_PATH) tx chains confirm-source-txs $(ARGS) --from broadcaster --keyring-backend $(SCALAR_KEYRING_BACKEND) --home .scalar/scalar/node1/scalard --chain-id $(SCALAR_CHAIN_ID) --gas 400000

cfs2:
	$(SCALAR_BIN_PATH) tx chains confirm-source-txs "bitcoin|4" 35326772ac35ccf5c0c46087238073af9c11cbe7557448101f18af7826342f67 --from broadcaster --keyring-backend $(SCALAR_KEYRING_BACKEND) --home .scalar/scalar/node1/scalard --chain-id $(SCALAR_CHAIN_ID) --gas 500000
cftf:
	$(SCALAR_BIN_PATH) tx chains create-pending-transfers "evm|11155111" --from broadcaster --keyring-backend $(SCALAR_KEYRING_BACKEND) --home .scalar/scalar/node1/scalard --chain-id $(SCALAR_CHAIN_ID) --gas 203590

cfs3:
	$(SCALAR_BIN_PATH) tx chains confirm-source-txs "bitcoin|4" 5188eea7ceb9f585f5ba8a2abebb82f9850dd671b6e2926263674af6882fd3f6 fc702634fc245b254d44585dfaa6527871899dbfe38d5585abecbb738f8865bf --from broadcaster --keyring-backend $(SCALAR_KEYRING_BACKEND) --home .scalar/scalar/node1/scalard --chain-id $(SCALAR_CHAIN_ID) --gas 537803

cfd2:
	$(SCALAR_BIN_PATH) tx chains confirm-source-txs "evm|11155111" 8062e7c01545e6bac064e675b5bec41d9536dbc3b849368385279e5b2c6e96fd --from broadcaster --keyring-backend $(SCALAR_KEYRING_BACKEND) --home .scalar/scalar/node1/scalard --chain-id $(SCALAR_CHAIN_ID) --gas 507368

.PHONY: cfd2b
cfd2b:
	$(SCALAR_BIN_PATH) tx chains confirm-source-txs "evm|11155111" 71b4e2cae2e93b45586469febcc56cabc554b964e96ff788b2887d9070c8b7a7 --from broadcaster --keyring-backend $(SCALAR_KEYRING_BACKEND) --home .scalar/scalar/node1/scalard --chain-id $(SCALAR_CHAIN_ID) --gas 507368

	
cfd3:
	$(SCALAR_BIN_PATH) tx chains confirm-source-txs "evm|11155111" 691b0cc5eb8abcec4b56e3aea8044ddda8a4152ab00e87dfaecf3123bd7290a0 --from broadcaster --keyring-backend $(SCALAR_KEYRING_BACKEND) --home .scalar/scalar/node1/scalard --chain-id $(SCALAR_CHAIN_ID) --gas 600000

cfd4:
	$(SCALAR_BIN_PATH) tx chains confirm-source-txs "evm|11155111" 4a21de48f14ae787a11cce77f6232fe52308590791829b54ecc2f82f36a2468f --from broadcaster --keyring-backend $(SCALAR_KEYRING_BACKEND) --home .scalar/scalar/node1/scalard --chain-id $(SCALAR_CHAIN_ID) --gas 400000

.PHONY: sign-btc
sign-btc:
	$(SCALAR_BIN_PATH) tx chains sign-btc-commands "bitcoin|4" --from broadcaster --keyring-backend $(SCALAR_KEYRING_BACKEND) --home .scalar/scalar/node1/scalard --chain-id $(SCALAR_CHAIN_ID) --gas 1000535

.PHONY: batch-command
batch-command:
	$(SCALAR_BIN_PATH) query chains batched-commands "bitcoin|4" $(BATCH_ID)

.PHONY: params
params:
	@if [ -z "$(CHAIN)" ]; then \
		$(SCALAR_BIN_PATH) query chains params "bitcoin|4"; \
	else \
		$(SCALAR_BIN_PATH) query chains params '$(CHAIN)'; \
	fi

.PHONY: add-protocol
add-protocol:
	$(SCALAR_BIN_PATH) tx protocol add '{"attribute":{"model":0},"avatar":"nVJ5","bitcoin_pubkey":"03620a0b56223990b75c28dded4f30c854e0918e52179d5a7dfbf79df89fce7bcd","custodian_group_uid":"2689d0b8-6a61-45fa-a39f-53b8f9057aad","name":"lalaa","tag":"pepeS","asset":{"chain_name":"bitcoin|4","asset_name":"Zkkk"}}' --from broadcaster --keyring-backend $(SCALAR_KEYRING_BACKEND) --home .scalar/scalar/node1/scalard --chain-id $(SCALAR_CHAIN_ID) --gas 1000535

.PHONY: open-docs
open-docs:
	open client/docs/static/openapi/index.html

.PHONY: mnemonic
mnemonic:
	$(eval user := $(filter-out $@,$(MAKECMDGOALS)))
	$(BIN_PATH) keys export $(user) --keyring-backend $(SCALAR_KEYRING_BACKEND) --unsafe --unarmored-hex --home $(SCALAR_DIR)


.PHONY: lib lib-clean
lib:
	mkdir -p ./lib
	cp ../bitcoin-vault/target/release/libbitcoin_vault_ffi.* ./lib

lib-clean:
	rm -rf ./lib/*