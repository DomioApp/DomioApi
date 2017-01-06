#!/usr/bin/env bash

ssh root@104.131.183.100 "export DOMIO_DB_USER=postgres; export DOMIO_DB_PASSWORD=pass; export DOMIO_DB_NAME=domio_dev; /domio/domio &"