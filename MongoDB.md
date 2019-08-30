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
