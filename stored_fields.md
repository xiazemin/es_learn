也可以使用 store 映射选项来存储单个字段的值。 你可以使用 stored_fields 参数将这些存储的值包括在搜索响应中。

警告：stored_fields 参数用于显式标记为存储在映射中的字段，该字段默认情况下处于关闭状态，通常不建议使用。 而是使用源过滤来选择要返回的原始源文档的子集。

允许有选择地为 search hit 表示的每个文档加载特定的 store 字段。

PUT my_index
{
  "mappings": {
    "properties": {
      "title": {
        "type": "text",
        "store": true 
      },
      "date": {
        "type": "date",
        "store": true 
      },
      "content": {
        "type": "text"
      },
      "city": {
        "type": "keyword"
      }
    }
  }
}
 
PUT my_index/_doc/1
{
  "title": "Some short title",
  "date": "2015-01-01",
  "content": "A very long content field...",
  "city": "Beijing"
}
我们可以通过如下的方式来进行搜索：

GET my_index/_search
{
  "stored_fields": [
    "title",
    "date"
  ],
  "query": {
    "term": {
      "city": "Beijing"
    }
  }
}
上面的搜索结果为：

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
    "max_score" : 0.2876821,
    "hits" : [
      {
        "_index" : "my_index",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.2876821,
        "fields" : {
          "date" : [
            "2015-01-01T00:00:00.000Z"
          ],
          "title" : [
            "Some short title"
          ]
        }
      }
    ]
  }
}
在上面，我们可以使用 * 来返回所有的 stored fields:

GET my_index/_search
{
  "stored_fields": "*",
  "query": {
    "term": {
      "city": "Beijing"
    }
  }
}
空数组将导致每次匹配仅返回 _id 和 _type，例如：

GET my_index/_search
{
  "stored_fields": [],
  "query": {
    "term": {
      "city": "Beijing"
    }
  }
}
上面的搜索将返回：

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
    "max_score" : 0.2876821,
    "hits" : [
      {
        "_index" : "my_index",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.2876821
      }
    ]
  }
}
如果未设定为 store 请求的字段（将存储映射设置为false），则将忽略它们。

从文档本身获取的 store 字段值始终以数组形式返回。 相反，诸如 _routing 之类的元数据字段从不作为数组返回。

另外，只能通过 stored_fields 选项返回 leaf field。 如果指定了对象字段，它将被忽略。

注意：就其本身而言，stored_fields 不能用于加载 nested 对象中的字段 - 如果字段在其路径中包含 nested 对象，则不会为该存储字段返回任何数据。 要访问 nested 字段，必须在 inner_hits 块内使用 stored_fields。

禁止 stored fields
要完全禁用 store 字段（和元数据字段），请使用：_none_ ：

GET my_index/_search
{
  "stored_fields": "_none_",
  "query": {
    "term": {
      "city": "Beijing"
    }
  }
}