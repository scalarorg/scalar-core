# syntax=docker/dockerfile:experimental

FROM rust:1.82-alpine3.20 as libbuilder
RUN apk add --no-cache git libc-dev
# Build bitcoin-vault lib
# Todo: select a specific feature, eg ffi
RUN git clone https://github.com/scalarorg/bitcoin-vault.git
WORKDIR /bitcoin-vault
RUN cargo build --release

# Buil scalar-core

FROM golang:1.23.3-alpine3.20

ARG ARCH=x86_64
ARG WASM=true
ARG IBC_WASM_HOOKS=false
ARG WASMVM_VERSION=v2.1.3
ARG USER_ID=1000
ARG GROUP_ID=1001
RUN apk add --no-cache --update \
  ca-certificates \
  git \
  make \
  build-base \
  linux-headers

# Copy the bitcoin-vault lib
COPY --from=libbuilder /bitcoin-vault/target/release/libbitcoin_vault_ffi.* /usr/lib/

WORKDIR scalar

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# Use a compatible libwasmvm
# Alpine Linux requires static linking against muslc: https://github.com/CosmWasm/wasmd/blob/v0.33.0/INTEGRATION.md#prerequisites
RUN if [[ "${WASM}" == "true" ]]; then \
  wget https://github.com/CosmWasm/wasmvm/releases/download/${WASMVM_VERSION}/libwasmvm_muslc.${ARCH}.a \
  -O /lib/libwasmvm_muslc.a && \
  wget https://github.com/CosmWasm/wasmvm/releases/download/${WASMVM_VERSION}/checksums.txt -O /tmp/checksums.txt && \
  sha256sum /lib/libwasmvm_muslc.a | grep $(cat /tmp/checksums.txt | grep libwasmvm_muslc.${ARCH}.a | cut -d ' ' -f 1); \
  fi
RUN addgroup -S -g ${GROUP_ID} scalard && adduser -S -u ${USER_ID} scalard -G scalard
USER scalard
WORKDIR /home/scalard
ENTRYPOINT ["sleep", "infinity"]

