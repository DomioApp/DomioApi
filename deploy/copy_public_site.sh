#!/usr/bin/env bash
set -e

echo Copying public site files...

rm -rf /usr/share/nginx/html/public
yes | mv -f ~/domioapi/public /usr/share/nginx/html

echo Public site files copied!