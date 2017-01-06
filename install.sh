#!/usr/bin/env bash

cd ~/domioapi
export GOPATH=$PWD
echo $GOPATH

sh ~/domioapi/deploy/apt_update.sh

if ! [ -x "$(command -v go)" ]; then
   echo 'go is not installed.' >&2
   sh ~/domioapi/deploy/install_go.sh
  else
   echo "Go is already installed!" >&2
fi

sh ~/domioapi/deploy/install_deps.sh

if ! [ -x "$(command -v psql)" ]; then
   echo 'Postgres is not installed.' >&2
   sh ~/domioapi/deploy/install_pg.sh
  else
   echo "Postgres is already installed!" >&2
fi

sh ~/domioapi/deploy/build.sh
