package main

import (
	"fmt"
	arrayshashing "learn-algoritm/arraysHashing"
)

func main() {
	/** Arrays & Hashing */
	fmt.Println("### Arrays & Hashing Start ###")

	arrayshashing.ExecuteContainsDuplicate([]int{1, 2, 3, 1})
	arrayshashing.ExecuteValidAnagram("anagram", "nagaram")
	arrayshashing.ExecuteTwoSum([]int{2, 7, 11, 15}, 9)
	arrayshashing.ExecuteGroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	arrayshashing.ExecuteTopKFrequentElements([]int{1,1,1,2,2,3}, 2)

	fmt.Println("### Arrays & Hashing End ###")
}
