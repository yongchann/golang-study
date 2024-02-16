package main

import (
	"fmt"
	"sync"
)

var count = 0

func add(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < n; i++ {
		count++
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
count= 1138181

[Exp]
공유자원인 count 에 여러 고루틴이 동시 접근함에 따라 의도한 결과(4000000)가 나오지 않음
*/
