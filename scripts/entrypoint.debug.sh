#!/bin/bash

#!/usr/bin/env bash
if [ -f .env ]; then
  export $(cat .env | grep -v '#' | sed 's/\r$//' | xargs)
else
  echo ".env file not found"
  exit 1
fi

echo ""
echo "🌳 Script Environment Variables 🌳"
echo "SCALAR_BIN_NAME: ${SCALAR_BIN_NAME}"
echo "SCALAR_BIN_PATH: ${SCALAR_BIN_PATH}"
echo "SCALAR_HOME_DIR: ${SCALAR_HOME_DIR}"
echo "SCALAR_CHAIN_ID: ${SCALAR_CHAIN_ID}"
echo "SCALAR_KEYRING_BACKEND: ${SCALAR_KEYRING_BACKEND}"
echo "SCALAR_CHAINS_DIR: ${SCALAR_CHAINS_DIR}"
echo "SCALAR_NUM_VALIDATORS: ${SCALAR_NUM_VALIDATORS}"
echo ""

set -e

trap stop_gracefully SIGTERM SIGINT

stop_gracefully() {
  echo "Stopping all processes..."
  pkill -f ${SCALAR_BIN_NAME}
  sleep 10
  echo "All processes stopped."
}

if [ ! -d "${SCALAR_HOME_DIR}" ]; then
  echo "Please set SCALAR_HOME_DIR in env. It is the runtime directory for the node."
  echo "Example: SCALAR_HOME_DIR=.scalar/node1/scalard"
  exit 1
fi

startNodeProc() {
  $SCALAR_BIN_PATH start --home $SCALAR_HOME_DIR --log_level=debug
}

if [ -z "$1" ]; then
  echo "--- ✅ STARTING NODE 🚀 ---"
  startNodeProc &
  wait
else
  $@ &
  wait
fi

# fileCount() {
#   find "$1" -maxdepth 1 ! -iname ".*" ! -iname "$(basename "$1")" | wc -l
# }

# addPeers() {
#   sed "s/^seeds =.*/seeds = \"$1\"/g" "$CONFIG_DIR/config.toml" >"$CONFIG_DIR/config.toml.tmp" &&
#     mv "$CONFIG_DIR/config.toml.tmp" "$CONFIG_DIR/config.toml"
# }

# if [ -n "$CONFIG_PATH" ] && [ -d "$CONFIG_PATH" ]; then
#   mkdir -p "$CONFIG_DIR"
#   echo "Copying config files from $CONFIG_PATH to $CONFIG_DIR"
#   if [ -f "$CONFIG_PATH/config.toml" ]; then
#     cp "$CONFIG_PATH/config.toml" "$CONFIG_DIR/config.toml"
#   fi
#   if [ -f "$CONFIG_PATH/app.toml" ]; then
#     cp "$CONFIG_PATH/app.toml" "$CONFIG_DIR/app.toml"
#   fi
#   if [ -f "$CONFIG_PATH/vald.toml" ]; then
#     cp "$CONFIG_PATH/vald.toml" "$CONFIG_DIR/vald.toml"
#   fi
#   if [ -f "$CONFIG_PATH/genesis.json" ]; then
#     cp "$CONFIG_PATH/genesis.json" "$CONFIG_DIR/genesis.json"
#   fi
#   if [ -f "$CONFIG_PATH/seeds.toml" ]; then
#     cp "$CONFIG_PATH/seeds.toml" "$CONFIG_DIR/config/seeds.toml"
#   fi
# fi

# if [ -n "$PEERS_FILE" ]; then
#   PEERS=$(cat "$PEERS_FILE")
#   addPeers "$PEERS"
# fi
