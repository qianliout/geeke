#!/bin/zsh

sudo apt-get install -y \
  ca-certificates \
  curl \
  gnupg \
  lsb-release \
  wget

cd /root/work

wget https://studygolang.com/dl/golang/go1.19.6.linux-amd64.tar.gz

tar -C /usr/local -zxvf go1.19.6.linux-amd64.tar.gz

## 增加GOROOT
echo 'export GOROOT=/usr/local/go' >>/root/.profile
echo 'export GOPRIVATE=scm.tensorsecurity.cn' >>/root/.bash_profile
echo 'export GONOSUMDB=scm.tensorsecurity.cn' >>/root/.bash_profile
echo 'export GONOPROXY=scm.tensorsecurity.cn' >>/root/.bash_profile
echo 'export PATH=$PATH:$GOROOT/bin' >>/root/.profile

source /root/.profile
