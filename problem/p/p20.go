package main

import "fmt"

func main() {
	s := "([)]"
	fmt.Println('{')
	fmt.Println('(')
	fmt.Println('[')
	fmt.Println(']')
	fmt.Println(')')
	fmt.Println('}')
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var calArr []int32
	var relatedArr map[int32]int32
	relatedArr = make(map[int32]int32)
	relatedArr[')'] = '('
	relatedArr[']'] = '['
	relatedArr['}'] = '{'
	for _, i2 := range s {
		if i2 == '(' || i2 == '[' || i2 == '{' {
			calArr = append(calArr, i2)
		} else {
			if len(calArr) <= 0 || calArr[len(calArr)-1] != relatedArr[i2] {
				return false
			} else {
				calArr = calArr[:len(calArr)-1]
			}
		}
	}

	return len(calArr) == 0

}
