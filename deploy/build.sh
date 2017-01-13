#!/usr/bin/env bash
set -e

echo Building Domio...
rm -rf /domio
mkdir /domio

cd ~/domioapi

go build -o /domio/domio domio

/domio/domio init --aws-access-key-id=$AWS_ACCESS_KEY_ID \
                  --aws-secret-access-key=$AWS_SECRET_ACCESS_KEY \
                  --db-name=$DOMIO_DB_NAME \
                  --db-user=$DOMIO_DB_USER \
                  --db-password=$DOMIO_DB_PASSWORD

#/domio/domio init --aws-access-key-id=12 --aws-secret-access-key=23 --db-name=34 --db-user=45 --db-password=56

cd /
rm -rf ~/domioapi

echo Domio is built and ready!

logger -n logs5.papertrailapp.com -t deploy -P 18422 -p user.notice "Domio is built and ready!"