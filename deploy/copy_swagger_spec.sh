#!/usr/bin/env bash
set -e

echo Copying Swagger Domio Spec...

yes | mv -f ~/domioapi/swagger_schema/domio_api.json /usr/local/domio_api

echo Swagger Domio Spec copied!