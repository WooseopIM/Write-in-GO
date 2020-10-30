package main

import (
	"fmt"
	"strings"
)

// 반복하는 함수
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println(("I'm done")) // lenAndUpper 함수가 값을 return하고 난 뒤 실행된다.
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, uppercase := lenAndUpper("nico")
	fmt.Println(totalLength, uppercase)
}
