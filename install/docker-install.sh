
sudo apt-get -y remove docker docker-engine docker.io containerd runc

sudo apt-get -y update
sudo apt-get install -y \
  ca-certificates \
  curl \
  gnupg \
  lsb-release \
  curl

# 安装curl工具，若已安装可以跳过
sudo apt install -y curl

sudo mkdir -m 0755 -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list >/dev/null

# 安装docker
sudo apt-get -y update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# 安装完毕之后可以在如下路径查看到docker container runtime，这是k8s的缺省搜索路径
# https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/
# kubeadm automatically tries to detect an installed container runtime by scanning through a list of well known Unix domain sockets.
ll /var/run/docker.sock

# 配置 Docker daemon，设置国内源
#
sudo mkdir /etc/docker
cat <<EOF | sudo tee /etc/docker/daemon.json
{
  "registry-mirrors": [
    "https://hub-mirror.c.163.com",
    "https://mirror.baidubce.com"
  ]
}
EOF

# 重启 docker 后台服务
sudo systemctl daemon-reload
sudo systemctl restart docker

sudo systemctl enable docker

sudo usermod -aG docker $USER && newgrp docker

sudo curl -L "https://github.com/docker/compose/releases/download/2.20.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose