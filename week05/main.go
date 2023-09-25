package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year() // 위 변수 now 가져옴.
	fmt.Println(year)
	// fmt.Println(now.Month(), now.Minute(), now.Second())

	month := now.Month()
	fmt.Println(month)

	// Replace practice
	brokenWords := "in#atc #i"                      // 고장난 문자
	replaceWords := strings.NewReplacer("#", "h")   // # 을 h로 바꿀
	fixedWords := replaceWords.Replace(brokenWords) // 바꿈.
	fmt.Println(fixedWords)

	// get value from the uesr
	fmt.Print("Enter Your name >> ")
	reader := bufio.NewReader(os.Stdin)
	inputName, err := reader.ReadString('\n') // ReadString이 복수개의 값을 리턴함.

	// 솔루션 1
	//  inputName, _ := reader.ReadString('\n')

	//솔루션 2
	// inputName, err := reader.ReadString('\n')
	// log.Fatal(err)

	// 솔루션 3
	if err != nil {
		log.Fatal(err)
	} // err 값이 nil 이 아니면 log 파일 남김.

	fmt.Println(inputName)

	fmt.Print("Enter Your Score >> ")
	inputScore, _ := reader.ReadString('\n')
	inputScore = strings.TrimSpace(inputScore)       // 쓸모없는 공백 제거
	score, err := strconv.ParseFloat(inputScore, 64) // string -> Float64

	// if score >= 90 {
	// 	grade := "A grade"
	// } else {
	// 	grade := "under A grade"
	// }
	// fmt.Println("You will get", grade)
	// 에러남. 왜? if 괄호 안에서 변수를 선언했기 때문에 밖에서 출력X

	//고침.
	var grade string // 밖에서 변수 선언 후
	if score >= 90 {
		grade = "A grade" // 대입
	} else {
		grade = "under A grade" // 대입
	}
	fmt.Println("You will get", grade)

	// 난수 발생 코드
	fmt.Println("Random Number")
	randomNum := rand.Intn(6) + 1 // 6 의미 : 0~5 근데 +1 해서 1~6
	fmt.Println(randomNum)

	//
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(100) + 1 // 1 ~ 100
	fmt.Println("Guess Number !!(1~100)")
	fmt.Println(answer)

	for i := 1; i <= 3; i++ { // i는 횟수
		fmt.Println("남은 기회 : ", 4-i)
		fmt.Print("Guess number :")
		inputNumberString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNumberString = strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoi(inputNumberString)
		// inputNumberString은 내가 입력한거 / 이거를 int로 바꿈.
		if err != nil {
			log.Fatal(err)
		}
		if inputNumber == answer {
			fmt.Println("Correct !! Good job")
			break
		} else if inputNumber < answer {
			fmt.Println("Guess Number is lower than answer.")
		} else if inputNumber > answer {
			fmt.Println("Guess Number is higher than answer.")
		}

	}

}
