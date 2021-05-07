1. match 查询
无论你在任何字段上进行的是全文搜索还是精确查询，match 查询是你可用的标准查询。

如果你在一个全文字段上使用 match 查询，在执行查询前，它将用正确的分析器去分析查询字符串

2. multi_match 查询
multi_match 查询可以在多个字段上执行相同的 match 查询：

3.range 查询
range 查询找出那些落在指定区间内的数字或者时间：

4.term 查询
term 查询被用于精确值匹配，这些精确值可能是数字、时间、布尔或者那些 not_analyzed 的字符串：
5. terms 查询
terms 查询和 term 查询一样，但它允许你指定多值进行匹配。如果这个字段包含了指定值中的任何一个值，那么这个文档满足条件

6. exists 查询和 missing 查询
exists 查询和 missing 查询被用于查找那些指定字段中有值 (exists) 或无值 (missing) 的文档。这与SQL中的 IS_NULL (missing) 和 NOT IS_NULL (exists) 在本质上具有共性：


https://rourou.blog.csdn.net/article/details/108003267