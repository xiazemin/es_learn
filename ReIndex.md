https://zhuanlan.zhihu.com/p/29360527

https://www.cnblogs.com/Ace-suiyuan008/p/9985249.html


第一步，新建索引：

curl -XPUT 'http://127.0.0.1:9200/your_index_name/' -d '{

"settings":{

"index":{

"number_of_shards":3,

"number_of_replicas":2

}

} }'

第二步，索引迁移。

curl -PUT 'http://127.0.0.1:9200/_reindex' -d '

{

"source": {

"index": "old_index"

},

"dest": {

"index": "your_index_name",

"op_type": "create"

}

}'

当es返回

{"took":55,"timed_out":false,"total":11,"updated":0,"created":11,"batches":1,"version_conflicts":0,"noops":0,"retries":0,"failures":[]}

表明迁移成功。

第三步，删除原始index。curl -XDELETE 'http://127.0.0.1:9200/old_index',删除就的索引，是为了让别名可以使用。