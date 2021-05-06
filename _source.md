Soure filtering
你可以使用 _source 参数选择返回源的哪些字段。 这称为源过滤。

以下搜索 API 请求将 _source 请求主体参数设置为 false。 该文档来源不包含在响应中。

GET developers/_search
{
  "_source": false,
  "query": {
    "match": {
      "city": "Beijing"
    }
  }
}
要仅返回 source 字段的子集，请在 _source参 数中指定通配符（*）模式。 以下搜索 API 请求仅返回 ski 为开头的所有字段以及 name.fir 为开头的所有字段：

GET developers/_search
{
  "_source": ["ski*", "name.fir*"],
  "query": {
    "match": {
      "city": "Beijing"
    }
  }
}
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
    "max_score" : 0.6931471,
    "hits" : [
      {
        "_index" : "developers",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.6931471,
        "_source" : {
          "skills" : [
            "Java",
            "C++"
          ],
          "name" : {
            "firstname" : "san"
          }
        }
      }
    ]
  }
}
为了更好地控制，你可以在 _source 参数中指定一个includes 和  excludes 模式的数组的对象。

如果指定了 includes 属性，则仅返回与其模式之一匹配的源字段。 你可以使用 excludes 属性从此子集中排除字段。

如果未指定 include 属性，则返回整个文档源，不包括与 excludes 属性中与模式匹配的任何字段。

以下搜索 API 请求仅返回 sk* 和 name 字段及其属性的源，不包括 name.secondname 字段。

GET developers/_search
{
  "query": {
    "match_all": {}
  },
  "_source": {
    "includes": [ "sk*", "name" ],
    "excludes": [ "name.secondname"]
  }
}