package main

import (
	"fmt"
	"interview-quiz-heshuo/quiz5"
)

func main() {
	fmt.Println("Hello!")
	var num float64 = 12345
	fmt.Printf("Unit Converted %g -> %s", num, quiz5.Format(num))
}
