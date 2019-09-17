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

备份恢复

```
mongorestore /mongodbjump
```