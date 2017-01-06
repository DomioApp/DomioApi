#!/usr/bin/env bash

set -e

sh ~/domioapi/deploy/apt_update.sh

if ! [ -x "$(command -v go)" ]; then
   echo 'go is not installed.' >&2
  else
   sh ~/domioapi/deploy/install_go.sh
fi


sh ~/domioapi/deploy/install_deps.sh

echo ======================================================
echo ~/
echo ======================================================

echo ------------------------------------------------------
echo ~/domioapi/
echo ------------------------------------------------------
sh ~/domioapi/deploy/buld.sh

sh ~/domioapi/deploy/install_pg.sh
