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