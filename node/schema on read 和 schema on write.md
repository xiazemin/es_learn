https://elasticstack.blog.csdn.net/article/details/113813915

Elasticsearch 依靠 schema on write 的模式来快速搜索数据。现在，我们向 Elasticsearch 添加了 schema on read 模式，以便用户可以灵活地在摄取后更改文档的 schema，还可以生成仅作为搜索查询一部分存在的字段。schema on read 和 schema on write 一起为用户提供了选择，可以根据他们的需求来平衡性能和灵活性。
我们的 schema on read 解决方案是 runtime fields，它们仅在查询时进行评估。它们在索引映射或查询中定义，一旦定义，它们立即可用于搜索请求，聚合，过滤和排序。由于未对 runtime fields 进行索引，因此添加运行时字段不会增加索引的大小。实际上，它们可以降低存储成本并提高摄取速度。

但是，需要权衡取舍。对运行时字段的查询可能会很昂贵，因此你通常搜索或筛选所依据的数据仍应映射到索引字段。即使你的索引大小较小，runtime fields 也会降低搜索速度。我们建议结合使用 runtime fields 和索引字段，以在用例的摄取速度，索引大小，灵活性和搜索性能之间找到合适的平衡。

