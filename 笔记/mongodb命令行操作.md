
进入 mongodb 容器后

# 登录
```
mongo -u <username> -p <password>
```

# 查询使用数据库
```
show dbs
use <db_name>
```

# 查看表

```
show tables;
```

# 查看表

```
db.<表名>.find()
db.<表名>.find().count()
```

# 获取最早的一条记录
```
db.<表名>.find({}).limit(1).pretty()
```

# 更新某条文档
```
db.<表名>.update({'step':'SUCCESS'}, {$set: {'step': 'FAILURE'}})
```

# 删除某条文档
```
db.<表名>.remove({'step':'SUCCESS'})
```