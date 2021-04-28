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