#!/bin/zsh

sudo apt-get install -y \
  ca-certificates \
  curl \
  gnupg \
  lsb-release \
  wget

cd ~/work

wget https://studygolang.com/dl/golang/go1.19.6.linux-amd64.tar.gz

tar -C /usr/local -zxvf go1.19.6.linux-amd64.tar.gz

## 增加GOROOT
echo 'export GOROOT=/usr/local/go' >>~/.profile
echo 'export GOPRIVATE=scm.tensorsecurity.cn' >>~/.profile
echo 'export GONOSUMDB=scm.tensorsecurity.cn' >>~/.profile
echo 'export GONOPROXY=scm.tensorsecurity.cn' >>~/.profile
echo 'export PATH=$PATH:$GOROOT/bin' >>~/.profile

source ~/.profile
