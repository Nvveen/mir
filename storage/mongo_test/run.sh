#!/bin/bash

start() {
  echo "starting test database"
  mkdir db
  cd db
  cp ../supervisord.conf .
  supervisord -c supervisord.conf || (echo "failed executing supervisord" && exit 1)
}

stop() {
  if [ -d db ]; then
    echo "shutting down test database"
    (cd db && supervisorctl shutdown &>/dev/null)
    rm -rf db
  fi
}

case "$1" in
  start)
    start;
    ;;
  stop)
    stop
    ;;
esac
