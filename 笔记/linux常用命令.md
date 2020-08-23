修改所属用户和所属组
```
chown -R root:root /tmp
```

# 查看Linux配置
查看服务器有几个cpu
```
cat /proc/cpuinfo |grep "physical id"|sort |uniq|wc -l
```

1个cpu中有几个核
```
cat /proc/cpuinfo |grep "cores"|uniq
```

1个物理cpu中有几个逻辑CPU
```
cat /proc/cpuinfo| grep "processor"| wc -l
```

# 查看cpu使用率
输入`top`后按`1`

# CI中使用公私钥

1. 生成当前用户的ssh: 
```
ssh-keygen -t rsa
```
2. 创建authorized_keys并设置权限: 
```
touch authorized_keys
chmod 600 authorized_keys 
```
3. 添加公钥:  
```
cat id_rsa.pub >> ~/.ssh/authorized_keys
```
4. 复制id.rsa的内容，添加到CI/CD里面的 SSH_PRIVATE_KEy

# 查看文件夹大小
```
du -h
或
du -sh *

# 查看 root 目录下各个文件夹大小
du -h /root --max-depth=1
```

# 查看磁盘使用情况
```
df -h

Filesystem      Size  Used Avail Use% Mounted on
devtmpfs        2.0G     0  2.0G   0% /dev
tmpfs           2.0G     0  2.0G   0% /dev/shm
/dev/vda1        40G   23G   15G  61% /
/dev/vdb         50G   23G   25G  48% /data
overlay          40G   23G   15G  61% /var/lib/docker/overlay2/8e2e1ef400b95321a2370b3fe652605b5870384c7d632768ce619e2732e6d72d/merged
shm              64M     0   64M   0% /var/lib/docker/containers/f6090219cc71517c7318570ae786bc65431efbd8e48e5396c44fcbab764f0718/mounts/shm
...
```

`/dev/vda` 和 `/dev/vdb` 表示有两个磁盘，后面的数字表示磁盘的分区。分区类似把一个盘分成C盘和D盘的意思差不多。
`tmpfs` 是一个临时文件系统（默认最大为内存的一半大小），驻留在内存中，我们后面的目录 `/dev/shm` 不是在硬盘上，而是在内存中，所以这个目录下读写非常快，但是这个目录里的数据在断电后会丢失。了解了tmpfs这个特性可以用来提高服务器性能，把一些对读写性能要求较高，但是数据又可以丢失的这样的数据保存在 `/dev/shm` 中，来提高访问速度。


# 查找文件

```
# 查找当前目录下以 .txt 结尾的文件
find . -name "*.txt"
```