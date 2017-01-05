#!/usr/bin/env bash

echo Building Domio...
rm -rf /domio
gb build domio
mkdir /domio
mv bin/domio /domio
cd /
rm -rf ~/domioapi
