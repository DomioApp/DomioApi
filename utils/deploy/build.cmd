cd ../../


set GOOS=linux
set GOARCH=386
gb build domio

set GOOS=linux
set GOARCH=amd64
gb build domio

set GOOS=darwin
set GOARCH=amd64
gb build domio

set GOOS=windows
set GOARCH=amd64
gb build domio


cd utils/deploy