package main

import (
	_ "net/http/pprof"
	"sync"
)

func main() {
	//#file,_:=os.Create("trace.heap")trace.Start(file)
	//defer trace.Stop()
	//sort.Float64s()
	var wg sync.WaitGroup

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			a := 0

			for i := 0; i < 1e6; i++ {
				a += 1
			}

			wg.Done()
		}()
	}

	wg.Wait()

	//http.ListenAndServe("0.0.0.0:9999", nil)
}
