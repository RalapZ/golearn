package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	id := getGID()
	return func() { fmt.Printf("g[%02d]: exit %s\n", id, name) }
}

func A1() {
	defer trace()()
	B1()
}

func B1() {
	defer trace()()
	C1()
}

func C1() {
	defer trace()()
	D()
}

func D() {
	defer trace()()
}

func A2() {
	defer trace()()
	B2()
}
func B2() {
	defer trace()()
	C2()
}
func C2() {
	defer trace()()
	D()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()

	time.Sleep(time.Millisecond * 50)
	A1()
	wg.Wait()
}
