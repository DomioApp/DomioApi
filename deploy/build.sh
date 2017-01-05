#!/usr/bin/env bash

echo Building Domio...
mkdir /domio

gb build domio
mv bin/domio /domio
cd /
rm -rf ~/domioapi
