https://blog.csdn.net/UbuntuTouch/article/details/104865306

https://github.com/liu-xiao-guo/elasticzipcodes

https://blog.csdn.net/UbuntuTouch/article/details/99655350

https://www.elastic.co/cn/downloads/logstash

https://www.elastic.co/cn/downloads/past-releases/logstash-7-9-2

https://elasticsearch.cn/download/

/Users/xiazemin/software/elasticsearch-7.9.3/logstash-7.9.2/logstash-core/lib/logstash/java_pipeline.rb:369:in `block in start_input'

修改配置
		path => "/Users/xiazemin/source/es_learn/Geo\ Search/zipcodes/zipcodes.csv"

	elasticsearch {
		index => "zipcodes"
		hosts => ["http://elastic:password@localhost:9200"]
	}



output {
    stdout {
        codec => rubydebug
    }
    elasticsearch {
        hosts => ["{{ cbx_logstash_es_server }}"]
        index => "%{indexName}"
        action => "index"
    }
}
1
2
3
4
5
6
7
8
9
10
根据配置，并结合堆栈信息来分析，可以认为是Logstash的stdout插件在高并发状态下使用rubydebug进行编解码时抛出了异常。

其实这里的stdout插件是不必要的，之前只是在本地测试使用到的。而在测试环境下，并发量远非本地测试能比，此外将大量的message输出到console上也会对性能产生影响。可以说，这种配置等同于在Java代码中频繁使用System.out.print()语句来输出信息，并不推荐这种做法。

解决方案
将配置文件里的stdout插件去掉，最终output的配置如下：

output {
    elasticsearch {
        hosts => ["{{ cbx_logstash_es_server }}"]
        index => "%{indexName}"
        action => "index"
    }
}
https://blog.csdn.net/lewky_liu/article/details/99496420

换7.9.3
~/software/elasticsearch-7.9.3/logstash-7.9.3/bin/logstash -f ./zipcodes/logstash_zipcodes.conf


https://discuss.elastic.co/t/block-in-start-input/166474
sincedb_path => "NUL"
not /dev/null

[2021-04-30T13:24:34,270][INFO ][logstash.agent           ] Successfully started Logstash API endpoint {:port=>9600}

[2021-04-30T13:24:58,598][WARN ][logstash.outputs.elasticsearch][main][8d7e71af2e0e5beb2c2297fb800e751d57b217320990afcf8ee71594d82a0e2a] Could not index event to Elasticsearch. {:status=>400, :action=>["index", {:_id=>nil, :_index=>"zipcodes", :routing=>nil, :_type=>"_doc"}, #<LogStash::Event:0x55bbe9be>], :response=>{"index"=>{"_index"=>"zipcodes", "_type"=>"_doc", "_id"=>"kuQ8IXkBHw7XwuHSua9W", "status"=>400, "error"=>{"type"=>"mapper_parsing_exception", "reason"=>"failed to parse field [location] of type [geo_point]", "caused_by"=>{"type"=>"parse_exception", "reason"=>"latitude must be a number"}}}}}

GET zipcodes/_count
{
  "count" : 42358,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  }
}





