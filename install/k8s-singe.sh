#!/bin/bash

# 安装kubelet/kubeadm/kubectl
sudo apt-get install -y kubelet kubeadm kubectl

# 启动kubelet
sudo systemctl daemon-reload
sudo systemctl restart kubelet

# kube-adm-images.sh
# 如下镜像列表和版本，请运行kubeadm config images list命令获取
#

images=(
  kube-apiserver:v1.24.0
  kube-controller-manager:v1.24.0
  kube-scheduler:v1.24.0
  kube-proxy:v1.24.0
  pause:3.2
  etcd:3.4.13-0
  coredns:1.7.0
)

for imageName in ${images[@]}; do
  docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
  docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
done

# 主机节点网段：10.0.2.0/8
# k8s service网段：10.1.0.0/16
# k8s pod网段：10.244.0.0/16
# 启动k8s集群
sudo kubeadm init \
  --image-repository registry.aliyuncs.com/google_containers \
  --kubernetes-version v1.24.0 \
  --service-cidr=10.1.0.0/16 \
  --pod-network-cidr=10.244.0.0/16

# 配置环境变量KUBECONFIG
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 查看k8s集群状态
kubectl cluster-info
