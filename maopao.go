package main

import (
	"fmt"
)

func main() {

	a := [10]int{8, 5, 6, 2, 1, 0, 3, 4, 7, 9}
	var porint *[10]int
	porint = &a
	bubble_sort(porint, 10)
	print_result(porint, 10)

}

func bubble_sort(p *[10]int, n int) {
	var j int
	var tmpmun int
	for i := 0; i < n; i++ {
		for j := j + i; j < n; j++ {
			if p[j] > p[i] {
				tmpmun = p[i]
				p[i] = p[j]
				p[j] = tmpmun
			}
		}
	}
}

func print_result(p *[10]int, n int) {

	for k := 0; k < n; k++ {
		fmt.Println("数据是：", p[k])
	}
}
