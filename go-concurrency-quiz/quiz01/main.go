package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("goroutine 1 start")
		ch <- 1
		fmt.Println("goroutine 1 end")

	}()

	for i := 0; i < 3; i++ {
		fmt.Println("main goroutine loop [", i+1, "]")
		time.Sleep(time.Second)
	}
}

/*
[Result]
goroutine 1 start
main goroutine loop [ 1 ]
main goroutine loop [ 2 ]
main goroutine loop [ 3 ]

[Exp]
goroutine1 에서 ch 채널에 값을 보내고자하지만 이를 받아줄 고루틴이 없음
따라서 goroutine1 은 채널에 값을 전송하지 못하고 block 됨
main 의 for loop 가 종료되면서 모든 프로세스가 종료됨
*/
