package arrayshashing

import (
	"fmt"
	"sort"
)

/*　問題
Level：Medium
URL：https://leetcode.com/problems/group-anagrams/description/

▪️問題 翻訳
文字列の配列 strs が与えられた場合、アナグラムをグループ化します。答えはどのような順番でも構いません。
アナグラムとは、通常、すべての元の文字を正確に一度使用して、異なる単語やフレーズの文字を並べ替えて形成される単語またはフレーズです。
*/

/*
*
■ Solution-1：配列要素を（String型->byte型にキャスト）し、Sortした結果をMapに格納
・MapのKeyはSort後の文字列、ValueはSort前の値のスライス

実行結果
・Runtime：23ms
・Memory：7.3MB
*/
func ExecuteGroupAnagrams(strs []string) {
	fmt.Println("### 4.GroupAnagrams Start ###")
	fmt.Println(groupAnagramsSortMap(strs))
	fmt.Println("### 4.GroupAnagrams End ###")
}

func groupAnagramsSortMap(strs []string) [][]string {
	var res [][]string
	anagram := newGroupAnagram()

	for _, v := range strs {
		b := []byte(v)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		anagram.addItem(string(b), v)
	}

	for _, v := range anagram.items {
		res = append(res, v)
	}

	return res
}

type groupAnagram struct {
	items map[string][]string
}

func newGroupAnagram() *groupAnagram {
	return &groupAnagram{
		items: make(map[string][]string),
	}
}

func (a *groupAnagram) addItem(s1 string, s2 string) {
	a.items[s1] = append(a.items[s1], s2)
}
