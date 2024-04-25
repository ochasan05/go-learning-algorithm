package arrayshashing

import (
	"fmt"
	"sort"
	"strings"
)

/*　問題
Level：Easy
URL：https://leetcode.com/problems/contains-duplicate/description/

▪️問題 翻訳
与えられた二つの文字列 s と t について、t が s のアナグラムであれば true を返し、そうでなければ false を返します。
※アナグラムとは、通常、すべての元の文字を正確に一度使用して、異なる単語やフレーズの文字を並べ替えて形成される単語またはフレーズです。

▪️ Solution（解決策）
Sort：https://leetcode.com/problems/valid-anagram/solutions/4410240/three-lines-of-code-by-prodonik-java-c-python-go-js-ruby-c
Map：https://qiita.com/RyeWiskey/items/7b9eeec0bf35ac1434d7
*/

func ExecuteValidAnagram(t string, s string) {
	fmt.Println("### 1.ValidAnagram Start ###")
	fmt.Println(isAnagramSort(t, s))
	fmt.Println(isAnagramMap(t, s))
	fmt.Println(isAnagramMap2(t, s))
	fmt.Println("### 1.ValidAnagram End ###")
}

/*
*
■ Solution-1：スライス化してSort
・s,tをそれぞれスライス化してSort -> 文字列に戻し比較する
・Sortがボトルネック

実行結果
・Runtime：15ms
・Memory：6.75MB
*/
func isAnagramSort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss := strings.Split(s, "")
	tt := strings.Split(t, "")

	sort.Strings(ss)
	sort.Strings(tt)

	return strings.Join(ss, "") == strings.Join(tt, "")
}

/*
* ■ Solution-2：Map
・chatGPTに教えてもらった

実行結果
・Runtime：7ms
・Memory：3.06MB
*/
func isAnagramMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	// 文字列tの各文字の出現回数を減算
	// 出現回数が負になった場合、アナグラムではない
	// 片方にしかない値が存在した場合、100%負の値になる
	for _, char := range t {
		charCount[char]--

		if charCount[char] < 0 {
			return false
		}
	}

	// すべての文字の出現回数が一致した場合、アナグラムである
	return true
}

/*
* ■ Solution-3：Map2
・s,t用にMapを2つ作る
・s,tの文字列を1文字ずつKeyにし、Valueを出現回数にする
・s,tのMapを比較し、同じKeyのValueが全て同じ値かどうか比較する

実行結果
・Runtime：3ms
・Memory：3MB
*/
func isAnagramMap2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	anagram := newAnagram()
	anagram.addItems(s, t)

	for k, v := range anagram.sItems {
		v2, ok := anagram.tItems[k]
		if !ok {
			return false
		}

		if v != v2 {
			return false
		}
	}

	return true
}

type anagram struct {
	sItems map[rune]int
	tItems map[rune]int
}

func newAnagram() *anagram {
	return &anagram{
		sItems: make(map[rune]int),
		tItems: make(map[rune]int),
	}
}

func (anagram *anagram) addItems(s string, t string) {
	for _, v := range s {
		anagram.sItems[v]++
	}

	for _, v := range t {
		anagram.tItems[v]++
	}
}
