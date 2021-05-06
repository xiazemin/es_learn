

Doc_values
默认情况下，大多数字段都已编入索引，这使它们可搜索。反向索引允许查询在唯一的术语排序列表中查找搜索词，并从中立即访问包含该词的文档列表。

sort，aggregtion 和访问脚本中的字段值需要不同的数据访问模式。除了查找术语和查找文档外，我们还需要能够查找文档并查找其在字段中具有的术语。

Doc values 是在文档索引时构建的磁盘数据结构，这使这种数据访问模式成为可能。它们存储与 _source 相同的值，但以面向列（column）的方式存储，这对于排序和聚合而言更为有效。几乎所有字段类型都支持Doc值，但对字符串字段除外 （text 及annotated_text）。Doc values 告诉你对于给定的文档 ID，字段的值是什么。比如，当我们向Elasticsearch中加入如下的文档：

PUT cities
{
  "mappings": {
    "properties": {
      "city": {
        "type": "keyword"
      }
    }
  }
}
 
PUT cities/_doc/1
{
  "city": "Wuhan"
}
 
PUT cities/_doc/2
{
  "city": "Beijing"
}
 
PUT cities/_doc/3
{
  "city": "Shanghai"
}
那么将在在 Elasticsearch 中将创建像如下的 doc_values 的一个列存储（Columnar store）表格:

doc id	city
1	Wuhan
2	Beijing
3	Shanghai
默认情况下，所有支持 doc 值的字段均已启用它们。如果您确定不需要对字段进行排序或汇总，也不需要通过脚本访问字段值，则可以禁用 doc 值以节省磁盘空间：

比如我们可以通过如下的方式来使得 city 字段不可以做 sort 或 aggregation：

DELETE twitter
PUT twitter
{
  "mappings": {
    "properties": {
      "city": {
        "type": "keyword",
        "doc_values": false,
        "ignore_above": 256
      },
      "address": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "age": {
        "type": "long"
      },
      "country": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "location": {
        "properties": {
          "lat": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "lon": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          }
        }
      },
      "message": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "name": {
        "properties": {
          "firstname": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "surname": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          }
        }
      },
      "province": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "uid": {
        "type": "long"
      },
      "user": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      }
    }
  }
}
在上面，我们把 city 字段的 doc_values 设置为 false。

      "city": {
        "type": "keyword",
        "doc_values": false,
        "ignore_above": 256
      },
 

我们通过如下的方法来创建一个文档：

PUT twitter/_doc/1
{
  "user" : "双榆树-张三",
  "message" : "今儿天气不错啊，出去转转去",
  "uid" : 2,
  "age" : 20,
  "city" : "北京",
  "province" : "北京",
  "country" : "中国",
  "name": {
    "firstname": "三",
    "surname": "张"
  },
  "address" : [
    "中国北京市海淀区",
    "中关村29号"
  ],
  "location" : {
    "lat" : "39.970718",
    "lon" : "116.325747"
  }
}
那么，当我们使用如下的方法来进行 aggregation 时：

GET twitter/_search
{
  "size": 0,
  "aggs": {
    "city_bucket": {
      "terms": {
        "field": "city",
        "size": 10
      }
    }
  }
}
在我们的 Kibana 上我们可以看到：

{
  "error": {
    "root_cause": [
      {
        "type": "illegal_argument_exception",
        "reason": "Can't load fielddata on [city] because fielddata is unsupported on fields of type [keyword]. Use doc values instead."
      }
    ],
    "type": "search_phase_execution_exception",
    "reason": "all shards failed",
    "phase": "query",
    "grouped": true,
    "failed_shards": [
      {
        "shard": 0,
        "index": "twitter",
        "node": "IyyZ30-hRi2rnOpfx4n1-A",
        "reason": {
          "type": "illegal_argument_exception",
          "reason": "Can't load fielddata on [city] because fielddata is unsupported on fields of type [keyword]. Use doc values instead."
        }
      }
    ],
    "caused_by": {
      "type": "illegal_argument_exception",
      "reason": "Can't load fielddata on [city] because fielddata is unsupported on fields of type [keyword]. Use doc values instead.",
      "caused_by": {
        "type": "illegal_argument_exception",
        "reason": "Can't load fielddata on [city] because fielddata is unsupported on fields of type [keyword]. Use doc values instead."
      }
    }
  },
  "status": 400
}
显然，我们的操作是失败的。尽管我们不能做 aggregation 及 sort，但是我们还是可以通过如下的命令来得到它的 source：

GET twitter/_doc/1
显示结果为：

{
  "_index" : "twitter",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 1,
  "_seq_no" : 0,
  "_primary_term" : 1,
  "found" : true,
  "_source" : {
    "user" : "双榆树-张三",
    "message" : "今儿天气不错啊，出去转转去",
    "uid" : 2,
    "age" : 20,
    "city" : "北京",
    "province" : "北京",
    "country" : "中国",
    "name" : {
      "firstname" : "三",
      "surname" : "张"
    },
    "address" : [
      "中国北京市海淀区",
      "中关村29号"
    ],
    "location" : {
      "lat" : "39.970718",
      "lon" : "116.325747"
    }
  }
}
更多阅读请参阅 “Mapping parameters: doc_values”。

其实在实际的 Elasticsearch 存储中，还有一类存储。它就是 store。请详细阅读我的另外一篇文章 “Elasticsearch: 理解 mapping 中的 store 属性”。