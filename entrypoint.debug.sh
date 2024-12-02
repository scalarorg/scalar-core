#!/bin/bash
set -e

trap stop_gracefully SIGTERM SIGINT

stop_gracefully() {
  echo "stopping all processes"
  pkill "scalard"
  sleep 10
  echo "all processes stopped"
}

HOME_DIR="$PWD"
CONFIG_PATH="$HOME_DIR/config"
D_HOME_DIR="$HOME_DIR/.scalar"
HOME=${HOME:-$HOME_DIR}

echo "--- ℹ️ Environment variables set ---  ℹ️"
echo "HOME_DIR=$HOME_DIR"
echo "CONFIG_PATH=$CONFIG_PATH"
echo "D_HOME_DIR=$D_HOME_DIR"
echo "HOME=$HOME"

fileCount() {
  find "$1" -maxdepth 1 ! -iname ".*" ! -iname "$(basename "$1")" | wc -l
}

addPeers() {
  sed "s/^seeds =.*/seeds = \"$1\"/g" "$D_HOME_DIR/config/config.toml" >"$D_HOME_DIR/config/config.toml.tmp" &&
    mv "$D_HOME_DIR/config/config.toml.tmp" "$D_HOME_DIR/config/config.toml"
}

cont() {
  if [ "$1" = true ]; then
    "--continue"
  else
    ""
  fi
}

startValdProc() {
  DURATION=${SLEEP_TIME:-"10s"}
  sleep $DURATION

  if [ "$VALD_CONTINUE" != true ]; then
    unset VALD_CONTINUE
  fi

  dlv --listen=:2346 --headless=true ${VALD_CONTINUE:+--continue} --api-version=2 --accept-multiclient exec \
    /usr/local/bin/scalard -- vald-start ${TOFND_HOST:+--tofnd-host "$TOFND_HOST"} ${VALIDATOR_HOST:+--node "$VALIDATOR_HOST"} --validator-addr "${VALIDATOR_ADDR:-$(scalard keys show validator -a --bech val)}"
}

startNodeProc() {
  # if [ "$CORE_CONTINUE" != true ]; then
  #   unset CORE_CONTINUE
  # fi

  # dlv --listen=:2345 --headless=true ${CORE_CONTINUE:+--continue} --api-version=2 --accept-multiclient exec \
  #   /usr/local/bin/scalard -- start
  ./bin/scalard start
}

if [ -n "$PRESTART_SCRIPT" ] && [ -f "$PRESTART_SCRIPT" ]; then
  echo "Running pre-start script at $PRESTART_SCRIPT"
  source "$PRESTART_SCRIPT"
fi

if [ -n "$CONFIG_PATH" ] && [ -d "$CONFIG_PATH" ]; then
  mkdir -p "$D_HOME_DIR/config"
  echo "Copying config files from $CONFIG_PATH to $D_HOME_DIR/config"
  if [ -f "$CONFIG_PATH/config.toml" ]; then
    cp "$CONFIG_PATH/config.toml" "$D_HOME_DIR/config/config.toml"
  fi
  if [ -f "$CONFIG_PATH/app.toml" ]; then
    cp "$CONFIG_PATH/app.toml" "$D_HOME_DIR/config/app.toml"
  fi
  if [ -f "$CONFIG_PATH/vald.toml" ]; then
    cp "$CONFIG_PATH/vald.toml" "$D_HOME_DIR/config/vald.toml"
  fi

  if [ -f "$CONFIG_PATH/genesis.json" ]; then
    cp "$CONFIG_PATH/genesis.json" "$D_HOME_DIR/config/genesis.json"
  fi

  if [ -f "$CONFIG_PATH/seeds.toml" ]; then
    cp "$CONFIG_PATH/seeds.toml" "$D_HOME_DIR/config/seeds.toml"
  fi
fi

if [ -n "$PEERS_FILE" ]; then
  PEERS=$(cat "$PEERS_FILE")
  addPeers "$PEERS"
fi

if [ "$REST_CONTINUE" != true ]; then
  unset REST_CONTINUE
fi

if [ -z "$1" ]; then

  # startValdProc &

  # echo "--- ✅ STARTING VALD --- 🚀 🚀 🚀"
  # startValdProc &

  echo "--- ✅ STARTING NODE 🚀 🚀 🚀 ---"
  startNodeProc &

  wait

else

  $@ &

  wait

fi
