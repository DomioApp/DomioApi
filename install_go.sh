#!/usr/bin/env bash
sudo apt-get update
sudo apt-get -y upgrade

rm -rf /godist
rm /usr/bin/go
rm /usr/bin/go-fmt

mkdir /godist
cd /godist
wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
sudo tar -xvf go1.7.4.linux-amd64.tar.gz
sudo mv go /usr/local

ln -s /usr/local/go/bin/go /usr/bin/go
ln -s /usr/local/go/bin/go-fmt /usr/bin/go-fmt

export GOROOT=/usr/local/go

rm go1.7.4.linux-amd64.tar.gz