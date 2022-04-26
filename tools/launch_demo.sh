#!/usr/bin/env bash

scriptdir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Random port numbers:
HTTP_PORT=`shuf -i 30000-40000 -n 1`
DB_PORT=`shuf -i 30000-40000 -n 1`
export TC_FRONTEND_ADDR=":$HTTP_PORT"
export TC_REDIS_ADDR=":$DB_PORT"
echo "TC_FRONTEND_ADDR: $TC_FRONTEND_ADDR"
echo "TC_REDIS_ADDR: $TC_REDIS_ADDR"

echo "RUNNING DATABASE IN BACKGROUND..."
docker run \
    --rm \
    -p $DB_PORT:6379 \
    -d \
    eqalpha/keydb keydb-server /etc/keydb/keydb.conf --appendonly yes


echo "BUILDING..."
bazel build //src/cmd/web_frontend

echo "RUNNING FRONTEND IN BACKGROUND..."
binpath=$scriptdir/../dist/bin/src/cmd/web_frontend/web_frontend_/web_frontend
# $binpath
$binpath &

echo "RUNNING ngrok PROXY..."
sleep 0.3
ngrok http $HTTP_PORT
