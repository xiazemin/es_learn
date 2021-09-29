索引API
当对具有特定映射的相应索引进行请求时，它有助于在索引中添加或更新JSON文档。
PUT schools/_doc/5
{
   name":"City School", "description":"ICSE", "street":"West End"
}

自动索引创建
当请求将JSON对象添加到特定索引时，如果该索引不存在，则此API会自动创建该索引以及该特定JSON对象的基础映射。可以通过将elasticsearch.yml文件中存在的以下参数的值更改为false来禁用此功能。

action.auto_create_index:false
index.mapper.dynamic:false


操作类型
操作类型用于强制执行创建操作。这有助于避免覆盖现有文档。

PUT chapter/_doc/1?op_type=create
{
   "Text":"this is chapter one"
}

自动ID生成
如果在索引操作中未指定ID，则Elasticsearch会自动为该文档生成ID。

POST chapter/_doc/
{
   "user" : "tpoint"
}

获取API
API通过对特定文档执行get请求来帮助提取类型JSON对象。
GET schools/_doc/5

此操作是实时的，不受索引刷新率的影响。

您还可以指定版本，然后Elasticsearch将仅获取该文档的版本。

您还可以在请求中指定_all，以便Elasticsearch可以按每种类型搜索该文档ID，它将返回第一个匹配的文档。

您还可以在特定文档的结果中指定所需的字段。

GET schools/_doc/5?_source_includes=name,fees

您还可以通过在get请求中添加_source部分来获取结果中的源部分。

GET schools/_doc/5?_source


删除API
您可以通过向Elasticsearch发送HTTP DELETE请求来删除特定的索引，映射或文档。

DELETE schools/_doc/4


更新API
脚本用于执行此操作，版本控制用于确保在获取和重新编制索引期间未发生任何更新。例如，您可以使用脚本更新学费-

POST schools/_update/4
{
   "script" : {
      "source": "ctx._source.name = params.sname",
      "lang": "painless",
      "params" : {
         "sname" : "City Wise School"
      }
   }
 }


 
