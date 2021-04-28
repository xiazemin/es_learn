我们可以使用 Metrics 来统计我们的数值数据，比如我们想知道所有用户的平均年龄是多少？我们可以用下面的聚合：

GET twitter/_search
{
  "size": 0,
  "aggs": {
    "average_age": {
      "avg": {
        "field": "age"
      }
    }
  }
}


POST twitter/_search
{
  "size": 0,
  "query": {
    "match": {
      "city": "北京"
    }
  },
  "aggs": {
    "average_age_beijing": {
      "avg": {
        "field": "age"
      }
    }
  }
}

 Elasticsearch 提供了一个特殊的 global 聚合，该全局全局对所有文档执行，而不受查询的影响。

POST twitter/_search
{
  "size": 0,
  "query": {
    "match": {
      "city": "北京"
    }
  },
  "aggs": {
    "average_age_beijing": {
      "avg": {
        "field": "age"
      }
    },
    "average_age_all": {
      "global": {},
      "aggs": {
        "age_global_avg": {
          "avg": {
            "field": "age"
          }
        }
      }
    }
  }
}
