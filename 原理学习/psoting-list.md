前面的部分叫做dictionary（字典），里面的每个单词叫做term，后面的文档列表叫做psoting-list，list中记录了所有含有该term的文档id，两个组合起来就是一个完成的倒排索引（Inverted Index

创建Inverted Index是最关键也是最耗时的过程，而且真正的Inverted Index结构也远比图中展示的复杂，不仅需要对文档进行分词（ES里中文可以自定义分词器），还要计算TF-IDF，方便评分排序（当查找you时，评分决定哪个doc显示在前面，也就是所谓的搜索排名），压缩等操作。每接收一个document，ES就会将其信息更新在倒排索引中。

https://zhuanlan.zhihu.com/p/42776873?tt_from=weixin
https://www.cnblogs.com/qiaoyihang/p/6262806.html