#!/bin/bash

# install docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

docker-compose --version

# install development tools
sudo yum group install -y "Development Tools"
sudo yum install -y wget nc unzip python3 python3-devel git java-11-openjdk-devel

wget https://go.dev/dl/go1.17.5.linux-amd64.tar.gz
tar zxpf go1.17.5.linux-amd64.tar.gz

wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.2/protoc-3.19.2-linux-x86_64.zip
unzip protoc-3.19.2-linux-x86_64.zip -d protoc

python3 --version
pip3 --version
java -version

~/go/bin/go version
~/protoc/bin/protoc --version

# install consul
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://rpm.releases.hashicorp.com/RHEL/hashicorp.repo
sudo yum -y install consul

# install etcd
wget https://github.com/etcd-io/etcd/releases/download/v3.5.1/etcd-v3.5.1-linux-amd64.tar.gz
tar zxpf etcd-v3.5.1-linux-amd64.tar.gz

# update PATH env
echo 'export PATH="$PATH:~/go/bin:~/protoc/bin:~/etcd-v3.5.1-linux-amd64"' >> ~/.bashrc
