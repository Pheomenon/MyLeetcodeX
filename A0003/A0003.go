/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
通过次数722,330提交次数2,015,753

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

//func lengthOfLongestSubstring(s string) int {
//	m := map[byte]int{}
//	right, ans := -1, 0
//	for i := 0; i < len(s); i++ {
//		if i != 0 {
//			delete(m, s[i-1])
//		}
//		for right+1 < len(s) && m[s[right+1]] == 0 {
//			m[s[right+1]]++
//			right++
//		}
//		ans = max(ans, right-i+1)
//	}
//	return ans
//}
//
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	window := map[rune]int{}
	left, right := 0, 0
	res := 0
	for _, c := range s {
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[rune(d)]--
		}
		res = max(right-left, res)
	}
	return res
}
