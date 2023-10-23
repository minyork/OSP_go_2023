package main

import (
	"fmt"
	"log"
)

func sayHello() { // 리턴값없는 일반적인 사,정,함
	fmt.Println("Hello")
}

// 리턴타입 ★
func processScore(kor int, eng int, math int) (int, int) {
	totalScore := kor + eng + math
	average := totalScore / 3 // 3.0으로 나누려면 캐스팅 필요 주의

	//fmt.Printf("%s님의 총점은 %d점, 평균은%d 입니다.\n", name, totalScore, average)
	// C언어 스타일 포멧은 printf 써야됨.
	return totalScore, average
}

func isDobuleNum(a int, b int) {
	// a에서 b까지 숫자중 짝수만 출력하는 함수
	val1 := a
	val2 := b

	if a > b { // 2 ~ 10 이런경우말고   10 ~ 2 이런거시 변수값 서로 바꿈.
		temp := a
		val1 = val2
		val1 = temp
	}

	for i := val1; i <= val2; i++ {
		if i%2 == 0 { // 짝수면 출력
			fmt.Print(i, " ")
		}
	}
}

// 메인 --
func main() {

	var number int
	fmt.Print("숫자를 입력하시오. : ")
	_, err := fmt.Scanln(&number) // 줄 단위로 입력 받는 Scanln
	// num,err := fmt.Scanln(&number) 여러 조합으로 가능.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n 입력 받은 숫자 : ", number)

	//os.Exit(0) // 정상종료
	//os.Exit(1) // 비정상 종료

	// == 섀도우 배우기 ==
	var float64 float64 = 9.1
	//var test float64 = 7.9 윗줄만 출력은 괜찮으나 다른 변수에 영향
	fmt.Println(float64)

	//패키지를 변수명으로 사용X    var fmt string
	//함수를 변수명으로 사용X      var append string

	//사용자 지정 함수
	sayHello() // 사용!

	// kor := 100
	// eng := 90
	// math := 100
	// name := "김민혁"
	// processScore(name, kor, eng, math)

	// kor2 := 95
	// eng2 := 80
	// math2 := 100
	// name2 := "노수댕"
	// processScore(name2, kor2, eng2, math2)

	t, a := processScore(100, 93, 90) // 리턴 스타일
	fmt.Printf("%s님의 총점은 %d점, 평균은 %d점 입니다\n", "미뇩", t, a)

	isDobuleNum(10, 2)
}
