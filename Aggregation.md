Bucketing
构建存储桶的一系列聚合，其中每个存储桶与密钥和文档标准相关联。执行聚合时，将在上下文中的每个文档上评估所有存储桶条件，并且当条件匹配时，文档被视为“落入”相关存储桶。在聚合过程结束时，我们最终会得到一个桶列表 - 每个桶都有一组“属于”它的文档。

Metric
聚合可跟踪和计算一组文档的指标。

Martrix
一系列聚合，它们在多个字段上运行，并根据从请求的文档字段中提取的值生成矩阵结果。与度量标准和存储区聚合不同，此聚合系列尚不支持脚本。

Pipeline
聚合其他聚合的输出及其关联度量的聚合


由于每个存储桶( bucket )有效地定义了一个文档集（属于该 bucket 的所有文档），因此可以在 bucket 级别上关联聚合，并且这些聚合将在该存储桶的上下文中执行。这就是聚合的真正力量所在：聚合可以嵌套！

注意一：bucketing聚合可以具有子聚合（bucketing 或 metric）。 将为其父聚合生成的桶计算子聚合。 嵌套聚合的级别/深度没有硬性限制（可以在“父”聚合下嵌套聚合，“父”聚合本身是另一个更高级聚合的子聚合）。

注意二：聚合可以操作于 double 类型的上限的数据。 因此，当在绝对值大于2 ^ 53的 long 上运行时，结果可能是近似的。

Aggregation 请求是搜索 API 的一部分，它可以带有一个 query 的结构或者不带。


DELETE twitter
 
PUT twitter
{
  "mappings": {
    "properties": {
      "DOB": {
        "type": "date"
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
      "city": {
        "type": "keyword"
      },
      "country": {
        "type": "keyword"
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
        "type": "keyword"
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


POST _bulk
{"index":{"_index":"twitter","_id":1}}
{"user":"张三","message":"今儿天气不错啊，出去转转去","uid":2,"age":20,"city":"北京","province":"北京","country":"中国","address":"中国北京市海淀区","location":{"lat":"39.970718","lon":"116.325747"}, "DOB": "1999-04-01"}
{"index":{"_index":"twitter","_id":2}}
{"user":"老刘","message":"出发，下一站云南！","uid":3,"age":22,"city":"北京","province":"北京","country":"中国","address":"中国北京市东城区台基厂三条3号","location":{"lat":"39.904313","lon":"116.412754"}, "DOB": "1997-04-01"}
{"index":{"_index":"twitter","_id":3}}
{"user":"李四","message":"happy birthday!","uid":4,"age":25,"city":"北京","province":"北京","country":"中国","address":"中国北京市东城区","location":{"lat":"39.893801","lon":"116.408986"}, "DOB": "1994-04-01"}
{"index":{"_index":"twitter","_id":4}}
{"user":"老贾","message":"123,gogogo","uid":5,"age":30,"city":"北京","province":"北京","country":"中国","address":"中国北京市朝阳区建国门","location":{"lat":"39.718256","lon":"116.367910"}, "DOB": "1989-04-01"}
{"index":{"_index":"twitter","_id":5}}
{"user":"老王","message":"Happy BirthDay My Friend!","uid":6,"age":26,"city":"北京","province":"北京","country":"中国","address":"中国北京市朝阳区国贸","location":{"lat":"39.918256","lon":"116.467910"}, "DOB": "1993-04-01"}
{"index":{"_index":"twitter","_id":6}}
{"user":"老吴","message":"好友来了都今天我生日，好友来了,什么 birthday happy 就成!","uid":7,"age":28,"city":"上海","province":"上海","country":"中国","address":"中国上海市闵行区","location":{"lat":"31.175927","lon":"121.383328"}, "DOB": "1991-04-01"}


并不是所有的字段都可以做聚合的。一般来说，具有 keyword 或者数值类型的字段是可以做聚合的。我们可以通过 _field_caps 接口来进行查询：

GET twitter/_field_caps?fields=country
{
  "indices" : [
    "twitter"
  ],
  "fields" : {
    "country" : {
      "keyword" : {
        "type" : "keyword",
        "searchable" : true,
        "aggregatable" : true
      }
    }
  }
}

http://localhost:5601/app/management/kibana/indexPatterns/patterns/c0b02770-a7f8-11eb-aaf7-4b8fcce27efc#/?_a=(tab:indexedFields)



聚合的语法是这样的：

"aggregations" : {
    "<aggregation_name>" : {
        "<aggregation_type>" : {
            <aggregation_body>
        }
        [,"meta" : {  [<meta_data_body>] } ]?
        [,"aggregations" : { [<sub_aggregation>]+ } ]?
    }
    [,"<aggregation_name_2>" : { ... } ]*
}
通常，我们也可以使用 aggs 来代替上面的 “aggregations”。


range聚合
GET twitter/_search
{
  "size": 0,
  "aggs": {
    "age": {
      "range": {
        "field": "age",
        "ranges": [
          {
            "from": 20,
            "to": 22
          },
          {
            "from": 22,
            "to": 25
          },
          {
            "from": 25,
            "to": 30
          }
        ]
      }
    }
  }
}

我们可以在 bucket 聚合之下，做 sub-aggregation：
GET twitter/_search
{
  "size": 0,
  "aggs": {
    "age": {
      "range": {
        "field": "age",
        "ranges": [
          {
            "from": 20,
            "to": 22
          },
          {
            "from": 22,
            "to": 25
          },
          {
            "from": 25,
            "to": 30
          }
        ]
      },
      "aggs": {
        "avg_age": {
          "avg": {
            "field": "age"
          }
        }
      }
    }
  }
}


GET twitter/_search
{
  "size": 0,
  "aggs": {
    "age": {
      "range": {
        "field": "age",
        "ranges": [
          {
            "from": 20,
            "to": 22
          },
          {
            "from": 22,
            "to": 25
          },
          {
            "from": 25,
            "to": 30
          }
        ]
      },
      "aggs": {
        "avg_age": {
          "avg": {
            "field": "age"
          }
        },
        "min_age": {
          "min": {
            "field": "age"
          }
        },
        "max_age": {
          "max": {
            "field": "age"
          }
        }
      }
    }
  }
}



