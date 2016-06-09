#!/bin/sh

trap 'echo "Exit signal received"; kill $PID; exit 0' INT TERM;

CONF=/app/config/config.json

USER=${MYSQL_USER:-root}
PASS=${MYSQL_PASS:-}
DB=${MYSQL_DB:-gowebapp}
HOST=${MYSQL_HOST:-localhost}
PORT=${MYSQL_PORT:-3306}

for i in USER PASS DB HOST PORT; do
  eval sub=\$$i
  sed -i "s/MYSQL_$i/$sub/" $CONF
done

/app/gowebapp &
PID=$!
while true; do sleep 1; done
