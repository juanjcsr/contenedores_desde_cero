#!/usr/bin/env bash

sudo apt-get update
sudo apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

sudo apt-get update
sudo apt-get install -y docker-ce

curl -LO https://storage.googleapis.com/golang/go1.10.linux-amd64.tar.gz
sudo tar -C /usr/local -xvzf go1.10.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >>/home/vagrant/.bashrc
echo "export GOROOT=/usr/local/go" | tee -a /root/.bashrc
echo "export PATH=$PATH:/usr/local/go/bin" | tee -a /root/.bashrc
mkdir -p /home/vagrant/containers/
echo "export GOPATH=/vagrant/demo" | tee -a /home/vagrant/.bashrc
echo "export GOPATH=/vagrant/demo" | tee -a /root/.bashrc


mkdir -p /home/vagrant/containers/fs/rootfs-alpine/
mkdir -p /home/vagrant/containers/fs/rootfs-ubuntu/

tar -xzvf /vagrant/alpine-rootfs.tar.gz -C /home/vagrant/containers/fs/rootfs-alpine/
tar -xzvf /vagrant/ubuntu-rootfs.tar.gz -C /home/vagrant/containers/fs/rootfs-ubuntu/
