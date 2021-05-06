默认情况下，对字段值进行索引以使其可搜索，但不存储它们 (store)。 这意味着可以查询该字段，但是无法检索原始字段值。在这里我们必须理解的一点是: 如果一个字段的 mapping 中含有 store 属性为 true，那么有一个单独的存储空间为这个字段做存储，而且这个存储是独立于 _source 的存储的。它具有更快的查询。存储该字段会占用磁盘空间。如果需要从文档中提取（即在脚本中和聚合），它会帮助减少计算。在聚合时，具有store属性的字段会比不具有这个属性的字段快。 此选项的可能值为 false 和 true。

通常这无关紧要。 该字段值已经是 _source 字段的一部分，默认情况下已存储。 如果你只想检索单个字段或几个字段的值，而不是整个 _source 的值，则可以使用 source filtering 来实现。

在某些情况下，存储字段可能很有意义。 例如，如果你有一个带有标题，日期和很大的内容字段的文档，则可能只想检索标题和日期，而不必从较大的 _source 字段中提取这些字段。

接下来我们还是通过一个具体的例子来解释这个，虽然上面的描述有点绕口。

首先我们来创建一个叫做 my_index 的索引：

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
      }
    }
  }
}
在上面的 mapping 中，我们把 title 及 date 字段里的 store 属性设置为 true，表明有一个单独的 index fragement 是为它们而配备的，并存储它们的值。我们来写入一个文档到 my_index 索引中：

PUT my_index/_doc/1
{
  "title": "Some short title",
  "date": "2015-01-01",
  "content": "A very long content field..."

}

接下来，我们来做一个搜索：

GET my_index/_search
{
  "took" : 4,
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
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "my_index",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 1.0,
        "_source" : {
          "title" : "Some short title",
          "date" : "2015-01-01",
          "content" : "A very long content field..."
        }
      }
    ]
  }
}

以在 _source 中看到这个文档的 title，date 及 content 字段。

我们可以通过 source filtering 的方法提前我们想要的字段：

GET my_index/_search
{
  "_source": ["title", "date"]
}


{
  "took" : 1,
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
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "my_index",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 1.0,
        "_source" : {
          "date" : "2015-01-01",
          "title" : "Some short title"
        }
      }
    ]
  }
}

显然上面的结果显示我们想要的字段 date 及 title 是可以从 _source 里获取的。

我们也可以通过如下的方法来获取这两个字段的值：

GET my_index/_search
{
  "stored_fields": [
    "title",
    "date"
  ]
}

{
  "took" : 14,
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
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "my_index",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 1.0,
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

我们可以看出来在 fields 里有一个 date 及 title 的数组返回查询的结果。

也许我们很多人想知道到底这个 store 到底有什么用途呢？如果都能从 _source 里得到字段的值。

有一种就是我们在开头我们已经说明的情况：我们有时候并不想存下所有的字段在 _source 里，因为该字段的内容很大，或者我们根本就不想存 _source，但是有些字段，我们还是想要获取它们的内容。那么在这种情况下，我们就可以使用 store 来实现。

我们还是用一个例子来说明。首先创建一个叫做 my_index1 的索引：
DELETE my_index1
PUT my_index1
{
  "mappings": {
    "_source": {
      "enabled": false
    },
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
        "type": "text",
        "store": false
      }
    }
  }
}
PUT my_index1/_doc/1
{
  "title": "Some short title",
  "date": "2015-01-01",
  "content": "A very long content field..."
}


同样我们来做一个搜索：

GET my_index1/_search
{
  "query": {
    "match": {
      "content": "content"
    }
  }
}
我们可以看到搜索的结果：
{
  "took" : 4,
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
        "_index" : "my_index1",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.2876821
      }
    ]
  }
}


我们没有看到 _source 字段，这是因为我们已经把它给 disabled 了。但是我们可以通过如下的方法来获取那些store 字段：

GET my_index1/_search
{
  "stored_fields": [
    "title",
    "date"
  ],
  "query": {
    "match": {
      "content": "content"
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
      "value" : 1,
      "relation" : "eq"
    },
    "max_score" : 0.2876821,
    "hits" : [
      {
        "_index" : "my_index1",
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




