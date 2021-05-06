你可以使用 script_fields 参数为每个 hit 检索进行脚本运算（基于不同的字段）。 例如：

GET developers/_search
{
  "query": {
    "match_all": {}
  },
  "script_fields": {
    "2_times_age": {
      "script": {
        "source": """
          doc['age'].value * 2
        """
      }
    },
    "3_times_age": {
      "script": {
        "source": """
          doc['age'].value * 3
        """
      }
    }
  }
}

Scripted fields 可以在未 store 的字段上工作（在上述情况下为 age），并允许返回要返回的自定义值（脚本的计算值）。

脚本字段还可以使用 params['_ source'] 访问实际的 _source 文档并提取要从中返回的特定元素。 这是一个例子：

GET developers/_search
{
  "query": {
    "match_all": {}
  },
  "script_fields": {
    "my_city": {
      "script": {
        "source": """
          "I am living in " + params["_source"]["city"]
        """
      }
    }
  }
}

请注意此处的 _source 关键字，以浏览类似 json 的模型。

重要的是要了解 doc['my_field'].value 和 params['_ source'] ['my_field'] 之间的区别。 第一个使用 doc 关键字，将导致将该字段的术语加载到内存中（缓存），这将导致执行速度更快，但会占用更多内存。 此外，doc [...]表示法仅允许使用简单值字段（你无法从中返回 json 对象），并且仅对未分析或基于单个术语的字段有意义。 但是，从文档访问值的方式来说，仍然建议使用doc（即使有可能），因为 _source 每次使用时都必须加载和解析。 使用 _source 非常慢。