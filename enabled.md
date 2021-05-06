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

{
  "_index" : "twitter",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 2,
  "result" : "updated",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 9,
  "_primary_term" : 1
}




GET twitter/_search
{
  "query": {
    "match": {
      "user": "张三"
    }
  }
}


DELETE twitter
PUT twitter
{
  "mappings": {
    "properties": {
      "city": {
        "type": "keyword",
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
        "type": "object",
        "enabled": false
      }
    }
  }
}


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

{
  "_index" : "twitter",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 1,
  "result" : "created",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 0,
  "_primary_term" : 1
}


通过 mapping 对 user 字段进行了修改：

 "user": {
        "type": "object",
        "enabled": false
  }
也就是说这个字段将不被建立索引，我们如果使用这个字段进行搜索的话，不会产生任何的结果：

GET twitter/_search
{
  "query": {
    "match": {
      "user": "张三"
    }
  }
}

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
      "value" : 0,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  }
}


但是如果我们对这个文档进行查询的话：

GET twitter/_doc/1

user 的信息是存放于 source 里的。只是它不被我们所搜索而已。


如果我们不想我们的整个文档被搜索，我们甚至可以直接采用如下的方法：

DELETE twitter
 
PUT twitter 
{
  "mappings": {
    "enabled": false 
  }
}
那么整个 twitter 索引将不建立任何的 inverted index，那么我们通过如下的命令：
 

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
 
GET twitter/_search
{
  "query": {
    "match": {
      "city": "北京"
    }
  }
}

# GET twitter/_search
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
      "value" : 0,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  }
}


我们也可以使用如下的方式来使得我们禁止对一个字段进行查询：

{
  "mappings": {
    "properties": {
      "http_version": {
        "type": "keyword",
        "index": false
      }
     ...
    }
  }
}
上面的设置使得 http_version 不被索引。上面的 mapping 使得我们不能对 http_version 字段进行搜索，从而节省磁盘空间，但是它并不妨碍我们对该字段进行 aggregation 及对 source 的访问。


