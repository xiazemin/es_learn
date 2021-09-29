逗号分隔符号
POST /index1,index2,index3/_search
{
   "query":{
      "query_string":{
         "query":"any_string"
      }
   }
}

通配符（*，+，–）
POST /school*,-schools_gov /_search
JSON对象来自所有以“ school”开头的索引，但不是来自school_gov并包含CBSE的索引。

响应过滤
通过将它们添加到field_path参数中，我们可以过滤对较少字段的响应。例如，

POST /schools/_search?filter_path = hits.total



