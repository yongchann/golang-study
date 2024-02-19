package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int32 = 0

func add(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < n; i++ {
		atomic.AddInt32(&count, 1)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4)

	go add(1000000, &wg)
	go add(1000000, &wg)
	go add(1000000, &wg)
	go add(1000000, &wg)

	wg.Wait()

	fmt.Println("count=", count)
}

/*
[Result]
count= 4000000

[Exp]
"sync/atomic" 패키지를 사용하여 add 연산을 atio
*/
