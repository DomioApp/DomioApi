#!/usr/bin/env bash

cd ~/domioapi
export GOPATH=$PWD
echo $GOPATH


go get -u github.com/dgrijalva/jwt-go
go get -u github.com/aws/aws-sdk-go/aws
go get -u github.com/go-ini/ini
go get -u github.com/jmespath/go-jmespath
go get -u github.com/fatih/color
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/context
go get -u github.com/jmoiron/sqlx
go get -u github.com/lib/pq
go get -u github.com/stripe/stripe-go
go get -u golang.org/x/crypto/bcrypt