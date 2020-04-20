# Minikube的安装

[Minikube安装教程](https://kubernetes.io/docs/tasks/tools/install-minikube/)

安装好后执行

```
minikube start
```

等待 minikube 启动...

# Minikube 升级

[下载最新的包](https://github.com/kubernetes/minikube/releases)

```
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

再用 `minikube version` 查看版本是不是最新的。

# Minikube 遇到的一些问题

- 有时候使用 `kubectl get po` 等命令发生错误，使用 `minikube status` 查看，发现 `apiserver` 挂掉了，但是不知道怎么重新启动 `apiserver` ，只能是 重启电脑然后再次执行 `minikube start` 。

- 升级完 minikube 版本后 `minikube start` 启动报错，不知道为什么。后来加上 `--vm-driver=virtualbox` 参数才可以。

# 安装helm（k8s包管理工具）
[下载helm](https://github.com/helm/helm/releases/latest)

找到对应系统下载后解压，执行：

```
helm init
```

该命令会创建一个 tiller , tiller 是 helmCli 工具的服务端，负责和 k8s 交互。

这里创建 tiller 这个 pod 的时候可能会有些问题，镜像拉不下来。这里的镜像不是电脑本地的镜像，而是执行 minikube ssh 后，登录到虚拟机中，执行 docker images 看到的镜像。

我的情况是 Minukube 虚拟机里 tiller 镜像拉不下来，本地却可以下载 tiller 镜像，所以我是把本地镜像打包成 tar 文件，再到虚拟机里解压，以下是操作步骤：

```
# 先搜索 dockerhub 上可用的镜像
docker search tiller
```

搜到的结果是：
```
NAME                                    DESCRIPTION                                     STARS               OFFICIAL            AUTOMATED
jessestuart/tiller                      Nightly multi-architecture (amd64, arm64, ar…   19                                      [OK]
sapcc/tiller                            Mirror of https://gcr.io/kubernetes-helm/til…   9                                       
ist0ne/tiller                           https://gcr.io/kubernetes-helm/tiller           3                                       [OK]
rancher/tiller                                                                          2                                       
jmgao1983/tiller                        from gcr.io/kubernetes-helm/tiller              2                                       [OK]
ibmcom/tiller                           Docker Image for IBM Cloud private-CE (Commu…   1                                       
itinerisltd/tiller-circleci                                                             1                   
```

`sapcc/tiller` 这个镜像，说明是 `Mirror of Mirror of https://gcr.io/kubernetes-helm/til…  ` 说明应该是和官方镜像一样的，然后我们在 dockerhub 上面搜 `sapcc/tiller` ，点进去查看tag，发现和官方的tag是一样的，我们下一个需要的版本。
```
docker pull sapcc/tiller:v2.16.6
```

pull 下来后给他改名，改成和官方镜像的名字一样

```
docker tag sapcc/tiller:v2.16.6 gcr.io/kubernetes-helm/tiller:v2.16.6
```

现在我们本地已经有这个 tiller 镜像了，现在我们把它导入到 minikube 的 docker 环境里：

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