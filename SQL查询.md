SQL查询
对于与很多已经习惯用 RDMS 数据库的工作人员，他们更喜欢使用 SQL 来进行查询。Elasticsearch 也对 SQL 有支持：

GET /_sql?
{
  "query": """
    SELECT * FROM twitter 
    WHERE age = 30
  """
}
通过这个查询，我们可以找到所有在年龄等于30的用户。在个搜索中，我们使用了SQL语句。利用 SQL 端点我们可以很快地把我们的 SQL 知识转化为 Elasticsearch 的使用场景中来。我们可以通过如下的方法得到它对应的 DSL 语句：

GET /_sql/translate
{
  "query": """
    SELECT * FROM twitter 
    WHERE age = 30
  """
}
我们得到的结果是：

{
  "size" : 1000,
  "query" : {
    "term" : {
      "age" : {
        "value" : 30,
        "boost" : 1.0
      }
    }
  },
  "_source" : {
    "includes" : [
      "address",
      "message",
      "region",
      "script.source",
      "user"
    ],
    "excludes" : [ ]
  },
  "docvalue_fields" : [
    {
      "field" : "age"
    },
    {
      "field" : "city"
    },
    {
      "field" : "country"
    },
    {
      "field" : "location"
    },
    {
      "field" : "province"
    },
    {
      "field" : "script.params.value"
    },
    {
      "field" : "uid"
    }
  ],
  "sort" : [
    {
      "_doc" : {
        "order" : "asc"
      }
    }
  ]
}
 如果你想了解更多关于Elasticsearch EQL，请参阅我的另外一篇文章 “Elasticsearch SQL介绍及实例”