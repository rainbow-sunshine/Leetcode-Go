package leetcode

import (
	"fmt"
)

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 示例 1: 
//
// 输入: 121
//输出: true
// 
//
// 示例 2: 
//
// 输入: -121
//输出: false
//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
// 
//
// 示例 3: 
//
// 输入: 10
//输出: false
//解释: 从右向左读, 为 01 。因此它不是一个回文数。
// 
//
// 进阶: 
//
// 你能不将整数转为字符串来解决这个问题吗？ 
// Related Topics 数学 
// 👍 1164 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	//解法1.转换成字符串然后翻转
	//if x <0{
	//	return false
	//}
	//s := strconv.Itoa(x)
	//s1 := reserve(s)
	//if s1 == s {
	//	return true
	//}
	//return false

	//解法2.翻转一半的数字
	if x < 0 ||(x%10 == 0 && x!=0){
		return false
	}
	revernum := 0
	for x > revernum {
		revernum = revernum*10 + x%10
		x = x/10
	}
	return x==revernum || x ==revernum/10
}

func reserve(str string) string{
	var res string
	length := len(str)
	for i ,_:= range str {
		j := length - i -1
		res += fmt.Sprintf("%c",str[j])
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
