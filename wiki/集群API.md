群集API用于获取有关群集及其节点的信息并在其中进行更改。要调用此API，我们需要指定节点名称，地址或_local。

GET /_nodes/_local


集群运行状况
API用于通过附加'health'关键字来获取集群运行状况的状态。

GET /_cluster/health

集群状态
该API用于通过附加'state'关键字URL来获取有关集群的状态信息。状态信息包含版本，主节点，其他节点，路由表，元数据和块。

GET /_cluster/state

集群统计
该API通过使用'stats'关键字来帮助检索有关群集的统计信息。该API返回分片号，存储大小，内存使用率，节点数，角色，操作系统和文件系统。

GET /_cluster/stats


节点统计
该API用于检索集群中另外一个节点的统计信息。节点统计信息与集群几乎相同。

GET /_nodes/stats

节点hot_threads
该API可帮助您检索有关群集中每个节点上的当前热线程的信息。

GET /_nodes/hot_threads

