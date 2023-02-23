#!/bin/bash

wget  https://gitee.com/pphh/blog/blob/master/210215_k8s_deployment/kube-flannel.yml

kubectl apply -f kube-flannel.yml

cat /etc/cni/net.d/10-flannel.conflist