package logtail

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

var (
	Config = tail.Config{
		ReOpen:    true,
		Location:  &tail.SeekInfo{0, 2},
		MustExist: false,
		Follow:    true,
		Poll:      true,
	}
	LogFile    = "logtail/test.log"
	ClientTail *tail.Tail
	err        error
	Lineschan  = make(chan string)
	workerdone = make(chan bool)
)

func InitTail(filename string, config tail.Config) error {
	ClientTail, err = tail.TailFile(filename, Config)
	return err
}

func TailLine() {
	//fmt.Println(Linescha)
	for true {
		select {
		case line := <-ClientTail.Lines:
			//fmt.Println(line.Text)
			Lineschan <- line.Text
			//fmt.Println(line.)
		case <-workerdone:
			return
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func readall() {
	for {
		select {
		case lin := <-Lineschan:
			fmt.Println(lin)
		default:
		}
	}

}
