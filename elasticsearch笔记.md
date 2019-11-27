最近公司搭建了 kafka 来收集 k8s 上所有的请求日志记录，然后同步数据到 elasticsearch 中，最后在kibana管理后台来查询日志。

因为 kibana 后台中不能直接显示出 kafka 同步过来的数据（现在只能查询容器中标准输出的日志），所以需要在 `dev tool` 中自己构建查询语句来查询。

# 几个概念

elastic中有几个概念需要了解：

- index（索引）：可以理解为数据库，1个 index 表示1个数据库。
- document：一条记录，使用 json 格式表示，许多 documents 构成一个 index。
- type：每个 document 都有一个 `_type` 字段，对 document 进行分组。

# 查询语句

查看所有的 index （列出所有数据库）

```
GET _cat/indices?v
```

列出某个数据库中的 document

```
GET /{index}/{type}/_search
或
GET /{index}/_search
```

查询可以带参数

```
GET /log/_search
{
  "query": {
    "match": { "request.uri": "home" } // 要匹配的字段和值
  }, 
  "size": 2           // 展示的条数，默认展示10条
  "from": 1           // 表示位移
}

# 如果要匹配多个字段
"query": {
  "bool": {
    "must": [
      {
        "match": {
          "request.uri": "home"
        }
      },
      {
        "match": {
          "request.body": "phone"
        }
      }
    ]
  }
}

# 根据时间范围查询
"query": {
  "bool": {
    "must": [
      {
        "match": {
          "request.uri": "home"
        }
      }
    ],
    "filter": {
      "range": {
        "start_time": {
          "gte": 1574784000,
          "lte": 1574835193
        }
      }
    }
  }
}
```



# 搜索结果说明

```
{
  "took" : 25,
  "timed_out" : false,
  "hits" : {
    "total" : 264,
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "log-0001.01.01",
        "_type" : "doc",
        "_id" : "dD-upm4B9_WPbZUqrKnW",
        "_score" : 1.0,
        "_source" : {
          "start_time" : 1.574754490209E9,
          "request" : {
            "method" : "POST",
            "headers" : {
              "accept" : "application/json"
            }
          },
          "cost_time" : 0.095999956130981
        }
      }
    ]
  }
}
```

返回的结果中， took 表示该操作的耗时（毫秒），timed_out 表示本次请求是否超时，hits 表示命中的 document。

hits 里面的字段说明，total 表示该 index 中所有 document 的数量。max_score 表示最高的匹配程度。hits 表示返回的记录数组。

