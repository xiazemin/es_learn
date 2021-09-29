PUT colleges
DELETE /colleges
GET colleges
HEAD colleges 索引存在

索引设置
您只需在网址末尾附加_settings关键字即可获取索引设置。
GET /colleges/_settings

索引统计
该API可帮助您提取有关特定索引的统计信息。您只需要在末尾发送带有索引URL和_stats关键字的get请求。

GET /_stats

冲洗(Flush)
索引的刷新过程可确保当前仅保留在事务日志中的所有数据也将永久保留在Lucene中。这减少了恢复时间，因为在打开Lucene索引之后，不需要从事务日志中重新索引数据。

POST colleges/_flush

