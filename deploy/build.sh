#!/usr/bin/env bash

echo Building Domio...
rm -rf /domio
mkdir /domio

cd ~/domioapi
export GOPATH=$PWD
echo $GOPATH

go build -o /domio/domio domio

cd /
rm -rf ~/domioapi
