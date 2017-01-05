#!/usr/bin/env bash

echo Building Domio...
rm -rf /domio
mkdir /domio

cd ~/domioapi
gb build domio
mv bin/domio /domio
cd /
rm -rf ~/domioapi
