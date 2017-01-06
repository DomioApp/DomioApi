#!/usr/bin/env bash
set -e

echo Building Domio...
rm -rf /domio
mkdir /domio

cd ~/domioapi

go build -o /domio/domio domio

cd /
rm -rf ~/domioapi

echo Domio is built and ready!