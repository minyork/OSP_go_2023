package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	fmt.Println(math.Floor(3.14)) // 소수점 내림.  결과 3
	fmt.Println(math.Ceil(3.14))  // 소수점 올림.  결과 4

	fmt.Println(strings.Title("오픈소스프로그래밍 \n 줄바꿈")) // 구버전에서 쓰던거라 앞으론X

	fmt.Println('A') // ACSII 코드 ( 아스키코드 ) 라 65라 뜸.
	fmt.Println('김')

	fmt.Println(strings.Title("open source programming go!"))
	// go 언어에서는 대문자를 넣어줌. open -> Open 처럼

	fmt.Println(strings.Title("============================"))
	// reflect 라는 패키지를 import 시 하는 것.
	// 내가 import 안써줘도 relflect. 쓰면 알아서 추가해줌.
	fmt.Println(reflect.TypeOf(3.124))
	fmt.Println(reflect.TypeOf(100))
	fmt.Println(reflect.TypeOf('B'))
	fmt.Println(reflect.TypeOf("미뇩"))
	fmt.Println(reflect.TypeOf(true))

	var a int = 9 // 변수 타입이 뒤로감. ★
	var b float32 = 3.14
	var c string = "수댄"
	var d, e bool = true, false

	fmt.Println(a, b, c, d, e)

	fmt.Println(strings.Title("============================"))
	// zero value : 할당 안한 값의 기본 값.

	var val1 int     // 0
	var val2 float32 //0
	var val3 string  //빈칸
	var val4 bool    //false

	fmt.Println(val1, val2, val3, val4)

	val5 := 'Z'
	val6 := "묙"
	val7 := 123
	val8 := 3.141592
	Val9 := "변수명이 대문자로 시작하면 다른 패키지서 사용가능함."
	fmt.Println(val5, val6, val7, val8, Val9)

	fmt.Println(strings.Title("============================"))

	// type 이 다르면 비교, 수학 연산, 비교(할당)이 불가능함.

	var val10 int = 10
	var val11 float32 = 3.14
	//fmt.Println(val10 > val11) type이 달라서 비교 불가.
	fmt.Println(val10 > int(val11))     // type 형변환(캐스팅)
	fmt.Println(float32(val10) > val11) // type 형변환(캐스팅)

	//fmt.Println(val10 * val11) type이 달라서 계산 불가.
	fmt.Println(val10 * int(val11))     // type 형변환(캐스팅)
	fmt.Println(float32(val10) * val11) // type 형변환(캐스팅)

}
