https://elasticstack.blog.csdn.net/article/details/108347813

must 和 must_not 这样的组合，他们表示必须满足以及禁止满足的意思。我们也同时看到了 should 这个 clause，它表示如果满足就可以增加分数，但是我们看到没有 should_not 这样的表达方法。在我们的实际的搜索中这个其实也是蛮有用的，比如我想对某些不满足一定条件的查询进行加分。


GET twitter/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "age": {
              "gte": 40
            }
          }
        }
      ],
      "should": [
        {
          "bool": {
            "must_not": [
              {
                "match": {
                  "city": "上海"
                }
              }
            ]
          }
        }
      ]
    }
  }
}


{
  "took" : 26,
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
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "5",
        "_score" : 1.0
      },
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "6",
        "_score" : 1.0
      }
    ]
  }
}
