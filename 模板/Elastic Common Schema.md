https://elasticstack.blog.csdn.net/article/details/108916961

如果，我们在一开始就遵循 Elastic Common Schema，那么我们就不会有任何的问题。但是在实际的生产环境中，有可能在一开始我们就没有这么做，那我们该如何解决这个问题呢？比如我们有如下的两个数据：

POST logs_server1/_doc/
{
  "level": "info"
}
 
POST logs_server2/_doc/
{
  "log_level": "info"
}
在上面的两个数据是来自两个不同的服务器，在当时设计的时候，表示 log 的级别分别用了不同的字段：level 及 log_level。显然这两个不同的字段不便于我们统计数据。安装 Elastic Common Schema 的要求，正确的字段应该是 log.level。

POST logs_server1/_doc/
{
  "level": "info"
}
 
POST logs_server2/_doc/
{
  "log_level": "info"
}


检查一下这两个索引的 mapping:

GET logs_server1/_mapping
{
  "logs_server1" : {
    "mappings" : {
      "properties" : {
        "level" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    }
  }
}
GET logs_server2/_mapping
{
  "logs_server2" : {
    "mappings" : {
      "properties" : {
        "log_level" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    }
  }
}
显然上面的两个索引的 mapping 都是不一样的。

如果我们想统计一下 logs 按照级别 level 进行统计的话，我们只能按照如下的方法来进行：

GET logs_server*/_search
{
  "size": 0,
  "aggs": {
    "levels": {
      "terms": {
        "script": {
          "source": """
             if (doc.containsKey('level.keyword')) {
               return doc['level.keyword'].value
             } else {
               return doc['log_level.keyword'].value
             }
             
          """
        }
      }
    }
  }
}
在上面，我使用了 script 来进行统计。在上面脚本中的 doc，其实就是 doc_values。
https://elasticstack.blog.csdn.net/article/details/108315145
GET logs_server*/_search
{
  "size": 0,
  "aggs": {
    "levels": {
      "terms": {
        "script": {
          "source": """
            Debug.explain(doc)
          """
        }
      }
    }
  }
}

使用 alias 数据类型把数据归一化
我们可以把 level 都按照 ECS 的要求，对应于 log.level。对上面的两个索引做如下的操作：

PUT logs_server1/_mapping
{
  "properties": {
    "log": {
      "properties": {
        "level": {
          "type": "alias",
          "path": "level.keyword"
        }
      }
    }
  }
}
GET logs_server1/_mapping

PUT logs_server2/_mapping
{
  "properties": {
    "log": {
      "properties": {
        "level": {
          "type": "alias",
          "path": "log_level.keyword"
        }
      }
    }
  }
}

经过上面的改造之后，我们可以看出来，这两个索引的 mapping 都有一个共同的字段 log.level，尽管它们是 alias 数据类型。

我们很容易使用如下的方法来对 level 进行统计了：

GET  logs_server*/_search
{
  "size": 0,
  "aggs": {
    "levels": {
      "terms": {
        "field": "log.level",
        "size": 10
      }
    }
  }
}
