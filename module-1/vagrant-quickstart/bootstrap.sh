#!/bin/bash

cd /home/vagrant

# install docker, docker-compose, minikube
sudo yum install -y yum-utils
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo

sudo yum install -y docker-ce docker-ce-cli containerd.io

sudo systemctl start docker

sudo usermod -aG docker vagrant
	
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube

# install development tools
sudo yum group install -y "Development Tools"
sudo yum update -y
sudo yum install -y nc unzip python3 python3-devel git java-11-openjdk-devel

curl -LO https://go.dev/dl/go1.17.5.linux-amd64.tar.gz
tar zxpf go1.17.5.linux-amd64.tar.gz

curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.19.2/protoc-3.19.2-linux-x86_64.zip
unzip protoc-3.19.2-linux-x86_64.zip -d protoc

# install consul, etcd
sudo yum-config-manager --add-repo https://rpm.releases.hashicorp.com/RHEL/hashicorp.repo
sudo yum install -y consul

curl -LO https://github.com/etcd-io/etcd/releases/download/v3.5.1/etcd-v3.5.1-linux-amd64.tar.gz
tar zxpf etcd-v3.5.1-linux-amd64.tar.gz

# change file ownership, set PATH env
chown -R vagrant:vagrant /home/vagrant/*
echo 'export PATH="$PATH:/home/vagrant/go/bin:/home/vagrant/protoc/bin:/home/vagrant/etcd-v3.5.1-linux-amd64"' >> /home/vagrant/.bashrc
