# Docker的安装

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

# Docker的基本命令

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

# Docker的network

```
# 创建network（最好是指定--subnet网段，以免和内网冲突）
docker network create <network_name> --subnet=172.120.0.0/16

# 删除network（删除的时候必须确认该network没有其他容器正在使用）
docker network rm <network_name>
```

创建好的network就可以给容器使用了，也可以在docker-compose.yml文件中使用。例子：

```
version: '2'

services:
  app:
    build: ./
    networks:
      - my_network
    ports:
      - "80:80"

networks:
  my_network:
    external: true
```