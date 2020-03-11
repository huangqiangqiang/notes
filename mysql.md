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

# inner join

内连接，用来获取两个表中字段匹配关系的记录，结果如下所示：

![inner join](./images/img_innerjoin.gif)

举个例子：有一张用户收藏表和商品表，收藏表里存的是用户id和对应的商品id，现在要查询某个用户收藏的商品信息。

我之前的做法会是这样，先查出收藏表中用户的所有收藏记录，在把记录中的商品id拼成一个数组，用 `in (?)` 方式查询出收藏的商品信息。

上面这种方法需要两次查询，我们使用 `inner join` 只要一次就够了，具体的写法如下：

```
select * from favorite_tbl inner join product_tbl on favorite_tbl.product_id = product_tbl.id
```

以上的 SQL 语句等价于：

```
select * from favorite_tbl, product_tbl where favorite_tbl.product_id = product_tbl.id
```

按照以上例子的需求，我们的 SQL 语句可以这样写：

```
select * from favorite_tbl, product_tbl where favorite_tbl.product_id = product_tbl.id and favorite_tbl.user_id = '1'
```

这样，就查找出来用户id为1的用户的收藏记录。

值得注意的是：如果 favorite_tbl 表中的某个 product_id 在 product_tbl 中找不到，则不会返回这条记录，也就是说返回的是两张表中都匹配的记录。

# left join

了解了以上 inner join 之后，了解 left 和 right 的 join 就容易多了，先来看下图：

![inner join](./images/img_leftjoin.gif)

left join 和 inner join 的区别是，left join 会把左表的全部记录查出，举个例子：

```
select * from favorite_tbl left join product_tbl on favorite_tbl.product_id = product_tbl.id
```

以上 SQL 语句的左表是 favorite_tbl ，右表是 product_tbl ，条件是 favorite_tbl.product_id = product_tbl.id ，即使左表的 product_id 在右表中对应的 id 不存在，也会查出来。而使用 inner join 的话，product_id 在右表不存在是不会返回的。

# right join

同 left join

![inner join](./images/img_rightjoin.gif)

# on 和 where

1. 对于left join，不管on后面跟什么条件，左表的数据全部查出来，因此要想过滤需把条件放到where后面

2. 对于inner join，满足on后面的条件表的数据才能查出，可以起到过滤作用。也可以把条件放到where后面。
