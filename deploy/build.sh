#!/usr/bin/env bash
set -e

echo Building Domio...
rm -rf /domio
mkdir /domio

cd ~/domioapi

go build -o /domio/domio domio
#cp ~/domioapi/utils/run.sh /domio/run.sh

rm -rf /usr/sbin/domio
ln -s /domio/domio /usr/sbin/domio

cd /
rm -rf ~/domioapi

echo Domio is built and ready!