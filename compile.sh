#!/usr/bin/env bash

PROJECT_NAME=domio

#buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S'`
buildstamp=`date -u '+%s'`
hash=`git rev-parse --short HEAD`
version=`git tag -l --points-at HEAD`

echo
echo ---------------------------
echo "  Buildstamp: ${buildstamp}"
echo "  Hash:       ${hash}"
echo "  Version:    ${version}"
echo ---------------------------
echo


#echo Compiling for Windows...
#export GOARCH=amd64
#export GOOS=windows
#go build -o /usr/local/bin/${PROJECT_NAME}_win.exe -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" ${PROJECT_NAME}


#echo Compiling for Linux...
#export GOARCH=amd64
#export GOOS=linux
#go build -o /usr/local/bin/${PROJECT_NAME}_linux -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" ${PROJECT_NAME}

echo Compiling for Mac...
export GOARCH=amd64
export GOOS=darwin
go build -o /usr/local/bin/${PROJECT_NAME}_mac -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" ${PROJECT_NAME}
