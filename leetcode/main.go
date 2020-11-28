package main

import (
	"fmt"
	"github.com/machenggong1996/golang-learn/leetcode/array"
)

func main() {

	//arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//array.UniqueOccurrences(arr)
	//res := array.SortByBits(arr)
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	res := array.RelativeSortArray(arr1, arr2)
	fmt.Println(res)
}
