#!/usr/bin/env bash

service nginx stop
letsencrypt certonly --standalone -d api.domio.in --agree-tos -m sergei@basharov.net
service nginx start