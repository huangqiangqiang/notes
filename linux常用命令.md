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
```

# 查看系统一级文件夹大小
```
df -h
```
