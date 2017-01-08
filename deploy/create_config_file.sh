#!/usr/bin/env bash
echo Creating config file...
echo "{ \"AWS_ACCESS_KEY_ID\": \"${AWS_ACCESS_KEY_ID}\", \"AWS_SECRET_ACCESS_KEY\": \"${AWS_SECRET_ACCESS_KEY}\", \"DOMIO_DB_USER\": \"${DOMIO_DB_USER}\", \"DOMIO_DB_PASSWORD\": \"${DOMIO_DB_PASSWORD}\", \"DOMIO_DB_NAME\": \"${DOMIO_DB_NAME}\", \"PORT\": ${PORT} }" > /domio/config.json
echo Config file created!