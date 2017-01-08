#!/usr/bin/env bash
set -e

echo Copying nginx config files...

yes | cp -rf ~/domioapi/deploy/config/confg.nginx /etc/nginx/nginx.conf
nginx -s reload

echo nginx config files copied!