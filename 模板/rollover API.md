https://elasticstack.blog.csdn.net/article/details/102670918

rollover API 使你可以根据索引大小，文档数或使用期限自动过渡到新索引。 当 rollover 触发后，将创建新索引，写别名（write alias) 将更新为指向新索引，所有后续更新都将写入新索引。

对于基于时间的 rollover 来说，基于大小，文档数或使用期限过渡至新索引是比较适合的。 在任意时间 rollover 通常会导致许多小的索引，这可能会对性能和资源使用产生负面影响。

Rollover历史数据

在大多数情况下，无限期保留历史数据是不可行的
         - 时间序列数据随着时间的流逝而失去价值，我们最终不得不将其删除

         - 但是其中一些数据对于分析仍然非常有用 

Elasticsearch 6.3 引入了一项新的 rollover 功能，该功能
         - 以紧凑的聚合格式保存旧数据

         - 仅保存您感兴趣的数据

我们定义了一个叫做 logs-alias 的alias，对于写操作来说，它总是会自动指向最新的可以用于写入index 的一个索引。针对我们上面的情况，它指向 logs-000002。如果新的 rollover 发生后，新的 logs-000003 将被生成，并对于写操作来说，它自动指向最新生产的 logs-000003 索引。而对于读写操作来说，它将同时指向最先的 logs-1，logs-000002 及 logs-000003。在这里我们需要注意的是：在我们最早设定 index 名字时，最后的一个字符必须是数字，比如我们上面显示的 logs-1。否则，自动生产 index 将会失败。


PUT /%3Clogs-%7Bnow%2Fd%7D-1%3E
{
  "aliases": {
    "log_alias": {
      "is_write_index": true
    }
  }
}

如果大家对于上面的字符串 “%3Clogs-%7Bnow%2Fd%7D-1%3E” 比较陌生的话，可以参考网站 https://www.urlencoder.io/。实际上它就是字符串 “<logs-{now/d}-1>” 的url编码形式。请注意上面的 is_write_index 必须设置为 true。


{
  "error" : {
    "root_cause" : [
      {
        "type" : "illegal_argument_exception",
        "reason" : "cannot create index with name [logs-2021.05.06-1], because it matches with template [logs] that creates data streams only, use create data stream api instead"
      }
    ],
    "type" : "illegal_argument_exception",
    "reason" : "cannot create index with name [logs-2021.05.06-1], because it matches with template [logs] that creates data streams only, use create data stream api instead"
  },
  "status" : 400
}


GET _cat/templates
ilm-history                      [ilm-history-2*]             2147483647 2       
.monitoring-beats                [.monitoring-beats-7-*]      0          7000199 
.slm-history                     [.slm-history-2*]            2147483647 2       
.ml-config                       [.ml-config]                 0          7090399 
.logstash-management             [.logstash]                  0                  
zipcodes                         [zipcodes*]                  10                 
.watches                         [.watches*]                  2147483647 11      
template_2                       [te*]                        1                  
.monitoring-alerts-7             [.monitoring-alerts-7]       0          7000199 
.ml-notifications-000001         [.ml-notifications-000001]   0          7090399 
.triggered_watches               [.triggered_watches*]        2147483647 11      
.kibana-event-log-7.9.2-template [.kibana-event-log-7.9.2-*]  0                  
.transform-notifications-000002  [.transform-notifications-*] 0          7090399 
.management-beats                [.management-beats]          0          70000   
.watch-history-11                [.watcher-history-11*]       2147483647 11      
logstash                         [logstash-*]                 0          60001   
.ml-state                        [.ml-state*]                 0          7090399 
.monitoring-es                   [.monitoring-es-7-*]         0          7000199 
logs_template                    [log_xzm-*]                  1                  
.monitoring-logstash             [.monitoring-logstash-7-*]   0          7000199 
.ml-inference-000002             [.ml-inference-000002]       0          7090399 
.ml-meta                         [.ml-meta]                   0          7090399 
.monitoring-kibana               [.monitoring-kibana-7-*]     0          7000199 
.transform-internal-005          [.transform-internal-005]    0          7090399 
.ml-anomalies-                   [.ml-anomalies-*]            0          7090399 
.ml-stats                        [.ml-stats-*]                0          7090399 
template_1                       [t*]                         0                  
metrics                          [metrics-*-*]                100        0       [metrics-mappings, metrics-settings]
logs                             [logs-*-*]                   100        0       [logs-mappings, logs-settings]
template_1                       [de*, bar*]                  200        3       [component_template1, other_component_template]



PUT /%3Cmy_logs-%7Bnow%2Fd%7D-1%3E
{
  "aliases": {
    "log_alias": {
      "is_write_index": true
    }
  }
}
运行上面的结果是：
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "my_logs-2021.05.06-1"
}

显然，它帮我们生产了一个叫做 my_logs-2021.05.06-1 的 index。


GET _cat/indices/kibana_sample_data_logs
green open kibana_sample_data_logs X-l9f22_Rbiwr-zF5yDcgw 1 0 14074 0 11.2mb 11.2mb
它显示 kibana_sample_data_logs 具有 11.1M 的数据，并且它有 14074 个文档：



我们接下来运行如下的命令：

POST _reindex
{
  "source": {
    "index": "kibana_sample_data_logs"
  },
  "dest": {
    "index": "log_alias"
  }
}

{
  "took" : 3047,
  "timed_out" : false,
  "total" : 14074,
  "updated" : 0,
  "created" : 14074,
  "deleted" : 0,
  "batches" : 15,
  "version_conflicts" : 0,
  "noops" : 0,
  "retries" : {
    "bulk" : 0,
    "search" : 0
  },
  "throttled_millis" : 0,
  "requests_per_second" : -1.0,
  "throttled_until_millis" : 0,
  "failures" : [ ]
}

这个命令的作用是把 kibana_sample_data_logs 里的数据 reindex 到 log_alias 所指向的 index。也就是把 kibana_sample_data_logs 的文档复制一份到我们上面显示的 my_logs-2021.05.06-1 索引里。我们做如下的操作查看一下结果：

GET my_logs-2021.05.06-1/_count
显示的结果是：
{
  "count" : 14074,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  }
}

显然，我们已经复制到所有的数据。那么接下来，我们来运行如下的一个指令：

POST /log_alias/_rollover?dry_run
{
  "conditions": {
    "max_age": "7d",
    "max_docs": 14000,
    "max_size": "5gb"
  }
}
在这里，我们定义了三个条件：

如果时间超过7天，那么自动 rollover，也就是使用新的 index
如果文档的数目超过 14000 个，那么自动 rollover
如果 index 的大小超过 5G，那么自动 rollover
在上面我们使用了 dry_run 参数，表明就是运行时看看，但不是真正地实施。显示的结果是：
{
  "acknowledged" : false,
  "shards_acknowledged" : false,
  "old_index" : "my_logs-2021.05.06-1",
  "new_index" : "my_logs-2021.05.06-000002",
  "rolled_over" : false,
  "dry_run" : true,
  "conditions" : {
    "[max_docs: 14000]" : true,
    "[max_size: 5gb]" : false,
    "[max_age: 7d]" : false
  }
}
根据目前我们的条件，我们的 logs-2019.10.21-1 文档数已经超过 14000 个了，所以会生产新的索引 logs-2019.10.21-000002。因为我使用了 dry_run，也就是演习，所以显示的  rolled_over 是 false。

为了能真正地 rollover，我们运行如下的命令：

POST /log_alias/_rollover
{
  "conditions": {
    "max_age": "7d",
    "max_docs": 1400,
    "max_size": "5gb"
  }
}
显示的结果是：
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "old_index" : "my_logs-2021.05.06-1",
  "new_index" : "my_logs-2021.05.06-000002",
  "rolled_over" : true,
  "dry_run" : false,
  "conditions" : {
    "[max_docs: 1400]" : true,
    "[max_size: 5gb]" : false,
    "[max_age: 7d]" : false
  }
}

说明它已经 rolled_ovder了。我们可以通过如下写的命令来检查：

GET _cat/indices/my_slogs-2021*
显示的结果为：
yellow open my_logs-2021.05.06-1      nyqua6U0SRusnqPCY32Xvw 1 1 14074 0 9.3mb 9.3mb
yellow open my_logs-2021.05.06-000002 e7r5hZpcSYi7wgnrVKuJqg 1 1     0 0  208b  208b


我们现在可以看到有两个以 logs-2019.10.21 为头的 index，并且第二文档 logs-2019.10.21-000002 文档数为0。如果我们这个时候直接再想 log_alias 写入文档的话：

POST log_alias/_doc
{
  "agent": "Mozilla/5.0 (X11; Linux x86_64; rv:6.0a1) Gecko/20110421 Firefox/6.0a1",
  "bytes": 6219,
  "clientip": "223.87.60.27",
  "extension": "deb",
  "geo": {
    "srcdest": "IN:US",
    "src": "IN",
    "dest": "US",
    "coordinates": {
      "lat": 39.41042861,
      "lon": -88.8454325
    }
  },
  "host": "artifacts.elastic.co",
  "index": "kibana_sample_data_logs",
  "ip": "223.87.60.27",
  "machine": {
    "ram": 8589934592,
    "os": "win 8"
  },
  "memory": null,
  "message": """          
  223.87.60.27 - - [2018-07-22T00:39:02.912Z] "GET /elasticsearch/elasticsearch-6.3.2.deb_1 HTTP/1.1" 200 6219 "-" "Mozilla/5.0 (X11; Linux x86_64; rv:6.0a1) Gecko/20110421 Firefox/6.0a1"
  """,
  "phpmemory": null,
  "referer": "http://twitter.com/success/wendy-lawrence",
  "request": "/elasticsearch/elasticsearch-6.3.2.deb",
  "response": 200,
  "tags": [
    "success",
    "info"
  ],
  "timestamp": "2019-10-13T00:39:02.912Z",
  "url": "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-6.3.2.deb_1",
  "utc_time": "2019-10-13T00:39:02.912Z"
}
显示的结果：

{
  "_index" : "my_logs-2021.05.06-000002",
  "_type" : "_doc",
  "_id" : "EeXIQHkBHw7XwuHSEjUT",
  "_version" : 1,
  "result" : "created",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 0,
  "_primary_term" : 1
}
显然它写入的是 logs-2019.10.21-000002 索引。我们再次查询 log_alias 的总共文档数：

GET log_alias/_count
显示的结果是：

{
  "count" : 14075,
  "_shards" : {
    "total" : 2,
    "successful" : 2,
    "skipped" : 0,
    "failed" : 0
  }
}
显然它和之前的 14074 个文档多增加了一个文档，也就是说 log_alias 是同时指向 logs-2019.10.21-1 及 logs-2019.10.21-000002。
https://github.com/dadoonet/demo-index-split-shrink-rollover