#!/bin/zsh

sudo apt-get install -y \
  ca-certificates \
  curl \
  gnupg \
  lsb-release \
  wget

wget https://studygolang.com/dl/golang/go1.20.11.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.11.linux-amd64.tar.gz

## 增加GOROOT
echo 'export GOROOT=/usr/local/go' >>~/.bash_profile
echo 'export GOPRIVATE=scm.tensorsecurity.cn' >>~/.bash_profile
echo 'export GONOSUMDB=scm.tensorsecurity.cn' >>~/.bash_profile
echo 'export GONOPROXY=scm.tensorsecurity.cn' >>~/.bash_profile
echo 'export CC=/usr/bin/gcc' >>~/.bash_profile
echo 'export GO111MODULE=on' >>~/.bash_profile
echo 'export CGO_ENABLED=1' >>~/.bash_profile
echo 'export GOBIN=$GOPATH/bin' >>~/.bash_profile
echo 'export PATH=$PATH:$GOROOT/bin' >>~/.bash_profile

source ~/.bash_profile
