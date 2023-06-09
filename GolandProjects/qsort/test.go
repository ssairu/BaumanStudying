package main

import "fmt"

var a []int = []int{11, 2, 31, 4, 0, 61, 12, 10, 9, 2345, -234, 34, 345, 2, -43, 34, 6, 23, -8}

func less1(i, j int) bool {
	return a[i] < a[j]
}

func swap1(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func mid(a, b, c int, less func(i, j int) bool) int {
	if (!less(a, b) && !less(b, c)) || (!less(c, b) && !less(b, a)) {
		return b
	}
	if (!less(b, a) && !less(a, c)) || (!less(c, a) && !less(a, b)) {
		return a
	}
	if (!less(a, c) && !less(c, b)) || (!less(b, c) && !less(c, a)) {
		return c
	}
	return -10000
}

func sort1(left, right int,
	less func(i, j int) bool,
	swap func(i, j int)) {

	n := right - left + 1

	if n <= 5 {
		for i := 1; i < n; i++ {
			for j := left; j <= right-i; j++ {
				if less(j+1, j) {
					swap(j, j+1)
				}
			}
		}
	} else {
		var l, r int = left, right
		m := (left + right) / 2
		m = mid(left, right, m, less)
		for l < r {
			if !less(m, l) {
				l++
			} else if !less(r, m) {
				r--
			}
			if less(m, l) && less(r, m) {
				swap(l, r)
			}
		}
		sort1(left, l-1, less, swap)
		sort1(l, right, less, swap)
	}

}

func qsort(n int,
	less func(i, j int) bool,
	swap func(i, j int)) {
	sort1(0, n-1, less, swap)
}

func main() {
	fmt.Scan(&a)
	qsort(len(a), less1, swap1)
	fmt.Println(a)
}
