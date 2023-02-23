cd /root/work/install/jenkins

# git clone https://github.com/scriptcamp/kubernetes-jenkins

kubectl create namespace devops-tools

kubectl apply -f serviceAccount.yaml

kubectl apply -f volume.yaml

kubectl apply -f deployment.yaml

kubectl apply -f service.yaml
