lasticsearch 依据我们提供的数据，根据自己的估计，并自动生成相应的 mapping。这在很多的时候是非常有用的。 它可以帮我们自动省很多的手动操作，而且在大多数的情况下，Elasticsearch 帮我们自动生成的 mapping 也是有效的。一般来讲，如果我们想自己定义自己的 mapping 的话，如下的步骤将是可取的，也是推荐的方法：

把自己的一个数据输入到 Elasticsearch 中
得到上面输入数据的 mapping，并在此基础上进行调整，从而得出适合自己数据的 mapping
动态的 mapping 并不总是优化的
针对一个浮点数来说：

PUT my_index/_doc/1
{
  "price": 1.99
}
我们可以得到它的 mapping:

GET my_index/_mapping
{
  "my_index" : {
    "mappings" : {
      "properties" : {
        "price" : {
          "type" : "float"
        }
      }
    }
  }
}
从上面我们可以看出来，price 的数据类型是一个 float 类型。对于大多数的情况来说，这个应该没有问题。但是在实际的应用中，我们可以把这个 float 数据类型转换为 scaled float数据类型。Scaled float 由 long 数据类型来支持。long 数据类型在 Lucene 里可以被更加有效地压缩，从而节省存储的空间。在使用 scaled float 数据类型时，你必须使用scaling_factore来配置它的精度：

PUT my_index1/_doc/1
{
  "mappings": {
    "properties": {
      "price": {
        "type": "scaled_float",
        "scaling_factor": 100
      }
    }
  }
}
在上面我们定义 price 类型为 scaled_float 数据类型，并定义 scaling_factor 为 100。这样我们的数据，比如 1.99 刚好可以在乘以 100 变为 199，从而成为一个整型数。

经过这样的改造后，我们可以试试重新在 my_index1 里输入一个文档：

PUT my_index1/_doc/1
{
  "price": 1.99

}