package main

import (
	"fmt"
	"github.com/machenggong1996/golang-learn/leetcode/array"
)

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//array.UniqueOccurrences(arr)
	res := array.SortByBits(arr)
	fmt.Println(res)
}
