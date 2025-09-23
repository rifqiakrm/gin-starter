#!/bin/sh

set -euo pipefail

echo "wait for pgsql @" $POSTGRES_HOST:$POSTGRES_PORT
./wait-for -t 15 $POSTGRES_HOST:$POSTGRES_PORT
ls
./gin-api