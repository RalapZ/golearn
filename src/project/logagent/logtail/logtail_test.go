package logtail

import (
	"fmt"
	"testing"
	"time"
)

func TestTailLine(t *testing.T) {
	//初始化tail文件
	err := InitTail(LogFile, Config)
	if err != nil {
		fmt.Println(err)
	}
	//获取变化line
	go TailLine()
	go readall()
	time.Sleep(10 * time.Second)

	workerdone <- true

}
