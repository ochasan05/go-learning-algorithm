package main

import (
	"fmt"
	arrayshashing "learn-algoritm/arraysHashing"
)

func main() {
	/** Arrays & Hashing */
	fmt.Println("### Arrays & Hashing Start ###")

	nums := []int{1, 2, 3, 1}
	arrayshashing.ExecuteContainsDuplicate(nums)
	arrayshashing.ExecuteValidAnagram("anagram", "nagaram")
	arrayshashing.ExecuteTwoSum(nums, 5)
	fmt.Println("### Arrays & Hashing End ###")
}
