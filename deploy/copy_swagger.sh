#!/usr/bin/env bash
set -e

echo Copying Swagger...
rm -rf ~/swagger
yes | mv -f ~/domioapi/swagger ~/

echo Swagger copied!