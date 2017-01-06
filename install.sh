#!/usr/bin/env bash

#sh ~/domioapi/deploy/apt_update.sh

if ! [ -x "$(command -v go)" ]; then
   echo 'go is not installed.' >&2
   sh ~/domioapi/deploy/install_go.sh
  else
   echo "Go is already installed!" >&2
fi

cat ~/domioapi/deploy/build.sh

#sh ~/domioapi/deploy/install_deps.sh

sh ~/domioapi/deploy/build.sh

#sh ~/domioapi/deploy/install_pg.sh
