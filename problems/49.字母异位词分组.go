/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */
package main

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int8][]string, 10010)
	for _, str := range strs {
		temp := [26]int8{}
		for i := 0; i < len(str); i++ {
			temp[str[i]-'a']++
		}
		hash[temp] = append(hash[temp], str)
	}
	ans := make([][]string, 0, len(hash))
	for _, v := range hash {
		ans = append(ans, v)
	}
	return ans
}

// @lc code=end
