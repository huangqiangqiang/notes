MongoDB默认是无密码直接连接的

# docker-compose.yml
```
version: '2'
services:
  mongodb:
    image: mongo:latest
    environment:
      - MONGO_INITDB_ROOT_USERNAME=${username}
      - MONGO_INITDB_ROOT_PASSWORD=${password}
      - MONGO_INITDB_DATABASE=${database_name}
    ports:
      - "27017:27017"
    volumes:
      - ${PWD}/db:/data/db
```

# mongodb connect
```
mongodb://${username}:${password}@SERVER_IP:27017/${database_name}?authSource=admin
```
`authSource`表示用哪个数据库来验证用户名和密码，默认是`admin`

# mongo 命令行连接

```
mongo 192.168.1.2 -u xxx -p xxx --authenticationDatabase admin
```
或
```
mongo --host mongodb://${username}:${password}@SERVER_IP:27017/${database_name}?authSource=admin
```

# 按时间倒序

```
{timestamp:-1}
```

# 模糊查询

```
# 搜索name字段包含haha字符串的结果
{name:/haha/}
```

# 大于/小于某个时间

```
# gte: 大于某个时间
# lt: 小于某个时间
{timestamp: {$gte: '1553617031966', $lt: '1553616593572'}}
```

# 数据备份

最新有一个需求，因为有很多项目都有搜集请求日志。导致 mongodb 中的日志很多，但又是不怎么重要的。所以现在的优化方案是每周做一次本周到上周时间范围内的备份，然后删除3个月以前的数据。因为日志主要是查询用，还原当时的请求，一般时间太久的话是不需要的，所以直接备份成文件，从数据库中删除。

我们采用的是 mongodump 命令备份 mongodb 数据，然后放入 crontab 每周定时执行一次。

mongo 数据库的备份主要使用 mongodump 命令

```
mongodump -h 127.0.0.1:27017 -d test -o /mongodbjump
```
- -h 表示 `--host`        host地址 
- -d 表示 `--db`          数据库的名字
- -c 表示 `--collection`  数据表的名字
- -o 表示 `--out`         备份文件输出目录
- -q 表示 `--query`       指定备份条件

如果要备份指定的时间范围，如果要指定时间范围，必须指定 -c 参数

```
mongodump -d lltest -c test  -q '{"created_at":{$gte:Date(1568649600000),$lt:Date(1568736000000)}}' -o /mongodbjump
```

注：
```
gt: greater than 大于
gte: greater than or equal 大于等于
lt: less than 小于
lte: less than or equal 小于等于
```

话不多说，直接贴我们现在在用的备份脚本：

```
# mongo_bak.sh

#!/bin/sh
# 没有设置 PATH 会报错：docker command not found
PATH=/usr/local/bin:/usr/local/sbin:~/bin:/usr/bin:/bin:/usr/sbin:/sbin

# 当前时间毫秒
today_stamp=`expr $(date +%s) \* 1000`
offset=`expr 3600 \* 24 \* 7 \* 1000`
lastweek_stamp=`expr $today_stamp - $offset`

# 备份的文件夹名称
DATE=`date +%Y_%m_%d_%H_%M_%S`   
back_path="/mongodump/mongo_bak_$DATE"

# 执行容器内的命令，又不用进入容器
query="{\"created_at\":{\$gte:Date($lastweek_stamp),\$lt:Date($today_stamp)}}"
bakCommand="mongodump -h 127.0.0.1:27017 -d lltest -c test -q '$query' -o $back_path"
echo $bakCommand
docker exec test_mongodb /bin/sh -c "$bakCommand"
```

以上脚本为每周日0点执行，生成的备份文件在 `/mongodump/mongo_bak_$DATE` 目录。

备份恢复

```
mongorestore /mongodbjump
```