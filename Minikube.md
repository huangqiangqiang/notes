# Minikube的安装

[Minikube安装教程](https://kubernetes.io/docs/tasks/tools/install-minikube/)

安装好后执行

```
minikube start
```

等待 minikube 启动...

# 安装helm（k8s包管理工具）
[下载helm](https://github.com/helm/helm/releases/latest)

找到对应系统下载后解压，执行：

```
helm init
```

该命令会创建一个 tiller , tiller 是 helmCli 工具的服务端，负责和 k8s 交互。

这里创建 tiller 这个 pod 的时候可能会有些问题，镜像拉不下来。这里的镜像不是电脑本地的镜像，而是执行 minikube ssh 后，登录到虚拟机中，执行 docker images 看到的镜像。

我的情况是 Minukube 虚拟机里 tiller 镜像拉不下来，本地却可以下载 tiller 镜像，所以我是把本地镜像打包成 tar 文件，再到虚拟机里解压，一下是操作步骤：

```
#打包
docker save -o xxx.tar <ImageName>

# 登录到虚拟机
minikube ssh

# 解压（载入镜像）
docker load --input xxx.tar
```

# 更换源

装好 helm 后，你会发现很多软件都搜不到，我们需要更换源，这里有一个微软的源比较好，更换方式是：

```
# 查看当前源
helm repo list

# 删除源
helm repo remove stable
helm repo remove incubator

# 添加源
helm repo add stable http://mirror.azure.cn/kubernetes/charts/
helm repo add incubator http://mirror.azure.cn/kubernetes/charts-incubator/
```