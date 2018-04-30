#!/usr/bin/env bash

sudo yum -y update
curl -LO https://storage.googleapis.com/golang/go1.10.linux-amd64.tar.gz
sudo tar -C /usr/local -xvzf go1.10.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >>/home/vagrant/.bashrc

mkdir -p /home/vagrant/containers/
echo "export GOPATH=/vagrant/demo" | tee -a /home/vagrant/.bashrc


mkdir -p /home/vagrant/containers/fs/rootfs-alpine/
mkdir -p /home/vagrant/containers/fs/rootfs-ubuntu/

tar -xzvf /vagrant/alpine-rootfs.tar.gz -C /home/vagrant/containers/fs/rootfs-alpine/
tar -xzvf /vagrant/ubuntu-rootfs.tar.gz -C /home/vagrant/containers/fs/rootfs-ubuntu/
