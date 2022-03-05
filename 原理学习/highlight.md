ES提供了一个highlight属性，和query同级别的。

fragment_size ：指定高亮数据展示多少个字符回来；
pre_tag：指定前缀标签，如 <font color="red">
post_tags：指定后缀标签，如 </font>
field：指定那个字段为高亮字段

https://www.cnblogs.com/chengbao/p/14993673.html

ES使用highlight来实现搜索结果中一个或多个字段突出显示。

高亮显示需要字段中的内容，如果没有存储字段store=true，则加载实际的_source并从_source提取相关字段。

三种不同的高亮策略的区别

unified（通用高亮策略）
其使用的是Lucene的Unified Highlighter。此高亮策略将文本分解成句子，并使用BM25算法对单个句子进行评分，支持精确的短语和多术语(模糊、前缀、正则表达式)突出显示。这个是默认的高亮策略。

plain （普通高亮策略）
其使用的是Lucene的standard Lucene highlighter。它试图在理解词的重要性和短语查询中的任何词定位标准方面反映查询匹配逻辑。此高亮策略是和在单个字段中突出显示简单的查询匹配。如果想用复杂的查询在很多文档中突出显示很多字段，还是使用unified

Fast vector highlighter（快速向量策略）
其使用的是Lucene的Fast Vector highlighter。使用此策略需要在映射中将对应字段中属性term_vector设置为with_positions_offsets。这个策略以后会单独介绍。


https://blog.csdn.net/qq330983778/article/details/103690377

highlighter如何确定高亮内容
为了从查询的词汇中获得搜索片段位置，高亮策略显示需要知道原始文本中每个单词的起始和结束字符偏移量。目前根据模式不同获取这些数据途径不同

检索列表，如果在映射中index_options设置了offsets，unified会将其中数据应用在文档中，而不会重新分析文本。它直接对文档进行原始查询，并从索引中提取匹配的偏移数据。在字段内容很大的时候，使用此配置很重要，因为它不需要重新分析文本内容。和term_vectors相比，它还需要更少的磁盘空间。
术语向量，如果在映射中term_vector设置为with_positions_offsets则unified highlighter使用term_vector来突出显示字段。对于大字段（大于1MB）和多术语查询它的速度会比较快。而fvh highlighter总是使用term_vector。
普通的高亮策略（
Plain highlighting），当没有其他选择的时候，unified highlighter使用此模式，他在内存中创建一个小的索引（index），通过运行Lucene的查询执行计划来访问文档的匹配信息，对需要高亮显示的每个字段和每个文档进行处理。plain highlighter总是使用此策略。注意此方式在大型文本上可能需要大量的时间和内存。在使用此策略时候可以设置分析的文本字符的最大数量限制为1000000。这个数值可以通过修改索引的index.highlight.max_analyzed_offset参数来改变。

lucene支持三种高亮显示方式highlighter, fast-vector-highlighter， postings-highlighter

highlighter 高亮是缺省配置高亮方式。

highlighter 高亮也叫plain高亮
搜索引擎查询到了目标数据docid后，将需要高亮的字段数据提取到内存，再调用该字段的分析器进行处理，分析器对文本进行分析处理，分析完成后采用相似度算法计算得分最高的前n组并高亮段返回数据。

highlighter高亮器是实时分析高亮器，这种实时分析机制会让ES占用较少的IO资源同时也占用较少的存储空间（词库较全的话相比fvh方式能节省一半的存储空间），其实时计算高亮是采用cpu资源来缓解io压力，在高亮字段较短（比如高亮文章的标题）时候速度较快，同时因io访问的次数少，io压力较小，有利于提高系统吞吐量。


为解决 highlighter 高亮器质大文本字段上高亮速度跟不上的问题，lucene高亮模块提供了基于向量的高亮方式 fast-vector-highlighter（也称为fvh）。fast-vector-highlighter（fvh）高亮器利用建索引时候保存好的词向量来直接计算高亮段落，在高亮过程中比plain高亮方式少了实时分析过程，取而代之的是直接从磁盘中将分词结果直接读取到内存中进行计算。故要使用fvh的前置条件就是在建索引时候，需要配置存储词向量，词向量需要包含词位置信息、词偏移量信息。

fvh在高亮时候的逻辑如下：

    1.分析高亮查询语法，提取表达式中的高亮词集合
    2.从磁盘上读取该文档字段下的词向量集合
    3.遍历词向量集合，提取自表达式中出现的词向量
    4.根据提取到目标词向量读取词频信息，根据词频获取每个位置信息、偏移量
    5.通过相似度算法获取得分较高的前n组高亮信息
    6.读取字段内容（多字段用空格隔开），根据提取的词向量直接定位截取高亮字段

，fvh 省去了实时分析过程，但是多了词条向量信息存储和读取，在词库丰富的系统中，存储词向量往往要比不存储词向量多占用一倍的空间，同时在高亮时候会比plain高亮多出至少一倍的io操作次数，读取的字节大小也多出至少一倍，大量的io请求会让搜索引擎并发能力降低。

当实时分词速度小于磁盘读随机取速度的时候，从磁盘读取词词条向量结果的的fast-vector-highlighter高亮有明显优势，


postings-highlighter ：

         默认plain高亮方式占用空间小，但是对大字段高亮慢，fvh对大字段高亮快，但占用空间过大，为此，lucene还提供了一占用空间不是太大，高亮速度不是太慢的的折中方案-postings-highlighter（也称postings）。postings 高亮方式与fvh相似，采用词量向量的方式进行高亮，与fvh高亮不同的是postings高亮只存储了词向量的位置信息，并未存储词向量的偏移量，故中大字段存储中，postings其比fvh节省约20-30%的存储空间，速度与fvh基本相当。

https://blog.csdn.net/kjsoftware/article/details/76293204

