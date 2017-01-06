#!/usr/bin/env bash

sh ~/domioapi/deploy/apt_update.sh

if ! [ -x "$(command -v go)" ]; then
   echo 'go is not installed.' >&2
   sh ~/domioapi/deploy/install_go.sh
  else
   echo "Go is already installed!" >&2
fi


#sh ~/domioapi/deploy/install_deps.sh

echo ======================================================
echo ~/
echo ======================================================

echo ------------------------------------------------------
echo ~/domioapi/
ls ~/domioapi/
cat ~/domioapi/deploy/buld.sh
echo ------------------------------------------------------
sh ~/domioapi/deploy/buld.sh

sh ~/domioapi/deploy/install_pg.sh
