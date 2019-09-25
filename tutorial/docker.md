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


### 安装 docker-compose

Docker Compose 存放在Git Hub，不太稳定。 
你可以也通过执行下面的命令，高速安装Docker Compose。
```
curl -L https://get.daocloud.io/docker/compose/releases/download/1.23.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
```
```
chmod +x /usr/local/bin/docker-compose
```

官网安装
```
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```
https://docs.docker.com/compose/install/#install-compose

---

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


# 一些注意细节讲解

### 其实 docker 包括了客户端和守护进程两部分

docker 包括客户端和后台运行的守护进程两个部分，在构建镜像的时候（build）都是由客户端把文件传到守护进程来构建的，而且客户端和守护经常不要求在同一台机器上，如果在一台非 linux 操作系统中使用docker，客户端就运行在你的操作系统上，但守护进程运行在虚拟机内，由于构建的文件都会传到守护进程内，如果包含了大文件而且守护进程不在本地运行，上传过程会花很多时间。

### 镜像不是一个大的二进制块，内部是分层的

这个可能很多人都知道镜像是分层的，但是你或许会认为一个Dockerfile只创建一个新的层，但是并不是这样，构建镜像时，Dockerfile中的每行指令都会创建一个新的层。

不通镜像会共享分层，所有相同的分层只会被存储一次，在拉取镜像（pull image）的时候也一样，docker 会独立下载每一层，一些分层可能已经存储在机器上了，所以 docker 只会下载未被存储的分层。

### 向 dockerhub 推送镜像

推送到 dockerhub 的镜像需要一定的命名格式，就是 `dockerhub_ID/image_name:tag` 这种格式，如果要推送的镜像不是这样格式的，可以重命名镜像，执行：

```
docker tag image_name:tag dockerhub_ID/image_name:tag
```