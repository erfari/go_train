package main

import "fmt"

/*
Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

var nums = []int{2, 7, 11, 15}

var target = 9

func main() {
	fmt.Printf("%v", twoSum(nums, target))
}

/*
Runtime 31ms   Memory 3.54 MB
*/

func twoSum(nums []int, target int) []int {
	var fvalue int
	var response = make([]int, 0)
	for i := 0; i < len(nums); i++ {
		fvalue = target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if fvalue-nums[j] == 0 {
				response = append(response, i)
				response = append(response, j)
				return response
			}
		}
	}
	return response
}
