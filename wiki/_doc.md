Elasticsearch 6.x

5.x 创建的索引在 6.x 版中将继续可以使用
6.x 中将只允许单个类型单个索引， 比较推荐的类型名字为 _doc, 这样可以让索引的API具备相同的路径 PUT {index}/_doc/{id} 和 POST {index}/_doc
_type 字段名称将不再与 _id 字段合并生成 _uid字段， _uid 字段将作为 _id 的别名。
新的索引将不再支持父子关系，应该采用 join 字段类进行替代
_default 映射类型将不推荐使用
Elasticsearch 7.x

URL 中的 type 参数做变为可选。例如，所以文档将不再需要 type。指定 id 的 URL 将变为 PUT {index}/_doc/{id}, 自动生成 id 的 URL 为：POST {index}/_doc
GET | PUT _mapping API 支持查询字符串参数（include_type_name），该参数指示主体是否应包含类型名称。 它默认为 true， 7.x 没有显式类型的索引将使用虚拟类型名称 _doc。

_default 映射类型将被移除

https://www.do1618.com/archives/1276/elasticsearch-%E4%B8%AD%E7%9A%84%E7%B4%A2%E5%BC%95%E4%B8%8E%E7%B1%BB%E5%9E%8B%E7%9A%84%E5%89%8D%E7%94%9F%E4%BB%8A%E4%B8%96/
