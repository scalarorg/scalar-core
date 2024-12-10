#!/usr/bin/env bash
if [ -f .env ]; then
    export $(cat .env | grep -v '#' | sed 's/\r$//' | xargs)
else
    echo ".env file not found"
    exit 1
fi

echo "SCALAR_BIN_NAME: ${SCALAR_BIN_NAME}"
echo "SCALAR_BIN_PATH: ${SCALAR_BIN_PATH}"
echo "SCALAR_HOME_DIR: ${SCALAR_HOME_DIR}"
echo "SCALAR_CHAIN_ID: ${SCALAR_CHAIN_ID}"
echo "SCALAR_KEYRING_BACKEND: ${SCALAR_KEYRING_BACKEND}"
echo "SCALAR_CHAINS_DIR: ${SCALAR_CHAINS_DIR}"
echo "SCALAR_NUM_VALIDATORS: ${SCALAR_NUM_VALIDATORS}"
# Clean up existing data
rm -rf ${SCALAR_HOME_DIR}
pkill -f ${SCALAR_BIN_NAME}

# Clean up existing data
rm -rf ${SCALAR_HOME_DIR}
pkill -f ${SCALAR_BIN_NAME}

if [ ! -f "$SCALAR_BIN_PATH" ]; then
    echo "Please verify ${SCALAR_BIN_NAME} is installed"
    exit 1
fi

# Set the denomination and amounts
DENOM=scal
STAKING_AMOUNT=1000000000000 # The amount to be staked (1 trillion)
ALICE_BALANCE=2000000000000  # Initial balance for alice to cover fees and staking
BOB_BALANCE=3000000000000    # Initial balance for bob

# Step 1: Initialize the Blockchain
$SCALAR_BIN_PATH init test --chain-id $SCALAR_CHAIN_ID --home $SCALAR_HOME_DIR

# Optional: Update the denomination in genesis.json if needed
jq '.app_state.staking.params.bond_denom="'${DENOM}'"' $SCALAR_HOME_DIR/config/genesis.json >$SCALAR_HOME_DIR/config/genesis.json.tmp && mv $SCALAR_HOME_DIR/config/genesis.json.tmp $SCALAR_HOME_DIR/config/genesis.json

# Step 2: Add Keys
$SCALAR_BIN_PATH keys add alice --keyring-backend test --output json --home $SCALAR_HOME_DIR
$SCALAR_BIN_PATH keys add bob --keyring-backend test --output json --home $SCALAR_HOME_DIR

# Step 3: Add Accounts to Genesis
$SCALAR_BIN_PATH add-genesis-account $($SCALAR_BIN_PATH keys show alice -a --keyring-backend test --home $SCALAR_HOME_DIR) $ALICE_BALANCE$DENOM --home $SCALAR_HOME_DIR
$SCALAR_BIN_PATH add-genesis-account $($SCALAR_BIN_PATH keys show bob -a --keyring-backend test --home $SCALAR_HOME_DIR) $BOB_BALANCE$DENOM --home $SCALAR_HOME_DIR

# Step 4: Create a Genesis Validator
$SCALAR_BIN_PATH gentx alice $STAKING_AMOUNT$DENOM --chain-id $SCALAR_CHAIN_ID --keyring-backend test --home $SCALAR_HOME_DIR

# Step 5: Collect Genesis Transactions
$SCALAR_BIN_PATH collect-gentxs --home $SCALAR_HOME_DIR

# Step 6: Validate the Genesis File
$SCALAR_BIN_PATH validate-genesis --home $SCALAR_HOME_DIR

# Step 7: Start the Blockchain
echo "✅ Blockchain initialized successfully on Cosmos SDK v0.46!"
