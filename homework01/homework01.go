package homework01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	var count = make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	var result int = 0
	for num, count := range count {
		if count == 1 {
			result = num
			break
		}
	}
	return result
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	var str = strconv.Itoa(x)
	var chars = strings.Split(str, "")

    var restr string = ""
	for i :=len(chars)-1 ; i >=0 ; i-- {
		restr = restr + chars[i] 
	}

	if str == restr {
		return true
	} else {
		return false
	}
	
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	stack := []byte{}

    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for i := 0; i < len(s); i++ {
        ch := s[i]

        if ch == '(' || ch == '[' || ch == '{' {
            stack = append(stack, ch)
        } else {
            if len(stack) == 0 {
                return false  
            }
            top := stack[len(stack)-1]
            if pairs[ch] != top {
                return false  
            }
            stack = stack[:len(stack)-1]  
        }
    }

    return len(stack) == 0  // 栈空则全部匹配成功
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != ch {
				return strs[0][:i]
			}
		}
	}

	return strs[0]  
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	var result int 
	for _,v  := range digits {
		result = result * 10 + v
	}
	result += 1
	var slice []int 
	for result > 0 {
		slice = append( []int{result % 10}, slice...)
		result = result / 10
	}

	return slice
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
    i := 0
    for j := 1; j < len(nums); j++ {
        if nums[i] != nums[j] {
            i++
            nums[i] = nums[j]  
        }
    }

    return i + 1 
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    merged := [][]int{intervals[0]}

    for i := 1; i < len(intervals); i++ {
        current := intervals[i]
        last := merged[len(merged)-1]  

        if current[0] <= last[1] {
            last[1] = max(last[1], current[1])
        } else {
            merged = append(merged, current)
        }
    }

	return merged
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	var res []int
	for i:= 0; i <= len(nums)-1; i++ {
		for j:=i+1; j <= len(nums)-1; j++ {
			if nums[i] + nums[j] == target {
				res = append(res, i)
				res = append(res, j)
				fmt.Println(res)
				return res
			}
		}
	}
	return nil
}