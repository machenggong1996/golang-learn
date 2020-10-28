package array

//1207. 独一无二的出现次数
func UniqueOccurrences(arr []int) bool {
	cnts := map[int]int{}
	for _, v := range arr {
		cnts[v]++
	}
	times := map[int]struct{}{}
	for _, c := range cnts {
		times[c] = struct{}{}
	}
	return len(times) == len(cnts)
}
