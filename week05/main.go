package main

import (
	"bufio"
	"fmt"
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
	inputName := reader.ReadString('\n')
	fmt.Println(inputName)
}
