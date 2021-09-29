Elasticsearch中提供了cat APIs功能，有助于使结果的打印格式更易于阅读和理解。c

详细
详细的输出可以很好地显示cat命令的结果。在下面给出的示例中，我们获得了集群中存在的各种索引的详细信息。

GET /_cat/indices?v

标头
h参数（也称为标头）仅用于显示命令中提到的那些列。

GET /_cat/nodes?h=ip,port


