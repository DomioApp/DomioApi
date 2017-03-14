#!/usr/bin/env bash
set -e

echo Copying Swagger Domio Spec...

yes | mv -f ~/domioapi/swagger_schema/domio_api.json /usr/share/nginx/html

echo Swagger Domio Spec copied!