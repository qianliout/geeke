#!/bin/bash

# https://minikube.sigs.k8s.io/docs/start/

cd ~/work/install

curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb

sudo dpkg -i minikube_latest_amd64.deb
