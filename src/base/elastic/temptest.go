package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
	"strconv"
)

var (
	subject   Subject
	indexName = "subject"
	servers   = []string{"http://10.10.8.151:9200/"}
)

type Subject struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Genres []string `json:"genres"`
}

const mapping = `
{
    "mappings": {
        "properties": {
            "id": {
                "type": "long"
            },
            "title": {
                "type": "text"
            },
            "genres": {
                "type": "keyword"
            }
        }
    }
}`

func Search(client *elastic.Client, ctx context.Context, genre string) {
	fmt.Printf("Search: %s", genre)
	// Term搜索
	termQuery := elastic.NewTermQuery("genres", genre)
	//elastic.NewTerm
	searchResult, err := client.Search().
		Index(indexName).
		Query(termQuery).
		Sort("id", true). // 按id升序排序
		From(0).Size(10). // 拿前10個結果
		Pretty(true).
		Do(ctx) // 執行
	if err != nil {
		panic(err)
	}
	//searchResult.
	total := searchResult.TotalHits()
	//fmt.Printf("Found %d subjects\n", total)
	if total > 0 {
		for _, item := range searchResult.Each(reflect.TypeOf(subject)) {
			if t, ok := item.(Subject); ok {
				fmt.Printf("Found: Subject(id=%d, title=%s,Genres=%v)\n", t.ID, t.Title, t.Genres)
			}
		}

	} else {
		fmt.Println("Not found!")
	}
}

func Search1(client *elastic.Client, ctx context.Context, Term string) {
	Termq := elastic.NewTermQuery("genres", Term)
	result, _ := client.Search().
		Index(indexName).
		Query(Termq).
		Sort("id", true).
		Pretty(true).
		Do(ctx)
	fmt.Println(result)
}

func main() {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(servers...),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(exists)
	if !exists {
		_, err := client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
	subject = Subject{
		ID:     1,
		Title:  "肖恩克的救贖",
		Genres: []string{"犯罪", "劇情"},
	}

	// 寫入
	doc, err := client.Index().
		Index(indexName).
		Id(strconv.Itoa(subject.ID)).
		BodyJson(subject).
		Refresh("wait_for").
		Do(ctx)

	//client.Index().Index(indexName).

	if err != nil {
		panic(err)
	}

	//fmt.Printf("Indexed with id=%v, type=%v, status=%v\n", doc.Id, doc.Type,doc.Status)
	subject = Subject{
		ID:     2,
		Title:  "千與千尋",
		Genres: []string{"劇情", "喜劇", "愛情", "戰爭"},
	}
	fmt.Println(subject.ID, doc)
	doc, err = client.Index().
		Index(indexName).
		Id(strconv.Itoa(subject.ID)).
		BodyJson(subject).
		Refresh("wait_for").
		Do(ctx)

	if err != nil {
		panic(err)
	}
	result, _ := client.Get().Index(indexName).Id(strconv.Itoa(subject.ID)).Do(ctx)
	fmt.Printf("%#v\n", string(result.Source))

	fmt.Println("================\n")
	Search1(client, ctx, "犯罪")
}
