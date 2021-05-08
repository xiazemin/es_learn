ES笔记七：filter和match的区别
filter与query对比大解密

filter，仅仅只是按照搜索条件过滤出需要的数据而已，不计算任何相关度分数，对相关度没有任何影响
query，会去计算每个document相对于搜索条件的相关度，并按照相关度进行排序

一般来说，如果你是在进行搜索，需要将最匹配搜索条件的数据先返回，那么用query；如果你只是要根据一些条件筛选出一部分数据，不关注其排序，那么用filter
除非是你的这些搜索条件，你希望越符合这些搜索条件的document越排在前面返回，那么这些搜索条件要放在query中；如果你不希望一些搜索条件来影响你的document排序，那么就放在filter中即可

filter与query性能

filter，不需要计算相关度分数，不需要按照相关度分数进行排序，同时还有内置的自动cache最常使用filter的数据
query，相反，要计算相关度分数，按照分数进行排序，而且无法cache结果

https://www.cnblogs.com/lovezhr/p/14438224.html


https://www.cnblogs.com/dongma/p/13611224.html

{
    "query":{
        "bool":{
            "filter":[
                {"term":
                    {"itemID":"id100124"}
                }
            ]
        }
    }
}

{"term":
{"itemID":"id100124"}
}
之间是且的关系

https://www.cnblogs.com/zhangchenliang/p/4209011.html

