package main

import "fmt"

func exists(items []int, value int) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}
func superman(n int, k int, chicken []int) int {
	max := 0
	for i := 0; i < n; i++ {
		current := 0

		for j := chicken[i]; j < k+chicken[i]; j++ {

			if exists(chicken, j) {
				current++
			}
		}
		if current > max {
			max = current
		}
	}
	return max
}
func main() {
	n1 := 5
	k1 := 5
	chicken1 := []int{2, 5, 10, 12, 15}
	maxSave1 := superman(n1, k1, chicken1)
	fmt.Println(maxSave1)
	n2 := 6
	k2 := 10
	chicken2 := []int{2, 5, 11, 12, 15, 20}
	maxSave2 := superman(n2, k2, chicken2)
	fmt.Println(maxSave2)
}
