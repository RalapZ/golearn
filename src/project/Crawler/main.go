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
//   Date  : 2020/11/05                                                                             //
//##################################################################################################//
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"regexp"
)

const city = `<a cityid="([0-9]*)" href="//(www\.autohome\.com\.cn/[0-9a-zA-Z]*/cheshi/)" target="_blank">([^<>]+)</a>`

func ParserCity(htmlstr string) {
	re := regexp.MustCompile(city)
	k := re.FindAllStringSubmatch(htmlstr, -1)
	for _, v := range k {
		fmt.Println(v)
	}
	fmt.Println(k)
}

func Fatcher(httpurl string) {
	client := http.Client{}
	Url, err := http.NewRequest("GET", httpurl, nil)
	if err != nil {
		panic(err)
	}
	resp, _ := client.Do(Url)
	defer resp.Body.Close()
	BodyReader := bufio.NewReader(resp.Body)
	BodyReader.ReadLine()
	e := GetEncodeFormat(BodyReader)
	fmt.Println(e)

	k := transform.NewReader(BodyReader, e.NewDecoder())
	str, _ := ioutil.ReadAll(k)
	//fmt.Println(string(str))
	ParserCity(string(str))

}

func GetEncodeFormat(r *bufio.Reader) encoding.Encoding {
	bytest, _ := r.Peek(1024)
	EncodeType, Name, Cer := charset.DetermineEncoding(bytest, "")
	fmt.Println(EncodeType, Name, Cer)
	return EncodeType
}

func main() {
	Fatcher("https://www.autohome.com.cn/AreaList.html")
}
