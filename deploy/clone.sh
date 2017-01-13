#!/usr/bin/env bash

cd ~/
rm -rf ~/domioapi
git clone git@gitlab.com:basharov/DomioApi.git ~/domioapi
cd ~/domioapi
git tag -l --points-at HEAD
