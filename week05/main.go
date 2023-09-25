package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
}
