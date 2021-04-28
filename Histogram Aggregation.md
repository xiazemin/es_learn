GET twitter/_search
{
  "size": 0,
  "aggs": {
    "age_distribution": {
      "histogram": {
        "field": "age",
        "interval": 2
      }
    }
  }
}



GET twitter/_search
{
  "size": 0,
  "aggs": {
    "age_distribution": {
      "histogram": {
        "field": "age",
        "interval": 2,
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


GET twitter/_search
{
  "size": 0,
  "aggs": {
    "age_distribution": {
      "date_histogram": {
        "field": "DOB",
        "interval": "year"
      }
    }
  }
}

GET twitter/_search
{
  "size": 0,
  "aggs": {
    "number_of_cities": {
      "cardinality": {
        "field": "city"
      }
    }
  }
}

