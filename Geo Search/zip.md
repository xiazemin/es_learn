PUT _template/zipcodes
{
  "order": 10,
  "index_patterns": [
    "zipcodes*"
  ],
  "settings": {
    "number_of_replicas": 0,
    "number_of_shards": 1
  },
  "mappings": {
    "properties": {
      "zipcode": {
        "type": "text"
      },
      "location": {
        "type": "geo_point"
      }
    }
  },
  "aliases": {}
}

{
  "acknowledged" : true
}



~/software/elasticsearch-7.9.3/logstash-7.9.2/bin/logstash -f ./zipcodes/logstash_zipcodes.conf

~/software/elasticsearch-7.9.3/logstash-7.9.2/bin/logstash -e 'input { stdin { } } output { stdout {} }'