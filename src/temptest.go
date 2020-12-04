//##################################################################################################//
//                   			         ┌─┐       ┌─┐ + +                                          //
//                   			      ┌──┘ ┴───────┘ ┴──┐++                                         //
//                   			      │       ───       │++ + + +                                   //
//                   			      ███████───███████ │+                                          //
//                   			      │       ─┴─       │+                                          //
//                   			      └───┐         ┌───┘                                           //
//                   			          │         │   + +                                         //
//                   			          │         └──────────────┐                                //
//                   			          │                        ├─┐                              //
//                   			          │                        ┌─┘                              //
//                   			          └─┐  ┐  ┌───────┬──┐  ┌──┘  + + + +                       //
//                   			            │ ─┤ ─┤       │ ─┤ ─┤                                   //
//                   			            └──┴──┘       └──┴──┘  + + + +                          //
//                   			      神兽出没               永无BUG                                 //
//   Author: Ralap                                                                                  //
//   Date  : 2020/11/13                                                                             //
//##################################################################################################//
package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var ESClient *elastic.Client
var IndexName = "skywalkinginfor1"
var ctx = context.Background()

const Mapping = `
{
    "mappings": {
        "properties": {
            "StartTime": {
                "type": "date"
            },
			"Starttimerecv": {
                "type": "text"
            },
            "ScopeId": {
                "type": "text"
            },
			"Scope": {
                "type": "text"
            },
			"BInfo":{
				"properties": {
					"BusinesType": {
						"type": "text"
					},
					"BusinesName": {
						"type": "text"
					}
				}
			},
			"Name": {
                "type": "text"
            },
			"id0": {
                "type": "text"
            },
			"id1": {
                "type": "text"
            },
			"RuleName": {
                "type": "text"
            },
            "AlarmMessage": {
                "type": "text"
            }
        }
    }
}`

func main() {
	var ctx = context.Background()
	Url := []string{"http://10.10.8.151:9200/"}
	//ESClientConn(Url)
	ESClient2, err := elastic.NewClient(elastic.SetURL(Url...), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	ESClient = ESClient2
	fmt.Println("TEst", ESClient)
	exists, err := ESClient.IndexExists(IndexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := ESClient.CreateIndex(IndexName).BodyString(Mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}

}
