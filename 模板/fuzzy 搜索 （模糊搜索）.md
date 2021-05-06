在实际的搜索中，我们有时候会打错字，从而导致搜索不到。在 Elasticsearch 中，我们可以使用fuzziness属性来进行模糊查询，从而达到搜索有错别字的情形。

match 查询具有 “fuziness” 属性。它可以被设置为 “0”， “1”， “2”或 “auto”。“auto” 是推荐的选项，它会根据查询词的长度定义距离。在实际的使用中，当我们使用 auto 时，如果字符串的长度大于5，那么 funziness 的值自动设置为2，如果字符串的长度小于2，那么 fuziness 的值自动设置为 0。

Fuzzy query
返回包含与搜索词相似的词的文档，以 Levenshtein编辑距离 测量。

编辑距离是将一个术语转换为另一个术语所需的一个字符更改的次数。 这些更改可以包括：

更改字符（box→fox）
删除字符（black→lack）
插入字符（sic→sick）
转置两个相邻字符（act→cat）
为了找到相似的词，模糊查询会在指定的编辑距离内创建搜索词的所有可能变化或扩展的集合。 查询然后返回每个扩展的完全匹配。

PUT fuzzyindex/_doc/1
{
  "content": "I like blue sky"
}

{
  "_index" : "fuzzyindex",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 1,
  "result" : "created",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 0,
  "_primary_term" : 1
}


GET fuzzyindex/_search
{
  "query": {
    "match": {
      "content": "ski"
    }
  }
}

{
  "took" : 3,
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



GET fuzzyindex/_search
{
  "query": {
    "match": {
      "content": {
        "query": "ski",
        "fuzziness": "1"
      }
    }
  }
}


{
  "took" : 152,
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
    "max_score" : 0.19178805,
    "hits" : [
      {
        "_index" : "fuzzyindex",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.19178805,
        "_source" : {
          "content" : "I like blue sky"
        }
      }
    ]
  }
}

显然是找到我们需要的结果了。这是因为 sky 和 ski 时间上是只差别一个字母。

同样，如果我们选用“auto”选项看看：

GET fuzzyindex/_search
{
  "query": {
    "match": {
      "content": {
        "query": "ski",
        "fuzziness": "auto"
      }
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
      "value" : 1,
      "relation" : "eq"
    },
    "max_score" : 0.19178805,
    "hits" : [
      {
        "_index" : "fuzzyindex",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.19178805,
        "_source" : {
          "content" : "I like blue sky"
        }
      }
    ]
  }
}


如果我们进行如下的匹配：

GET fuzzyindex/_search
{
  "query": {
    "match": {
      "content": {
        "query": "bxxe",
        "fuzziness": "auto"
      }
    }
  }
}

{
  "took" : 20,
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


那么它不能匹配任何的结果，但是，如果我们进行如下的搜索：

GET fuzzyindex/_search
{
  "query": {
    "match": {
      "content": {
        "query": "bxxe",
        "fuzziness": "2"
      }
    }
  }
}

{
  "took" : 12,
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
    "max_score" : 0.14384104,
    "hits" : [
      {
        "_index" : "fuzzyindex",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.14384104,
        "_source" : {
          "content" : "I like blue sky"
        }
      }
    ]
  }
}

我们也可以使用如下的格式：

GET /_search
{
    "query": {
        "fuzzy": {
            "content": {
                "value": "bxxe",
                "fuzziness": "2"
            }
        }
    }
}
那么它可以显示搜索的结果，这是因为我们能够容许两个编辑的错误。
我们接着再做一个试验：

GET fuzzyindex/_search
{
  "query": {
    "match": {
      "content": {
        "query": "bluo ski",
        "fuzziness": 1
      }
    }
  }
}

{
  "took" : 2,
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
    "max_score" : 0.40754962,
    "hits" : [
      {
        "_index" : "fuzzyindex",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.40754962,
        "_source" : {
          "content" : "I like blue sky"
        }
      }
    ]
  }
}


在上面的搜索中 “bluo ski”，这个词语有两个错误。我们想，是不是超过了我们定义的 "funziness": 1。其实不是的。 fuziness 为1，表示是针对每个词语而言的，而不是总的错误的数值。

在 Elasticsearch 中，有一个单独的 fuzzy 搜索，但是这个只针对一个 term 比较有用。其功能和上面的是差不多的：

GET fuzzyindex/_search
{
  "query": {
    "fuzzy": {
      "content": {
        "value": "ski",
        "fuzziness": 1
      }
    }
  }
}

模糊性是拼写错误的简单解决方案，但具有很高的 CPU 开销和非常低的精度。



