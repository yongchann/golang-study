package main

import (
	"fmt"
	"time"
)

func pingpong(name string, ch chan int) {
	for {
		i := <-ch
		fmt.Println(name, ":", i)

		ch <- i
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)

	go pingpong("goroutine A", ch)
	go pingpong("goroutine B", ch)
}

/*
[Result]
 ~/golang-study/go-concurrency-quiz/quiz03 > go run main.go

[Exp]
고루틴 실행 후 main 함수가 종료되어 프로세스가 종료됨
*/
