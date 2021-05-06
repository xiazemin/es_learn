GET my_index/_search
{
  "query": {
    "match": {
      "content": {
        "query": "Elastic speed scalability relevance",
        "minimum_should_match": 2
      }
    }
  }
}
GET my_index/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "match": {
            "content": "Elastic"
          }
        },
        {
          "match": {
            "content": "speed"
          }
        },
        {
          "match": {
            "content": "scalability"
          }
        },
        {
          "match": {
            "content": "relevance"
          }
        }
      ],
      "minimum_should_match": 2
    }
  }
}

两种方法查询的结果是完全一样的，而且它们的得分也是一样的。

我们可以通过 _validate API 接口来进行比较：

GET my_index/_validate/query?rewrite=true
{
  "query": {
    "bool": {
      "should": [
        {
          "match": {
            "content": "Elastic"
          }
        },
        {
          "match": {
            "content": "speed"
          }
        },
        {
          "match": {
            "content": "scalability"
          }
        },
        {
          "match": {
            "content": "relevance"
          }
        }
      ],
      "minimum_should_match": 2
    }
  }
}


{
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "failed" : 0
  },
  "valid" : true,
  "explanations" : [
    {
      "index" : "my_index",
      "valid" : true,
      "explanation" : "(content:elastic content:speed content:scalability content:relevance)~2"
    }
  ]
}
"explanation" : "(content:elastic content:speed content:scalability content:relevance)~2"
这个部分是真正要在 Apache Lucene 的部分进行查询的方法。

我们可以使用同样的方法来对 match 查询来进行验证：

GET my_index/_validate/query?rewrite=true
{
  "query": {
    "match": {
      "content": {
        "query": "Elastic speed scalability relevance",
        "minimum_should_match": 2
      }
    }
  }
}
上面的方法返回的结果是：

{
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "failed" : 0
  },
  "valid" : true,
  "explanations" : [
    {
      "index" : "my_index",
      "valid" : true,
      "explanation" : "(content:elastic content:speed content:scalability content:relevance)~2"
    }
  ]
}
从上面的结果可以看出来，这两种方法的查询的结果是完全一样的。针对 Apache Lucene 的查询完全是一样的，虽然它们的 DSL 的写法完全不同。

GET my_index/_validate/query?rewrite=true
{
  "query": {
    "query_string": {
      "default_field": "content",
      "query": "Elastic speed scalability relevance",
      "minimum_should_match": "50%"
    }
  }
}

我们也可以利用 explain 参数来对查询进行解释，比如：

GET my_index/_validate/query?explain=true
{
  "query": {
    "match": {
      "content": {
        "query": "Elastic speed scalability relevance",
        "minimum_should_match": 2
      }
    }
  }
}


如果我们不加任何的参数，我们并没有执行这个查询，只是验证一下查询是否为有效的查询：

GET my_index/_validate/query
{
  "query": {
    "match": {
      "content": {
        "query": "Elastic speed scalability relevance",
        "minimum_should_match": 2
      }
    }
  }
}
