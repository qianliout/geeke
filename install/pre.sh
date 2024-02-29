#!/bin/bash

sudo apt-get update -y
sudo apt-get full-upgrade -y
sudo apt-get install -y wget vim
sudo apt-get install -y openssh-client openssh-server
sudo apt-get install -y libclamav-dev libgpgme-dev libdevmapper1.02.1 zip unzip
sudo apt-get install -y ca-certificates curl gnupg lsb-release wget
sudo apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common iptables
#sudo apt-get install -y  docker-ce-cli


sudo /etc/init.d/ssh restart

sudo apt-get install -y \
  ca-certificates \
  curl \
  gnupg \
  lsb-release \
  wget

# 查看系统CPU
cat /proc/cpuinfo
# 查看系统memory
cat /proc/meminfo

uname -a

cat /etc/lsb-release

# 加载netfilter
#modprobe br_netfilter

#cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
#br_netfilter
#EOF

#cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
#net.bridge.bridge-nf-call-ip6tables = 1
#net.bridge.bridge-nf-call-iptables = 1
#EOF

#sysctl --system

# 查看内存中的swap分配情况
#free -m

# 临时关闭swap
swapoff -a

# 查看内存中的swap分配为0
#free -m

# 永久关闭swap（请使用root用户运行如下命令）
#echo "vm.swappiness=0" >>/etc/sysctl.conf
#sudo sysctl -p /etc/sysctl.conf

# 为了方便演示快速部署，建议直接关闭防火墙（若有的话）。
#sudo systemctl stop firewalld && systemctl disable firewalld

# 添加k8s apt国内源（请使用root用户运行如下命令）
#curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -
#cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
#deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
#EOF
#sudo apt-get update -y
