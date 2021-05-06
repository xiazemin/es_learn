在 Elasticsearch 中，通常每个文档的每一个字段都会被存储在 shard 里存放 source 的地方

_source 里我们可以看到 Elasticsearch 为我们所存下的所有的字段。如果我们不想存储任何的字段，那么我们可以做如下的设置：

DELETE twitter
 
PUT twitter
{
  "mappings": {
    "_source": {
      "enabled": false
    }
  }
}
那么我们使用如下的命令来创建一个 id 为1的文档：

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
那么同样地，我们来查询一下这个文档：

GET twitter/_doc/1
显示的结果为：

{
  "_index" : "twitter",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 1,
  "_seq_no" : 0,
  "_primary_term" : 1,
  "found" : true
}
显然我们的文档是被找到了，但是我们看不到任何的 source。那么我们能对这个文档进行搜索吗？尝试如下的命令：

GET twitter/_search
{
  "query": {
    "match": {
      "city": "北京"
    }
  }
}
显示的结果为：

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
    "max_score" : 0.5753642,
    "hits" : [
      {
        "_index" : "twitter",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.5753642
      }
    ]
  }
}
显然这个文档 id 为1的文档可以被正确地搜索，也就是说它有完好的 inverted index 供我们查询，虽然它没有它的 source。

那么我们如何有选择地进行存储我们想要的字段呢？这种情况适用于我们想节省自己的存储空间，只存储那些我们需要的字段到source里去。我们可以做如下的设置：

DELETE twitter
 
PUT twitter
{
  "mappings": {
    "_source": {
      "includes": [
        "*.lat",
        "address",
        "name.*"
      ],
      "excludes": [
        "name.surname"
      ]
    }    
  }
}
在上面，我们使用 include 来包含我们想要的字段，同时我们通过 exclude 来去除那些不需要的字段。我们尝试如下的文档输入：

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
通过如下的命令来进行查询，我们可以看到：

GET twitter/_doc/1
结果是：

{
  "_index" : "twitter",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 1,
  "_seq_no" : 0,
  "_primary_term" : 1,
  "found" : true,
  "_source" : {
    "address" : [
      "中国北京市海淀区",
      "中关村29号"
    ],
    "name" : {
      "firstname" : "三"
    },
    "location" : {
      "lat" : "39.970718"
    }
  }
}
显然，我们只有很少的几个字段被存储下来了。通过这样的方法，我们可以有选择地存储我们想要的字段。

在实际的使用中，我们在查询文档时，也可以有选择地进行显示我们想要的字段，尽管有很多的字段被存于source中：

GET twitter/_doc/1?_source=name,location
在这里，我们只想显示和name及location相关的字段，那么显示的结果为：

{
  "_index" : "twitter",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 1,
  "_seq_no" : 0,
  "_primary_term" : 1,
  "found" : true,
  "_source" : {
    "name" : {
      "firstname" : "三"
    },
    "location" : {
      "lat" : "39.970718"
    }
  }
}
更多的阅读，可以参阅文档“Mapping meta-field: _source”
