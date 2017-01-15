#!/usr/bin/env bash
set -e

echo Copying Domio Api Service Config...

yes | cp -rf ~/domioapi/deploy/config/domio_service.sh /etc/init.d/domio_api

systemctl daemon-reload

echo Domio Api Service Config copied!