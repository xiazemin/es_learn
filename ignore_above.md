ignore_above
对超过 ignore_above 的字符串，analyzer 不会进行处理；所以就不会索引起来。导致的结果就是最终搜索引擎搜索不到了。这个选项主要对 not_analyzed 字段有用，这些字段通常用来进行过滤、聚合和排序。而且这些字段都是结构化的，所以一般不会允许在这些字段中索引过长的项。

https://www.jianshu.com/p/133a0f49311a