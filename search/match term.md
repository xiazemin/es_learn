term是代表完全匹配，即不进行分词器分析，文档中必须包含整个搜索的词汇
match和term的区别是,match查询的时候,elasticsearch会根据你给定的字段提供合适的分析器,而term查询不会有分析器分析的过程，match查询相当于模糊匹配,只包含其中一部分关键词就行

https://blog.csdn.net/z8756413/article/details/85068970
https://www.elastic.co/guide/cn/elasticsearch/guide/current/_finding_multiple_exact_values.html