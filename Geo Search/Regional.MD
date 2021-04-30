https://elasticstack.blog.csdn.net/article/details/104950484


我们将使用 Logstash 定期地去 opensky network 网站上去抓取数据，并导入到 Elasticsearch 中。Logstash 的配置文件如下：

fligths_logstash.conf

input {
    http_poller {
        codec => "json"
        schedule => { every => "15s" }
        urls => {
            url => "https://opensky-network.org/api/states/all"
        }
    }
}
 
filter {
    split {
        field => "states"
        add_field => {
            "icao" => "%{[states][0]}"
            "callsign" => "%{[states][1]}"
            "origin_country" => "%{[states][2]}"
            "time_position" => "%{[states][3]}"
            "last_contact" => "%{[states][4]}"
            "location" => "%{[states][6]},%{[states][5]}"
            "baro_altitude" => "%{[states][7]}"
            "on_ground" => "%{[states][8]}"
            "velocity" => "%{[states][9]}"
            "true_track" => "%{[states][10]}"
            "vertical_rate" => "%{[states][11]}"
            "geo_altitude" => "%{[states][13]}"
            "squawk" => "%{[states][14]}"
            "spi" => "%{[states][15]}"
            "position_source" => "%{[states][16]}"
        }
    }
    mutate {
        strip => ["callsign"]
        rename => { "time" => "request_time" }
        remove_field => ["states", "@version"]
    }
    translate {
        field => "[position_source]"
        destination => "position_source"
        override => "true"
        dictionary => {
          "0" => "ADS-B"
          "1" => "ASTERIX"
          "2" => "MLAB"
        }
    }
 
    if [time_position] =~ /^%{*/ {
        drop { }
    }
    if [callsign] =~ /^%{*/ {
        mutate { remove_field => ["callsign"] }
    }
    if [location] =~ /^%{*/ {
        mutate { remove_field => ["location"] }
    }
    if [baro_altitude] =~ /^%{*/ {
        mutate { remove_field => ["baro_altitude"] }
    }
    if [velocity] =~ /^%{*/ {
        mutate { remove_field => ["velocity"] }
    }
    if [true_track] =~ /^%{*/ {
        mutate { remove_field => ["true_track"] }
    }
    if [vertical_rate] =~ /^%{*/ {
        mutate { remove_field => ["vertical_rate"] }
    }
    if [sensors] =~ /^%{*/ {
        mutate { remove_field => ["sensors"] }
    }
    if [geo_altitude] =~ /^%{*/ {
        mutate { remove_field => ["geo_altitude"] }
    }
    if [squawk] =~ /^%{*/ {
        mutate { remove_field => ["squawk"] }
    }
 
    mutate {
		convert => { 
            "baro_altitude" => "float" 
		    "geo_altitude" => "float"
            "last_contact" => "integer"
            "on_ground" => "boolean"
            "request_time" => "integer"
            "spi" => "boolean"
            "squawk" => "integer"
            "time_position" => "integer"
            "true_track" => "float"
            "velocity" => "float"
            "vertical_rate" => "float"
        }
	}
}
 
output {
    stdout { 
        codec => rubydebug
    }
 
    elasticsearch {
        manage_template => "false"
        index => "flights"
        # pipeline => "flights_aircraft_enrichment"
     	hosts => "localhost:9200"
    }


https://elasticstack.blog.csdn.net/article/details/106531939
https://elasticstack.blog.csdn.net/article/details/115727243
https://elasticstack.blog.csdn.net/article/details/109292893


https://elasticstack.blog.csdn.net/article/details/108706278
https://elasticstack.blog.csdn.net/article/details/107634339