package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {
	strArr := strings.Split(str, "")
	size := len(strArr)
	if size < 2 {
		return str
	}

	for i := size - 1; i >= size/2; i-- {
		strArr[i], strArr[size-i-1] = strArr[size-i-1], strArr[i]
	}

	return strings.Join(strArr, "")
}

func main() {
	fmt.Println(Reverse("My name is Rishi"))
}
