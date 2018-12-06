
# docker和docker-compose安装

1. 删除旧版

```
sudo yum remove docker \
          docker-common \
          docker-selinux \
          docker-engine
```

2.  安装库

```
sudo yum install -y yum-utils \
  device-mapper-persistent-data \
  lvm2
```

3. 配置stable repo
```
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
```
4. 安装
```
sudo yum install docker-ce
```
5. 启动
```
sudo systemctl start docker
```

6. 开机启动

```
sudo systemctl enable docker
```

7 . hello world

```
sudo docker run hello-world
```

### 安装 Docker Compose

Docker Compose 存放在Git Hub，不太稳定。 
你可以也通过执行下面的命令，高速安装Docker Compose。
```
curl -L https://get.daocloud.io/docker/compose/releases/download/1.23.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
```
```
chmod +x /usr/local/bin/docker-compose
```

https://get.daocloud.io/#install-compose

## 官网安装
```
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```
https://docs.docker.com/compose/install/#install-compose

---


# 删除none镜像
```
docker rmi $(docker images -f "dangling=true" -q)
```

# 查看容器的配置，比如容器的ip，容器的挂载目录
```
docker inspect [container_name]
```