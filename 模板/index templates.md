Index template 定义在创建新 index 时可以自动应用的 settings 和 mappings。 Elasticsearch 根据与 index 名称匹配的 index 模式将模板应用于新索引。这个对于我们想创建的一系列的 Index 具有同样的 settings 及 mappings。比如我们希望每一天/月的日志的index都具有同样的设置。

https://elasticstack.blog.csdn.net/article/details/100553185

Index template 仅在 index 创建期间应用。 对 index template 的更改不会影响现有索引。 create index API 请求中指定的设置和映射会覆盖索引模板中指定的任何设置或映射。

你可以在代码中加入像 C 语言那样的 block 注释。你可以把这个注释放在出来开头 “{” 和结尾的 “}” 之间的任何地方。


定义一个Index template
我们可以使用如下的接口来定义一个 index template：

PUT /_template/<index-template>
我们可以使用_template这个终点来创建，删除，查看一个 index template。

PUT _template/logs_template
{
  "index_patterns": "logs-*",
  "order": 1, 
  "settings": {
    "number_of_shards": 4,
    "number_of_replicas": 1
  },
  "mappings": { 
    "properties": {
      "@timestamp": {
        "type": "date"
      }
    }
  }
}

#! Deprecation: legacy template [logs_template] has index patterns [logs-*] matching patterns from existing composable templates [logs] with patterns (logs => [logs-*-*]); this template [logs_template] may be ignored in favor of a composable template at index creation time
{
  "acknowledged" : true
}


https://www.elastic.co/guide/en/elasticsearch/reference/current/index-templates.html
Elasticsearch has built-in index templates, each with a priority of 100, for the following index patterns:

logs-*-*
metrics-*-*
synthetics-*-*



任何以 “logs-” 为开头的任何一个 index 将具有在该 template 里具有的 settings 及 mappings 属性。这里的 “order” 的意思是：如果索引与多个模板匹配，则 Elasticsearch 应用此模板的顺序。该值为1，表明有最先合并，如果有更高 order 的 template，这个 settings 或 mappings 有可能被其它的 template 所覆盖。


PUT _template/logs_template
{
  "index_patterns": "log_xzm-*",
  "order": 1, 
  "settings": {
    "number_of_shards": 4,
    "number_of_replicas": 1
  },
  "mappings": { 
    "properties": {
      "@timestamp": {
        "type": "date"
      }
    }
  }
}

{
  "acknowledged" : true
}



GET _template/logs


PUT logs-2019-03-01

{
  "error" : {
    "root_cause" : [
      {
        "type" : "illegal_argument_exception",
        "reason" : "cannot create index with name [logs-2019-03-01], because it matches with template [logs] that creates data streams only, use create data stream api instead"
      }
    ],
    "type" : "illegal_argument_exception",
    "reason" : "cannot create index with name [logs-2019-03-01], because it matches with template [logs] that creates data streams only, use create data stream api instead"
  },
  "status" : 400
}


PUT log_xzm-2019-03-01

{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "log_xzm-2019-03-01"
}


GET log_xzm-2019-03-01
{
  "log_xzm-2019-03-01" : {
    "aliases" : { },
    "mappings" : {
      "properties" : {
        "@timestamp" : {
          "type" : "date"
        }
      }
    },
    "settings" : {
      "index" : {
        "creation_date" : "1620272907871",
        "number_of_shards" : "4",
        "number_of_replicas" : "1",
        "uuid" : "L98-LarRS5iOhy3H3ywYug",
        "version" : {
          "created" : "7090399"
        },
        "provided_name" : "log_xzm-2019-03-01"
      }
    }
  }
}


我们甚至可以为我们的 index template 添加 index alias：
PUT _template/logs_template
{
  "index_patterns": "log_xzm-*",
  "order": 1, 
  "settings": {
    "number_of_shards": 4,
    "number_of_replicas": 1
  },
  "mappings": { 
    "properties": {
      "@timestamp": {
        "type": "date"
      }
    }
  },
  "aliases": {
    "{index}-alias" : {}
  }
}

PUT log_xzm-2019-04-01

{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "log_xzm-2019-04-01"
}

GET log_xzm-2019-04-01-alias
{
  "log_xzm-2019-04-01" : {
    "aliases" : {
      "log_xzm-2019-04-01-alias" : { }
    },
    "mappings" : {
      "properties" : {
        "@timestamp" : {
          "type" : "date"
        }
      }
    },
    "settings" : {
      "index" : {
        "creation_date" : "1620273139720",
        "number_of_shards" : "4",
        "number_of_replicas" : "1",
        "uuid" : "OAr71DkARkeW2n9cETXnkQ",
        "version" : {
          "created" : "7090399"
        },
        "provided_name" : "log_xzm-2019-04-01"
      }
    }
  }
}


多个索引模板可能与索引匹配，在这种情况下，设置和映射都合并到索引的最终配置中。 可以使用 order 参数控制合并的顺序，首先应用较低的顺序，并且覆盖它们的较高顺序。

PUT /_template/template_1
{
    "index_patterns" : ["t*"],
    "order" : 0,
    "settings" : {
        "number_of_shards" : 1
    },
    "mappings" : {
        "_source" : { "enabled" : false }
    }
}
 
PUT /_template/template_2
{
    "index_patterns" : ["te*"],
    "order" : 1,
    "settings" : {
        "number_of_shards" : 1
    },
    "mappings" : {
        "_source" : { "enabled" : true }
    }
}

# PUT /_template/template_1
{
  "acknowledged" : true
}

# PUT /_template/template_2
{
  "acknowledged" : true
}


以上的 template_1 将禁用存储 _source，但对于以 te * 开头的索引，仍将启用 _source。 注意，对于映射，合并是 “深度” 的，这意味着可以在高阶模板上轻松添加/覆盖特定的基于对象/属性的映射，而较低阶模板提供基础。

我们可以来创建一个例子看看：
PUT test10
GET test10

# PUT test10
#! Deprecation: index [test10] matches multiple legacy templates [template_1, template_2], composable templates will only match a single template
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "test10"
}

# GET test10
{
  "test10" : {
    "aliases" : { },
    "mappings" : { },
    "settings" : {
      "index" : {
        "creation_date" : "1620273374167",
        "number_of_shards" : "1",
        "number_of_replicas" : "1",
        "uuid" : "UQmbgYLXREuB95xe-d42Ig",
        "version" : {
          "created" : "7090399"
        },
        "provided_name" : "test10"
      }
    }
  }
}



PUT tast10
GET tast10

# PUT tast10
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "tast10"
}

# GET tast10
{
  "tast10" : {
    "aliases" : { },
    "mappings" : {
      "_source" : {
        "enabled" : false
      }
    },
    "settings" : {
      "index" : {
        "creation_date" : "1620273439539",
        "number_of_shards" : "1",
        "number_of_replicas" : "1",
        "uuid" : "-dQLEc3fTvyIBxCVpXeUXw",
        "version" : {
          "created" : "7090399"
        },
        "provided_name" : "tast10"
      }
    }
  }
}


如果对于两个 templates 来说，如果 order 是一样的话，我们可能陷于一种不可知论的合并状态。在实际的使用中必须避免。

我们可以通过如下的接口来查询已经被创建好的 index template:

GET /_template/<index-template>
比如：

GET _template/logs_template

检查一个 index template 是否存在
我们可以使用如下的命令来检查一个 index template 是否存在：

HEAD _template/logs_template
如果存在就会返回：

200 - OK

删除一个 index template
 

在之前的练习中，我们匹配 “*”，也就是我们以后所有的创建的新的 index 将不存储 source，这个显然不是我们所需要的。我们需要来把这个 template 进行删除。删除一个 template 的接口如下：

DELETE /_template/<index-template>

 











