
复合查询（compound query）
 
什么是复合查询呢？如果说上面的查询是 leaf 查询的话，那么复合查询可以把很多个 leaf 查询组合起来从而形成更为复杂的查询。它一般的格式是：
POST _search
{
  "query": {
    "bool" : {
      "must" : {
        "term" : { "user" : "kimchy" }
      },
      "filter": {
        "term" : { "tag" : "tech" }
      },
      "must_not" : {
        "range" : {
          "age" : { "gte" : 10, "lte" : 20 }
        }
      },
      "should" : [
        { "term" : { "tag" : "wow" } },
        { "term" : { "tag" : "elasticsearch" } }
      ],
      "minimum_should_match" : 1,
      "boost" : 1.0
    }
  }
}
从上面我们可以看出，它是由 bool 下面的 must, must_not, should 及 filter 共同来组成的。针对我们的例子，

GET twitter/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "city": "北京"
          }
        },
        {
          "match": {
            "age": "30"
          }
        }
      ]
    }
  }
}
这个查询的是必须是 北京城市的，并且年刚好是30岁的。

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
      "value" : 2,
      "relation" : "eq"
    },
    "max_score" : 1.4823241,
    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "2",
        "_score" : 1.4823241,
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
        "_id" : "3",
        "_score" : 1.4823241,
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
  }
}
如果我们想知道为什么得出来这样的结果，我们可以在搜索的指令中加入"explained" : true。

GET twitter/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "city": "北京"
          }
        },
        {
          "match": {
            "age": "30"
          }
        }
      ]
    }
  },
  "explain": true
}
这样在我们的显示的结果中，可以看到一些一些解释：



我们的显示结果有2个。同样，我们可以把一些满足条件的排出在外，我们可以使用 must_not。

GET twitter/_search
{
  "query": {
    "bool": {
      "must_not": [
        {
          "match": {
            "city": "北京"
          }
        }
      ]
    }
  }
}
我们想寻找不在北京的所有的文档：

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
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "6",
        "_score" : 0.0,
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
      }
    ]
  }
}
我们显示的文档只有一个。他来自上海，其余的都北京的。

接下来，我们来尝试一下 should。它表述“或”的意思，也就是有就更好，没有就算了。比如：

GET twitter/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "age": "30"
          }
        }
      ],
      "should": [
        {
          "match_phrase": {
            "message": "Happy birthday"
          }
        }
      ]
    }
  }
}
这个搜寻的意思是，age 必须是30岁，但是如果文档里含有 “Hanppy birthday”，相关性会更高，那么搜索得到的结果会排在前面：

{
  "took" : 8,
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
    "max_score" : 2.641438,
    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "3",
        "_score" : 2.641438,
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
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "2",
        "_score" : 1.0,
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
      }
    ]
  }
}
在上面的结果中，我们可以看到：同样是年龄30岁的两个文档，第一个文档由于含有 “Happy birthday” 这个字符串在 message 里，所以它的结果是排在前面的，相关性更高。我们可以从它的 _score 中可以看出来。第二个文档里 age 是30，但是它的 message 里没有 “Happy birthday” 字样，但是它的结果还是有显示，只是得分比较低一些。

在使用上面的复合查询时，bool 请求通常是 must，must_not, should 及 filter 的一个或其中的几个一起组合形成的。我们必须注意的是：

查询类型对 hits 及 _score 的影响
Clause	影响 #hits	影响 _score
must	Yes	Yes
must_not	Yes	No
should	No*	Yes
filter	Yes	No
如上面的表格所示，should 只有在特殊的情况下才会影响 hits。在正常的情况下它不会影响搜索文档的个数。那么在哪些情况下会影响搜索的结果呢？这种情况就是针对只有 should 的搜索情况，也就是如果你在 bool query 里，不含有 must, must_not 及 filter 的情况下，一个或更多的 should 必须有一个匹配才会有结果，比如：

GET twitter/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "match": {
            "city": "北京"
          }
        },
        {
          "match": {
            "city": "武汉"
          }
        }
      ]
    }
  }
}
上面的查询显示结果为：

  "hits" : {
    "total" : {
      "value" : 5,
      "relation" : "eq"
    },
    "max_score" : 0.48232412,
    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.48232412,
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
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "2",
        "_score" : 0.48232412,
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
  ...
}
在这种情况下，should 是会影响查询的结果的。