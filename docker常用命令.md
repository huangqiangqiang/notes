
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


# 删除none镜像
```
docker rmi $(docker images -f "dangling=true" -q)
```

# 查看容器的配置，比如容器的ip，容器的挂载目录
```
docker inspect [container_name]
```

# 快速把一个项目打包成镜像上传到服务器
- 项目根目录添加dockerfile，把所有项目文件都拷贝进去。如：
```
FROM node:alpine
WORKDIR /usr/src/app
COPY . .
CMD [ "npm", "start" ]
```

- 项目正常跑起来后，打包成镜像
```
docker build -t <ImageName> .
```

- 把镜像存成tar文件
```
docker save -o xxx.tar <ImageName>
```

- 把tar文件上传到服务器
```
scp xxx.tar root@127.0.0.1:/root
```

- 把tar文件载入镜像
```
docker load --input xxx.tar
```

# Push 自己做的镜像到 dockerhub

首先，自己做的镜像名字有要求，必须是 `dockerhub的用户名/镜像名` ，比如 `hqqsk8/golang` ，

镜像重命名

```
docker tag <原来镜像的名称> <修改后的镜像名称>
```

如果带 tag 的话就是，`hqqsk8/golang:latest`，镜像做好了之后就是先登录dockerhub

```
docker login
```

输入用户名和密码，登录成功后，执行 push

```
docker push hqqsk8/golang
```

# docker内部时区改成东8区

dockerfile 加入以下命令
```
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
```