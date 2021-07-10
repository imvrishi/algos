package reverse_string

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world!"
	fmt.Println(ReverseString(str))
}

func ReverseString(str string) string {
	arrayStr := strings.Split(str, "")
	reversedArray := &[]string{}
	AddToArray(arrayStr, reversedArray)
	return strings.Join(*reversedArray, "")
}

func AddToArray(array []string, reversedArray *[]string) {
	size := len(array)
	if size > 0 {
		*reversedArray = append(*reversedArray, array[size-1])
		array = array[:size-1]
		AddToArray(array, reversedArray)
	}
}

func ReverseStringRecursive(str string) string {
	if str == "" {
		return ""
	}

	return ReverseStringRecursive(str[1:]) + string(str[0])
}
