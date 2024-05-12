package arrayshashing

import (
	"fmt"
	"sort"
)

/*　問題
Level：Medium
URL：https://leetcode.com/problems/top-k-frequent-elements/description/

▪️問題 翻訳
整数配列 nums と整数 k が与えられた場合、最も頻度の高い k 個の要素を返します。答えはどのような順番でも構いません。
*/

func ExecuteTopKFrequentElements(nums []int, k int) {
	fmt.Println("### 5.Top_K_FrequentElements Start ###")
	fmt.Println(topKFrequentElementsSortMap(nums, k))
	fmt.Println("### 5.Top_K_FrequentElements End ###")
}

func topKFrequentElementsSortMap(nums []int, k int) []int {
	frequentElements := newFrequentElements()

	for _, v := range nums {
		frequentElements.addItem(v)
	}

	var numbers = []int{}
	for _, v := range frequentElements.item {
		numbers = append(numbers, v)
	}

	sort.Ints(numbers)

	var res = []int{}
	for num, count := range frequentElements.item {
		if count >= numbers[len(numbers)-k] {
			res = append(res, num)
		}
	}

	return res
}

type frequentElements struct {
	item map[int]int
}

func newFrequentElements() *frequentElements {
	return &frequentElements{item: make(map[int]int)}
}

func (frequentElements *frequentElements) addItem(num int) {
	frequentElements.item[num]++
}
