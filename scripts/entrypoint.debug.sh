#!/bin/bash

APP_NAME=scalar
BIN_NAME=scalard

set -e

trap stop_gracefully SIGTERM SIGINT

stop_gracefully() {
  echo "Stopping all processes..."
  pkill -f ${BIN_NAME}
  sleep 10
  echo "All processes stopped."
}

HOME_DIR="$PWD"
CONFIG_DIR="$HOME_DIR/.${APP_NAME}"
CONFIG_PATH="$HOME_DIR/config"
EXEC_PATH="$HOME_DIR/bin/${BIN_NAME}"

fileCount() {
  find "$1" -maxdepth 1 ! -iname ".*" ! -iname "$(basename "$1")" | wc -l
}

addPeers() {
  sed "s/^seeds =.*/seeds = \"$1\"/g" "$CONFIG_DIR/config.toml" >"$CONFIG_DIR/config.toml.tmp" &&
    mv "$CONFIG_DIR/config.toml.tmp" "$CONFIG_DIR/config.toml"
}

startNodeProc() {
  $EXEC_PATH start --home $CONFIG_DIR
}

if [ -n "$CONFIG_PATH" ] && [ -d "$CONFIG_PATH" ]; then
  mkdir -p "$CONFIG_DIR"
  echo "Copying config files from $CONFIG_PATH to $CONFIG_DIR"
  if [ -f "$CONFIG_PATH/config.toml" ]; then
    cp "$CONFIG_PATH/config.toml" "$CONFIG_DIR/config.toml"
  fi
  if [ -f "$CONFIG_PATH/app.toml" ]; then
    cp "$CONFIG_PATH/app.toml" "$CONFIG_DIR/app.toml"
  fi
  if [ -f "$CONFIG_PATH/vald.toml" ]; then
    cp "$CONFIG_PATH/vald.toml" "$CONFIG_DIR/vald.toml"
  fi
  if [ -f "$CONFIG_PATH/genesis.json" ]; then
    cp "$CONFIG_PATH/genesis.json" "$CONFIG_DIR/genesis.json"
  fi
  if [ -f "$CONFIG_PATH/seeds.toml" ]; then
    cp "$CONFIG_PATH/seeds.toml" "$CONFIG_DIR/config/seeds.toml"
  fi
fi

if [ -n "$PEERS_FILE" ]; then
  PEERS=$(cat "$PEERS_FILE")
  addPeers "$PEERS"
fi

if [ -z "$1" ]; then
  echo "--- ✅ STARTING NODE 🚀 ---"
  startNodeProc &
  wait
else
  $@ &
  wait
fi
