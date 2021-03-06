package main

import "fmt"

/*
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false。

示例 1：

输入：flowerbed = [1,0,0,0,1], n = 1
输出：true
示例 2：

输入：flowerbed = [1,0,0,0,1], n = 2
输出：false

提示：

1 <= flowerbed.length <= 2 * 104
flowerbed[i] 为 0 或 1
flowerbed 中不存在相邻的两朵花
0 <= n <= flowerbed.length
通过次数65,795提交次数193,042

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/can-place-flowers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	flowerbed := []int{1, 0, 0, 0, 0, 1}
	n := 2
	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append(flowerbed, 0)
	tmp := make([]int, len(flowerbed)+1)
	tmp[0] = 0
	for i := 1; i < len(flowerbed); i++ {
		tmp[i] = flowerbed[i-1]
	}
	blank := 0

	for i := 1; i < len(tmp)-1; i++ {
		if tmp[i] == 0 && tmp[i+1] == 0 && tmp[i-1] == 0 {
			tmp[i] = 1
			blank++
		}
	}

	if blank >= n {
		return true
	}

	return false
}
