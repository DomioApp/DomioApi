#!/usr/bin/env bash

set -e

sh deploy/apt_update.sh
sh deploy/install_go.sh
sh deploy/install_gb.sh
sh deploy/install_deps.sh
sh deploy/buld.sh

sh deploy/install_pg.sh
