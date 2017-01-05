#!/usr/bin/env bash

echo Building Domio...
gb build domio
mkdir /domio
mv bin/domio /domio
cd /
rm -rf ~/domioapi
