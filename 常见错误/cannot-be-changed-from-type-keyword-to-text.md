https://stackoverflow.com/questions/65031413/elasticsearch-7-9-0-cannot-be-changed-from-type-keyword-to-text

You need to add _doc in URL while posting a document to Elasticsearch, change the URL to POST /my-demo1/_doc/1

