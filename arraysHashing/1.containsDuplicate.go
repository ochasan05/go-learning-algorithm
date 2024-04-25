package arrayshashing

import (
	"fmt"
	"sort"
)

/*　問題
Level：Easy
URL：https://leetcode.com/problems/valid-anagram/description/
概要：
・スライス内の値で重複している値がどうか判定したい
・最低2つ以上あればTrue、それ以外はFalesを返却

▪️問題 翻訳
与えられた整数配列 nums について、配列内に少なくとも1つの値が2回以上現れる場合は true を返し、すべての要素が異なる場合は false を返します。

▪️ Solution（解決策）
URL：https://leetcode.com/problems/contains-duplicate/solutions/2552285/go-multiple-solutions-in-go-golang/
*/

func ExecuteContainsDuplicate(nums []int) {
	fmt.Println("### 1.containsDuplicate Start ###")

	// ■ Solution-1：　Brute-Force（総当たり）
	fmt.Println(containsDuplicateBruteForce(nums))

	// ■ Solution-2：　Sorting
	fmt.Println(containsDuplicateSorting(nums))

	// ■ Solution-3：Using a Multiset
	fmt.Println(containsDuplicateUsingMultiSet(nums))

	// ■ Solution-4：Using a Multiset2
	fmt.Println(containsDuplicateUsingMultiSet2(nums))

	fmt.Println("### 1.containsDuplicate End ###")
}

/*
■ Solution-1：　Brute-Force（総当たり）
・forループで現在のインデックスの値と次インデックスの値を比較する
・一番シンプルだが、計算量が最も多くなる

実行結果：Time Limit Exceeded
*/
func containsDuplicateBruteForce(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

/*
■ Solution-2：　Sorting
・すべての数字をソートし、隣り合ったペアを比較する
・ソート処理がボトルネックになる

実行結果
・Runtime：77ms
・Memory：8.02MB
*/
func containsDuplicateSorting(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

/*
■ Solution-3：Using a Multiset
・Mapを作成し、与えられる値をKey、出現回数をValueにする
・GoはSetが存在しないため、Mapを実装し再現する

実行結果
・Runtime：83ms
・Memory：11.09MB
*/
type MultiSet struct {
	items map[int]int
}

func newMultiSet() *MultiSet {
	return &MultiSet{
		items: make(map[int]int),
	}
}

func (s *MultiSet) addNums(nums []int) {
	for _, v := range nums {
		s.items[v]++
	}
}

// ここはnumで
func (s *MultiSet) hasDuplicate(num int) bool {
	occ, ok := s.items[num]
	if !ok {
		return false
	}

	return occ > 1
}

func containsDuplicateUsingMultiSet(nums []int) bool {
	set := newMultiSet()
	set.addNums(nums)

	for _, v := range nums {
		if set.hasDuplicate(v) {
			return true
		}
	}

	return false
}

/*
■ Solution-4：Using a Multiset2
・Mapを作成し、与えられる値をKey、空のStructをValueにする
・GoはSetが存在しないため、Mapを実装し再現する

実行結果
・Runtime：87ms
・Memory：9.58MB
*/
type Set struct {
	items map[int]struct{}
}

func newSet() *Set {
	return &Set{
		items: make(map[int]struct{}),
	}
}

func (s *Set) has(num int) bool {
	_, ok := s.items[num]
	return ok
}

func (s *Set) add(num int) {
	s.items[num] = struct{}{}
}

func containsDuplicateUsingMultiSet2(nums []int) bool {
	set := newSet()
	
	for _, v := range nums {
		if set.has(v) {
			return true
		}

		set.add(v)
	}

	return false
}