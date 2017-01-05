#!/usr/bin/env bash

set -e

export GOPATH=$PWD
echo $GOPATH

sh deploy/apt_update.sh
sh deploy/install_go.sh
sh deploy/install_gb.sh
sh deploy/install_deps.sh
./deploy/buld.sh

sh deploy/install_pg.sh
