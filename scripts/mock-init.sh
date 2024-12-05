#!/usr/bin/env bash

APP_NAME=scalar
BIN_NAME=scalard

# Clean up existing data
rm -rf ./.${APP_NAME}
pkill -f ${BIN_NAME}

BIN=./bin/${BIN_NAME}

if [ ! -f "$BIN" ]; then
    echo "Please verify ${BIN_NAME} is installed"
    exit 1
fi

CONFIG_DIR="./.${APP_NAME}"
CHAIN_ID=scalar-testnet

# Set the denomination and amounts
DENOM=scal
STAKING_AMOUNT=1000000000000 # The amount to be staked (1 trillion)
ALICE_BALANCE=2000000000000  # Initial balance for alice to cover fees and staking
BOB_BALANCE=3000000000000    # Initial balance for bob

# Step 1: Initialize the Blockchain
$BIN init test --chain-id $CHAIN_ID --home $CONFIG_DIR

# Optional: Update the denomination in genesis.json if needed
jq '.app_state.staking.params.bond_denom="'${DENOM}'"' $CONFIG_DIR/config/genesis.json >$CONFIG_DIR/config/genesis.json.tmp && mv $CONFIG_DIR/config/genesis.json.tmp $CONFIG_DIR/config/genesis.json

# Step 2: Add Keys
$BIN keys add alice --keyring-backend test --output json --home $CONFIG_DIR
$BIN keys add bob --keyring-backend test --output json --home $CONFIG_DIR

# Step 3: Add Accounts to Genesis
$BIN add-genesis-account $($BIN keys show alice -a --keyring-backend test --home $CONFIG_DIR) $ALICE_BALANCE$DENOM --home $CONFIG_DIR
$BIN add-genesis-account $($BIN keys show bob -a --keyring-backend test --home $CONFIG_DIR) $BOB_BALANCE$DENOM --home $CONFIG_DIR

# Step 4: Create a Genesis Validator
$BIN gentx alice $STAKING_AMOUNT$DENOM --chain-id $CHAIN_ID --keyring-backend test --home $CONFIG_DIR

# Step 5: Collect Genesis Transactions
$BIN collect-gentxs --home $CONFIG_DIR

# Step 6: Validate the Genesis File
$BIN validate-genesis --home $CONFIG_DIR

# Step 7: Start the Blockchain
echo "✅ Blockchain initialized successfully on Cosmos SDK v0.46!"
