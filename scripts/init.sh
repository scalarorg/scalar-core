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
echo "SCALAR_NODE1: ${SCALAR_NODE1}"

# Clean up existing data
rm -rf ${SCALAR_HOME_DIR}
pkill -f ${SCALAR_BIN_NAME}

if [ ! -f "${SCALAR_BIN_PATH}" ]; then
    echo "Please verify ${SCALAR_BIN_NAME} is installed"
    exit 1
fi

init() {
    ${SCALAR_BIN_PATH} testnet init-files --v 5 -o ${SCALAR_HOME_DIR} \
        --keyring-backend=${SCALAR_KEYRING_BACKEND} \
        --chain-id=${SCALAR_CHAIN_ID} \
        --node-dir-prefix node \
        --node-domain=${SCALAR_NODE_DOMAIN} \
        --supported-chains=${SCALAR_CHAINS_DIR} \
        --env-file=${SCALAR_ENV_FILE}
}

verify() {
    ${SCALAR_BIN_PATH} validate-genesis --home ${SCALAR_NODE1}
    if [ $? -ne 0 ]; then
        echo "❌ Genesis file is invalid"
        exit 1
    fi
    echo "✅ Successfully verified genesis file"
}

init
verify
