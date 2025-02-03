FROM golang:1.23.3-alpine3.20 as build

ARG ARCH=x86_64
ARG WASM=true
ARG IBC_WASM_HOOKS=false
ARG WASMVM_VERSION=v1.3.1
RUN apk add --no-cache --update \
  ca-certificates \
  git \
  make \
  build-base \
  linux-headers

COPY --from=scalar/bitcoin-vault /bitcoin-vault/target/release/libbitcoin_vault_ffi.* /usr/lib/

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

COPY . .

RUN make LOCAL_LIB_PATH="/usr/lib"  MUSLC="${WASM}" WASM="${WASM}" IBC_WASM_HOOKS="${IBC_WASM_HOOKS}" build

FROM alpine:3.20

# Install libgcc and libstdc++ for bitcoin-vault ffi
RUN apk add --no-cache \
    libgcc \
    libstdc++

ARG USER_ID=1000
ARG GROUP_ID=1001
RUN apk add jq bash

COPY --from=build /go/scalar/bin/* /usr/local/bin/
COPY --from=build /usr/lib/libbitcoin_vault_ffi.* /usr/lib/

RUN addgroup -S -g ${GROUP_ID} scalard && adduser -S -u ${USER_ID} scalard -G scalard
USER scalard
COPY ./entrypoint.sh /entrypoint.sh

# The home directory of scalar-core where configuration/genesis/data are stored
ENV HOME_DIR /home/scalard
# Host name for tss daemon (only necessary for validator nodes)
ENV TOFND_HOST ""
# The keyring backend type https://docs.cosmos.network/master/run-node/keyring.html
ENV SCALARD_KEYRING_BACKEND file
# The chain ID
ENV SCALARD_CHAIN_ID scalar-testnet-1
# The file with the peer list to connect to the network
ENV PEERS_FILE ""
# Path of an existing configuration file to use (optional)
ENV CONFIG_PATH ""
# A script that runs before launching the container's process (optional)
ENV PRESTART_SCRIPT ""
# The Scalar node's moniker
ENV NODE_MONIKER ""

# Create these folders so that when they are mounted the permissions flow down
RUN mkdir /home/scalard/.scalar && chown scalard /home/scalard/.scalar
RUN mkdir /home/scalard/shared && chown scalard /home/scalard/shared
RUN mkdir /home/scalard/genesis && chown scalard /home/scalard/genesis
RUN mkdir /home/scalard/scripts && chown scalard /home/scalard/scripts
RUN mkdir /home/scalard/conf && chown scalard /home/scalard/conf

ENTRYPOINT ["/entrypoint.sh"]
