#!/bin/bash
set -e

# .envファイルを読み込み
if [ -f .env ]; then
  export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
  envsubst < .air.template.toml > .air.toml
fi
