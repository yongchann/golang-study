package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	i := 0

	go func() {
		for {
			ch <- i
			i++
		}
	}()

	for {
		fmt.Println(<-ch)
	}
}

/*
[Result]
0
1
2
3
4
...

[Exp]
goroutine1 과 main 에서 각각 무한루프로 송수신을 처리하므로 i 가 증가하며 정상동작함
*/
