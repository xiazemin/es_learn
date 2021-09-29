搜索是通过使用基于JSON的查询来进行的。查询由两个子句组成-
叶子查询子句——这些子句是匹配的、术语或范围，它们在特定字段中查找特定的值。
复合查询子句—这些查询是叶查询子句和其他复合查询的组合，用于提取所需的信息。



匹配所有查询
这是最基本的查询；它返回所有内容，每个对象的得分为1.0。

POST /schools/_search
{
   "query":{
      "match_all":{}
   }
}

匹配查询
此查询将文本或短语与一个或多个字段的值匹配。

POST /schools*/_search
{
   "query":{
      "match" : {
         "rating":"4.5"
      }
   }
}

多重比对查询
此查询将一个或多个字段匹配的文本或短语匹配。

POST /schools*/_search
{
   "query":{
      "multi_match" : {
         "query": "paprola",
         "fields": [ "city", "state" ]
      }
   }
}

查询字符串查询
该查询使用查询解析器和query_string关键字。

POST /schools*/_search
{
   "query":{
      "query_string":{
         "query":"beautiful"
      }
   }
}

词级查询
这些查询主要处理结构化数据，例如数字，日期和枚举。

POST /schools*/_search
{
   "query":{
      "term":{"zip":"176115"}
   }
}


范围查询
该查询用于查找具有给定值范围之间的值的对象。为此，我们需要使用运算符，例如-

gte −大于等于

gt −大于

lte −小于等于

lt −小于


复合查询
这些查询是不同查询的集合，这些查询通过使用布尔运算符（例如和/或，或不）或针对不同的索引或具有函数调用等彼此合并。

POST /schools/_search
{
   "query": {
      "bool" : {
         "must" : {
            "term" : { "state" : "UP" }
         },
         "filter": {
            "term" : { "fees" : "2200" }
         },
         "minimum_should_match" : 1,
         "boost" : 1.0
      }
   }
}


地理查询
这些查询处理地理位置和地理位置。这些查询有助于找出学校或任何其他地理位置附近的地理对象。您需要使用地理位置数据类型。

PUT /geo_example
{
   "mappings": {
      "properties": {
         "location": {
            "type": "geo_shape"
         }
      }
   }
}



