在实际的搜索返回数据中，我们经常会用选择地返回所需要的字段或部分的 source。这在某些情况下非常有用，因为对于大规模的数据来说，返回的数据大下直接影响网路带宽的使用以及内存的使用。默认情况下，搜索响应中的每个匹配都包含文档 _source，这是在为文档建立索引时提供的整个 JSON 对象。 要检索搜索响应中的特定字段，可以使用 fields 参数：

POST my-index-000001/_search
{
  "query": {
    "match": {
      "message": "foo"
    }
  },
  "fields": ["user.id", "@timestamp"],
  "_source": false
}
fields 参数同时查询文档的 _source 和索引 mapping，以加载和返回值。因为它利用了 mapping，所以与直接引用 _source 相比，字段具有一些优点：它接受 multi-fields 和字段别名，并且还以一致的方式设置诸如日期之类的字段值的格式。

文档的 _source 存储在 Lucene 中的单个字段中。因此，即使仅请求少量字段，也必须加载并解析整个 _source 对象。为避免此限制，你可以尝试另一种加载字段的方法：

使用 docvalue_fields 参数获取选定字段的值。当返回相当少量的支持 doc 值的字段（例如关键字和日期）时，这是一个不错的选择。
使用 stored_fields 参数获取特定存储字段（使用 store 映射选项的字段）的值。
你还可以使用 script_field 参数通过脚本来转换响应中的字段值。

你可以在以下各节中找到有关这些方法的更多详细信息：

Fields
Doc value fields
Stored fields
Source filtering
Script fields
 

Fields
fields 参数允许检索搜索响应中的文档字段列表。 它同时查阅文档 _source 和索引 mapping，以符合其映射类型的标准化方式返回每个值。 默认情况下，日期字段是根据其 mapping 中的日期格式参数设置格式的。 你还可以使用 fields 参数来检索运行时字段值。我们使用如下的文档作为例子：

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
在上面，我们创建了一个叫做 developers 的索引。其中的 DOB 字段指的是 date of birth。上面的 developer 的 mapping 是：

GET developers/_mapping

{
  "developers" : {
    "mappings" : {
      "properties" : {
        "DOB" : {
          "type" : "date"
        },
        "age" : {
          "type" : "long"
        },
        "city" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "name" : {
          "properties" : {
            "firstname" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "secondname" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            }
          }
        },
        "skills" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    }
  }
}

以下搜索请求使用 fields 参数来检索 city 字段，以 skil*开头的所有字段， name.firstname 以及 DOB 字段的值：

GET developers/_search
{
  "query": {
    "match": {
      "city": "Beijing"
    }
  },
  "fields": [
    "name.firstname",
    "ski*",
    "nianling",
    "city",
    {
      "field": "DOB",
      "format": "epoch_millis"
    }
  ]
}
在上面，我也使用了一个 alias 字段 nianling。
{
  "error" : {
    "root_cause" : [
      {
        "type" : "parsing_exception",
        "reason" : "Unknown key for a START_ARRAY in [fields].",
        "line" : 7,
        "col" : 13
      }
    ],
    "type" : "parsing_exception",
    "reason" : "Unknown key for a START_ARRAY in [fields].",
    "line" : 7,
    "col" : 13
  },
  "status" : 400
}

You should be using stored_fields or _source instead of fields, i.e.

stored_fields: ['snippet.publishedAt']
or

_source: ['snippet.publishedAt']

https://stackoverflow.com/questions/54647047/unknown-key-for-a-start-array-in-fields-in-elasticsearch


https://zyc88.blog.csdn.net/article/details/83059040

如果我们不想要 _source，我们也可以直接使用如下的查询：

GET developers/_search
{
  "query": {
    "match": {
      "city": "Beijing"
    }
  },
  "fields": [
    "name.firstname",
    "ski*",
    "nianling",
    "city",
    {
      "field": "DOB",
      "format": "epoch_millis"
    }
  ],
  "_source": false
}

在上面我们看到 DOB 是以我们想要的格式进行显示的。我们也可以使用 ski* 来显示 multi-fields 字段 skills 以及 skill.keyword。fields 不允许返回整个对象。它只能返回 leaf field。

在这里特别指出的是，我们可以直接可以通过  source filtering 的方法来返回 _source 中的部分字段：

GET developers/_search
{
  "query": {
    "match": {
      "city": "Beijing"
    }
  },
  "_source": ["city", "age", "name"]
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
          "city" : "Beijing",
          "name" : {
            "secondname" : "zhang",
            "firstname" : "san"
          },
          "age" : 20
        }
      }
    ]
  }
}


fields 参数处理字段类型，例如字段 aliases 和 constant_keyword，其值并不总是出现在 _source 中。 还应其他映射选项也被考虑，包括 ignore_above，ignore_malformed 和null_value。

注意：即使 _source 中只有一个值，fields 响应也总是为每个字段返回一个值数组。 这是因为 Elasticsearch 没有专用的数组类型，并且任何字段都可以包含多个值。 fields 参数也不能保证以特定顺序返回数组值。



https://stackoverflow.com/questions/54647047/unknown-key-for-a-start-array-in-fields-in-elasticsearch

https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-get.html#docs-get-api-query-params
