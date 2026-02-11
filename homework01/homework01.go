package homework01

import "fmt"

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	var result int = 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func IsPalindrome(x int) bool {
	return isPalindrome_0(x)
}

// 2. 回文数
// 判断一个整数是否是回文数
func isPalindrome_0(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	var cache []int
	for x > 0 {
		last := x % 10
		x = x / 10
		cache = append(cache, last)
	}

	size := len(cache)
	if size == 0 {
		return false
	}

	for i, v := range cache {
		if v != cache[size-1-i] {
			return false
		}
	}

	return true
}

func isPalindrome_1(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func isMatch(left, right byte) bool {
	switch {
	case left == '(' && right == ')':
		return true
	case left == '[' && right == ']':
		return true
	case left == '{' && right == '}':
		return true
	default:
		return false
	}
}

func IsValid(s string) bool {
	var cache []byte
	for i := 0; i < len(s); i++ {
		cur := s[i]
		var next byte = '0'
		if i < len(s)-1 {
			next = s[i+1]
		}
		if isMatch(cur, next) {
			i++
		} else {
			if len(cache) != 0 && isMatch(cache[len(cache)-1], cur) {
				if len(cache) > 1 {
					cache = cache[:len(cache)-1]
				} else {
					cache = []byte{}
				}
			} else {
				cache = append(cache, cur)
			}
			fmt.Printf("cache: %v\n", cache)
		}
	}

	if len(cache) == 0 {
		return true
	}

	return false
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// TODO: implement
	return ""
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// TODO: implement
	return nil
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// TODO: implement
	return 0
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// TODO: implement
	return nil
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// TODO: implement
	return nil
}
