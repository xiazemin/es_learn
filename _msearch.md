
 

Multi Search API
使用单个 API 请求执行几次搜索。这个 API 的好处是节省 API 的请求个数，把多个请求放到一个 API 请求中来实现。

为了说明问题的方便，我们可以多加一个叫做 twitter1 的 index。它的内容如下：

POST _bulk
{"index":{"_index":"twitter1","_id":1}}
{"user":"张庆","message":"今儿天气不错啊，出去转转去","uid":2,"age":20,"city":"重庆","province":"重庆","country":"中国","address":"中国重庆地区","location":{"lat":"39.970718","lon":"116.325747"}}
这样在我们的 Elasticsearch 中就有两个索引了。我们可以做如下的 _msearch。

GET twitter/_msearch
{"index":"twitter"}
{"query":{"match_all":{}},"from":0,"size":1}
{"index":"twitter"}
{"query":{"bool":{"filter":{"term":{"city.keyword":"北京"}}}}, "size":1}
{"index":"twitter1"}
{"query":{"match_all":{}}}
上面我们通过 _msearch 终点来实现在一个 API 请求中做多个查询，对多个 index 进行同时操作。显示结果为：



 

多个索引操作
在上面我们引入了另外一个索引 twitter1。在实际的操作中，我们可以通过通配符，或者直接使用多个索引来进行搜索：

GET twitter*/_search
上面的操作是对所有的以 twitter 为开头的索引来进行搜索，显示的结果是在所有的 twitter 及 twitter1 中的文档：



GET /twitter,twitter1/_search
也可以做同样的事。在写上面的查询的时候，在两个索引之间不能加入空格，比如：

GET /twitter, twitter1/_search
上面的查询并不能返回你所想要的结果。
