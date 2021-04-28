
Profile API
Profile API 是调试工具。 它添加了有关执行的详细信息搜索请求中的每个组件。 它为用户提供有关搜索的每个步骤的洞察力
请求执行并可以帮助确定某些请求为何缓慢。

GET twitter/_search
{
  "profile": "true", 
  "query": {
    "match": {
      "city": "北京"
    }
  }
}
在上面，我们加上了 "profile":"true" 后，除了显示搜索的结果之外，还显示 profile 的信息：

  "profile" : {
    "shards" : [
      {
        "id" : "[ZXGhn-90SISq1lePV3c1sA][twitter][0]",
        "searches" : [
          {
            "query" : [
              {
                "type" : "BooleanQuery",
                "description" : "city:北 city:京",
                "time_in_nanos" : 1390064,
                "breakdown" : {
                  "set_min_competitive_score_count" : 0,
                  "match_count" : 5,
                  "shallow_advance_count" : 0,
                  "set_min_competitive_score" : 0,
                  "next_doc" : 31728,
                  "match" : 3337,
                  "next_doc_count" : 5,
                  "score_count" : 5,
                  "compute_max_score_count" : 0,
                  "compute_max_score" : 0,
                  "advance" : 22347,
                  "advance_count" : 1,
                  "score" : 16639,
                  "build_scorer_count" : 2,
                  "create_weight" : 342219,
                  "shallow_advance" : 0,
                  "create_weight_count" : 1,
                  "build_scorer" : 973775
                },
                "children" : [
                  {
                    "type" : "TermQuery",
                    "description" : "city:北",
                    "time_in_nanos" : 107949,
                    "breakdown" : {
                      "set_min_competitive_score_count" : 0,
                      "match_count" : 0,
                      "shallow_advance_count" : 3,
                      "set_min_competitive_score" : 0,
                      "next_doc" : 0,
                      "match" : 0,
                      "next_doc_count" : 0,
                      "score_count" : 5,
                      "compute_max_score_count" : 3,
                      "compute_max_score" : 11465,
                      "advance" : 3477,
                      "advance_count" : 6,
                      "score" : 5793,
                      "build_scorer_count" : 3,
                      "create_weight" : 34781,
                      "shallow_advance" : 18176,
                      "create_weight_count" : 1,
                      "build_scorer" : 34236
                    }
                  },
                  {
                    "type" : "TermQuery",
                    "description" : "city:京",
                    "time_in_nanos" : 49929,
                    "breakdown" : {
                      "set_min_competitive_score_count" : 0,
                      "match_count" : 0,
                      "shallow_advance_count" : 3,
                      "set_min_competitive_score" : 0,
                      "next_doc" : 0,
                      "match" : 0,
                      "next_doc_count" : 0,
                      "score_count" : 5,
                      "compute_max_score_count" : 3,
                      "compute_max_score" : 5162,
                      "advance" : 15645,
                      "advance_count" : 6,
                      "score" : 3795,
                      "build_scorer_count" : 3,
                      "create_weight" : 13562,
                      "shallow_advance" : 1087,
                      "create_weight_count" : 1,
                      "build_scorer" : 10657
                    }
                  }
                ]
              }
            ],
            "rewrite_time" : 17930,
            "collector" : [
              {
                "name" : "CancellableCollector",
                "reason" : "search_cancelled",
                "time_in_nanos" : 204082,
                "children" : [
                  {
                    "name" : "SimpleTopScoreDocCollector",
                    "reason" : "search_top_hits",
                    "time_in_nanos" : 23347
                  }
                ]
              }
            ]
          }
        ],
        "aggregations" : [ ]
      }
    ]
  }
从上面我们可以看出来，这个搜索是搜索了“北”及“京”，而不是把北京作为一个整体来进行搜索的。我们可以在以后的文档中可以学习使用中文分词器来进行分词搜索。有兴趣的同学可以把上面的搜索修改为 city.keyword 来看看。如果你对分词感兴趣的话，请参阅我的文章 “Elastic：菜鸟上手指南” 中的分词器部分。

除了上面的通过命令来进行 profile 以外，我们也可以通过 Kibana 的 UI 对我们的搜索进行 profile：