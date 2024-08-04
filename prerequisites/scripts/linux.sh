#!/bin/bash

sudo apt update
sudo apt install -y wget curl tar git

GO_VERSION="1.22.4"
wget https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz
rm go$GO_VERSION.linux-amd64.tar.gz

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

sudo apt install -y git

if ! command -v code &> /dev/null; then
    wget -O code.deb https://go.microsoft.com/fwlink/?LinkID=760868
    sudo apt install -y ./code.deb
    rm code.deb
fi

go version
git --version
