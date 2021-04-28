Exists 查询
我们可以通过 exists 来查询一个字段是否存在。比如我们再增加一个文档：

PUT twitter/_doc/20
{
  "user" : "王二",
  "message" : "今儿天气不错啊，出去转转去",
  "uid" : 20,
  "age" : 40,
  "province" : "北京",
  "country" : "中国",
  "address" : "中国北京市海淀区",
  "location" : {
    "lat" : "39.970718",
    "lon" : "116.325747"
  }
}
在这个文档里，我们的 city 这一个字段是不存在的，那么一下的这个搜索将不会返回上面的这个搜索。

GET twitter/_search
{
  "query": {
    "exists": {
      "field": "city"
    }
  }
}
如果文档里只要 city 这个字段不为空，那么就会被返回。反之，如果一个文档里city这个字段是空的，那么就不会返回。

如果查询不含 city 这个字段的所有的文档，可以这样查询：

GET twitter/_search
{
  "query": {
    "bool": {
      "must_not": {
        "exists": {
          "field": "city"
        }
      }
    }
  }
}
假如我们创建另外一个索引 twitter1，我们打入如下的命令：

PUT  twitter10/_doc/1
{
  "locale": null
}
然后，我们使用如下的命令来进行查询：

GET twitter10/_search
{
  "query": {
    "exists": {
      "field": "locale"
    }
  }
}
上面查询的结果显示：

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
      "value" : 0,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  }
}
也就是没有找到。

如果你想找到一个 missing 的字段，你可以使用如下的方法：

GET twitter10/_search
{
  "query": {
    "bool": {
      "must_not": [
        {
          "exists": {
            "field": "locale"
          }
        }
      ]
    }
  }
}
上面的方法返回的数据是：

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
      "value" : 1,
      "relation" : "eq"
    },
    "max_score" : 0.0,
    "hits" : [
      {
        "_index" : "twitter1",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.0,
        "_source" : {
          "locale" : null
        }
      }
    ]
  }
}
显然这个就是我们想要的结果。
