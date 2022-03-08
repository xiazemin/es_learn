叶子查询语句（Leaf Query)
用于查询某个特定的字段，如 match , term 或 range 等


复合查询语句 (Compound query clauses)
用于合并其他的叶查询或复合查询语句，也就是说复合语句之间可以嵌套，用来表示一个复杂的单一查询


Query and filter context

一个查询语句究竟具有什么样的行为和得到什么结果，主要取决于它到底是处于查询上下文(Query Context) 还是过滤上下文(Filter Context)。两者有很大区别

Query context 查询上下文

这种语句在执行时既要计算文档是否匹配，还要计算文档相对于其他文档的匹配度有多高，匹配度越高，*_score* 分数就越高

Filter context 过滤上下文

过滤上下文中的语句在执行时只关心文档是否和查询匹配，不会计算匹配度，也就是得分。



query 参数表示整个语句是处于 query context 中

bool 和 match 语句被用在 query context 中，也就是说它们会计算每个文档的匹配度（_score)

filter 参数则表示这个子查询处于 filter context 中

filter 语句中的 term 和 range 语句用在 filter context 中，它们只起到过滤的作用，并不会计算文档的得分。


Match All Query
这个查询最简单，所有的 _score 都是 1.0。
GET /_search{"query": {"match_all": {} }}

它的反面就是 Match None Query， 匹配不到任何文档（不知道用它来做什么……）

GET /_search{"query": {"match_none": {} }}



全文查询 Full text queries

全文本查询的使用场合主要是在出现大量文字的场合，例如 email body 或者文章中搜寻出特定的内容。
match query
全文查询中最主要的查询，包括模糊查询(fuzzy matching) 或者临近查询(proximity queries)。
match_phrase query
和 match 查询比较类似，但是它会保留包含所有搜索词项，且位置与搜索词项相同的文档。
match_phrase_prefix query
是一种输入即搜索(search-as-you-type) 的查询，它和 match_phrase 比较类似，区别就是会将查询字符串的最后一个词作为前缀来使用。
multi_match query
多字段版本的 match query
common terms query
只知道是一种特殊的查询，具体干什么还不清楚，后面弄明白后会再来补充。
query_string query
支持复杂的 Lucene query String 语法，除非你是专家用户，否则不推荐使用。
simple_query_string query
简化版的 query_string ，语法更适合用户操作。






https://www.cnblogs.com/yb38156/p/13755671.html
https://zhuanlan.zhihu.com/p/50592855
https://www.iteye.com/blog/llb-code-2307801


DSL基本用法
模糊查询时使用match，精准查询时使用term。

term query：直接对关键词准确查找，该查询只适合keyword、numeric、date。

term：查询某个字段中含有某个关键词的文档。
terms：查询某个字段中含有多个关键词的文档。
match query：对所查找的关键词进行分词，在根据分词匹配查找。

match_all：查询所有文档。
multi_match：指定多个字段。
match_phrase：短语匹配查询。

https://blog.csdn.net/qq1021979964/article/details/106096726