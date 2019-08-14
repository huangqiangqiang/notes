# 在记录不存在时插入，存在时则更新

需要用到 `on duplicate key update` 语句

示例:
```
insert into teacher(id, name, phone, avatar) values('%d', '%s', '%s', '%s') on duplicate key update name='%s', phone='%s', avatar='%s';
```

前半句是在执行插入操作，如果前半句中的insert操作中不能重复的字段重复了，比如id，就会执行update操作。

# mysql支持表情

1. mysql版本必须5.5或以上

2. 数据库编码格式设置为utf8mb4，理论上向下兼容utf8。

3. 连接数据库的url带上charset='utf8mb4'参数