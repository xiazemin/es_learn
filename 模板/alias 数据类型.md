https://elasticstack.blog.csdn.net/article/details/100179703
https://www.elastic.co/guide/en/ecs/current/ecs-getting-started.html

就像其他的很多语言一样，我们可以给已有的变量取一个别名（alias）。即便是对高级语言一样，比如我们定义不同的指针变量，指向同一个内存空间。这个有些类似别名的概念。

在Elasticsearch中，我们也可以为index中的一个字段（field）取一个另外的名字：

它可以用来代替搜索请求中的目标（target）字段
以及其它的被选定的API中
通常alias可以用来帮助我们重新命名一个字段，并让这个字段的名称符合我们的命名规则。我们可以参考ECS。通过alias的使用，可以使得我们的字段根据符合ECS标准。一个字段的alias只能有一个目标字段。

在使用alias时，字段别名的目标有一些限制：

它必须是一个具体的字段（不是一个对象或者是另外一个alias）
它必须在alias被创建时已经存在
如果是一个nested的对象，那么alias必须具有和它的目标具有同样的nested scope
例子 一
下面，我们来用一个具体的例子来说说明。我们首先来定义一个index的mapping如下：

PUT trips
{
  "mappings": {
    "properties": {
      "distance": {
        "type": "long"
      },
      "route_length_miles": {
        "type": "alias",
        "path": "distance" 
      },
      "transit_mode": {
        "type": "keyword"
      }
    }
  }
}
现在我们输入一下的两个文档，并搜索：

PUT trips/_doc/1
{
  "distance": 100,
  "transit_mode": "mode1"
}
 
PUT trips/_doc/2
{
  "distance": 50,
  "transit_mode": "mode2"
}
 
GET _search
{
  "query": {
    "range" : {
      "route_length_miles" : {
        "gte" : 60
      }
    }
  }
}

{
  "took" : 29,
  "timed_out" : false,
  "_shards" : {
    "total" : 46,
    "successful" : 46,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 1,
      "relation" : "eq"
    },
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "trips",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 1.0
      }
    ]
  }
}


PUT logs_server
{
  "mappings": {
    "properties": {
      "http": {
        "properties": {
          "request": {
            "properties": {
              "method": {
                "type": "alias",
                "path": "method.keyword"
              }
            }
          }
        }
      },
      "method": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword"
          }
        }
      }
    }
  }
}
在上面，我们定义了两个字段，其中的一个字段是 alias:

method.keyword
http.request.method
其中 http.request.method 被定义为 alias 指向 method.keyword。运行上面的指令，并执行如下的操作：

PUT logs_server/_doc/1
{
  "method": "GET"
}
我们可以通过如的方法来进行搜索：

GET logs_server/_search
{
  "query": {
    "match": {
      "http.request.method": "GET"
    }
  }
}


不被支持的API
不支持写入字段别名：尝试在索引或更新请求中使用别名将导致失败。 同样，别名不能用作copy_to的目标或多字段。

由于文档源中不存在别名，因此在执行源过滤时不能使用别名。 例如，以下请求将返回_source的空结果：
 

不支持写入字段别名：尝试在索引或更新请求中使用别名将导致失败。 同样，别名不能用作copy_to的目标或多字段。
 
由于文档源中不存在别名，因此在执行源过滤时不能使用别名。 例如，以下请求将返回_source的空结果：