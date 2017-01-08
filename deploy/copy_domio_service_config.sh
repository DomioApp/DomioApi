#!/usr/bin/env bash
set -e

echo Copying Domio Service Config...

yes | cp -rf ~/domioapi/deploy/config/domio_service.sh /etc/init.d/domio

systemctl daemon-reload

echo Domio Service Config copied!