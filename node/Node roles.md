https://elasticstack.blog.csdn.net/article/details/110947372

默认情况下，群集中的每个节点都可以处理 HTTP 和 Transport 流量。 Transport 层专门用于节点之间的通信。 HTTP 层由 REST 客户端使用。

所有节点都知道群集中的所有其他节点，并且可以将客户端请求转发到适当的节点。

默认情况下，节点为以下所有类型：master-eligible, data, ingest 和（如果可用）machine learning。 所有 data 节点也是 transform 节点。

你可以通过设置 node.roles 来定义节点的角色。 如果你未配置此设置，则该节点默认具有以下角色：

master
data
data_content
data_hot
data_warm
data_cold
ingest
ml
remote_cluster_client