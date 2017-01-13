#!/usr/bin/env bash
echo Creating config file...

/domio/domio init --aws-access-key-id=$AWS_ACCESS_KEY_ID \
                  --aws-secret-access-key=$AWS_SECRET_ACCESS_KEY \
                  --db-name=$DOMIO_DB_NAME \
                  --db-user=$DOMIO_DB_USER \
                  --db-password=$DOMIO_DB_PASSWORD


echo Config file created!