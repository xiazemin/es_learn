query 可以帮我们进行全文搜索，而  aggregation 可以帮我们对数据进行统计及分析

GET /_search
在这里我们没有指定任何index，我们将搜索在该 cluster 下的所有的 index。目前默认的返回个数是10个，除非我们设定 size:

GET /_search?size=20

上面的命令也等同于：

GET /_all/_search
我们也可以这样对多个 index 进行搜索：

POST /index1,index2,index3/_search
上面，表明，我们可以针对 index1，index2，index3 索引进行搜索。当然，我们甚至也可以这么写：

POST /index*,-index3/_search
上面表明，我们可以针对所有以 index 为开头的索引来进行搜索，但是排除 index3 索引。

如果我们只想搜索我们特定的 index，比如 twitter，我们可以这么做：

GET twitter/_search

_score 的项。它表示我们搜索结果的相关度。这个分数值越高，表明我们搜索匹配的相关度越高。在默认没有 sort 的情况下，所有搜索的结果的是按照分数由大到小来进行排列的。

在默认的情况下，我们可以得到10个结果。我们可以通过设置size参数得到我们想要的个数。同时，我们可以也配合 from 来进行分页。

GET twitter/_search?size=2&from=2

上面的查询类似于 DSL 查询的如下语句：

GET twitter/_search
{
  "size": 2,
  "from": 2, 
  "query": {
    "match_all": {}
  }
}
我们可以通过 filter_path 来控制输出的较少的字段，比如：

GET twitter/_search?filter_path=hits.total
{
  "hits" : {
    "total" : {
      "value" : 5,
      "relation" : "eq"
    }
  }
}


source filtering
我们可以通过 _source 来定义返回想要的字段：

GET twitter/_search
{
  "_source": ["user", "city"],
  "query": {
    "match_all": {
    }
  }
}

{
  "took" : 13,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 5,
      "relation" : "eq"
    },
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "3",
        "_score" : 1.0,
        "_source" : {
          "city" : "北京",
          "user" : "东城区-李四"
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "4",
        "_score" : 1.0,
        "_source" : {
          "city" : "北京",
          "user" : "朝阳区-老贾"
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "5",
        "_score" : 1.0,
        "_source" : {
          "city" : "北京",
          "user" : "朝阳区-老王"
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "6",
        "_score" : 1.0,
        "_source" : {
          "city" : "上海",
          "user" : "虹桥-老吴"
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "2",
        "_score" : 1.0,
        "_source" : {
          "city" : "长沙",
          "user" : "东城区-老刘"
        }
      }
    ]
  }
}


我们也可以使用如下的方法：

GET twitter/_search
{
  "_source": {
    "includes": ["user", "city"]
  },
  "query": {
    "match_all": {
    }
  }
}

我们可以看到只有 user 及 city 两个字段在 _source 里返回。我们可以可以通过设置  _source 为 false，这样不返回任何的 _source信息：

GET twitter/_search
{
  "_source": false,
  "query": {
    "match": {
      "user": "张三"
    }
  }
}

{
  "took" : 4,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 0,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  }
}


我们可以看到只有 _id 及 _score 等信息返回。其它任何的 _source 字段都没有被返回。它也可以接收通配符形式的控制，比如：

GET twitter/_search
{
  "_source": {
    "includes": [
      "user*",
      "location*"
    ],
    "excludes": [
      "*.lat"
    ]
  },
  "query": {
    "match_all": {}
  }
}


如果我们把 _source 设置为[]，那么就是显示所有的字段，而不是不显示任何字段的功能。

GET twitter/_search
{
  "_source": [],
  "query": {
    "match_all": {
    }
  }
}

有些时候，我们想要的 field 可能在 _source 里根本没有，那么我们可以使用 script field 来生成这些 field。允许为每个匹配返回script evaluation（基于不同的字段


GET twitter/_search
{
  "query": {
    "match_all": {}
  },
  "script_fields": {
    "years_to_100": {
      "script": {
        "lang": "painless",
        "source": "100-doc['age'].value"
      }
    },
    "year_of_birth":{
      "script": "2019 - doc['age'].value"
    }
  }
}

{
  "took" : 110,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 5,
      "relation" : "eq"
    },
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "3",
        "_score" : 1.0,
        "fields" : {
          "years_to_100" : [
            70
          ],
          "year_of_birth" : [
            1989
          ]
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "4",
        "_score" : 1.0,
        "fields" : {
          "years_to_100" : [
            65
          ],
          "year_of_birth" : [
            1984
          ]
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "5",
        "_score" : 1.0,
        "fields" : {
          "years_to_100" : [
            50
          ],
          "year_of_birth" : [
            1969
          ]
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "6",
        "_score" : 1.0,
        "fields" : {
          "years_to_100" : [
            10
          ],
          "year_of_birth" : [
            1929
          ]
        }
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "2",
        "_score" : 1.0,
        "fields" : {
          "years_to_100" : [
            70
          ],
          "year_of_birth" : [
            1989
          ]
        }
      }
    ]
  }
}


这种使用 script 的方法来生成查询的结果对于大量的文档来说，可能会占用大量资源。 在这里大家一定要注意的是：doc 在这里指的是 doc value。否则的话，我们需要使用 ctx._source 来做一些搜索的动作。


Count API
我们经常会查询我们的索引里到底有多少文档，那么我们可以使用_count重点来查询：

GET twitter/_count
如果我们想知道满足条件的文档的数量，我们可以采用如下的格式：

GET twitter/_count
{
  "query": {
    "match": {
      "city": "北京"
    }
  }
}

{
  "count" : 3,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  }
}


我们可以通过如下的接口来获得一个 index 的 settings

GET twitter/_settings


PUT twitter
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  }
}
一旦我们把 number_of_shards 定下来了，我们就不可以修改了，除非把 index 删除，并重新 index 它。这是因为每个文档存储到哪一个 shard 是和 number_of_shards这 个数值有关的。一旦这个数值发生改变，那么之后寻找那个文档所在的 shard 就会不准确。


