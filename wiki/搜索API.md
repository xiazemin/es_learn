多索引
Elasticsearch允许我们搜索所有索引或某些特定索引中存在的文档。例如，如果我们需要搜索名称包含“ central”的所有文档，则可以执行以下操作：

GET /_all/_search?q=city:paprola

请求正文搜索
我们还可以在请求正文中使用查询DSL来指定查询，并且在前面的章节中已经给出了很多示例。这里给出一个这样的实例-

POST /schools/_search
{
   "query":{
      "query_string":{
         "query":"up"
      }
   }
}