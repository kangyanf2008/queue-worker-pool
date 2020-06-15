package main

import (
	"fmt"
	"time"
	"worker"
)

func main() {

	start := time.Now()
	w := worker.NewWokerPool(1000).Run()

	for i:=0; i<100000000; i++  {
		a:= &worker.PrintNum{Num: i}
		w.PutJob(a)
	}
	//w.Stop()
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)

}
