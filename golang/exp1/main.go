package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	elastic "gopkg.in/olivere/elastic.v5"
)

//https://studygolang.com/articles/27657?fr=sidebar

func GetEsClient() *elastic.Client {
	file := "./eslog.log"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766) // 应该判断error，此处简略
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200/"),
		elastic.SetBasicAuth("elastic", "password"),
		//docker
		elastic.SetSniff(false),
		elastic.SetInfoLog(log.New(logFile, "ES-INFO: ", 0)),
		elastic.SetTraceLog(log.New(logFile, "ES-TRACE: ", 0)),
		elastic.SetErrorLog(log.New(logFile, "ES-ERROR: ", 0)),
	)

	if err != nil {
		return nil
	}

	return client
}

func IsDocExists(id int, typ, index string) bool {
	client := GetEsClient()
	defer client.Stop()
	exist, _ := client.Exists().Index(index).Type(typ).Id(strconv.Itoa(id)).Do(context.Background())
	if !exist {
		log.Println("ID may be incorrect! ", id)
		return false
	}
	return true
}

func GetDoc(id int, typ, index string) (*elastic.GetResult, error) {
	client := GetEsClient()
	defer client.Stop()
	if !IsDocExists(id, typ, index) {
		return nil, fmt.Errorf("id不存在")
	}
	esResponse, err := client.Get().Index(index).Type(typ).Id(strconv.Itoa(id)).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return esResponse, nil
}

func AddDoc(id int, doc string, typ, index string) (*elastic.IndexResponse, error) {
	client := GetEsClient()
	defer client.Stop()
	if IsDocExists(id, typ, index) {
		return nil, fmt.Errorf("id不存在")
	}

	rsp, err := client.Index().Index(index).Type(typ).Id(strconv.Itoa(id)).BodyJson(doc).Do(context.Background())

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UpdateDoc(updateField *map[string]interface{}, id int, typ, index string) (*elastic.UpdateResponse, error) {
	client := GetEsClient()
	defer client.Stop()
	if !IsDocExists(id, typ, index) {
		return nil, fmt.Errorf("id不存在")
	}
	rsp, err := client.Update().Index(index).Type(typ).Id(strconv.Itoa(id)).Doc(updateField).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return rsp, nil

}

func DeleteDoc(id int, typ, index string) (*elastic.DeleteResponse, error) {
	client := GetEsClient()
	defer client.Stop()
	rsp, err := client.Delete().Index(index).Type(typ).Id(strconv.Itoa(id)).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func main() {
	var v map[string]interface{}
	s := `{"user":"双榆树-张三","message":"今儿天气不错啊，出去转转去","uid":2,"age":20,"city":"北京","province":"北京","country":"中国","address":"中国北京市海淀区","location":{"lat":"39.970718","lon":"116.325747"}}`
	_ = json.Unmarshal([]byte(s), &v)
	doc, err := GetDoc(1, "_doc", "twitter")
	ar, err := AddDoc(10, s, "_doc", "twitter")
	ur, err := UpdateDoc(&v, 10, "_doc", "twitter")
	dr, err := DeleteDoc(10, "_doc", "twitter")
	fmt.Println(IsDocExists(1, "_doc", "twitter"), "\n",
		doc, err, "\n",
		ar, "\n",
		ur, "\n",
		dr, "\n",
	)
}

//go mod init es_learn
//go mod tidy

//https://www.tizi365.com/archives/850.html

/*
go run golang/exp1/main.go
2021/04/28 17:12:33 ID may be incorrect!  10
true
 &{twitter _doc 1    0x14000202658 0x14000204228 true map[] <nil>} <nil>
 &{twitter _doc 10 1 created 0x1400007a360 7 1 0 false false}
 &{twitter _doc 10 1 0x14000114ba0 noop false <nil>}
 &{twitter _doc 10 2 deleted 0x1400009dcb0 8 1 0 false false}
*/

/**
GET twitter/_doc/1

{
  "_index" : "twitter",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 1,
  "_seq_no" : 0,
  "_primary_term" : 1,
  "found" : true,
  "_source" : {
    "user" : "张三",
    "message" : "今儿天气不错啊，出去转转去",
    "uid" : 2,
    "age" : 20,
    "city" : "北京",
    "province" : "北京",
    "country" : "中国",
    "address" : "中国北京市海淀区",
    "location" : {
      "lat" : "39.970718",
      "lon" : "116.325747"
    },
    "DOB" : "1999-04-01"
  }
}



*/
