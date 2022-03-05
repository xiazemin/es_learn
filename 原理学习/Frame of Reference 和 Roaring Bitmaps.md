保存倒排索引（docid列表）ES发明了一些有意思的编码方法，包括Frame of Reference 和 Roaring Bitmaps

检索的时候，在做一些过滤操作的时候（与或非），实际上是对倒排索引的一个拉链归并操作。


为了有效地计算交集和并集，我们需要这些倒排索引是排序的。排好序的数组还可以使用delta-encoding进一步压缩


例子，拉链是[73,300,302,332,343,372]，那delta list就是[73,227,2,30,11,29]
有以下特点

所有delta都是0-255（只需要一个byte）
这是Lucene使用的在硬盘上保存倒排索引的方法：将拉链切分成blocks，每个block是256个docID，每个block单独使用上面的方式编码：Lucene计算一个block里要存储这些delta需要的空间，将这个信息加到block header，然后编码所有deltas。


搜索的时候也是用的相同的抽象：queries和filters返回一个排序迭代器，表示他们匹配的文档列表。

到底block是怎么切分的呢？
step 1、将排序的整数列表转换成delta列表
step 2、切分成blocks。具体是怎么做的呢？Lucene是规定每个block是256个delta，这里为了简化一下，搞成3个delta。
step 3、看下每个block最大的delta是多少。上图的第一个block，最大的delta是227，最接近的2次幂是256(8bits)，于是规定这个block里都用8bits来编码（看绿色的header就是8），第二个block，最大的delta是30，最接近的2次幂是32（5bits），于是规定这个block里都用5bit来编码（看绿色的header就是5）


Roaring bitmaps
在filter cache中，Lucene需要编码一个int列表。这是一个很受欢迎的技术，可以加速频繁使用的过滤器。这是一个简单的cache。

因为我们只需要cache那些常用的filter，压缩率就显得没那么重要了（xxx）
然而我们需要这些filter是比那些重复执行的filter要快，所以使用好的数据结构很重要
cached filters是保存在内存的，倒排索引是典型的保存在磁盘的


编码排序整数的方法：

选项1：Integer array
最简单的方法：直接保存成array。这样遍历很简单，然而压缩很糟糕。这个编码技术每个entry需要4个bytes，使得稠密的filters变得很消耗内存。如果你有一个segment包含100M的文档，而且有一个filter匹配了大多数文档（这儿有点抽象了，举个例子吧，拉链是“的”这个词对应的倒排索引，100M的文档全匹配上了）那么单纯地在这个segment缓存这一个filter就需要大约400MB的内存。我们还是希望有一些更好的方法来缓存这个稠密的sets

选项2：bitmap
当数字列表很稠密的时候，bitmaps就很好使了。还是刚才这个例子，100M的文档，用bitmap那就是100M/8=12.5MB。

选项3：roaring bitmaps
这家伙想要融合上面两种选项的优势。首先根据数字的高16位把拉链切分成不同的块（block）。

这就意味着，第一个block我们会编码0-65535的值，第二个block是65536-131071，以此类推。

然后在每个block我们分别编码低16位：如果列表数量小于4096，我们会使用选项一（Integer array），否则使用bitmap。

需要注意的是，使用Integer array的时候每个值实际上我们一般需要4个byte。但是在这里我们只需要使用2个byte因为block ID已经说明了高16位是啥

因为每个block的大小是65536(2^16)，如果是bitmap的话，需要65536/8 byte=8192 byte（不论实际的文档数有多少都需要这么多）

而如果是用Integer array的话，则真正包含的文档数越多的话，需要的内存越大，是一个线性增长的过程。

roaring bitmap有很多feature，但是只有两个是在Lucene中使用的：

遍历所有匹配文档。如果你在一个cached filter上运行constant_score query就会用到
在集合中找到第一个大于等于某个数字的doc id。典型的应用：filter取交集。

https://blog.csdn.net/waltonhuang/article/details/107397028

https://www.elastic.co/cn/blog/frame-of-reference-and-roaring-bitmaps

Roaring Bitmaps
ES会缓存频率比较高的filter查询，其中的原理也比较简单，即生成(fitler, segment)和id列表的映射，但是和倒排索引不同，我们只把常用的filter缓存下来而倒排索引是保存所有的，并且filter缓存应该足够快，不然直接查询不就可以了。ES直接把缓存的filter放到内存里面，映射的posting list放入磁盘中。

Roaring Bitmap是由int数组和bitmap这两个数据结构改良过的成果——int数组速度快但是空间消耗大，bitmap相对来说空间消耗小但是不管包含多少文档都需要12.5MB的空间，即使只有一个文件也要12.5MB的空间，这样实在不划算，所以权衡之后就有了下面的Roaring Bitmap。

Roaring Bitmap首先会根据每个id的高16位分配id到对应的block里面，比如第一个block里面id应该都是在0到65535之间，第二个block的id在65536和131071之间
对于每一个block里面的数据，根据id数量分成两类
如果数量小于4096，就是用short数组保存
数量大于等于4096，就使用bitmap保存

https://zhuanlan.zhihu.com/p/137574234?utm_source=wechat_session
https://blog.csdn.net/weixin_33746819/article/details/112219052