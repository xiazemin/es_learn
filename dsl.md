DSL（Domain Specifc Lanaguage）来帮我们进行查询。


match query

GET twitter/_search
{
  "query": {
    "match": {
      "city": "北京"
    }
  }
}

也可以使用 script query来完成：

GET twitter/_search
{
  "query": {
    "script": {
      "script": {
        "source": "doc['city.keyword'].contains(params.name)",
        "lang": "painless",
        "params": {
          "name": "北京"
        }
      }
    }
  }
}

相比较而言，script query 的方法比较低效。另外，假如我们的文档是几百万或者 PB 级的数据量，那么上面的运算可能被执行无数次，那么可能需要巨大的计算量。

搜索也可以这么实现：

GET twitter/_search?q=city:"北京"

如果我们不需要这个 score，我们可以选择 filter 来完成。

GET twitter/_search
{
  "query": {
    "bool": {
      "filter": {
        "term": {
          "city.keyword": "北京"
        }
      }
    }
  }
}


这里我们使用了city.keyword。对于一些刚接触 Elasticsearch的人来说，这个可能比较陌生。正确的理解是 city 在我们的 mapping 中是一个 multi-field 项。它既是 text 也是 keyword 类型。对于一个 keyword 类型的项来说，这个项里面的所有字符都被当做一个字符串。它们在建立文档时，不需要进行 index。keyword 字段用于精确搜索，aggregation 和排序（sorting）。
所以在我们的 filter 中，我们是使用了 term 来完成这个查询。
我们也可以使用如下的办法达到同样的效果：
GET twitter/_search
{
  "query": {
    "constant_score": {
      "filter": {
        "term": {
          "city.keyword": {
            "value": "北京"
          }
        }
      }
    }
  }
}
在我们使用 match query 时，默认的操作是 OR，我们可以做如下的查询：

GET twitter/_search
{
  "query": {
    "match": {
      "user": {
        "query": "朝阳区-老贾",
        "operator": "or"
      }
    }
  }
}
上面的查询也和如下的查询是一样的：

GET twitter/_search
{
 "query": {
   "match": {
     "user": "朝阳区-老贾"
   }
 }
}
这是因为默认的操作是 or 操作。上面查询的结果是任何文档匹配：“朝”，“阳”，“区”，“老”及“贾”这5个字中的任何一个将被显示：

    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "4",
        "_score" : 4.4209847,
        "_source" : {
          "user" : "朝阳区-老贾",
          "message" : "123,gogogo",
          "uid" : 5,
          "age" : 35,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市朝阳区建国门",
          "location" : {
            "lat" : "39.718256",
            "lon" : "116.367910"
          }
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "5",
        "_score" : 2.9019678,
        "_source" : {
          "user" : "朝阳区-老王",
          "message" : "Happy BirthDay My Friend!",
          "uid" : 6,
          "age" : 50,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市朝阳区国贸",
          "location" : {
            "lat" : "39.918256",
            "lon" : "116.467910"
          }
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "2",
        "_score" : 0.8713734,
        "_source" : {
          "user" : "东城区-老刘",
          "message" : "出发，下一站云南！",
          "uid" : 3,
          "age" : 30,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市东城区台基厂三条3号",
          "location" : {
            "lat" : "39.904313",
            "lon" : "116.412754"
          }
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "6",
        "_score" : 0.4753614,
        "_source" : {
          "user" : "虹桥-老吴",
          "message" : "好友来了都今天我生日，好友来了,什么 birthday happy 就成!",
          "uid" : 7,
          "age" : 90,
          "city" : "上海",
          "province" : "上海",
          "country" : "中国",
          "address" : "中国上海市闵行区",
          "location" : {
            "lat" : "31.175927",
            "lon" : "121.383328"
          }
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "3",
        "_score" : 0.4356867,
        "_source" : {
          "user" : "东城区-李四",
          "message" : "happy birthday!",
          "uid" : 4,
          "age" : 30,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市东城区",
          "location" : {
            "lat" : "39.893801",
            "lon" : "116.408986"
          }
        }
      }
    ]
我们也可以设置参数 minimum_should_match 来设置至少匹配的 term。比如：

GET twitter/_search
{
  "query": {
    "match": {
      "user": {
        "query": "朝阳区-老贾",
        "operator": "or",
        "minimum_should_match": 3
      }
    }
  }
}
上面显示我们至少要匹配“朝”，“阳”，“区”，“老” 及 “贾” 这5个中的3个字才可以。显示结果：

    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "4",
        "_score" : 4.4209847,
        "_source" : {
          "user" : "朝阳区-老贾",
          "message" : "123,gogogo",
          "uid" : 5,
          "age" : 35,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市朝阳区建国门",
          "location" : {
            "lat" : "39.718256",
            "lon" : "116.367910"
          }
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "5",
        "_score" : 2.9019678,
        "_source" : {
          "user" : "朝阳区-老王",
          "message" : "Happy BirthDay My Friend!",
          "uid" : 6,
          "age" : 50,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市朝阳区国贸",
          "location" : {
            "lat" : "39.918256",
            "lon" : "116.467910"
          }
        }
      }
    ]
我们也可以改为 and 操作看看：

GET twitter/_search
{
  "query": {
    "match": {
      "user": {
        "query": "朝阳区-老贾",
        "operator": "and"
      }
    }
  }
}
显示的结果是：

    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "4",
        "_score" : 4.4209847,
        "_source" : {
          "user" : "朝阳区-老贾",
          "message" : "123,gogogo",
          "uid" : 5,
          "age" : 35,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市朝阳区建国门",
          "location" : {
            "lat" : "39.718256",
            "lon" : "116.367910"
          }
        }
      }
    ]
在这种情况下，需要同时匹配索引的5个字才可以。显然我们可以通过使用 and 来提高搜索的精度。

 

Ids  query
我们可以通过 id 来进行查询，比如：

GET twitter/_search
{
  "query": {
    "ids": {
      "values": ["1", "2"]
    }
  }
}
上面的查询将返回 id 为 “1” 和 “2” 的文档。

 

multi_match
在上面的搜索之中，我们特别指明一个专有的 field 来进行搜索，但是在很多的情况下，我们并胡知道是哪一个 field 含有这个关键词，那么在这种情况下，我们可以使用 multi_match 来进行搜索：

GET twitter/_search
{
  "query": {
    "multi_match": {
      "query": "朝阳",
      "fields": [
        "user",
        "address^3",
        "message"
      ],
      "type": "best_fields"
    }
  }
}
在上面，我们可以看到这个 multi_search 的 type 为 best_fields，也就是说它搜索了3个字段。最终的分数 _score 是按照得分最高的那个字段的分数为准。更多类型的定义，请在链接查看。在上面，我们可以同时对三个 fields: user，adress 及 message进行搜索，但是我们对 address 含有 “朝阳” 的文档的分数进行3倍的加权。返回的结果：

    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "5",
        "_score" : 6.1777167,
        "_source" : {
          "user" : "朝阳区-老王",
          "message" : "Happy good BirthDay My Friend!",
          "uid" : 6,
          "age" : 50,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市朝阳区国贸",
          "location" : {
            "lat" : "39.918256",
            "lon" : "116.467910"
          }
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "4",
        "_score" : 5.9349246,
        "_source" : {
          "user" : "朝阳区-老贾",
          "message" : "123,gogogo",
          "uid" : 5,
          "age" : 35,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市朝阳区建国门",
          "location" : {
            "lat" : "39.718256",
            "lon" : "116.367910"
          }
        }
      }
    ]
 

Prefix query
返回在提供的字段中包含特定前缀的文档。
 

GET twitter/_search
{
  "query": {
    "prefix": {
      "user": {
        "value": "朝"
      }
    }
  }
}
查询 user 字段里以“朝”为开头的所有文档：

    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "4",
        "_score" : 1.0,
        "_source" : {
          "user" : "朝阳区-老贾",
          "message" : "123,gogogo",
          "uid" : 5,
          "age" : 35,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市朝阳区建国门",
          "location" : {
            "lat" : "39.718256",
            "lon" : "116.367910"
          }
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "5",
        "_score" : 1.0,
        "_source" : {
          "user" : "朝阳区-老王",
          "message" : "Happy BirthDay My Friend!",
          "uid" : 6,
          "age" : 50,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市朝阳区国贸",
          "location" : {
            "lat" : "39.918256",
            "lon" : "116.467910"
          }
        }
      }
    ]

Term query 
Term query 会在给定字段中进行精确的字词匹配。 因此，您需要提供准确的术语以获取正确的结果。 

GET twitter/_search
{
  "query": {
    "term": {
      "user.keyword": {
        "value": "朝阳区-老贾"
      }
    }
  }
}
在这里，我们使用 user.keyword 来对“朝阳区-老贾”进行精确匹配查询相应的文档：

    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "4",
        "_score" : 1.5404451,
        "_source" : {
          "user" : "朝阳区-老贾",
          "message" : "123,gogogo",
          "uid" : 5,
          "age" : 35,
          "city" : "北京",
          "province" : "北京",
          "country" : "中国",
          "address" : "中国北京市朝阳区建国门",
          "location" : {
            "lat" : "39.718256",
            "lon" : "116.367910"
          }
        }
      }
    ]
Terms query
如果我们想对多个 terms 进行查询，我们可以使用如下的方式来进行查询：

GET twitter/_search
{
  "query": {
    "terms": {
      "user.keyword": [
        "双榆树-张三",
        "东城区-老刘"
      ]
    }
  }
}
上面查询 user.keyword 里含有“双榆树-张三”或“东城区-老刘”的所有文档。

 

Terms_set query
查询在提供的字段中包含最少数目的精确术语的文档。除你可以定义返回文档所需的匹配术语数之外，terms_set 查询与术语查询相同。 例如：

PUT /job-candidates
{
  "mappings": {
    "properties": {
      "name": {
        "type": "keyword"
      },
      "programming_languages": {
        "type": "keyword"
      },
      "required_matches": {
        "type": "long"
      }
    }
  }
}
 
PUT /job-candidates/_doc/1?refresh
{
  "name": "Jane Smith",
  "programming_languages": [ "c++", "java" ],
  "required_matches": 2
}
 
 
PUT /job-candidates/_doc/2?refresh
{
  "name": "Jason Response",
  "programming_languages": [ "java", "php" ],
  "required_matches": 2
}
 
GET /job-candidates/_search
{
  "query": {
    "terms_set": {
      "programming_languages": {
        "terms": [ "c++", "java", "php" ],
        "minimum_should_match_field": "required_matches"
      }
    }
  }
}
在上面，我们为 job-candiates 索引创建了两个文档。我们需要找出在 programming_languages 中同时含有 c++, java 以及 php 中至少有两个 term 的文档。在这里，我们使用了一个在文档中定义的字段 required_matches 来定义最少满足要求的 term 个数。另外一种方式是使用 minimum_should_match_script 来定义，如果没有一个专有的字段来定义这个的话：

GET /job-candidates/_search
{
  "query": {
    "terms_set": {
      "programming_languages": {
        "terms": [ "c++", "java", "php" ],
        "minimum_should_match_script": {
          "source": "2"
        }
      }
    }
  }
}
上面标示需要至少同时满足有 2 个及以上的 term。上面搜索的结果为：

{
  "took" : 0,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 2,
      "relation" : "eq"
    },
    "max_score" : 1.1005894,
    "hits" : [
      {
        "_index" : "job-candidates",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 1.1005894,
        "_source" : {
          "name" : "Jane Smith",
          "programming_languages" : [
            "c++",
            "java"
          ],
          "required_matches" : 2
        }
      },
      {
        "_index" : "job-candidates",
        "_type" : "_doc",
        "_id" : "2",
        "_score" : 1.1005894,
        "_source" : {
          "name" : "Jason Response",
          "programming_languages" : [
            "java",
            "php"
          ],
          "required_matches" : 2
        }
      }
    ]
  }
}
也就是说之前的两个文档都同时满足条件。当然如果我们使用如下的方式来进行搜索：

GET /job-candidates/_search
{
  "query": {
    "terms_set": {
      "programming_languages": {
        "terms": [ "c++", "java", "nodejs" ],
        "minimum_should_match_script": {
          "source": "2"
        }
      }
    }
  }
}
我们将看到只有一个文档是满足条件的。