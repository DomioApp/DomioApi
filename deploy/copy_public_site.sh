#!/usr/bin/env bash
set -e

echo Copying public site files...

rm -rf /usr/share/nginx/html/public_site
yes | mv -f ~/domioapi/public_site /usr/share/nginx/html

echo Public site files copied!