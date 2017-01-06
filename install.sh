#!/usr/bin/env bash

set -e

sh ~/domioapi/deploy/apt_update.sh
sh ~/domioapi/deploy/install_go.sh
#sh deploy/install_gb.sh
sh ~/domioapi/deploy/install_deps.sh
echo ======================================================
echo ~/
echo ======================================================
echo ------------------------------------------------------
echo ~/domioapi/
echo ------------------------------------------------------
sh ~/domioapi/deploy/buld.sh

sh ~/domioapi/deploy/install_pg.sh
