#!/usr/bin/env bash

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
echo Compiling for Windows...
export GOARCH=amd64
export GOOS=windows
go build -o /usr/local/bin/domio.exe -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" domio


echo Compiling for Linux...
export GOARCH=amd64
export GOOS=linux
go build -o /usr/local/bin/domio -ldflags "-X main.Buildstamp=$buildstamp -X main.Hash=$hash  -X main.Version=$version" domio
