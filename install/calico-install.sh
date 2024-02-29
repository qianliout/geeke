cd /root/work/install

wget https://raw.githubusercontent.com/projectcalico/calico/v3.24.4/manifests/tigera-operator.yaml

kubectl apply -f tigera-operator.yaml

wget  https://raw.githubusercontent.com/projectcalico/calico/v3.24.4/manifests/custom-resources.yaml

kubectl apply -f custom-resources.yaml