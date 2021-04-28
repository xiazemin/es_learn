我们知道 Elasticsearch 可以实现秒级的搜索速度，其中很重要的一个原因就当一个文档被存储的时候，同时它也对文档的数据进行了索引（indexing）。这样在以后的搜索中，就可以变得很快。

中间的那部分就叫做 Analyzer。我们可以看出来，它分为三个部分：Char Filters, Tokenizer及 Token Filter。它们的作用分别如下：

Char Filter: 字符过滤器的工作是执行清除任务，例如剥离HTML标记。
Tokenizer: 下一步是将文本拆分为称为标记的术语。 这是由 tokenizer 完成的。 可以基于任何规则（例如空格）来完成拆分。 有关 tokennizer 的更多详细信息，请访问以下 URL：https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-tokenizers.html。
Token filter: 一旦创建了token，它们就会被传递给 token filter，这些过滤器会对 token 进行规范化。 Token filter 可以更改token，删除术语或向 token 添加术语。

Elasticsearch 已经提供了比较丰富的 analyzer。我们可以自己创建自己的 token analyzer，甚至可以利用已经有的 char filter，tokenizer 及 token filter 来重新组合成一个新的 analyzer，并可以对文档中的每一个字段分别定义自己的 analyzer。如果大家对analyzer 比较感兴趣的话，请参阅我们的网址https://www.elastic.co/guide/en/elasticsearch/reference/current/analysis-analyzers.html。

在默认的情况下，standard analyzer 是 Elasticsearch 的缺省分析器：

没有 Char Filter
使用 standard tokonizer
把字符串变为小写，同时有选择地删除一些 stop words 等。默认的情况下 stop words 为 _none_，也即不过滤任何 stop words。


https://blog.csdn.net/UbuntuTouch/article/details/99621105
