#!/usr/bin/env bash
set -e

echo Building Domio...
rm -rf /domio
mkdir /domio

cd ~/domioapi


#=====================================================================================================================

buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
buildstamp=`date -u '+%s'`
hash=`git rev-parse --short HEAD`
version=`git tag -l --points-at HEAD`

echo ------------------------------------------------------
echo "Buildstamp: ${buildstamp}"
echo "Hash:       ${hash}"
echo "Version:    ${version}"
echo ------------------------------------------------------

go build -o /domio/domio -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" domio

#=====================================================================================================================

echo ------------------------------------------------------
echo Initializing config file with these variables:
echo --aws-access-key-id=$AWS_ACCESS_KEY_ID \
echo --aws-secret-access-key=$AWS_SECRET_ACCESS_KEY \
echo --db-name=$DOMIO_DB_NAME \
echo --db-user=$DOMIO_DB_USER \
echo --db-password=$DOMIO_DB_PASSWORD
echo ------------------------------------------------------

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