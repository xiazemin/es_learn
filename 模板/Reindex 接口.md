它可以帮我们把数据从一个 index 到另外一个 index 进行重新reindex。这个对于特别适用于我们在修改我们数据的 mapping 后，需要重新把数据从现有的 index 转到新的 index 建立新的索引，这是因为我们不能修改现有的 index 的 mapping 一旦已经定下来了。

为了能够使用 reindex 接口，我们必须满足一下的条件：

_source 选项对所有的源 index 文档是启动的，也即源 index 的 source 是被存储的
reindex 不是帮我们尝试设置好目的地 index。它不拷贝源 index 的设置到目的地的 index 里去。你应该在做 reindex 之前把目的地的源的 index 设置好，这其中包括 mapping, shard 数目，replica 等

PUT twitter2/_doc/1
{
  "user" : "双榆树-张三",
  "message" : "今儿天气不错啊，出去转转去",
  "uid" : 2,
  "age" : 20,
  "city" : "北京",
  "province" : "北京",
  "country" : "中国",
  "address" : "中国北京市海淀区",
  "location" : {
    "lat" : "39.970718",
    "lon" : "116.325747"
  }
}

GET /twitter2/_mapping

系统帮我们生产的 location 数据类型是不对的，我们必须进行修改。一种办法是删除现有的 twitter2 索引，让后修改它的 mapping，再重新索引所有的数据。这对于一个两个文档还是可以的，但是如果已经有很多的数据了，这个方法并不可取。另外一种方式，是建立一个完全新的 index，使用新的 mapping 进行 reindex。

PUT twitter3
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
   "_source" : {
        "enabled" : true
    },
    "properties": {
      "address": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "age": {
        "type": "long"
      },
      "city": {
        "type": "text"
      },
      "country": {
        "type": "text"
      },
      "location": {
        "type": "geo_point"
      },
      "message": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "province": {
        "type": "text"
      },
      "uid": {
        "type": "long"
      },
      "user": {
        "type": "text"
      }
    }
  }
}


这里我们我们修改了 location 及其它的一些数据项的数据类型。运行上面的指令，我们就可以创建一个完全新的 twitter3 的 index。我们可以通过如下的命令来进行 reindex：

POST _reindex
{
  "source": {
    "index": "twitter2"
  },
  "dest": {
    "index": "twitter3"
  }
}

{
  "error" : {
    "root_cause" : [
      {
        "type" : "illegal_argument_exception",
        "reason" : "[twitter2][_doc][1] didn't store _source"
      }
    ],
    "type" : "illegal_argument_exception",
    "reason" : "[twitter2][_doc][1] didn't store _source"
  },
  "status" : 400
}

https://blog.csdn.net/lijingjingchn/article/details/105682618
https://github.com/elastic/elasticsearch/issues/22893

PUT twitter2
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
   "_source" : {
        "enabled" : true
    },
    "properties": {
      "address": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "age": {
        "type": "long"
      },
      "city": {
        "type": "text"
      },
      "country": {
        "type": "text"
      },
      "location": {
        "type": "geo_point"
      },
      "message": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "province": {
        "type": "text"
      },
      "uid": {
        "type": "long"
      },
      "user": {
        "type": "text"
      }
    }
  }
}

POST _reindex
{
  "source": {
    "index": "twitter2"
  },
  "dest": {
    "index": "twitter3"
  }
}
{
  "took" : 10,
  "timed_out" : false,
  "total" : 0,
  "updated" : 0,
  "created" : 0,
  "deleted" : 0,
  "batches" : 0,
  "version_conflicts" : 0,
  "noops" : 0,
  "retries" : {
    "bulk" : 0,
    "search" : 0
  },
  "throttled_millis" : 0,
  "requests_per_second" : -1.0,
  "throttled_until_millis" : 0,
  "failures" : [ ]
}


GET /twitter3/_search
{
  "took" : 1,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
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
        "_index" : "twitter3",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 1.0,
        "_source" : {
          "user" : "双榆树-张三",
          "message" : "今儿天气不错啊，出去转转去",
          "uid" : 2,
          "age" : 20,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市海淀区",
          "location" : {
            "lat" : "39.970718",
            "lon" : "116.325747"
          }
        }
      }
    ]
  }
}


Reindex执行
Reindex 是一个时间点的副本
就像上面返回的结果显示的那样，它是以 batch（批量）的方式来执行的。默认的批量大小为1000
你也可以只拷贝源 index 其中的一部分数据
         -  通过加入 query 到 source 中

         -  通过定义 max_docs 参数


POST _reindex
{
  "max_docs": 100,
  "source": {
    "index": "twitter2",
    "query": {
      "match": {
        "city": "北京"
      }
    }
  },
  "dest": {
    "index": "twitter3"
  }
}

{
  "took" : 67,
  "timed_out" : false,
  "total" : 1,
  "updated" : 1,
  "created" : 0,
  "deleted" : 0,
  "batches" : 1,
  "version_conflicts" : 0,
  "noops" : 0,
  "retries" : {
    "bulk" : 0,
    "search" : 0
  },
  "throttled_millis" : 0,
  "requests_per_second" : -1.0,
  "throttled_until_millis" : 0,
  "failures" : [ ]
}


这里，我们定义最多不超过100个文档，同时，我们只拷贝来自“北京”的 twitter 记录。

设置 op_type to create 将导致 _reindex 仅在目标索引中创建缺少的文档。 所有现有文档都会导致版本冲突，比如：

POST _reindex
{
  "source": {
    "index": "twitter2"
  },
  "dest": {
    "index": "twitter3",
    "op_type": "create"
  }
}
如果我们之前已经做过 reindex，那么我们可以看到如下的结果：

{
  "took" : 103,
  "timed_out" : false,
  "total" : 1,
  "updated" : 0,
  "created" : 0,
  "deleted" : 0,
  "batches" : 1,
  "version_conflicts" : 1,
  "noops" : 0,
  "retries" : {
    "bulk" : 0,
    "search" : 0
  },
  "throttled_millis" : 0,
  "requests_per_second" : -1.0,
  "throttled_until_millis" : 0,
  "failures" : [
    {
      "index" : "twitter3",
      "type" : "_doc",
      "id" : "1",
      "cause" : {
        "type" : "version_conflict_engine_exception",
        "reason" : "[1]: version conflict, document already exists (current version [2])",
        "index_uuid" : "g9kstkUAQAi5j-KcgjAEWg",
        "shard" : "0",
        "index" : "twitter3"
      },
      "status" : 409
    }
  ]
}


它表明我们之前的文档 id 为1的有版本上的冲突。

默认情况下，版本冲突会中止 _reindex 进程。 “conflict” 请求 body 参数可用于指示 _reindex 继续处理版本冲突的下一个文档。 请务必注意，其他错误类型的处理不受 “conflict” 参数的影响。 当 “conflict”：在请求正文中设置 “proceed” 时， _reindex 进程将继续发生版本冲突并返回遇到的版本冲突计数：

POST _reindex
{
  "conflicts": "proceed",
  "source": {
    "index": "twitter"
  },
  "dest": {
    "index": "new_twitter",
    "op_type": "create"
  }
}


{
  "error" : {
    "root_cause" : [
      {
        "type" : "illegal_argument_exception",
        "reason" : "[twitter][_doc][1] didn't store _source"
      }
    ],
    "type" : "illegal_argument_exception",
    "reason" : "[twitter][_doc][1] didn't store _source"
  },
  "status" : 400
}


POST _reindex
{
  "conflicts": "proceed",
  "source": {
    "index": "twitter2"
  },
  "dest": {
    "index": "twitter3",
    "op_type": "create"
  }
}

{
  "took" : 11,
  "timed_out" : false,
  "total" : 1,
  "updated" : 0,
  "created" : 0,
  "deleted" : 0,
  "batches" : 1,
  "version_conflicts" : 1,
  "noops" : 0,
  "retries" : {
    "bulk" : 0,
    "search" : 0
  },
  "throttled_millis" : 0,
  "requests_per_second" : -1.0,
  "throttled_until_millis" : 0,
  "failures" : [ ]
}



Throttling
重新索引大量文档可能会使你的群集泛滥甚至崩溃。requests_per_second 限制索引操作速率。

POST _reindex?requests_per_second=500 
{
  "source": {
    "index": "blogs",
    "size": 500
  },
  "dest": {
    "index": "blogs_fixed"
  }
}
 

运用 index 别名来进行 reindex
我们可以通过如下的方法来实现从一个 index 到另外一个 index 的数据转移：
PUT test     
PUT test_2   
POST /_aliases
{
    "actions" : [
        { "add":  { "index": "test_2", "alias": "test" } },
        { "remove_index": { "index": "test" } }  
    ]

}

从远处进行 reindex
_reindex 也支持从一个远处的 Elasticsearch 的服务器进行 reindex，它的语法为：

POST _reindex
{
  "source": {
    "index": "blogs",
    "remote": {
      "host": "http://remote_cluster_node1:9200",
      "username": "USERNAME",
      "password": "PASSWORD"
    }
  },
  "dest": {
    "index": "blogs"
  }
}
这里显示它从一个在 http://remote_cluster_node1:9200 的服务器来拷贝文件从一个 index 到另外一个 index。

Update by Query
虽然这个不在我们的 reindex 介绍范围，但是在有些情况下，我们可以可以通过 _update_by_query API 来让我们轻松地更新一个字段的值：

POST blogs_fixed/_update_by_query
{
  "query": {
    "match": {
      "category.keyword": ""
    }
  },
  "script": {
    "source": """
       ctx._source['category'] = "None"
     """
  }
}
在上面，把 category.keyword 项为空的所有文档的 category 通过脚本设置为默认的 "None" 字符串。它和 reindex 的作用相似。

 
为 mapping 添加新的 mulit-field
假设我们要向 twitter_new 索引的 mapping 添加一个多字段（multi-field）

具体来说，假设我们要用新的方法分析 “content” 字段
PUT new_new/_mapping
{
  "properties": {
    "content": {
      "type": "text",
      "fields": {
        "english": {
          "type": "text",
          "analyzer": "english"
        }
      }
    }
  }
}
在上面我们为 content 字段添加了一个新的 english 字段，并且它的 analyzer 为 english。

由于 mapping 已经发生改变，但是索引中已经有的文档没有这个新的字段 english，如果这个时候我们进行如下的搜索，将不会找到任何的结果：

GET twitter_new/_search
{
  "query": {
    "match": {
      "content.english": "performance tips"
    }
  }
}
那么我们该如何使得索引中现有的文档都有 content.english 这个字段呢？运行 _update_by_query 以拥有现有文档选择新的 “content.english” 字段：

POST twitter_new/_update_by_query
当我们完成上面的请求后，然后再执行如下的操作，将会在twitter_new 索引中搜索到想要的文档：

GET twitter_new/_search
{
  "query": {
    "match": {
      "content.english": "performance tips"
    }
  }

}
