GET index_test001/_mapping?pretty

{
  "index_test001" : {
    "mappings" : {
      "properties" : {
        "createDate" : {
          "type" : "date"
        },
        "description" : {
          "type" : "text",
          "index" : false
        },
        "onSale" : {
          "type" : "boolean"
        },
        "price" : {
          "type" : "double"
        },
        "title" : {
          "type" : "text",
          "store" : true
        },
        "type" : {
          "type" : "integer"
        }
      }
    }
  }
}


修改
POST index_test001/_mapping?pretty
{
     "product": {
                "properties": {
                     "amount":{
                        "type":"integer"
                   }
                }
            }
}

{
  "acknowledged" : true
}


为什么不能修改一个字段的type？原因是一个字段的类型修改以后，那么该字段的所有数据都需要重新索引。Elasticsearch底层使用的是lucene库，字段类型修改以后索引和搜索要涉及分词方式等操作，不允许修改类型在我看来是符合lucene机制的。 
POST index_test001/_mapping?pretty
{
    "product": {
            "properties": {
                    "amount":{
                    "type":"string"
                }
            }
        }
}

{
  "error" : {
    "root_cause" : [
      {
        "type" : "mapper_parsing_exception",
        "reason" : "No handler for type [string] declared on field [amount]"
      }
    ],
    "type" : "mapper_parsing_exception",
    "reason" : "No handler for type [string] declared on field [amount]"
  },
  "status" : 400
}

https://www.cnblogs.com/sandea/p/10557125.html

删除索引
如果直接DELETE my_index，这样操作elasticseach并不会释放磁盘空间
通过kibana可以看到索引还在，但是状态变成了Unknown，
按照elasticsearch清理索引后硬盘空间不会马上释放？, 正确的做法是

POST my_index/_forcemerge?only_expunge_deletes=true
DELETE my_index




修改索引的 mapping
 

Elasticsearch 号称是 schemaless，在实际所得应用中，每一个 index 都有一个相应的 mapping。这个 mapping 在我们生产第一个文档时已经生产。它是对每个输入的字段进行自动的识别从而判断它们的数据类型。我们可以这么理解 schemaless：

不需要事先定义一个相应的 mapping 才可以生产文档。字段类型是动态进行识别的。这和传统的数据库是不一样的
如果有动态加入新的字段，mapping 也可以自动进行调整并识别新加入的字段
自动识别字段有一个问题，那就是有的字段可能识别并不精确，比如对于我们例子中的位置信息。那么我们需要对这个字段进行修改。

我们可以通过如下的命令来查询目前的 index 的 mapping:

GET twitter/_mapping
{
  "twitter" : {
    "mappings" : {
      "properties" : {
        "address" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
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
        "country" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "location" : {
          "properties" : {
            "lat" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "lon" : {
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
        "message" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "province" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "uid" : {
          "type" : "long"
        },
        "user" : {
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


注意：我们不能为已经建立好的 index 动态修改 mapping。这是因为一旦修改，那么之前建立的索引就变成不能搜索的了。一种办法是 reindex 从而重新建立我们的索引。如果在之前的 mapping 加入新的字段，那么我们可以不用重新建立索引。

为了能够正确地创建我们的 mapping，我们必须先把之前的 twitter 索引删除掉，并同时使用 settings 来创建这个 index。
DELETE twitter
PUT twitter
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  }
}
PUT twitter/_mapping
{
  "properties": {
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
    "city": {
      "type": "text",
      "fields": {
        "keyword": {
          "type": "keyword",
          "ignore_above": 256
        }
      }
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
      "type": "geo_point"
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


通常，一个索引中的所有类型共享相同的字段和设置。 _default_ 映射更加方便地指定通用设置，而不是每次创建新类型时都要重复设置。 _default_ 映射是新类型的模板。在设置 _default_ 映射之后创建的所有类型都将应用这些缺省的设置，除非类型在自己的映射中明确覆盖这些设置。

https://www.elastic.co/guide/cn/elasticsearch/guide/current/default-mapping.html


elasticsearch "dynamic": "strict",
https://stackoverflow.com/questions/65166024/how-to-update-index-mapping-with-dynamic-as-strict-in-future

动态映射（dynamic：true）：动态添加新的字段（或缺省）。
静态映射（dynamic：false）：忽略新的字段。在原有的映射基础上，当有新的字段时，不会主动的添加新的映射关系，只作为查询结果出现在查询中。
严格模式（dynamic： strict）：如果遇到新的字段，就抛出异常。

https://www.cnblogs.com/Neeo/articles/10585035.html
https://www.elastic.co/guide/cn/elasticsearch/guide/current/dynamic-mapping.html


doc_values和fielddata就是用来给文档建立正排索引的。他俩一个很显著的区别是，前者的工作地盘主要在磁盘，而后者的工作地盘在内存。

对于非text字段类型，doc_values默认情况下是打开的
https://blog.csdn.net/pony_maggie/article/details/104135289

ES允许对每一个字段配置得分算法或者相似算法，similarity就可以让我们选择不同于TF/IDF的相似算法.BM25

https://blog.csdn.net/zhanglh046/article/details/78529208