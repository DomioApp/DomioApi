#!/usr/bin/env bash

echo Building Domio...
rm -rf /domio
mkdir /domio

cd ~/domioapi
export GOPATH=$PWD
echo $GOPATH

gb build domio
mv bin/domio /domio
cd /
rm -rf ~/domioapi
