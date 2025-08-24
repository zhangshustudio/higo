package leetcode

func twoSum(nums []int, target int) []int {
	visited := map[int]int{}
	for i, num := range nums {
		if index, ok := visited[target-num]; ok {
			return []int{index, i}
		}
		visited[num] = i
	}
	return []int{}
}

func longestCommonPrefix(strs []string) string {
	s0 := strs[0]
	for i, c := range s0 { // 从左到右
		for _, s := range strs { // 从上到下
			if i == len(s) || s[i] != byte(c) { // 这一列有字母缺失或者不同
				return s0[:i] // 0 到 j-1 是公共前缀
			}
		}
	}
	return s0
}
