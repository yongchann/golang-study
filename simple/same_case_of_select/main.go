package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var case1cnt int32 = 0
var case2cnt int32 = 0

func recv(ch chan struct{}) {
	for {
		select {
		case <-ch:
			// fmt.Println("case1 recv")
			atomic.AddInt32(&case1cnt, 1)
		case <-ch:
			// fmt.Println("case2 recv")
			atomic.AddInt32(&case2cnt, 1)
		}
	}

}

func main() {
	ch := make(chan struct{})

	go recv(ch)
	wg := sync.WaitGroup{}
	cnt := 10000000
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			defer wg.Done()
			ch <- struct{}{}
		}()
	}

	wg.Wait()

	fmt.Println(case1cnt)
	fmt.Println(case2cnt)

}
