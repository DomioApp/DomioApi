#!/usr/bin/env bash
set -e

echo Installing Go dependendcies...
cd ~/domioapi

go get -u -v github.com/dgrijalva/jwt-go
go get -u -v github.com/aws/aws-sdk-go/aws
go get -u -v github.com/go-ini/ini
go get -u -v github.com/jmespath/go-jmespath
go get -u -v github.com/fatih/color
go get -u -v github.com/gorilla/mux
go get -u -v github.com/gorilla/context
go get -u -v github.com/jmoiron/sqlx
go get -u -v github.com/lib/pq
go get -u -v github.com/stripe/stripe-go
go get -u -v golang.org/x/crypto/bcrypt
go get -u -v github.com/RackSec/srslog

echo Go dependencies installed!