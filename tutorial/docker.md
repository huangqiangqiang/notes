# 安装

### Mac

docker在mac下是一个客户端，[下载](https://download.docker.com/mac/stable/Docker.dmg)

### Ubuntu
```
wget -qO- https://get.docker.com/ | sh
```

### CenterOS
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

# 基本命令

```
# 查看镜像
docker images

# 查看当前正在运行的容器
docker ps

# 查看所有的容器（包括正在运行和已停止的容器）
docker ps -a

# 查看容器的详细信息
docker inspect <dockerName|dockerId>
```