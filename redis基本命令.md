连接redis

```
redis-cli -h 127.0.0.1 -p 6379 -a pwd
```

选择数据库

```
select 0
```

查看所有的key

```
keys *
```

查看某个key的值

```
get key
```

获取key的过期时间（如key有设置过期时间，则返回过期的秒数，如没设置则返回-1，key不存在返回-2）

```
# 5秒后过期(TTL, time to live)）
ttl key
```

设置key的过期时间

```
# 5秒后过期
EXPIRE key 5
```