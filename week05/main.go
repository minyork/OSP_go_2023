package main

import (
	"fmt"
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
	brokenWords := "in#atc #i"
	replaceWords := strings.NewReplacer("#", "h")
	fixedWords := replaceWords.Replace(brokenWords)
	fmt.Println(fixedWords)
}
