#!/usr/bin/env bash
set -e

echo Copying Swagger...

rm -rf /usr/share/nginx/html/swagger
yes | mv -f ~/domioapi/swagger /usr/share/nginx/html

echo Swagger copied!