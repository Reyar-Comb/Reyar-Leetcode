/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */
package main

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	chars := make([]int, 26)
	for _, char := range p {
		chars[char-'a']++
	}

	differ := 0
	for _, char := range chars {
		if char != 0 {
			differ++
		}
	}
	left := 0
	right := len(p)
	ans := []int{}
	for i := left; i < right; i++ {
		chars[s[i]-'a']--
		if chars[s[i]-'a'] == 0 {
			differ--
		} else if chars[s[i]-'a'] == -1 {
			differ++
		}
	}

	if differ == 0 {
		ans = append(ans, 0)
	}

	if right >= len(s) {
		return ans
	}

	for {
		chars[s[right]-'a']--
		if chars[s[right]-'a'] == 0 {
			differ--
		} else if chars[s[right]-'a'] == -1 {
			differ++
		}

		chars[s[left]-'a']++
		if chars[s[left]-'a'] == 0 {
			differ--
		} else if chars[s[left]-'a'] == 1 {
			differ++
		}

		left++
		right++

		if differ == 0 {
			ans = append(ans, left)
		}

		if right == len(s) {
			break
		}

	}
	return ans
}

// @lc code=end
