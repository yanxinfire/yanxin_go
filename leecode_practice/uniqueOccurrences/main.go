package main

func main() {
	
}

func uniqueOccurrences(arr []int) bool {
	cnt:=[1001]int{}
	arrcount:=make(map[int]int)
	for _,v :=range arr{
		arrcount[v]++
	}
	for _,v:=range arrcount{
		cnt[v]++
		if cnt[v]>=2{
			return false
		}
	}
	return true
}