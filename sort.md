
POST /nba/_search
{
  "query": {
    "match": {
      "teamNameEn": "Rockets"
    }
  },
  "sort": [
    {
      "playYear": {
        "order": "desc"
      }
    }
  ]
}

https://www.cnblogs.com/juncaoit/p/12717039.html