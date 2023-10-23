package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv" // ParseInt
	"strings" // TrimSpace
)

func main() {

	fmt.Println("Input First number")
	reader := bufio.NewReader(os.Stdin)

	num1Str, err := reader.ReadString('\n')
	num1Str = strings.TrimSpace(num1Str)
	// num2, err := strconv.ParseInt(num, 10, 32) // 10진수  32bit = 8byte size
	num1, err := strconv.Atoi(num1Str)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Input Second number")
	num2Str, err := reader.ReadString('\n')
	num2Str = strings.TrimSpace(num2Str)
	num2, err := strconv.Atoi(num2Str)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("num1 = ", num1)
	fmt.Println("num2 = ", num2)
}
