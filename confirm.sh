#! /bin/bash

SCALAR_HOME_DIR=.runtime/scalar/node1/scalard 
SCALAR_BIN_PATH=./bin/scalard
WALLET=broadcaster
SCALAR_KEYRING_BACKEND=test
SCALAR_CHAIN_ID=scalar-testnet-1
# make cfst WALLET=broadcaster ARGS="bitcoin-testnet4 18fa2be86b54d9ff7e35aba97d57483f05500cd9301547607f67ea5b47fa1c87"
evm() {
    ${SCALAR_BIN_PATH} tx chains confirm-source-txs 'evm|11155111' 0x983bff649adbdc2766948e280f5c7c11d07081f02234dd254d06b3c02f21d5fd \
    --from $WALLET \
    --keyring-backend $SCALAR_KEYRING_BACKEND \
    --home $SCALAR_HOME_DIR \
    --chain-id $SCALAR_CHAIN_ID \
    --gas 500000 \
    --gas-adjustment 1.0
}

btc() {
    ${SCALAR_BIN_PATH} tx chains confirm-source-txs 'bitcoin|4' 18fa2be86b54d9ff7e35aba97d57483f05500cd9301547607f67ea5b47fa1c87 \
    --from $WALLET \
    --keyring-backend $SCALAR_KEYRING_BACKEND \
    --home $SCALAR_HOME_DIR \
    --chain-id $SCALAR_CHAIN_ID \
    --gas 500000 \
    --gas-adjustment 1.0
}
$@