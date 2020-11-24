package main

import (
	"fmt"
	"time"
)

func main() {
	var unixTime int64 = 1606113368
	t := time.Unix(unixTime, 0)
	//strDate := t.Format(time.UnixDate)
	fmt.Println(t)
}
