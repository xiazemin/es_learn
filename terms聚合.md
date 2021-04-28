我们也可以通过 term 聚合来查询某一个关键字出现的频率。在如下的 term 聚合中，我们想寻找在所有的文档出现 ”Happy birthday” 里按照城市进行分类的一个聚合。

GET twitter/_search
{
  "query": {
    "match": {
      "message": "happy birthday"
    }
  },
  "size": 0,
  "aggs": {
    "city": {
      "terms": {
        "field": "city",
        "size": 10
      }
    }
  }
}

在正常的情况下，聚合是按照 doc_count 来进行排序的，也就是说哪一个 key 的 doc_count 越多，那么它就排在第一位，以后依次排序。如果你想按照 key 进行排序的话，你可以尝试如下的方法：

GET twitter/_search
{
  "size": 0,
  "aggs": {
    "top_cities": {
      "terms": {
        "field": "city",
        "order": {
          "_key": "asc"
        }
      }
    }
  }
}



GET twitter/_search
{
  "size": 0,
  "aggs": {
    "top_cities": {
      "terms": {
        "field": "city",
        "order": {
          "_count": "asc"
        }
      }
    }
  }
}

GET twitter/_search
{
  "size": 0,
  "aggs": {
    "top_cities": {
      "terms": {
        "field": "city",
        "order": {
          "avg_age": "desc"
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


POST twitter/_search
{
  "size": 0,
  "aggs": {
    "birth_year": {
      "terms": {
        "script": {
          "source": "2019 - doc['age'].value"
        }, 
        "size": 10
      }
    }
  }
}