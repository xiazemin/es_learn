百分位数（percentile）表示观察值出现一定百分比的点。 例如，第95个百分位数是大于观察值的95％的值。该聚合针对从聚合文档中提取的数值计算一个或多个百分位数。 这些值可以从文档中的特定数字字段中提取，也可以由提供的脚本生成。

百分位通常用于查找离群值。 在正态分布中，第0.13和第99.87个百分位数代表与平均值的三个标准差。 任何超出三个标准偏差的数据通常被视为异常。这在统计的角度是非常有用的。


GET twitter/_search
{
  "size": 0,
  "aggs": {
    "age_quartiles": {
      "percentiles": {
        "field": "age",
        "percents": [
          25,
          50,
          75,
          100
        ]
      }
    }
  }
}

{
  "took" : 19,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 6,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  },
  "aggregations" : {
    "age_quartiles" : {
      "values" : {
        "25.0" : 22.0,
        "50.0" : 25.5,
        "75.0" : 28.0,
        "100.0" : 30.0
      }
    }
  }
}


实际的应用中，我们有时也很希望知道满足我们的 SLA (Service Level Aggreement) 百分比是多少，这样你可以找到自己服务的差距，比如达到一个标准的百分比是多少。针对我们的例子，我们可以使用 Percentile Ranks Aggregation