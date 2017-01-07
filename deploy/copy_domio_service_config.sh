#!/usr/bin/env bash
set -e

echo Copying Domio Service Config...
rm /etc/init.d/domio_service

cp ~/domioapi/deploy/config/domio_service.sh /etc/init.d/domio

echo Swagger copied!