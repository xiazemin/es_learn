https://zhuanlan.zhihu.com/p/35283785
在ES中，Master节点是通过发布ClusterState来通知其他节点的。Master会将新的ClusterState发布给其他的所有节点，当节点收到新的ClusterState后，会把新的ClusterState发给相关的各个模块，各个模块根据新的ClusterState判断是否要做什么事情，比如创建Shard等。即这是一种通过Meta数据来驱动各个模块工作的方式。

1. Meta：ClusterState、MetaData、IndexMetaData
Meta是用来描述数据的数据。在ES中，Index的mapping结构、配置、持久化状态等就属于meta数据，集群的一些配置信息也属于meta。这类meta数据非常重要，假如记录某个index的meta数据丢失了，那么集群就认为这个index不再存在了。ES中的meta数据只能由master进行更新，master相当于是集群的大脑。

ClusterState
集群中的每个节点都会在内存中维护一个当前的ClusterState，表示当前集群的各种状态。ClusterState中包含一个MetaData的结构，MetaData中存储的内容更符合meta的特征，而且需要持久化的信息都在MetaData中，此外的一些变量可以认为是一些临时状态，是集群运行中动态构建出来的。

ClusterState内容包括：
    long version: 当前版本号，每次更新加1
    String stateUUID：该state对应的唯一id
    RoutingTable routingTable：所有index的路由表
    DiscoveryNodes nodes：当前集群节点
    MetaData metaData：集群的meta数据
    ClusterBlocks blocks：用于屏蔽某些操作
    ImmutableOpenMap<String, Custom> customs: 自定义配置
    ClusterName clusterName：集群名



MetaData
上面提到，MetaData更符合meta的特征，而且需要持久化，那么我们看下这个MetaData中主要包含哪些东西：

MetaData中需要持久化的包括：
    String clusterUUID：集群的唯一id。
    long version：当前版本号，每次更新加1
    Settings persistentSettings：持久化的集群设置
    ImmutableOpenMap<String, IndexMetaData> indices: 所有Index的Meta
    ImmutableOpenMap<String, IndexTemplateMetaData> templates：所有模版的Meta
    ImmutableOpenMap<String, Custom> customs: 自定义配置


IndexMetaData
IndexMetaData指具体某个Index的Meta，比如这个Index的shard数，replica数，mappings等。

IndexMetaData中需要持久化的包括：
    long version：当前版本号，每次更新加1。
    int routingNumShards: 用于routing的shard数, 只能是该Index的numberOfShards的倍数，用于split。
    State state: Index的状态, 是个enum，值是OPEN或CLOSE。
    Settings settings：numbersOfShards，numbersOfRepilicas等配置。
    ImmutableOpenMap<String, MappingMetaData> mappings：Index的mapping
    ImmutableOpenMap<String, Custom> customs：自定义配置。
    ImmutableOpenMap<String, AliasMetaData> aliases： 别名
    long[] primaryTerms：primaryTerm在每次Shard切换Primary时加1，用于保序。
    ImmutableOpenIntMap<Set<String>> inSyncAllocationIds：处于InSync状态的AllocationId，用于保证数据一致性，下一篇文章会介绍。


每次需要更新ClusterState时提交一个Task给MasterService，MasterService中只使用一个线程来串行处理这些Task，每次处理时把当前的ClusterState作为Task中execute函数的参数。即保证了所有的Task都是在currentClusterState的基础上进行更改，然后不同的Task是串行执行的。


早期的ES版本没有解决这个问题，后来引入了两阶段提交的方式(Add two phased commit to Cluster State publishing)。所谓的两阶段提交，是把Master发布ClusterState分成两步，第一步是向所有节点send最新的ClusterState，当有超过半数的master节点返回ack时，再发送commit请求，要求节点commit接收到的ClusterState。如果没有超过半数的节点返回ack，那么认为本次发布失败，同时退出master状态，执行rejoin重新加入集群。



ES中，Master发送commit的原则是只要有超过半数MasterNode(master-eligible node)接收了新的ClusterState就发送commit。那么实际上就是认为只要超过半数节点接收了新的ClusterState，这个ClusterState就一定可以被commit，不会在各种场景下回退。

