#!/usr/bin/env bash
set -e

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

ENV_FILE="$DIR/.env"

if [ ! -f "$ENV_FILE" ]; then
  echo "Error: env file '$ENV_FILE' not found"
  exit 1
fi

source "$ENV_FILE"

# Now, run docker-compose from the script's directory
cd "$DIR"
docker-compose up -d
