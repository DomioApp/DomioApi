#!/usr/bin/env bash

echo Installing Postgres...
sudo apt-get -y install postgresql postgresql-contrib
psql --version