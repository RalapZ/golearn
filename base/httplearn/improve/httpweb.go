package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func WebResponse(Resp http.ResponseWriter, Req *http.Request) {
	//fmt.Println(Req.RequestURI(("/list")))
	WebPath := Req.URL.Path[len("/list/"):]
	//defer func(){
	//	r:=recover()
	//	if Errinf,ok:=r.(*os.PathError);ok{
	//		log.Println(Errinf.Path,Errinf.Err,Errinf.Op)
	//	}
	//}()
	fmt.Println(string(WebPath))
	fileR, err := os.OpenFile(string(WebPath), os.O_EXCL|os.O_RDONLY, 0755)
	defer fileR.Close()
	if err != nil {
		log.Println(err)
	}
	str, _ := ioutil.ReadAll(fileR)
	Resp.Write(str)
	//WebReq:=bufio.NewReader(fileR)
	//for {
	//	str,_:=WebReq.ReadString('\n')
	//	Resp.Write([]byte(str))
	//	//fmt.Fprintf(Resp,WebReq.Re)
	//}

	fmt.Fprintf(Resp, "hello")
}

func main() {
	http.HandleFunc("/list/", WebResponse)
	http.ListenAndServe(":9000", nil)
}
