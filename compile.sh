#!/usr/bin/env bash

#buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S'`
buildstamp=`date -u '+%s'`
hash=`git rev-parse --short HEAD`
version=`git tag -l --points-at HEAD`

echo ------------------------------------------------------
echo "Buildstamp: ${buildstamp}"
echo "Hash:       ${hash}"
echo "Version:    ${version}"
echo ------------------------------------------------------

echo Compiling for Linux...
set GOARCH=amd64
set GOOS=linux
go build -o /usr/local/bin/domio -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" domio

echo Compiling for Windows...
set GOARCH=amd64
set GOOS=windows
go build -o /usr/local/bin/domio.exe -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" domio
