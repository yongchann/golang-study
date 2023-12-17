package main

import "fmt"

/*
defer 문은 선언 시점에서 defer 내 사용된 파라미터 및 값을 평가한뒤 스택에 저장함
case1) 3번의 defer fmt.Println(*pInt) 에서는 해당 주소에 참조에 값을 직접 평가함

case2) 3번의 printPtrInt 에서 함수로 전달된 pItn은 동일한 주소값으로 평가되고, 실제 defer 실행시점에서 그 주소의 값을 읽음
따라서 main의 마지막에 pInt에 저장된 값이 500으로 바뀌었으므로 defer에서 동일한 500을 출력함.
*/

func printPtrInt(pInt *int) {
	fmt.Println(*pInt)
}

func main() {
	var pInt = new(int)

	*pInt = 10
	defer printPtrInt(pInt)  // 10
	defer fmt.Println(*pInt) // 500

	*pInt = 20
	defer printPtrInt(pInt)  // 20
	defer fmt.Println(*pInt) // 500

	*pInt = 30
	defer printPtrInt(pInt)  //30
	defer fmt.Println(*pInt) //500

	*pInt = 500
}
