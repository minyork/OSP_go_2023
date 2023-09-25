package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year() // 위 변수 now 가져옴.
	fmt.Println(year)
	// fmt.Println(now.Month(), now.Minute(), now.Second())

	month := now.Month()
	fmt.Println(month)
}
