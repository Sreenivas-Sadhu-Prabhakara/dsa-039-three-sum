package main

import (
	"fmt"
	"sort"
)

// ThreeSum returns all unique triplets of values summing to zero, sorted.
// Sort once; fix the smallest member; close each pair with two pointers.
func ThreeSum(nums []int) [][]int {
	a := append([]int{}, nums...)
	sort.Ints(a)
	out := [][]int{}
	for i := 0; i < len(a)-2; i++ {
		if a[i] > 0 {
			break // smallest member positive: no zero sum possible
		}
		if i > 0 && a[i] == a[i-1] {
			continue
		}
		l, r := i+1, len(a)-1
		for l < r {
			s := a[i] + a[l] + a[r]
			switch {
			case s == 0:
				out = append(out, []int{a[i], a[l], a[r]})
				l++
				r--
				for l < r && a[l] == a[l-1] {
					l++
				}
				for l < r && a[r] == a[r+1] {
					r--
				}
			case s < 0:
				l++
			default:
				r--
			}
		}
	}
	return out
}

func main() {
	fmt.Println(ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
}
