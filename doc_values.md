PUT users
{
    "mappings" : {
      "properties" : {
        "name" : {
          "type" : "text"
        },
        "mobile" : {
          "type" : "keyword"
        },
        "age" : {
          "type" : "integer"
        }
      }
    }
}

{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "users"
}


PUT users/_doc/1
{
  "name":"tom",
  "mobile": "15978866921",
  "age": 30
}

PUT users/_doc/2
{
  "name":"jerry",
  "mobile": "15978866920",
  "age": 35
}

PUT users/_doc/3
{
  "name":"jack",
  "mobile": "15978866922",
  "age": 20
}


POST users/_search
{
  "query": {
    "match_all": {}
  },
  "sort": [
    {
      "age": {
        "order": "desc"
      }
    }
  ]
}



DELETE users
PUT users
{
    "mappings" : {
      "properties" : {
        "name" : {
          "type" : "text"
        },
        "mobile" : {
          "type" : "keyword"
        },
        "age" : {
          "type" : "integer",
          "doc_values": false
        }
      }
    }
}


{
  "error" : {
    "root_cause" : [
      {
        "type" : "illegal_argument_exception",
        "reason" : "Can't load fielddata on [age] because fielddata is unsupported on fields of type [integer]. Use doc values instead."
      }
    ],
    "type" : "search_phase_execution_exception",
    "reason" : "all shards failed",
    "phase" : "query",
    "grouped" : true,
    "failed_shards" : [
      {
        "shard" : 0,
        "index" : "users",
        "node" : "BGxKsGwBTHiHvGSMuVmFLA",
        "reason" : {
          "type" : "illegal_argument_exception",
          "reason" : "Can't load fielddata on [age] because fielddata is unsupported on fields of type [integer]. Use doc values instead."
        }
      }
    ],
    "caused_by" : {
      "type" : "illegal_argument_exception",
      "reason" : "Can't load fielddata on [age] because fielddata is unsupported on fields of type [integer]. Use doc values instead.",
      "caused_by" : {
        "type" : "illegal_argument_exception",
        "reason" : "Can't load fielddata on [age] because fielddata is unsupported on fields of type [integer]. Use doc values instead."
      }
    }
  },
  "status" : 400
}

搜索的时候需要使用倒排索引
所以我们要查找包含brown的文档，先在词项列表中找到 brown，然后扫描所有列，可以快速找到包含 brown 的文档。

但是如果是要对搜索结果进行排序或者其它聚合操作，倒排索引这种方式就没真这么容易了，反而是类下面这种正排索引更方便。doc_values其实是Lucene在构建倒排索引时，会额外建立一个有序的正排索引（基于document => field value的映射列表）。


https://blog.csdn.net/pony_maggie/article/details/104135289

