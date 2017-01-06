#!/usr/bin/env bash

echo Building Domio...
rm -rf /domio
mkdir /domio

cd ~/domioapi

go build -o /domio/domio domio

cd /
rm -rf ~/domioapi
