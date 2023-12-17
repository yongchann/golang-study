package main

import (
	"fmt"
)

/*
Newton's method 컴퓨터가 제곱근을 구하는 방법
제곱근을 구하고자 하는 수를 k, 그 제곱근을 x 라고 할 때, k = x^2 이고, x에 대한 이차방정식 x^2 - k = 0의 양의 실근이 구하고자하는 값이 된다.
이 이차함수에 대해 접선의 x 절편이 구하고자 하는 해와 매우 근접해지는 특성을 이용한 계산법이다.
알고리즘의 효율이 메우 좋아 초기 z의 값이 커도 동작을 반복하며 빠르게 제곱근의 값에 가까워진다.
아래 풀이에서는 z = 1 로 설정힘
*/
func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 0; i < 10; i++ {
		diff := (z*z - x) / (2 * z)
		z -= diff
		fmt.Println(i, " ", z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
