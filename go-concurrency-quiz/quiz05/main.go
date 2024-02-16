package main

import (
	"fmt"
	"time"
)

func pingpong(name string, ch chan int) {
	for {
		i := <-ch
		fmt.Println(name, ":", i)

		i++
		ch <- i
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	go pingpong("goroutine A", ch)
	go pingpong("goroutine B", ch)

	ch <- 1

	for {
		time.Sleep(time.Second)
	}
}

/*
[Result]
goroutine B : 1
goroutine A : 2
goroutine B : 3
goroutine A : 4
goroutine B : 5
goroutine A : 6

[Exp]
main 에서 ch 에 값을 전송했고 이를 수신할 2개의 고루틴이 존재함
고루틴 A, B 중 먼저 실행되는 고루틴부터 수신하여 서로 pingpong 동작을 이어감
*/
