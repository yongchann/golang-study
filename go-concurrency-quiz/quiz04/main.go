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
	pingpong("goroutine B", ch) // quiz03 에서 go 키워드만 하나 빠짐

}

/*
[Result]
~/golang-study/go-concurrency-quiz/quiz04 > go run main.go
fatal error: all goroutines are asleep - deadlock!

[Exp]
main 함수가 종료되진 않지만 모든 고루틴이 수신 대기상태가 되어 데드락 발생
*/
