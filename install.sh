#!/usr/bin/env bash

set -e

sh deploy/apt_update.sh
sh deploy/install_go.sh
#sh deploy/install_gb.sh
sh deploy/install_deps.sh
#sh deploy/buld.sh


echo Building Domio...
rm -rf /domio
mkdir /domio

cd ~/domioapi
export GOPATH=$PWD
echo $GOPATH

go build -o /domio/domio domio

cd /
rm -rf ~/domioapi


sh deploy/install_pg.sh
