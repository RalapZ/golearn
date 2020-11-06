package main

import (
	"net/http"
	"net/url"
)

func HttpPostForm(httpurl string) {
	a := url.Values{}
	a.Add("name", "Ralap")
	http.PostForm()
}

func main() {

}
