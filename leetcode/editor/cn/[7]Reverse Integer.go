package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

//Given a 32-bit signed integer, reverse digits of an integer.
//
// Example 1: 
//
// 
//Input: 123
//Output: 321
// 
//
// Example 2: 
//
// 
//Input: -123
//Output: -321
// 
//
// Example 3: 
//
// 
//Input: 120
//Output: 21
// 
//
// Note: 
//Assume we are dealing with an environment which could only store integers with
//in the 32-bit signed integer range: [−231, 231 − 1]. For the purpose of this pro
//blem, assume that your function returns 0 when the reversed integer overflows. 
// Related Topics 数学 
// 👍 2063 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

//先转换成字符串，字符串翻转后取出
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
	if ressult > math.MaxInt32  {
		fmt.Print(ressult)
		return 0
	}
	if ressult < math.MinInt32{
		return 0
	}
	if err != nil{
		fmt.Printf("error id %v",err)
	}
	return int(ressult)
}

//leetcode submit region end(Prohibit modification and deletion)
