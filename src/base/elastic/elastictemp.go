package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	//client, err := elastic.NewClient(elastic.SetURL("http://10.10.8.151:9200/"))
	//file := "./eslog.log"
	//logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	context.Background()
	client, err := elastic.NewClient(
		elastic.SetURL("http://10.10.8.151:9200/"),
		//docker
		elastic.SetSniff(true),
		//elastic.SetInfoLog(log.New(logFile, "ES-INFO: ", 0)),
		//elastic.SetTraceLog(log.New(logFile, "ES-TRACE: ", 0)),
		//elastic.SetErrorLog(log.New(logFile, "ES-ERROR: ", 0)),
	)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connect to es success")
	p1 := Person{Name: "lmh", Age: 18, Married: false}

	put1, err := client.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	client.CreateIndex("myzone")

	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
