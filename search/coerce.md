https://elasticstack.blog.csdn.net/article/details/104091682

字符串将被强制转换为数字，比如 "5" 转换为整型数值5
浮点将被截断为整数值，比如 5.0 转换为整型值5
例如：

PUT my_index
{
  "mappings": {
    "properties": {
      "number_one": {
        "type": "integer"
      },
      "number_two": {
        "type": "integer",
        "coerce": false
      }
    }
  }
}
 
PUT my_index/_doc/1
{
  "number_one": "10" 
}
 
PUT my_index/_doc/2
{
  "number_two": "10" 

}

我们定义 number_one 为 integer 数据类型，但是它没有属性 coerce 为 false，那么当我们把 number_one 赋值为"10"，也就是一个字符串，那么它自动将"10"转换为整型值10。针对第二字段 number_two，它同样被定义为证型值，但是它同时也设置 coerce 为 false，也就是说当字段的值不匹配的时候，就会出现错误。

Index 级默认设置
可以在索引级别上设置 index.mapping.coerce 设置，以在所有映射类型中全局禁用强制：

PUT my_index
{
  "settings": {
    "index.mapping.coerce": false
  },
  "mappings": {
    "properties": {
      "number_one": {
        "type": "integer",
        "coerce": true
      },
      "number_two": {
        "type": "integer"
      }
    }
  }
}