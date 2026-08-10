package basic

import "strings"

func RepeatStr(repetitions int, str string) string {
	// result := ""
	// for i := 0; i < repetitions; i++ {
	// 	result += str
	// }
	// return result
	return strings.Repeat(str, repetitions)
}
