package main

import "fmt"

//=====忽略我，只是做测试用=======
func main() {
	reserve("xeth")

}
func reserve(str string) {
	var res string
	length := len(str)
	for i, _ := range str {
		j := length - i - 1
		res += fmt.Sprintf("%c", str[j])
	}
}
