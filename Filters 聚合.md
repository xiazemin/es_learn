Filters 聚合
在上面，我们使用 ranges 把数据分成不同的 bucket。通常这样的方法只适合字段为数字的字段。我们按照同样的思路，可以使用 filter 来对数据进行分类。在这种方法中，我们甚至可以针对非数字字段来进行建立不同的 bucket。这类聚合我们称之为 Filter aggregagation。定义一个多存储桶聚合，其中每个存储桶都与一个过滤器相关联。 每个存储桶将收集与其关联的过滤器匹配的所有文档。我们可以使用如下的例子：

GET twitter/_search
{
  "size": 0,
  "aggs": {
    "by_cities": {
      "filters": {
        "filters": {
          "beijing": {
            "match": {
              "city": "北京"
            }
          },
          "shanghai": {
            "match": {
              "city": "上海"
            }
          }
        }
      }
    }
  }
}
在上面的例子中，我们使用 filters 来分别针对 “北京” 和 “上海” 两地的文档进行统计：

GET twitter/_search
{
  "size": 0,
  "aggs": {
    "beijing": {
      "filter": {
        "match": {
          "city": "北京"
        }
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

date_range 聚合
我们可以使用 date_range 来统计在某个时间段里的文档数：
POST twitter/_search
{
  "size": 0,
  "aggs": {
    "birth_range": {
      "date_range": {
        "field": "DOB",
        "format": "yyyy-MM-dd",
        "ranges": [
          {
            "from": "1989-01-01",
            "to": "1990-01-01"
          },
          {
            "from": "1991-01-01",
            "to": "1992-01-01"
          }
        ]
      }
    }
  }
}

