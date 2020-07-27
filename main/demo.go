package main

import (
	"fmt"
	"math"
	"strconv"
)
//=====忽略我，只是做测试用=======
func main()  {
	fmt.Print(reverse(-123))
}
func reverse(x int) int {

	flag := false
	if x <0{
		x = -x
		flag = true
	}else {
		flag = false
	}
	s := strconv.Itoa(x)
	var res string
	len := len(s)
	for i,_ := range s {
		res =res + string(s[len-i-1])
	}
	if true == flag{
		res = "-"+res
	}
	ressult ,err:= strconv.ParseInt(res,10,64)
	if float64(ressult) >math.Exp2(float64(232)) || float64(ressult) < math.Exp2(float64(-231)) {
		fmt.Print(ressult)
		return 0
	}
	if err != nil{
		fmt.Printf("error id %v",err)
	}
	return int(ressult)
}