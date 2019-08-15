# Docker的安装

### Mac

docker在mac下是一个客户端，[下载](https://download.docker.com/mac/stable/Docker.dmg)

### Windows
[下载](https://docs.docker.com/docker-for-windows/install/)

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
docker inspect <container_name|container_id>

# 配合dockerfile使用
docker build -t <image_name>:<image_tag> .

# 进入容器内部
docker exec -it <container_name> /bin/sh

# 删除容器
docker rm <container_name>

# 删除镜像
docker rmi <image_name>
```

# docker login

docker login功能主要是用于登录，因为有些image（镜像）需要登录后才能pull，不是谁都能拉下来的。比如，gitlab自带一个docker仓库，你可以把自己的镜像push上去，也可以pull下来，当然，这些都是需要权限的。
就是需要先执行`docker login`。
```
docker login -u <gitlab_username> -p <gitlab_password>
```

# Docker的network

```
# 创建network（最好是指定--subnet网段，以免和内网冲突）
docker network create <network_name> --subnet=172.120.0.0/16

# 删除network（删除的时候必须确认该network没有其他容器正在使用）
docker network rm <network_name>

```

注：`172.120.0.0/16`代表的是`172.120.0.0~172.120.255.255`的范围，16是掩码位，ipv4是32位的，相当于把前16位给固定死了，只能改后16位。类似：
`11000000 10101000 00000000 00000000 - 11000000 10101000 11111111 11111111`

创建好的network就可以给容器使用了，也可以在docker-compose.yml文件中使用。例子：

```
version: '2'

services:
  app_name:
    build: ./
    networks:
      - my_network
    ports:
      - "80:80"

networks:
  my_network:
    external: true
```

# dockerfile

这是我公司一个go项目的dockerfile，`alpine`是一个很轻量的linux操作系统，只有5m大小

```
FROM alpine
WORKDIR /work
COPY . /work/
CMD ["./app"]
```