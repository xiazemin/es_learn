https://elasticstack.blog.csdn.net/article/details/104298141
POST /_cluster/reroute
{
  "commands": [
    {
      "move": {
        "index": "test",
        "shard": 0,
        "from_node": "node1",
        "to_node": "node2"
      }
    },
    {
      "allocate_replica": {
        "index": "test",
        "shard": 1,
        "node": "node3"
      }
    }
  ]
}
停用节点
另一个用例是从活动集群中停用节点。 这种情况下的主要挑战之一是在不导致群集停机或重启的情况下停用节点。 幸运的是，Elasticsearch 提供了一个选项，可以在不丢失数据或不会造成停机的情况下，优雅地删除/停用节点。 让我们看看如何实现它：

PUT _cluster/settings
{
  "transient": {
    "cluster.routing.allocation.exclude._ip": "IP of the node"
  }

}

