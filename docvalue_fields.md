PUT developers/_doc/1
{
  "name": {
    "firstname": "san",
    "secondname": "zhang"
  },
  "age": 20,
  "city": "Beijing",
  "skills": ["Java", "C++"],
  "DOB": "1989-06-04"
}
 
PUT developers/_doc/2
{
  "name": {
    "firstname": "si",
    "secondname": "li"
  },
  "age": 30,
  "city": "Shanghai",
  "skills": ["Ruby", "C++"],
  "DOB": "1999-07-08"

}

GET developers/_search
{
  "query": {
    "match": {
      "city": "Beijing"
    }
  },
  "docvalue_fields": [
    "age",
    "ski*.keyword",
    {
      "field": "date",
      "format": "epoch_millis"
    }
  ]
}

{
  "took" : 479,
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
    "max_score" : 0.6931471,
    "hits" : [
      {
        "_index" : "developers",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.6931471,
        "_source" : {
          "name" : {
            "firstname" : "san",
            "secondname" : "zhang"
          },
          "age" : 20,
          "city" : "Beijing",
          "skills" : [
            "Java",
            "C++"
          ],
          "DOB" : "1989-06-04"
        },
        "fields" : {
          "age" : [
            20
          ],
          "skills.keyword" : [
            "C++",
            "Java"
          ]
        }
      }
    ]
  }
}

你不能使用 docvalue_fields 参数来检索 nested 对象的 doc value。 如果指定 nested 对象，则搜索将为该字段返回一个空数组（[]）。 要访问 nested 字段，请使用inner_hits 参数的 docvalue_fields 属性。