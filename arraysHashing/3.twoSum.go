package arrayshashing

import "fmt"

/*
　問題
Level：Easy
URL：https://leetcode.com/problems/two-sum/description/

▪️問題 翻訳
整数の配列 nums と整数の target が与えられた場合、その中から合計が target になる二つの数のインデックスを返します。
各入力には正確に一つの解があると仮定しても構いません。また、同じ要素を二度使用することはできません。
答えはどのような順番でも構いません。
*/
func ExecuteTwoSum(nums []int, target int) {
	fmt.Println("### 3.Two Sum Start ###")
	fmt.Println(twoSum(nums, target))
	fmt.Println("### 3.Two Sum End ###")
}

/*
*
■ Solution-1：ターゲット整数からnumsの要素を引く -> 差分値を検索する

実行結果
・Runtime：23ms
・Memory：3.5MB
*/
func twoSum(nums []int, target int) []int {
	var answer = []int{}

outer:
	for i, v := range nums {
		// 差分値を出す
		minus := target - v

	inner:
		//現在のIndex以降の値で差分値と一致する値がないかをチェックする
		for j := i + 1; j < len(nums); j++ {
			if minus != nums[j] {
				continue inner
			}

			answer = append(answer, i)
			answer = append(answer, j)
			break outer
		}
	}

	return answer
}
