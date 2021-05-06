处理 nested 字段
nested 字段的字段响应与常规对象字段的字段响应略有不同。 常规对象字段内的 leaf value 作为扁平列表返回时，nested 字段内的值被分组以维护原始 nested 数组内每个对象的独立性。 对于 nested 字段数组内的每个条目，除非父 nested 对象内还有其他 nested 字段，否则值将再次作为一个扁平列表返回，在这种情况下，将对较深的 nested 字段再次重复相同的过程。

给定以下 mapping，其中 user 是一个 nested 字段，在索引以下文档并检索到 user 字段下的所有字段之后：

PUT my-index-000001
{
  "mappings": {
    "properties": {
      "group" : { "type" : "keyword" },
      "user": {
        "type": "nested",
        "properties": {
          "first" : { "type" : "keyword" },
          "last" : { "type" : "keyword" }
        }
      }
    }
  }
}
 
PUT my-index-000001/_doc/1?refresh=true
{
  "group" : "fans",
  "user" : [
    {
      "first" : "John",
      "last" :  "Smith"
    },
    {
      "first" : "Alice",
      "last" :  "White"
    }
  ]
}
 
POST my-index-000001/_search
{
  "fields": ["*"],
  "_source": false
}
响应会将 first 和 last 分组，而不是将它们作为扁平列表返回。


无论用于搜索它们的模式如何，nested 字段都将按其 nested 路径进行分组。 例如，在上面的示例中仅查询 user.first 字段：

POST my-index-000001/_search
{
  "fields": ["user.first"],
  "_source": false
}
将仅返回用户的 first 字段，但仍保持 nested 用户数组的结构：
