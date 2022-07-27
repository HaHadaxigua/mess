# Overview

- [install](#install)
    - [linux-amd64](#linux-amd64)
- [config go env](#config-go-env)
    - [basic](#basic)
    - [private](#private)

# install

## linux-amd64

```shell
#!/bin/bash 

version="1.18.4"
os="linux"
cpu="amd64"

sgo = go$version.$os.$cpu

wget https://go.dev/dl/$sgo.tar.gz
mkdir go$version-1
tar -zxvf $sgo.tar.gz
mv go$version-1/go go$version
rm -r go$version-1
mv go$version /usr/local/go$version
```

# config go env

## basic

```shell
export GOROOT=/usr/local/go
export GOPATH=/Users/dh/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
export GOPROXY=https://goproxy.cn
```

## private

```shell
git_username="xxx"
git_token="xxx"
git_host="git.xxxx.com:port/xxx"
git_http_host="http://git.xxxx.com:port/xxx"
git_email="xxx@xx.com"

git config --global url."https://$git_username:git_token@git_host".insteadOf "git_http_host"
git config --global user.email git_email
git config --global user.name git_username
```
