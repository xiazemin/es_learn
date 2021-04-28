
PUT twitter/_doc/7
{
  "user": "张三",
  "message": "今儿天气不错啊，出去转转去",
  "uid": 2,
  "city": "北京",
  "province": "北京",
  "country": "中国",
  "address": "中国北京市海淀区",
  "location": {
    "lat": "39.970718",
    "lon": "116.325747"
  },
  "DOB": "1999-04-01"
}


{
  "_index" : "twitter",
  "_type" : "_doc",
  "_id" : "7",
  "_version" : 1,
  "result" : "created",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 6,
  "_primary_term" : 1
}


GET twitter/_search
{
  "size": 0,
  "aggs": {
    "total_missing_age": {
      "missing": {
        "field": "age"
      }
    }
  }
}

{
  "took" : 524,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 7,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [ ]
  },
  "aggregations" : {
    "total_missing_age" : {
      "doc_count" : 1
    }
  }
}


