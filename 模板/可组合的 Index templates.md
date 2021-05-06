索引模板（Index template）是一种告诉 Elasticsearch 在创建索引时如何配置索引的方法。 对于数据流（data stream），索引模板会在创建流时支持其后备索引。 在创建索引之前先配置模板，然后在手动创建索引或通过对文档建立索引创建索引时，模板设置将用作创建索引的基础。

模板有两种类型：索引模板和组件模板。 组件模板是可重用的构建块，用于配置映射，设置和别名。 你使用组件模板来构造索引模板，但它们不会直接应用于一组索引。 索引模板可以包含组件模板的集合，也可以直接指定设置，映射和别名。

如果新数据流或索引与多个索引模板匹配，则使用优先级最高的索引模板。

重要提示：Elasticsearch 具有以下索引模板的内置索引模板，每个索引的优先级为 100：

logs-*-*
metrics-*-*
synthetics-*-*
Elastic Agent 使用这些模板来创建数据流。 如果使用 Elastic Agent，请为索引模板分配低于100的优先级，以避免覆盖内置模板。 否则，为避免意外应用内置模板，请执行以下一项或多项操作：

要禁用所有内置索引索引和组件模板，请使用集群更新设置 API 将 stack.templates.enabled 设置为 false。
使用非重叠索引模式。
为具有重叠模式的模板分配优先级高于100。例如，如果你不使用 Elastic Agent，并且想为 logs-* 索引模式创建模板，请为模板分配优先级 200。这样可以确保你的模板 应用而不是日志的内置模板 logs-*-*。 
当可组合模板与给定索引匹配时，它始终优先于旧模板。 如果没有可组合模板匹配，则旧版模板可能仍匹配并被应用。

如果使用显式设置创建索引并且该索引也与索引模板匹配，则创建索引请求中的设置将优先于索引模板及其组件模板中指定的设置。

 

创建可组合的 index template
下面，我们来使用一个例子来展示如何使用可组合的 index template。

我们首先来创建二个组件模板

PUT _component_template/component_template1
{
  "template": {
    "mappings": {
      "properties": {
        "@timestamp": {
          "type": "date"
        }
      }
    }
  }
}
 
PUT _component_template/other_component_template
{
  "template": {
    "mappings": {
      "properties": {
        "ip_address": {
          "type": "ip"
        }
      }
    }
  }
}

# PUT _component_template/component_template1
{
  "acknowledged" : true
}

# PUT _component_template/other_component_template
{
  "acknowledged" : true
}


PUT _index_template/template_1
{
  "index_patterns": ["de*", "bar*"],
  "template": {
    "settings": {
      "number_of_shards": 1
    },
    "mappings": {
      "_source": {
        "enabled": false
      },
      "properties": {
        "host_name": {
          "type": "keyword"
        },
        "created_at": {
          "type": "date",
          "format": "EEE MMM dd HH:mm:ss Z yyyy"
        }
      }
    },
    "aliases": {
      "mydata": { }
    }
  },
  "priority": 200,
  "composed_of": ["component_template1", "other_component_template"],
  "version": 3,
  "_meta": {
    "description": "my custom"
  }
}

{
  "acknowledged" : true
}


"composed_of": ["component_template1", "other_component_template"],
也就是说这个叫做 template_1 的 index template，它包含了两个可组合的 component templates：component_template 及 other_component_template。

模拟多组件模板
由于模板不仅可以由多个组件模板组成，还可以由索引模板本身组成，因此有两个模拟API可以确定最终的索引设置是什么。

要模拟将应用于特定索引名称的设置：

POST /_index_template/_simulate_index/test

https://elasticstack.blog.csdn.net/article/details/113751797