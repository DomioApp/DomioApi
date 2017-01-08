#!/usr/bin/env bash
set -e

echo Building Domio...
rm -rf /domio
mkdir /domio

cd ~/domioapi

go build -o /domio/domio domio

sh ~/domioapi/deploy/create_config_file.sh

cd /
rm -rf ~/domioapi

echo Domio is built and ready!

logger -n logs5.papertrailapp.com -t deploy -P 18422 -p user.notice "Domio is built and ready!"